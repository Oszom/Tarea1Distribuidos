
package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"omega/mediocres/pureba/chat"
	"os"
	"strings"
)

func main() {
	var conn *grpc.ClientConn
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Ingresa la IP del cliente:")

	ip, _ := reader.ReadString('\n')
	ip = strings.TrimSuffix(ip, "\n")
	ip = strings.TrimSuffix(ip, "\r")
	conn, err := grpc.Dial(ip+":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("no se pudo conectar: %s", err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	fmt.Printf("Jorge:")
	mensaje, _ := reader.ReadString('\n')
	var posicion int64 = 1
	for {

		message := chat.Cosita{
			Saludo:   mensaje,
			Posicion: posicion,
		}

		response, err := c.Saludar(context.Background(), &message)
		if err != nil {
			log.Fatalf("La polilla gigante ataco la conexion: %s", err)
		}

		fmt.Printf("Pablo: %s", response.Saludo)
		fmt.Printf("Pablo: %d", response.Posicion)
		fmt.Printf("Jorge:")

		mensaje, _ = reader.ReadString('\n')
		posicion++

	}
}
