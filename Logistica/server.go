package main

import (
	"Tarea1/Logistica/logistica"
	"log"
	"net"
	"sync"
	"google.golang.org/grpc"
)

//GetOutboundIP is
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func servirServidor(wg *sync.WaitGroup, logisticaServer *logistica.ServerLogistica, puerto string){
	lis, err := net.Listen("tcp", ":"+puerto)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", puerto, err)
	}
	grpcServer := grpc.NewServer()

	logistica.RegisterLogisticaServiceServer(grpcServer, logisticaServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s: %v", puerto, err)
	}
}

func main() {

	var wg sync.WaitGroup

	log.Printf("El IP del servidor es: %v", GetOutboundIP())

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Ingrese nombre de la maquina donde se encuentra finanzas: ")
	ipFinanzas, _ := reader.ReadString('\n')
	ipFinanzas = strings.TrimSuffix(ipFinanzas, "\n")
	ipFinanzas = strings.TrimSuffix(ipFinanzas, "\r")

	s := logistica.ServerLogistica{
		ipFinanzas = ipFinanzas
	}

	wg.Add(1)
	go servirServidor(&wg, &s, "9000")
	wg.Add(1)
	go servirServidor(&wg, &s, "9100")
	wg.Add(1)
	go servirServidor(&wg, &s, "9101")
	wg.Add(1)
	go servirServidor(&wg, &s, "9102")
	wg.Wait()

}
