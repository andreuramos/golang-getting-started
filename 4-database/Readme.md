# Database Access

[Database access](https://go.dev/doc/tutorial/database-access) section of the getting started golang tutorial, but dockerized

## Setup

build the containers with:
```
docker-compose up -d
```

setup the database with:
```
docker exec -it mysql-container mysql -uroot
mysql -uroot -p
test_db
```
then create the database:
```
mysql> create database recordings;
mysql> use recordings;
mysql> source /database/create-tables.sql
```

to enter the `golang` or `mysql` containers:
```
docker exec -it (golang-container or mysql-container) bash
```