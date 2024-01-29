
db-run:
	docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=password -p 3306:3306 -d mysql

db-exec: 
	docker exec -it mysql-container mysql -uroot -ppassword

run: 
	go run cmd/book-management-service/main.go

build-cli: 
	go build cmd/librarian-cli/librarian.go 

generate-books: 
	./librarian create-book "harry potter and the philosophers stone" "J.K.Rowling" "fantasy" 
	./librarian create-book "harry potter and the chamber of secrets" "J.K.Rowling" "fantasy" 
	./librarian create-book "harry potter and the prisoner of azkaban" "J.K.Rowling" "fantasy" 
	./librarian create-book "fellowship of the ring" "J.R.Tolkien" "fantasy" 
	./librarian create-book "how to win friends and influence people" "Dale Carnegie" "self help" 
	./librarian create-collection "J.K.Rowling"
	./librarian add-book 1 1
	./librarian add-book 1 2
	./librarian add-book 1 3 

get-books: 
	./librarian get-books
