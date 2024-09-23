# Multi-modules workspaces

Following [this](https://go.dev/doc/tutorial/workspaces) tutorial, dockerized

## Run

```
docker run -it -v $(pwd):/workspaces -w /workspaces --rm --user $(id -u):$(id -g) --name golang_workspces golang:latest bash
```