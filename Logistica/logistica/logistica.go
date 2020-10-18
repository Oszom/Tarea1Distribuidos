package logistica

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

//RegistroLogistica is
type RegistroLogistica struct {
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
}

//ServerCliente is
type ServerCliente struct {
	ListaEnvios      []*RegistroLogistica
	ColaRetail       []*RegistroLogistica
	ColaPrioritarios []*RegistroLogistica
	ColaNormales     []*RegistroLogistica
	currSeguimiento  int64
}

//newRegistro is
func newRegistro(idpaquete string, tipo string, nombre string, valor int64, origen string, destino string, currSeguimiento int64) *RegistroLogistica {
	registro := RegistroLogistica{
		idpaquete:    idpaquete,
		tipo:         tipo,
		origen:       origen,
		destino:      destino,
		fechaEntrega: "0",
		nombre:       nombre,
		seguimiento:  currSeguimiento + 1,
		timestamp:    time.Now().Format("02-01-2006 15:04"),
		estado:       "En Bodega",
	}
	return &registro
}

func tipoEnvio(prioridad int64) string {
	if prioridad == 0 {
		return "Normal"
	} else if prioridad == 1 {
		return "Prioritario"
	} else {
		return "Retail"
	}
}

//NuevaOrden is
func (s *ServerCliente) NuevaOrden(ctx context.Context, orden *OrdenCliente) (*SeguimientoCliente, error) {

	tipoEnvio := tipoEnvio(orden.Prioritario)
	nuevaOrden := newRegistro(orden.Id, tipoEnvio, orden.Producto, orden.Valor, orden.Tienda, orden.Destino, s.currSeguimiento)
	s.currSeguimiento++
	s.ListaEnvios = append(s.ListaEnvios, nuevaOrden)

	fmt.Printf("%s\n", s.ListaEnvios)

	return &SeguimientoCliente{
		Seguimiento: nuevaOrden.seguimiento,
		Estado:      nuevaOrden.estado,
		Producto:    nuevaOrden.nombre,
	}, nil
}

//InformarSeguimiento is
func (s *ServerCliente) InformarSeguimiento(ctx context.Context, codSeguimiento *SeguimientoCliente) (*SeguimientoCliente, error) {
	resultado := &SeguimientoCliente{
		Seguimiento: -1,
		Estado:      "No existe",
		Producto:    "-----",
	}
	for i := 0; i < len(s.ListaEnvios); i++ {
		if s.ListaEnvios[i].seguimiento == codSeguimiento.Seguimiento {
			resultado = &SeguimientoCliente{
				Seguimiento: s.ListaEnvios[i].seguimiento,
				Estado:      s.ListaEnvios[i].estado,
				Producto:    s.ListaEnvios[i].nombre,
			}
		}
	}

	return resultado, nil
}

func (s *ServerCliente) InformarEntrega(ctx context.Context, codSeguimiento *InformeCamion) (*InformeCamion, error) {

	return &InformeCamion{
		IdPaquete: 0,
		Estado:    "Muy lindo",
	}, nil

}
