package models

type Vertice struct {
    dato            string
    ListaAdyacentes []*Vertice
    estadoFinal    bool

    posicionX int
    posicionY int
}

func NewVertice(dato string) *Vertice {
    return &Vertice{
        dato:            dato,
        ListaAdyacentes: []*Vertice{},
        estadoFinal:    false,
        posicionX: 0,
        posicionY: 0,
    }
}

func (v *Vertice) GetEstadoFinal() bool {
    return v.estadoFinal
}

func (v *Vertice) SetEstadoFinal(estadoFinal bool) {
    v.estadoFinal = estadoFinal
}

func (v *Vertice) GetDato() string {
    return v.dato
}

func (v *Vertice) SetDato(dato string) {
    v.dato = dato
}

func (v *Vertice) GetAdyacentes() []*Vertice {
    return v.ListaAdyacentes
}

func (v *Vertice) SetAdyacentes(adyacentes []*Vertice) {
	v.ListaAdyacentes = adyacentes
}

func (v *Vertice) GetPosicionX() int {
    return v.posicionX
}

func (v *Vertice) SetPosicionX(posicionX int) {
    v.posicionX = posicionX
}

func (v *Vertice) GetPosicionY() int {
    return v.posicionY
}

func (v *Vertice) SetPosicionY(posicionY int) {
    v.posicionY = posicionY
}
