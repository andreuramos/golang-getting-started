# Getting started with Fuzzing

Golang [getting started with fuzzing]() tutorial, but dockerized

## setup

run a golang docker container:
```
docker run -it -v $(pwd):/app -w /app --user $(id -u):$(id -g) golang:latest bash
```

then execute any go commands such as `go run .` or `go test`