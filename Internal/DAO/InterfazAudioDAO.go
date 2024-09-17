package DAO

import "github.com/jcgr5/AudioInterface/Internal/VO"

type InterfazAudioDAO interface {
	GetInterfazAudio(id int) (*VO.InterfazAudioVO, error)
	CreateInterfazAudio(interfaz *VO.InterfazAudioVO) error
	UpdateInterfazAudio(interfaz *VO.InterfazAudioVO) error
	DeleteInterfazAudio(id int) error
	GetFrecuencia(id int) (*VO.FrecuenciaVO, error)
	SetFrecuencia(interfazID int, frecuencia *VO.FrecuenciaVO) error
	GetAll() ([]*VO.InterfazAudioVO, error)
}
