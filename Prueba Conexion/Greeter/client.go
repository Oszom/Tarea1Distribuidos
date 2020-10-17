
package main

import (
	"log"
	"bufio"
	"os"
	"strings"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"omega/mediocres/pureba/chat"
)

func main() {
	var conn *grpc.ClientConn
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Ingresa la IP del cliente:")

	ip, _ := reader.ReadString('\n')
	ip = strings.TrimSuffix(ip, "\n")
	conn, err := grpc.Dial(ip+":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("no se pudo conectar: %s", err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	fmt.Printf("Jorge:")
	mensaje, _ := reader.ReadString('\n')

	for {

		message := chat.Message{
			Body: mensaje,
		}

		response, err := c.SayHello(context.Background(), &message)
		if err != nil {
			log.Fatalf("La polilla gigante ataco la conexion: %s", err)
		}

		fmt.Printf("Pablo: %s", response.Body)

		fmt.Printf("Jorge:")

		mensaje, _ = reader.ReadString('\n')



	}
}