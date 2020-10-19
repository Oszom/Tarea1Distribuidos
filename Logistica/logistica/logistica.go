package logistica

import (
	"sync"
	"time"
	"log"
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

//ServerLogistica is
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
		return "normal"
	} else if prioridad == 1 {
		return "prioritario"
	} else {
		return "retail"
	}
}

//NuevaOrden is
func (s *ServerLogistica) NuevaOrden(ctx context.Context, orden *OrdenCliente) (*SeguimientoCliente, error) {

	tipoEnvio := tipoEnvio(orden.Prioritario)
	nuevaOrden := newRegistro(orden.Id, tipoEnvio, orden.Producto, orden.Valor, orden.Tienda, orden.Destino, s.currSeguimiento)
	s.currSeguimiento++
	s.ListaEnvios = append(s.ListaEnvios, nuevaOrden)
	
	if tipoEnvio == "retail"{
		s.ColaRetail = append(s.ColaRetail,nuevaOrden)
	} else if tipoEnvio == "prioritario"{
		s.ColaPrioritarios = append(s.ColaPrioritarios,nuevaOrden)
	} else if tipoEnvio == "normal"{
		s.ColaNormales = append(s.ColaNormales,nuevaOrden)
	}

	log.Printf("Llego una nueva orden con ID %s",nuevaOrden.idpaquete)

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

			log.Printf("Se pregunto por el seguimiento de la orden %d",resultado.Seguimiento)
		}
	}

	if resultado.Seguimiento == -1{
		log.Printf("Se pregunto por una orden inexistente con numero de seguimiento %d",codSeguimiento.Seguimiento)
	}

	return resultado, nil
}

//InformarEntrega is
func (s *ServerLogistica) InformarEntrega(ctx context.Context, codSeguimiento *InformeCamion) (*InformeCamion, error) {

	resultado := &InformeCamion{
		IdPaquete: "-1",
		Estado:    "Hola Leo",
	}

	for i := 0; i < len(s.ListaEnvios); i++ {
		if s.ListaEnvios[i].idpaquete == codSeguimiento.IdPaquete {
			resultado = &InformeCamion{
				IdPaquete: codSeguimiento.IdPaquete,
				Estado:    codSeguimiento.Estado,
			}
			s.ListaEnvios[i].estado = codSeguimiento.Estado
			/*Comunicacion con Finanzas






			 */
		}
	}

	return resultado, nil

}

//AsignarPaquete is
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

	if resultado.IdPaquete == ""{
		log.Printf("El camion %s se va con la cola entre las piernas :'(",presentacionCamion.Tipo)
	} else {
		log.Printf("El camion %s, se lleva la orden %s de tipo %s",presentacionCamion.Tipo,resultado.IdPaquete,resultado.Tipo)
	}

	s.muxCamiones.Unlock()

	return resultado, nil
}
