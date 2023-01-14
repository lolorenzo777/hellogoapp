package webserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sunraylab/verbose"
)

// ServeApiHealthj
func ApiServeHealth() func(http.ResponseWriter, *http.Request) {
	counter := 0
	return func(w http.ResponseWriter, r *http.Request) {
		counter++
		json.NewEncoder(w).Encode(map[string]string{"health": "live", "counter": strconv.Itoa(counter)})
	}
}

func RunWebServer() *http.Server {

	// get .env variables
	spa_staticfiledir := strings.ToLower(strings.Trim(os.Getenv("SPA_STATICFILEDIR"), " "))
	if spa_staticfiledir == "" {
		spa_staticfiledir = "./web/static"
	}
	http_port := strings.Trim(os.Getenv("HTTP_PORT"), " ")
	http_rwTimeout, _ := strconv.Atoi(os.Getenv("HTTP_RWTIMEOUT"))
	if http_rwTimeout <= 0 {
		http_rwTimeout = 15
	}
	http_idleTimeout, _ := strconv.Atoi(os.Getenv("HTTP_IDLETIMEOUT"))
	if http_idleTimeout <= 0 {
		http_idleTimeout = 15
	}
	http_cache_control := false
	if strings.ToLower(strings.Trim(os.Getenv("HTTP_CACHE_CONTROL"), " ")) == "true" {
		http_cache_control = true
	}

	// let's go
	fmt.Printf("Starting the SPA web server serving pages from %q and APIs on port %s\n", spa_staticfiledir, http_port)

	// configure the server, with or without trailing slash is the same route
	webrouter := mux.NewRouter().StrictSlash(true)
	//webrouter.HandleFunc("/", HomeHandler).Methods("GET")

	apirouter := webrouter.PathPrefix("/api").Subrouter()
	apirouter.HandleFunc("/health", ApiServeHealth())

	// the main handler serve spa files
	webrouter.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir(spa_staticfiledir)).ServeHTTP(w, r)
	})

	// add middleware to remove cache if requested in config file
	if !http_cache_control {
		fmt.Println("cache is off")
		webrouter.Use(middlewareNoCache)
	}
	webrouter.Use(middlewareJsonContent)

	// setup timeout
	srv := &http.Server{
		Addr:         http_port,
		WriteTimeout: time.Duration(http_rwTimeout) * time.Second,
		ReadTimeout:  time.Duration(http_rwTimeout) * time.Second,
		IdleTimeout:  time.Duration(http_idleTimeout) * time.Second,
	}

	// add middleware to log every request if in verbose mode
	if verbose.IsOn {
		fmt.Println("logging is on")
		srv.Handler = NewLogger(webrouter)
	} else {
		srv.Handler = webrouter
	}

	// listen and serve in a go routine to allow catching shutdown clean request in parallel
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL will not be caught.
	chansig := make(chan os.Signal, 1)
	signal.Notify(chansig, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	// Block until we receive a shutdown signal.
	<-chansig

	// Start the clean shutdown process.
	// Create a deadline to wait for, longer than the rwTimeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(http_rwTimeout)+time.Second*10)
	defer func() {
		signal.Stop(chansig) // clean stop listening os
		cancel()             // ensure clean cancel the context, so write ctx.Done()
	}()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	fmt.Println("SPA web server shutdown")
	return nil
}
