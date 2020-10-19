package main

import (
	"Tarea1/Logistica/logistica"
	"log"
	"net"

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

func main() {

	log.Printf("El IP del servidor es: %v", GetOutboundIP())

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	lis2,err2 := net.Listen("tcp", ":9100")
	if err2 != nil {
		log.Fatalf("Failed to listen on port 9100: %v", err)
	}

	lis3,err3 := net.Listen("tcp", ":9101")
	if err3 != nil {
		log.Fatalf("Failed to listen on port 9101: %v", err)
	}

	lis4,err4 := net.Listen("tcp", ":9102")
	if err4 != nil {
		log.Fatalf("Failed to listen on port 9102: %v", err)
	}

	s := logistica.ServerLogistica{}

	grpcServer := grpc.NewServer()

	logistica.RegisterLogisticaServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

	if err := grpcServer.Serve(lis2); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9100: %v", err)
	}

	if err := grpcServer.Serve(lis3); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9101: %v", err)
	}

	if err := grpcServer.Serve(lis4); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9102: %v", err)
	}

}
