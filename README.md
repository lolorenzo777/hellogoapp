# hellogoapp

getting started with [go-app](https://go-app.dev/) and wasm

> Go-app is a package for building progressive web apps (PWA) with the Go programming language (Golang) and WebAssembly (Wasm).

## Issue

``go-app`` looks very interesting but

1. there's few documentation available and the lifecycle is not trivial.
1. all the code used for the back is compiled and served with the webassembly
1. the directory structure, with `/web/` is mandatory, trying to change it cause many troubles
1. not been able to insert a component within a page after first rendering. (back to issue 1)


## Use taskfile to build and run

use [taskfile](https://taskfile.dev/) tool to build and run. see [Taskfile.yaml](Taskfile.yaml) config file for available tasks.

use `AIR` for livereloading the server side.

the backend but also the frontend (the wasm) need to be built, and only the server need to be launched.

In develomment mod, launch the front in a frst terminal, with the ``--watch`` for auto-rebuild
```bash
$ task dev_front -t /build/Taskfile.yaml --watch
```

and launch the back in a second terminal
```bash
$ task live_back -t /build/Taskfile.yaml 
```

This task will run air for watching any change on the server side and to re-run the server.

