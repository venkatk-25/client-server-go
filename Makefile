build:
	go build -o bin/client-server ./cmd/...

server:
	bin/client-server server

client-grpc:
	bin/client-server client grpc

client-tcp:
	bin/client-server client tcp