package MODELs

import "time"

type Configuracion struct {
	ID       int
	Interfaz InterfazAudio
	Usuario  Usuario
	Fecha    time.Time
	Canales  []Canal
	Entradas []Entrada
}
