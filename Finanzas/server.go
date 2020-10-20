package main

import (
	"Tarea1/Finanzas/finanza"
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

func servirServidor(wg *sync.WaitGroup, financiaServer *finanza.ServerFinzanza, puerto string) {
	lis, err := net.Listen("tcp", ":"+puerto)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", puerto, err)
	}
	grpcServer := grpc.NewServer()

	finanza.RegisterFinanzaServiceServer(grpcServer, financiaServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s: %v", puerto, err)
	}
}

func main() {

	var wg sync.WaitGroup

	log.Printf("El IP del servidor es: %v", GetOutboundIP())

	s := finanza.ServerFinzanza{}

	wg.Add(1)
	go servirServidor(&wg, &s, "9000")
	wg.Wait()

}
