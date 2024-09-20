package Interfaces

import "github.com/jcgr5/AudioInterface/Internal/MODELs/VO"

type Frecuencia_Int interface {
	GetFrecuencia(id int) (*VO.FrecuenciaVO, error)
	CreateFrecuencia(canal *VO.FrecuenciaVO) error
	UpdateFrecuencia(canal *VO.FrecuenciaVO) error
	DeleteFrecuencia(id int) error
	GetAll() ([]*VO.FrecuenciaVO, error)
}
