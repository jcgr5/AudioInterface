package VO

import "time"

type ConfiguracionVO struct {
	ID       int
	Interfaz InterfazAudioVO
	Usuario  UsuarioVO
	Fecha    time.Time
	Canales  []CanalVO
	Entradas []EntradaVO
}
