package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/lolorenzo777/hellogoapp/pkg/webserver"
	"github.com/sunraylab/verbose"
)

func main() {
	// get --verbose flag
	fverbose := flag.Bool("verbose", false, "verbose output")
	if fverbose != nil && *fverbose {
		verbose.IsOn = true
		fmt.Println("verbose is ON")
	}

	// get --env flag
	strenv := "dev"
	env := flag.String("env", "dev", "{file}.env environement file to load. dev by default.")
	if env != nil {
		strenv = *env
	}
	strenv = "./configs/" + strenv + ".env"

	// load environment variables
	err := godotenv.Load(strenv)
	if err != nil {
		log.Fatalf("Error loading .env variables: %s", err)
	}

	// run the web serverj
	webserver.RunWebServer()

}
