package logistica

import (
	"fmt"
	"sync"
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
type ServerLogistica struct {
	ListaEnvios      []*RegistroLogistica
	ColaRetail       []*RegistroLogistica
	ColaPrioritarios []*RegistroLogistica
	ColaNormales     []*RegistroLogistica
	currSeguimiento  int64
	muxCamiones      sync.Mutex
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
func (s *ServerLogistica) NuevaOrden(ctx context.Context, orden *OrdenCliente) (*SeguimientoCliente, error) {

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
func (s *ServerLogistica) InformarSeguimiento(ctx context.Context, codSeguimiento *SeguimientoCliente) (*SeguimientoCliente, error) {
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

func (s *ServerLogistica) InformarEntrega(ctx context.Context, codSeguimiento *InformeCamion) (*InformeCamion, error) {

	return &InformeCamion{
		IdPaquete: 0,
		Estado:    "Muy lindo",
	}, nil

}

func (s *ServerLogistica) AsignarPaquete(ctx context.Context, presentacionCamion *AsignacionCamion) (*PaqueteRegistro, error) {

	resultado := &PaqueteRegistro{
		IdPaquete:   "",
		Seguimiento: 0,
		Tipo:        "",
		Valor:       0,
		Intentos:    0,
		Estado:      "",
		Origen:      "",
		Destino:     "",
	}

	s.muxCamiones.Lock()
	//Asignacion de paquete y actualizacion de cola

	if presentacionCamion.Tipo == "retail" {
		if len(s.ColaRetail) > 0 {

			paqueteAEntregar := s.ColaRetail[0]

			resultado = &PaqueteRegistro{
				IdPaquete:   paqueteAEntregar.idpaquete,
				Seguimiento: paqueteAEntregar.seguimiento,
				Tipo:        paqueteAEntregar.tipo,
				Valor:       paqueteAEntregar.valor,
				Intentos:    0,
				Estado:      paqueteAEntregar.estado,
				Origen:      paqueteAEntregar.origen,
				Destino:     paqueteAEntregar.destino,
			}

			//elimino elemento
			s.ColaRetail[0] = nil
			s.ColaRetail = s.ColaRetail[1:]
		} else if len(s.ColaPrioritarios) > 0 && presentacionCamion.LastPaqueteEnviado == "retail" {
			paqueteAEntregar := s.ColaPrioritarios[0]

			resultado = &PaqueteRegistro{
				IdPaquete:   paqueteAEntregar.idpaquete,
				Seguimiento: paqueteAEntregar.seguimiento,
				Tipo:        paqueteAEntregar.tipo,
				Valor:       paqueteAEntregar.valor,
				Intentos:    0,
				Estado:      paqueteAEntregar.estado,
				Origen:      paqueteAEntregar.origen,
				Destino:     paqueteAEntregar.destino,
			}

			//elimino elemento
			s.ColaPrioritarios[0] = nil
			s.ColaPrioritarios = s.ColaPrioritarios[1:]
		}
	} else if presentacionCamion.Tipo == "normal" {
		if len(s.ColaPrioritarios) > 0 {
			paqueteAEntregar := s.ColaPrioritarios[0]

			resultado = &PaqueteRegistro{
				IdPaquete:   paqueteAEntregar.idpaquete,
				Seguimiento: paqueteAEntregar.seguimiento,
				Tipo:        paqueteAEntregar.tipo,
				Valor:       paqueteAEntregar.valor,
				Intentos:    0,
				Estado:      paqueteAEntregar.estado,
				Origen:      paqueteAEntregar.origen,
				Destino:     paqueteAEntregar.destino,
			}

			//elimino elemento
			s.ColaPrioritarios[0] = nil
			s.ColaPrioritarios = s.ColaPrioritarios[1:]
		} else if len(s.ColaNormales) > 0 {
			paqueteAEntregar := s.ColaNormales[0]

			resultado = &PaqueteRegistro{
				IdPaquete:   paqueteAEntregar.idpaquete,
				Seguimiento: paqueteAEntregar.seguimiento,
				Tipo:        paqueteAEntregar.tipo,
				Valor:       paqueteAEntregar.valor,
				Intentos:    0,
				Estado:      paqueteAEntregar.estado,
				Origen:      paqueteAEntregar.origen,
				Destino:     paqueteAEntregar.destino,
			}

			//elimino elemento
			s.ColaNormales[0] = nil
			s.ColaNormales = s.ColaNormales[1:]
		}
	}

	s.muxCamiones.Unlock()

	return resultado, nil
}
