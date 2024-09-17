package MODELs

type Canal struct {
	ID          int
	CodigoCanal int
	Etiqueta    string
	Volumen     float64
	Solo        bool
	Mute        bool
	Fuente      Fuente
}
