package main

import (
	logistica "Tarea1/Logistica/logistica"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Ingrese nombre de la maquina: ")
	ip, _ := reader.ReadString('\n')
	ip = strings.TrimSuffix(ip, "\n")
	ip = strings.TrimSuffix(ip, "\r")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ip+":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("no se pudo conectar: %s", err)
	}

	defer conn.Close()

	c := logistica.NewLogisticaServiceClient(conn)

	message1 := logistica.OrdenCliente{
		Id:          "ASS-1313",
		Producto:    "Un masajeador wink wink",
		Valor:       1313,
		Tienda:      "Solo Para Chicos Grandes",
		Destino:     "Tus Nalgas",
		Prioritario: -1,
	}

	message2 := logistica.SeguimientoCliente{
		Seguimiento: 1,
		Estado:      "Un masajeador wink wink",
		Producto:    "1313",
	}

	response1, err := c.NuevaOrden(context.Background(), &message1)
	response2, err := c.InformarSeguimiento(context.Background(), &message2)

	if err != nil {
		log.Fatalf("La polilla gigante ataco la conexion: %s", err)
	}

	log.Printf("El numero de seguimiento de la wea de producto %s es: %d", response1.Producto, response1.Seguimiento)
	log.Printf("El pedido ql que querí saber esta %s", response2.Estado)
}
