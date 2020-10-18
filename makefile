#················································Logistica·························································
.PHONY: runLogistica
runLogistica: 
	protoc -I Logistica/logistica/ Logistica/logistica/logistica.proto --go_out=plugins=grpc:Logistica/logistica/
	go run Logistica/server.go

.PHONY: compileLogistica
compileLogistica:
	protoc -I Logistica/logistica/ Logistica/logistica/logistica.proto --go_out=plugins=grpc:Logistica/logistica/

#··················································································································
#················································Cliente···························································
.PHONY: runCliente
runCliente: 
	go run Cliente/client.go
#··················································································································
#················································Camion····························································
.PHONY: runCamion
runCamion: 
	protoc -I Camiones/camion/ Camiones/camion/camion.proto --go_out=plugins=grpc:Camiones/camion/
	go run Camiones/server.go

.PHONY: compileCamion
compileCamion:
	protoc -I Camiones/camion/ Camiones/camion/camion.proto --go_out=plugins=grpc:Camiones/camion/
#··················································································································

.PHONY: server
server: ## Build and run server.
	go run -race -ldflags "-s -w" -o bin/server server/main.go
	bin/server
 
.PHONY: client
client: ## Build and run client.
	go run -race -ldflags "-s -w" -o Cliente/client.go

.PHONY: compilePrueba
compilePrueba: 
	protoc -I "Prueba Conexion/Greeter" "Prueba Conexion/Greeter/chat.proto" --go_out=plugins=grpc:"Prueba Conexion/Greeter/chat"