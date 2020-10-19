package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"omega/mediocres/pureba/chat"
	"os"
	"strings"

	"google.golang.org/grpc"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func llamadaAlServidor1(saludo string, posicion int) {

}

//Ramo is
type Ramo struct {
	posicion      int64
	elSaludo      string
	ipCliente     string
	puertoCliente string
}

//SalvarSemestre is
func (r *Ramo) SalvarSemestre(ctx context.Context, cosita *chat.Energetica) (*chat.Energetica, error) {

	conn, err := grpc.Dial(net.JoinHostPort(r.ipCliente, r.puertoCliente), grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to dial server:, %s", err)

	}
	defer conn.Close()

	client := chat.NewChatServiceClient(conn)
	response, err := client.Saludar(ctx, &chat.Cosita{
		Saludo:   cosita.RamoASalvar,
		Posicion: cosita.NEnergeticas,
	})
	fmt.Printf("AAAAAA %s", response.Saludo)
	if err != nil {
		log.Fatalf("La polilla gigante ataco la conexion: %s", err)
	}
	return &chat.Energetica{
		RamoASalvar:  cosita.RamoASalvar,
		NEnergeticas: cosita.NEnergeticas,
		Sabor:        cosita.Sabor,
	}, nil
}

func main() {

	log.Printf("El IP del servidor es: %v", GetOutboundIP())

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen on port 9001: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Ingresa la IP del servidor 1:")

	ip, _ := reader.ReadString('\n')
	ip = strings.TrimSuffix(ip, "\n")
	ip = strings.TrimSuffix(ip, "\r")

	r := Ramo{
		ipCliente:     ip,
		puertoCliente: "9000",
	}

	grpcServer := grpc.NewServer()

	chat.RegisterChatService2RevengeServer(grpcServer, &r)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9001: %v", err)
	}

}
