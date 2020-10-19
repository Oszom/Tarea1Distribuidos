package camion

import (
	logistica "Tarea1/Logistica/logistica"
	"bufio"
	context "context"
	"fmt"
	wr "github.com/mroth/weightedrand"
	"google.golang.org/grpc"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Registro is

/*
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
*/

// CamionServer is
type CamionServer struct {
	tipo            string
	capacidad       int
	informe         []*Registro
	enviosActuales  []*Registro
	tipoUltimoEnvio string
}

//Registro is
type Registro struct {
	idpaquete    string
	tipo         string
	valor        int64
	origen       string
	destino      string
	intentos     int64
	seguimiento  int64
	fechaEntrega string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Ingrese nombre de la maquina: ")
	ip, _ := reader.ReadString('\n')
	ip = strings.TrimSuffix(ip, "\n")
	ip = strings.TrimSuffix(ip, "\r")

	fmt.Printf("Ingrese tiempo de espera machucao: ")
	//tiempoEspera, _ := reader.ReadString('\n')
	//tiempoEspera = strings.TrimSuffix(ip, "\n")
	//tiempoEspera = strings.TrimSuffix(ip, "\r")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ip+":9000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("no se pudo conectar: %s\n", err)
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
		fmt.Printf("La polilla gigante ataco la conexion: %s\n", err)
	}

	fmt.Printf("El numero de seguimiento de la wea de producto %s es: %d\n", response1.Producto, response1.Seguimiento)
	fmt.Printf("El pedido ql que querí saber esta %s\n", response2.Estado)
}

//RecorridoCamiones is
func RecorridoCamiones(tipoCamion string, ip string, tiempo int32) {
	camion := newCamion(tipoCamion)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ip+":9000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("no se pudo conectar: %s\n", err)
	}

	defer conn.Close()
	for {

		message1 := logistica.AsignacionCamion{
			Tipo:               camion.tipo,
			LastPaqueteEnviado: camion.tipoUltimoEnvio,
		}

		c := logistica.NewLogisticaServiceClient(conn)
		response1, err1 := c.AsignarPaquete(context.Background(), &message1)

		if err1 != nil {
			fmt.Printf("no se pudo asignar paquete al camión: %s\n", err1)
		}

		time.Sleep(time.Duration(tiempo) * time.Millisecond)

		response2, err2 := c.AsignarPaquete(context.Background(), &message1)

		if err2 != nil {
			fmt.Printf("no se pudo asignar paquete al camión: %s\n", err2)
		}
		if response2.IdPaquete != "" {
			var registroNuevo1 = newRegistro(response2.IdPaquete, response2.Tipo, response2.Valor, response2.Origen, response2.Destino, 0, "0")
			camion.informe = append(camion.informe, registroNuevo1)
			camion.enviosActuales = append(camion.enviosActuales, registroNuevo1)

		}

		var registroNuevo2 = newRegistro(response1.IdPaquete, response1.Tipo, response1.Valor, response1.Origen, response1.Destino, 0, "0")
		camion.informe = append(camion.informe, registroNuevo2)
		camion.enviosActuales = append(camion.enviosActuales, registroNuevo2)
		var paqueteAEntregar *Registro
		var posicion int
		for i := 0; i < 3; i++ {
			for j := 0; j < 2; j++ {
				if len(camion.enviosActuales) == 2 {
					if camion.enviosActuales[0].valor >= camion.enviosActuales[1].valor && camion.enviosActuales[j].fechaEntrega == "0" {
						paqueteAEntregar = camion.enviosActuales[0]
						posicion = 0
					} else {
						paqueteAEntregar = camion.enviosActuales[1]
						posicion = 1
					}

				} else {
					if camion.enviosActuales[0].fechaEntrega == "0" {
						paqueteAEntregar = camion.enviosActuales[0]
						posicion = 0
					}

				}

				//Intentar entrega
				var intentoEntrega = EntregarPaquete()
				if intentoEntrega == "entregado" {
					registrarEntregaDePaquete(paqueteAEntregar.idpaquete, camion)
					sumarIntentoEntrega(paqueteAEntregar.idpaquete, camion)
					camion.enviosActuales = remove(camion.enviosActuales, posicion)

				} else {
					sumarIntentoEntrega(paqueteAEntregar.idpaquete, camion)

				}
			}
		}

	}

}

func remove(slice []*Registro, s int) []*Registro {
	return append(slice[:s], slice[s+1:]...)
}

//NuevoPaquete is
func (cam *CamionServer) NuevoPaquete(ctx context.Context, paquete *PaqueteRegistro) (*InformeCamion, error) {

	nuevoPaquete := Registro{
		idpaquete:    paquete.IdPaquete,
		seguimiento:  paquete.Seguimiento,
		tipo:         paquete.Tipo,
		valor:        paquete.Valor,
		origen:       paquete.Origen,
		destino:      paquete.Destino,
		intentos:     0,
		fechaEntrega: "0",
	}

	cam.enviosActuales = append(cam.enviosActuales, &nuevoPaquete)
	cam.informe = append(cam.informe, &nuevoPaquete)

	return &InformeCamion{
		IdPaquete: paquete.IdPaquete,
		Estado:    "En camino",
	}, nil
}

//registrarEntregaDePaquete
func registrarEntregaDePaquete(idpaquete string, camion *CamionServer) {
	registro := camion.informe
	for i := 0; i < len(registro); i++ {
		if registro[i].idpaquete == idpaquete {
			registro[i].fechaEntrega = time.Now().Format("02-01-2006 15:04")
		}
	}
}

//sumarIntentoEntrega is
func sumarIntentoEntrega(idpaquete string, camion *CamionServer) {
	registro := camion.informe
	for i := 0; i < len(registro); i++ {
		if registro[i].idpaquete == idpaquete {
			registro[i].intentos++

		}
	}
}

//newRegistro is
func newRegistro(idpaquete string, tipo string, valor int64, origen string, destino string, intentos int64, fechaEntrega string) *Registro {
	registro := Registro{idpaquete: idpaquete,
		tipo:         tipo,
		origen:       origen,
		destino:      destino,
		intentos:     intentos,
		fechaEntrega: fechaEntrega}
	return &registro
}

//newCamion is
func newCamion(tipo string) *CamionServer {
	camion := CamionServer{tipo: tipo}
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
