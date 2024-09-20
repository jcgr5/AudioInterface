package Interfaces

import VO2 "github.com/jcgr5/AudioInterface/Internal/MODELs/VO"

type EntradaDAO interface {
	GetEntrada(id int) (*VO2.EntradaVO, error)
	CreateEntrada(entrada *VO2.EntradaVO) error
	UpdateEntrada(entrada *VO2.EntradaVO) error
	DeleteEntrada(id int) error
	GetDispositivo(EntradaID int) (*VO2.DispositivoVO, error)
	SetDispositivo(EntradaID int, dispositivo *VO2.DispositivoVO) error
	GetAll() ([]*VO2.EntradaVO, error)
	GetAllByDispositivo(DispositivoID int) ([]*VO2.EntradaVO, error)
}
