# API Webservices with Gin

[RESTful api with gin]() section of the getting started golang tutorial, but dockerized.

## setup

To init the module, install dependencies etc run a golang container
```
docker run -it -v $(pwd):/app -w /app --user $(id -u):$(id -g) golang:latest bash
```

once inside the container

```
export GOCACHE=/app/.cache
go mod init example/webservice-with-gin
```

or any other go commands like `go run .` etc