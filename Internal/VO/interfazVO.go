package VO

type InterfazAudioVO struct {
	ID              int
	NombreCorto     string
	Modelo          string
	NombreComercial string
	Precio          float64
	Frecuencia      FrecuenciaVO
}
