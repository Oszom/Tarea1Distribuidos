package main

import (
	logistica "Tarea1/Logistica/logistica"
	"bufio"
	context "context"
	"fmt"
	wr "github.com/mroth/weightedrand"
	"google.golang.org/grpc"
	"log"
	"sync"
	"math/rand"
	"os"
	"strconv"
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

	var wg sync.WaitGroup

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Ingrese nombre de la maquina: ")
	ip, _ := reader.ReadString('\n')
	ip = strings.TrimSuffix(ip, "\n")
	ip = strings.TrimSuffix(ip, "\r")

	fmt.Printf("Ingrese tiempo de espera machucao: ")
	tiempoEspera, _ := reader.ReadString('\n')
	tiempoEspera = strings.TrimSuffix(tiempoEspera, "\n")
	tiempoEspera = strings.TrimSuffix(tiempoEspera, "\r")
	tiempoEsperaInt, errn := strconv.ParseInt(tiempoEspera, 10, 64)
	if errn != nil {
		fmt.Println("Problema de conversión del tiempo\n", errn)
	}
	wg.Add(1)
	go RecorridoCamiones(&wg, "retail", ip, tiempoEsperaInt,"9100", 1)
	wg.Add(1)
	go RecorridoCamiones(&wg, "retail", ip, tiempoEsperaInt,"9101", 2)
	wg.Add(1)
	go RecorridoCamiones(&wg, "normal", ip, tiempoEsperaInt,"9102", 1)

	wg.Wait()

}

//RecorridoCamiones is
func RecorridoCamiones(wg *sync.WaitGroup,tipoCamion string, ip string, tiempo int64, puerto string, numeroRet int) {
	defer wg.Done()
	camion := newCamion(tipoCamion)
	log.Printf("Generando camión %s %d, con un tiempo de espera de %d segundos", camion.tipo, numeroRet, tiempo)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ip+":"+puerto, grpc.WithInsecure())
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
		for {
			response1, err1 := c.AsignarPaquete(context.Background(), &message1)
			log.Printf("%s",response1)
			if err1 != nil {
				fmt.Printf("no se pudo asignar paquete al camión %s %d (puerto: %s): %s\n", camion.tipo, numeroRet, puerto,err1)
			}
			time.Sleep(time.Duration(tiempo) * time.Second)
			if response1.IdPaquete != ""{
				var registroNuevo1 = newRegistro(response1.IdPaquete, response1.Tipo, response1.Valor, response1.Origen, response1.Destino, 0, "0")
				camion.informe = append(camion.informe, registroNuevo1)
				camion.enviosActuales = append(camion.enviosActuales, registroNuevo1)
				log.Printf("Paquete recibido por camión %s %d, id seguimiento: %d", camion.tipo, numeroRet, response1.Seguimiento)
				break
			}
		}

		response2, err2 := c.AsignarPaquete(context.Background(), &message1)

		if err2 != nil {
			fmt.Printf("no se pudo asignar paquete al camión %s %d (puerto: %s): %s\n", camion.tipo, numeroRet, puerto,err2)
		}
		log.Printf("Paquete recibido por camión %s %d, id seguimiento: %d", camion.tipo, numeroRet, response2.Seguimiento)
		if response2.IdPaquete != "" {
			var registroNuevo2 = newRegistro(response2.IdPaquete, response2.Tipo, response2.Valor, response2.Origen, response2.Destino, 0, "0")
			camion.informe = append(camion.informe, registroNuevo2)
			camion.enviosActuales = append(camion.enviosActuales, registroNuevo2)
		}

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

				} else if len(camion.enviosActuales) == 1 {
					if camion.enviosActuales[0].fechaEntrega == "0" {
						paqueteAEntregar = camion.enviosActuales[0]
						posicion = 0
					}

				} else {
					break
				}

				//Intentar entrega
				var intentoEntrega = EntregarPaquete()
				if intentoEntrega == "entregado" {
					log.Printf("Paquete de camión %s %d, con id seguimiento: %d entregado", camion.tipo, numeroRet, paqueteAEntregar.seguimiento)
					sumarIntentoEntrega(paqueteAEntregar.idpaquete, camion)
					registrarEntregaDePaquete(paqueteAEntregar.idpaquete, camion)
					newMessage := logistica.InformeCamion{
						IdPaquete: paqueteAEntregar.idpaquete,
						Estado: "Recibido",
						intentos: paqueteAEntregar.intentos + 1
					}
					res, err := c.InformarEntrega(context.Background(), &newMessage)
					if err != nil {
						log.PrintF("Error al momento de avisar a logistica\n(%s)",err)
					}
					camion.enviosActuales = remove(camion.enviosActuales, posicion)

				} else {

					sumarIntentoEntrega(paqueteAEntregar.idpaquete, camion)
					log.Printf("Paquete de camión %s %d, con id seguimiento: %d NO entregado (intento numero %d)", camion.tipo, numeroRet, paqueteAEntregar.seguimiento, paqueteAEntregar.intentos)

				}
			}
		}

		if len(camion.enviosActuales == 1) {
			sumarIntentoEntrega(paqueteAEntregar.idpaquete, camion)
			newMessage := logistica.InformeCamion{
				IdPaquete: paqueteAEntregar.idpaquete,
				Estado: "No Recibido",
				intentos: paqueteAEntregar.intentos + 1
			}
			res, err := c.InformarEntrega(context.Background(), &newMessage)
			if err != nil {
				log.PrintF("Error al momento de avisar a logistica\n(%s)",err)
			}
			camion.enviosActuales = remove(camion.enviosActuales, posicion)
		}

		log.Printf("Fin ronda de entrega")

	}

}

func remove(slice []*Registro, s int) []*Registro {
	return append(slice[:s], slice[s+1:]...)
}
/*
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
*/
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
