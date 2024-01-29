# Book Management Service

A Programming exercise I performed to implement a simple book management software. This software provides a REST API, store its data in a relational database and has a CLI client called `librarian` to interact with it.

## OVERVIEW

## SETUP 

### PREREQUISITES

* Latest version of Docker

* Lates version of the Comand line utility `Make` 

* Go v1.17

* an OS that can support the rest of these requirements

### STEPS

Start by opening a terminal window and pulling the MySQL image from Docker Hub

```bash
docker pull mysql
```

Then use the following make target from this repository to get a dockerized MySQL instance up and running 

```bash
make db-run
```

Exec into the MySQL instance using this make target: 

```bash
sudo make db-exec
```

Create a new user for Book Management Service to interact with the database as:

```bash
CREATE USER 'admin'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO 'admin'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
```

Create book_management_service_database in the MySQL instance:

```bash
CREATE DATABASE `book_management_service_database`;
```

Now use the following comand to exit the MySQL REPL:

```bash
exit
```

Now use the following make target to spin up the Book Management Service

```bash
make run
```

Now open a second Terminal window and run the rest of the setup commands there.

Before we can continue with the rest of the SETUP we'll need to build the object for the CLI using the following make target: 

```bash
make build-cli
```
The rest of the commands all use the `librarian` command line utility that is included in this repository to interact with the Book Management Service. Execute the following make target to generate books for your Book Management Service instance:

```bash
make generate-books
```

Now you can use the librarian cli 

```bash
make get-books
```

## EXAMPLES

Please complete the previous SETUP section from start to finish or these examples aren't going to work.