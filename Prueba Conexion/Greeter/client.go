package main

import (
	"bufio"
	"fmt"
	"log"
	"omega/mediocres/pureba/chat"
	"os"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Ingresa la IP del servidor 2:")

	ip, _ := reader.ReadString('\n')
	ip = strings.TrimSuffix(ip, "\n")
	ip = strings.TrimSuffix(ip, "\r")
	conn, err := grpc.Dial(ip+":9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("no se pudo conectar: %s", err)
	}

	defer conn.Close()

	c := chat.NewChatService2RevengeClient(conn)

	fmt.Printf("Jorge:")
	mensaje, _ := reader.ReadString('\n')
	var posicion int64 = 1

	message := chat.Energetica{
		RamoASalvar:  mensaje,
		NEnergeticas: posicion,
		Sabor:        "Frutilla",
	}

	response, err := c.SalvarSemestre(context.Background(), &message)
	fmt.Printf("It's Working, it's working %s", response)
	if err != nil {
		log.Fatalf("La polilla gigante ataco la conexion: %s", err)
	}

}
