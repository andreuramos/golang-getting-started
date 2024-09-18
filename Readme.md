# Intro to golang

Code following the [getting started](https://go.dev/doc/tutorial/getting-started) tutorial, but dockerized.

## setup

In order to setup docker and init the module from a docker container run:

```
docker run -it -v $(pwd):/app -w /app --user $(id -u):$(id -g) golang:latest bash
```

and then inside the container

```
go mod init example/hello
```

## run

As the `/.cache` directory can't be created in the default location to specify a custom cache directory for this build, the `go run` command needs to define a `GOCACHE` value:

```
GOCACHE=/app/.cache go run .
```