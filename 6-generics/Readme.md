# Getting Started with Generics.

[Generics](https://go.dev/doc/tutorial/generics) section of the Getting started Golang tutorial, but dockerized.

## Setup

Execute any go commands (such as `mod init` or `run .`) through a docker container:

```
docker run -it -v $(pwd):/app -w /app --user $(id -u):$(id -g) golang:latest bash
```

