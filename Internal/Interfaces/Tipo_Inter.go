package Interfaces

import "github.com/jcgr5/AudioInterface/Internal/MODELs/VO"

type TipoDAO interface {
	Create(t *VO.TipoVO) error
	GetByID(id int) (*VO.TipoVO, error)
	Update(t *VO.TipoVO) error
	Delete(id int) error
	GetAll() ([]*VO.TipoVO, error)
}
