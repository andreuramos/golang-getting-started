# Database Access

[Database access](https://go.dev/doc/tutorial/database-access) section of the getting started golang tutorial, but dockerized

## Setup

build the containers with:
```
docker-compose up -d
```

to enter the `golang` or `mysql` containers:
```
docker exec -it (golang-container or mysql-container) bash
```