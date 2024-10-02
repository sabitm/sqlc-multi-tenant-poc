gen:
	# protoc -I ./proto --go_out=./ --go-grpc_out=./ --grpc-gateway_out=./ ./proto/main.proto
	# protoc -I ./proto --go_out=./ --go-grpc_out=./ --grpc-gateway_out=./ ./proto/account.proto
	sqlc generate

run:
	go run .

test:
	go test -v ./...

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

mysql-up: mysql-down
	docker run -d -p 3306:3306 --name mysql \
		-e MYSQL_ROOT_PASSWORD=123456 \
		-e MYSQL_USER=user \
		-e MYSQL_PASSWORD=123456 \
		mysql:latest

mysql-down:
	docker container stop mysql
	docker container rm mysql
