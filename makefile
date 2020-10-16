.PHONY: compile
compile: ## Compile the proto file.
	protoc -I pkg/proto/credit/ pkg/proto/credit/credit.proto --go_out=plugins=grpc:pkg/proto/credit/
 
.PHONY: server
server: ## Build and run server.
	go run -race -ldflags "-s -w" -o bin/server server/main.go
	bin/server
 
.PHONY: client
client: ## Build and run client.
	go run -race -ldflags "-s -w" -o Cliente/client.go

.PHONY: CodigoPrueba
compilePrueba: protoc -I Prueba\ Conexion/Greeter Prueba\ Conexion/Greeter/chat.proto --go_out=plugins=grpc:Prueba\ Conexion/Greeter/chat