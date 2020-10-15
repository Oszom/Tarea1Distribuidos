package chat

import (
	"bufio"
	"os"
	"fmt"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error){
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Jorge: %s\n", message.Body)
	fmt.Printf("Pablo:")
	mensaje, _ := reader.ReadString('\n')
	return &Message{Body: mensaje}, nil
}