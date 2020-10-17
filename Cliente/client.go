package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"omega/mediocres/pureba/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("no se pudo conectar: %s", err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Quiero casarme con una chica anime polilla!",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("La polilla gigante ataco la conexion: %s", err)
	}

	log.Printf("El servidor responde: %s", response.Body)
}