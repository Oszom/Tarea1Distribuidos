package main

import (
	logistica "Tarea1/Logistica/logistica"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("no se pudo conectar: %s", err)
	}

	defer conn.Close()

	c := logistica.NewClienteServiceClient(conn)

	message := logistica.OrdenCliente{
		Id:          "ASS-1313",
		Producto:    "Un masajeador wink wink",
		Valor:       1313,
		Tienda:      "Solo Para Chicos Grandes",
		Destino:     "Tus Nalgas",
		Prioritario: -1,
		Seguimiento: 0,
	}

	response, err := c.NuevaOrden(context.Background(), &message)
	if err != nil {
		log.Fatalf("La polilla gigante ataco la conexion: %s", err)
	}

	log.Printf("El numero de seguimiento de la wea de producto %s es: %d", response.Producto, response.Seguimiento)
}
