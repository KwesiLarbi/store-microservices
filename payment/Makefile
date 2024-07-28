run:
	DB_DRIVER=mysql \
	DATA_SOURCE_URL=root:verysecretpass@tcp(127.0.0.1:3306)/payments \
	APPLICATION_PORT=3001 \
	ENV=development \
	go run cmd/main.go

docker:
	docker run -p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=verysecretpass \
	-e MYSQL_DATABASE=payment mysql