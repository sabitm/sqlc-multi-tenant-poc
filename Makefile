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
