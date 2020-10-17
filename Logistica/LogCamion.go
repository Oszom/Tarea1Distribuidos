package LogisticaCamion

import (
	"bufio"
	"fmt"
	"golang.org/x/net/context"
	"os"
)

type Registro struct {
	idPaquete int64;
	seguimiento int64;
	tipo string;
    valor int64;
    intentos int64;
    estado string;
}
}

//Server is
type Server struct {
	listaPedidos []*Registro;
	
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
	fmt.Printf("El saludo anterior fue: %s\n", s.elSaludo)
	fmt.Printf("La posicion anterior fue: %d\n", s.posicion)
	s.posicion = cosita.Posicion
	s.elSaludo = cosita.Saludo
	return &Cosita{Saludo: "jorge", Posicion: 1}, nil
}
