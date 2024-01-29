
db-run:
	docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=password -p 3306:3306 -d mysql

db-exec: 
	docker exec -it mysql-container mysql -uroot -ppassword

run: 
	go run cmd/book-management-service/main.go

build-cli: 
	go build cmd/librarian-cli/librarian.go 

generate-books: 
	./librarian create-book "harry potter and the philosophers stone" "J.K.Rowling" "fiction" 
	./librarian create-book "harry potter and the chamber of secrets" "J.K.Rowling" "fiction" 
	./librarian create-book "harry potter and the prisoner of azkaban" "J.K.Rowling" "fiction" 
	./librarian create-book "fellowship of the ring" "J.R.Tolkien" "fiction" 
	./librarian create-book "how to win friends and influence people" "Dale Carnegie" "fiction" 

get-books: 
	./librarian get-books
