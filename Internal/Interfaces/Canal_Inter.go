package Interfaces

import VO2 "github.com/jcgr5/AudioInterface/Internal/MODELs/VO"

type CanalDAO interface {
	GetCanal(id int) (*VO2.CanalVO, error)
	CreateCanal(canal *VO2.CanalVO) error
	UpdateCanal(canal *VO2.CanalVO) error
	DeleteCanal(id int) error
	GetFuente(canalID int) (*VO2.FuenteVO, error)
	SetFuente(canalID int, fuente *VO2.FuenteVO) error
	GetAll() ([]*VO2.CanalVO, error)
}
