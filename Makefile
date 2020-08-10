.PHONY : pb server client

all: pb server client

pb:
	protoc -I=pb --go_out=paths=import:. --go-grpc_out=paths=import:. pb/*.proto

server:
	go build -o bin/server server/*.go

client:
	go build -o bin/client client/*.go

