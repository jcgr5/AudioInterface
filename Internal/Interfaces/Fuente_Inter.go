package Interfaces

import "github.com/jcgr5/AudioInterface/Internal/MODELs/VO"

type FuenteDAO interface {
	Create(f *VO.FuenteVO) error
	GetByID(id int) (*VO.FuenteVO, error)
	Update(f *VO.FuenteVO) error
	Delete(id int) error
	GetAllByTipo(tipoID int) ([]*VO.FuenteVO, error)
}
