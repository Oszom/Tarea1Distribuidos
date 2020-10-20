package finanza

import (
	"log"

	"golang.org/x/net/context"
)

type RegistroFinanza struct {
	idpaquete    string
	tipo         string
	nombre       string
	valor        int64
	origen       string
	destino      string
	fechaEntrega string
	seguimiento  int64
	estado       string
	timestamp    string
	ganancia     int64
}

type ServerFinzanza struct {
	ListaEnvios       []*PaqueteRegistro
	gananciasActuales float32
}

func calcularGanancias(valorProducto int64, intentosT int64, tipoPedido string, entregaExitosa bool) float32 {

	var gananciaEntrega float32
	gananciaEntrega = 10.0 * float32(intentosT)

	if tipoPedido == "retail" {
		gananciaEntrega = gananciaEntrega + float32(valorProducto)
	} else if tipoPedido == "prioritario" {
		if entregaExitosa {
			gananciaEntrega = gananciaEntrega + float32(valorProducto)
		} else {
			gananciaEntrega = gananciaEntrega + 0.3*float32(valorProducto)
		}
	} else if tipoPedido == "normal" {
		if entregaExitosa {
			gananciaEntrega = gananciaEntrega + float32(valorProducto)
		}
	}

	return gananciaEntrega

}

func (sf *ServerFinzanza) informarEntrega(ctx context.Context, paquete *PaqueteRegistro) (*Ack, error) {

	gananciaEnvio := calcularGanancias(paquete.Valor, paquete.Intentos, paquete.Tipo, paquete.FechaEntrega != "0")

	newRegistro := PaqueteRegistro{
		IdPaquete:    paquete.IdPaquete,
		Seguimiento:  paquete.Seguimiento,
		Tipo:         paquete.Tipo,
		Nombre:       paquete.Nombre,
		Valor:        paquete.Valor,
		Origen:       paquete.Origen,
		Destino:      paquete.Destino,
		FechaEntrega: paquete.FechaEntrega,
		Estado:       paquete.Estado,
		Timestamp:    paquete.Timestamp,
		Ganancia:     gananciaEnvio,
	}

	sf.ListaEnvios = append(sf.ListaEnvios, &newRegistro)

	sf.gananciasActuales = sf.gananciasActuales + gananciaEnvio
	log.Printf("El envio correspondiente al ID %s, ha dado unas ganancias de %f.\nLas ganancias de la empresa son: %f", paquete.IdPaquete, gananciaEnvio, sf.gananciasActuales)

	return &Ack{IdPaquete: paquete.IdPaquete}, nil
}
