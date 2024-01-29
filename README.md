# Book Management Service

A Programming exercise I performed to implement a simple book management software. This software provides a REST API, store its data in a relational database and has a CLI client called `librarian` to interact with it.

## OVERVIEW

## TODOs:

* all the tests

* include a Publication date field on the book model and schema, and also filtering and cli support for this new field 

* include updating for books, removal of books from collections, and deletion for books and collections

## SETUP 

### PREREQUISITES

* Latest version of Docker

* Lates version of the Comand line utility `Make` 

* Go v1.17

* an OS that can support the rest of these requirements

### STEPS

Note that some of these command may need to be ran slightly differently depending on the operating system you're executing them in as they've only been tested on MacOS. The application itself should be good to run on Mac OS, Linux, and Windows.

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

Please complete the previous SETUP section from start to finish or these examples aren't going to work. In additon to to the previous coonstraint some of these examples are dependent on being executed directly after the SETUP and in the order listed below. Sequentially dependent examples are Annotated as: `SEQUENTIALLY_DEPENDANT`

Create a book:

```bash
./librarian create-book "the running grave" "J.K.Rowling" "crime fiction" 
```

Filter all books by author:

```bash
./librarian filter-by-author --author="J.K.Rowling"
```

Filter all books by genre:

```bash
./librarian filter-by-genre --genre="fantasy"
```

Create a collection:

```bash
./librarian create-collection "J.R.Tolkien"
```

Get All Books From a Collection

```bash
./librarian get-collection-books 1
```

Add a book to a collection `SEQUENTIALLY_DEPENDANT`:

```bash
./librarian add-book 1 6
```

List all collections:

```bash
./librarian get-collections
```

Filter all books in a collection by genre:

```bash
./librarian filter-collection-by-genre 1 --genre="fantasy"
```

Add a book to the first collection `SEQUENTIALLY_DEPENDANT`(I know it's not J.K.Rowling but I don't realy care either):

```bash
./librarian add-book 1 4
```

Filter all books in a collection by author:

```bash
./librarian filter-collection-by-author 1 --author="J.K.Rowling"
```