# Golang modules

Golang [Create a Go Module](https://go.dev/doc/tutorial/create-module) tutorial, but dockerized

## setup

Run the golang container:

```
docker run -it -v $(pwd):/apps -w /apps --user $(id -u):$(id -g) --name golang_tutorial golang:latest bash
```
and once inside the container set the cache directory to `/apps/.cache`

```
export GOCACHE=/apps/.cache
```