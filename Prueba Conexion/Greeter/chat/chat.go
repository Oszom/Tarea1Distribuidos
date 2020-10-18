package chat

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/net/context"
)

//Server is
type Server struct {
	posicion int64
	elSaludo string
}

//SayHello is
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Jorge: %s\n", message.Body)
	fmt.Printf("Pablo:")
	mensaje, _ := reader.ReadString('\n')
	return &Message{Body: mensaje}, nil
}

//Saludar is
func (s *Server) Saludar(ctx context.Context, cosita *Cosita) (*Cosita, error) {
	fmt.Printf("El saludo anterior fue: %s\n", cosita.Saludo)
	fmt.Printf("La posicion anterior fue: %d\n", cosita.Posicion)
	s.posicion = cosita.Posicion
	s.elSaludo = cosita.Saludo
	return &Cosita{Saludo: cosita.Saludo, Posicion: cosita.Posicion}, nil
}
