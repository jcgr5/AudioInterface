package Interfaces

import VO2 "github.com/jcgr5/AudioInterface/Internal/MODELs/VO"

type InterfazAudioDAO interface {
	GetInterfazAudio(id int) (*VO2.InterfazAudioVO, error)
	CreateInterfazAudio(interfaz *VO2.InterfazAudioVO) error
	UpdateInterfazAudio(interfaz *VO2.InterfazAudioVO) error
	DeleteInterfazAudio(id int) error
	GetFrecuencia(id int) (*VO2.FrecuenciaVO, error)
	SetFrecuencia(interfazID int, frecuencia *VO2.FrecuenciaVO) error
	GetAll() ([]*VO2.InterfazAudioVO, error)
}
