package camion

import (
	"fmt"
	"math/rand"
	"time"

	wr "github.com/mroth/weightedrand"
)

// Camion is
type Camion struct {
	tipo      string
	capacidad int
	informe   []*Registro
}

//Registro is
type Registro struct {
	idpaquete    int
	tipo         string
	valor        int
	origen       string
	destino      string
	intentos     int
	fechaEntrega string
}

func main() {
	//lis, err := net.Listen("tcp", ":9000")
	//if err != nil {
	//	log.Fatalf("Failed to listen on port 9000: %v", err)
	//}

	//s := chat.Server{}

	//grpcServer := grpc.NewServer()

	//chat.RegisterChatServiceServer(grpcServer, &s)

	//if err := grpcServer.Serve(lis); err != nil {
	//	log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	//}

	camionNormal := newCamion("Normal")
	//camionRetail1 := newCamion("Retail")
	//camionRetail2 := newCamion("Retail")

	entrega := newRegistro(1, "retail", 30, "casa andres", "casa jorge", 1, "-")
	camionNormal.informe = append(camionNormal.informe, entrega)

	for n := 0; n < 1; n++ {
		fmt.Println("Antes de agregar intento de entrega: ")
		fmt.Println(camionNormal.informe[n])
	}
	sumarIntentoEntrega(1, camionNormal)
	registrarEntregaDePaquete(1, camionNormal)
	for n := 0; n < 1; n++ {
		fmt.Println("Despues de agregar intento de entrega: ")
		fmt.Println(camionNormal.informe[n])
	}

	fmt.Println(camionNormal)
	result := EntregarPaquete()
	fmt.Println(result)
}

//registrarEntregaDePaquete
func registrarEntregaDePaquete(idpaquete int, camion *Camion) {
	registro := camion.informe
	for i := 0; i < len(registro); i++ {
		if registro[i].idpaquete == idpaquete {
			registro[i].fechaEntrega = time.Now().Format("02-01-2006 15:04")
		}
	}
}

//sumarIntentoEntrega is
func sumarIntentoEntrega(idpaquete int, camion *Camion) {
	registro := camion.informe
	for i := 0; i < len(registro); i++ {
		if registro[i].idpaquete == idpaquete {
			registro[i].intentos++

		}
	}
}

//newRegistro is
func newRegistro(idpaquete int, tipo string, valor int, origen string, destino string, intentos int, fechaEntrega string) *Registro {
	registro := Registro{idpaquete: idpaquete,
		tipo:         tipo,
		origen:       origen,
		destino:      destino,
		intentos:     intentos,
		fechaEntrega: fechaEntrega}
	return &registro
}

//newCamion is
func newCamion(tipo string) *Camion {
	camion := Camion{tipo: tipo}
	camion.capacidad = 2
	return &camion
}

//EntregarPaquete is
func EntregarPaquete() string {
	rand.Seed(time.Now().UTC().UnixNano())
	eleccion := wr.NewChooser(
		wr.Choice{Item: "no_entregado", Weight: 2},
		wr.Choice{Item: "entregado", Weight: 8},
	)
	result := eleccion.Pick().(string)
	return result
}
