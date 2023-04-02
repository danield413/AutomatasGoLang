package models

type Arista struct {
    Origen  *Vertice
    Destino *Vertice
    Peso    string
}

func NewArista(Origen *Vertice, Destino *Vertice, Peso string) *Arista {
    return &Arista{
        Origen:  Origen,
        Destino: Destino,
        Peso:    Peso,
    }
}

func (a *Arista) GetOrigen() *Vertice {
    return a.Origen
}

func (a *Arista) GetDestino() *Vertice {
    return a.Destino
}

func (a *Arista) GetPeso() string {
    return a.Peso
}
