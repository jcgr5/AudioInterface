package DAO

import "github.com/jcgr5/AudioInterface/Internal/VO"

type ConfiguracionDAO interface {
	GetConfiguracion(id int) (*VO.ConfiguracionVO, error)
	CreateConfiguracion(config *VO.ConfiguracionVO) error
	UpdateConfiguracion(config *VO.ConfiguracionVO) error
	DeleteConfiguracion(id int) error
	GetCanales(configID int) ([]VO.CanalVO, error)
	GetEntradas(configID int) ([]VO.EntradaVO, error)
	GetAll() ([]*VO.ConfiguracionVO, error)
	GetAllByUser(usuarioID int) ([]*VO.ConfiguracionVO, error)
}
