package Interfaces

import "github.com/jcgr5/AudioInterface/Internal/MODELs/VO"

type UsuarioDAO interface {
	GetUsuario(id int) (*VO.UsuarioVO, error)
	CreateUsuario(usuario *VO.UsuarioVO) error
	UpdateUsuario(usuario *VO.UsuarioVO) error
	DeleteUsuario(id int) error
	GetUsuarioByEmail(email string) (*VO.UsuarioVO, error)
	GetAll() ([]*VO.UsuarioVO, error)
}
