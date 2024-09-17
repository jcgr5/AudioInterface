package DAO

import "github.com/jcgr5/AudioInterface/Internal/VO"

type EntradaDAO interface {
	GetEntrada(id int) (*VO.EntradaVO, error)
	CreateEntrada(entrada *VO.EntradaVO) error
	UpdateEntrada(entrada *VO.EntradaVO) error
	DeleteEntrada(id int) error
	GetDispositivo(EntradaID int) (*VO.DispositivoVO, error)
	SetDispositivo(EntradaID int, dispositivo *VO.DispositivoVO) error
	GetAll() ([]*VO.EntradaVO, error)
	GetAllByDispositivo(DispositivoID int) ([]*VO.EntradaVO, error)
}
