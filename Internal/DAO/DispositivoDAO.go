package DAO

import "github.com/jcgr5/AudioInterface/Internal/VO"

type DispositivoDAO interface {
	Create(d *VO.DispositivoVO) error
	GetByID(id int) (*VO.DispositivoVO, error)
	Update(d *VO.DispositivoVO) error
	Delete(id int) error
	GetAll() ([]*VO.DispositivoVO, error)
}
