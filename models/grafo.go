package models

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

type Grafo struct {
	listaVertices     []*Vertice
	listaAristas      []*Arista
	visitadosCp       []string
	visitadosCa       []string
	adyacencias       map[string][]string
	visitadosCKruskal []string
	repetidos         int
	obstruidos        []string
	aristasAmplitud   []*Arista
}

func NewGrafo() *Grafo {
	return &Grafo{
		listaVertices:     []*Vertice{},
		listaAristas:      []*Arista{},
		visitadosCp:       []string{},
		visitadosCa:       []string{},
		adyacencias:       map[string][]string{},
		visitadosCKruskal: []string{},
		repetidos:         0,
		obstruidos:        []string{},
		aristasAmplitud:   []*Arista{},
	}
}

func (g *Grafo) GetListaVertices() []*Vertice {
	return g.listaVertices
}

func (g *Grafo) GetListaAristas() []*Arista {
	return g.listaAristas
}

func (g *Grafo) IngresarVertices(dato string) {
	if !g.VerificarExisteVertice(dato, g.listaVertices) {
		g.listaVertices = append(g.listaVertices, &Vertice{dato, []*Vertice{}, false, 0, 0})
	}
}

func (g *Grafo) VerificarExisteVertice(dato string, lista []*Vertice) bool {
	for i := 0; i < len(lista); i++ {
		if dato == lista[i].dato {
			return true
		}
	}
	return false
}

func (g *Grafo) MostrarVertices() {
	for i := 0; i < len(g.listaVertices); i++ {
		fmt.Printf(">>> Vertice: " + g.listaVertices[i].dato)
		fmt.Println(" Adyacentes: ")
		for j := 0; j < len(g.listaVertices[i].ListaAdyacentes); j++ {
			fmt.Println(g.listaVertices[i].ListaAdyacentes[j].dato)
		}
		fmt.Println("---------------------------------")
		if len(g.listaVertices[i].ListaAdyacentes) == 0 {
			print("No hay \n")
		}
	}
}

// metodo que me recorrar las aristas y me retorne true si una vertice solo tiene una adyacencia
func (g *Grafo) ConvertirAutomataACompleto() {
	contadorEstadosIncompletos := 0
	listaEstadosIncompletos := []string{}
	for i := 0; i < len(g.listaVertices); i++ {
		if len(g.listaVertices[i].ListaAdyacentes) == 1 || len(g.listaVertices[i].ListaAdyacentes) == 0 {
			println("Estado incompleto: ", g.listaVertices[i].dato)
			listaEstadosIncompletos = append(listaEstadosIncompletos, g.listaVertices[i].dato)
			contadorEstadosIncompletos++
		}
	}

	fmt.Println("Estados incompletos: ", contadorEstadosIncompletos)
	if contadorEstadosIncompletos > 0 {

		//crear el sumidero
		g.IngresarVertices("Sumidero")
		g.IngresarArista("Sumidero", "Sumidero", "0")
		g.IngresarArista("Sumidero", "Sumidero", "1")

		// fmt.Println(g.GetVertice("Sumidero").GetAdyacentes())

		for i := 0; i < len(listaEstadosIncompletos); i++ {
			fmt.Println("Estado incompleto: ", listaEstadosIncompletos[i])

			aristasVertice := g.GetAristasDeVertice(listaEstadosIncompletos[i])
			if len(aristasVertice) == 1 {
				//tomo el peso
				peso := aristasVertice[0].Peso
				pesoAlSumidero := 0
				if peso == "0" {
					pesoAlSumidero = 1
				} else {
					pesoAlSumidero = 0
				}

				//ingresar arista
				g.IngresarArista(listaEstadosIncompletos[i], "Sumidero", fmt.Sprintf("%v", pesoAlSumidero))
			}
			if len(aristasVertice) == 0 {
				g.IngresarArista(listaEstadosIncompletos[i], "Sumidero", "0")
				g.IngresarArista(listaEstadosIncompletos[i], "Sumidero", "1")
			}

		}

		fmt.Println("Autómata convertido a completo.")
		fmt.Println("---------------------------------")
		g.MostrarVertices()
		fmt.Println("---------------------------------")
		g.MostrarAristas()
		fmt.Println("---------------------------------")

	}
}

// obtener arisstas de un vertice
func (g *Grafo) GetAristasDeVertice(vertice string) []*Arista {
	aristas := []*Arista{}
	for i := 0; i < len(g.listaAristas); i++ {
		if g.listaAristas[i].Origen.dato == vertice {
			aristas = append(aristas, g.listaAristas[i])
		}
	}
	return aristas
}

func (g *Grafo) MostrarAristas() {
	for i := 0; i < len(g.listaAristas); i++ {
		fmt.Printf("Origen: %v Destino: %v Peso: %v\n", g.listaAristas[i].Origen.GetDato(), g.listaAristas[i].Destino.GetDato(), g.listaAristas[i].Peso)
	}
}

func (g *Grafo) GetNombreVertices() []string {
	vertices := []string{}
	for i := 0; i < len(g.listaVertices); i++ {
		vertices = append(vertices, g.listaVertices[i].dato)
	}
	return vertices
}

// ingresar arista
func (g *Grafo) IngresarArista(Origen string, Destino string, Peso string) {
	// if !g.VerificarExisteArista(&Arista{g.GetVertice(Origen), g.GetVertice(Destino), Peso}) {
	g.listaAristas = append(g.listaAristas, &Arista{g.GetVertice(Origen), g.GetVertice(Destino), Peso})

	g.GetVertice(g.GetVertice(Origen).GetDato()).ListaAdyacentes = append(g.GetVertice(g.GetVertice(Origen).GetDato()).ListaAdyacentes, g.GetVertice(Destino))
	// }
}

// obtener vertice
func (g *Grafo) GetVertice(dato string) *Vertice {
	for i := 0; i < len(g.listaVertices); i++ {
		if dato == g.listaVertices[i].dato {
			return g.listaVertices[i]
		}
	}
	return nil
}

// recorrer el grafo usando los valores de la cadena de caracteres
func (g *Grafo) RecorrerAutomata(cadena string, ventana fyne.Window) {
	fmt.Println("Recorriendo el autómata con la cadena: " + cadena)
	// valida := false
	//convertir la cadena a un arreglo de caracteres

	cadenaArr := []string{}

	for i := 0; i < len(cadena); i++ {
		cadenaArr = append(cadenaArr, string(cadena[i]))
	}

	//recorrer el arreglo de caracteres

	estadoActual := g.GetVertice("A")

	g.CambiarEstado(estadoActual, 0, cadenaArr, ventana)

}

func (g *Grafo) CambiarEstado(estadoActual *Vertice, posicionCadena int, cadenaArrInt []string, ventana fyne.Window) {

	//si el estado actual es un estado de aceptacion
	if estadoActual.GetEstadoFinal() && posicionCadena == len(cadenaArrInt) {
		fmt.Println("La cadena cumple, acabó en un estado final!")
		dialog.ShowInformation("Resultado", "La cadena cumple, acabó en un estado final!", ventana)
		return
	} else if posicionCadena == len(cadenaArrInt) && !estadoActual.GetEstadoFinal() {
		fmt.Println("La cadena no cumple, no acabó un estado final...")
		dialog.ShowInformation("Resultado", "La cadena no cumple, no acabó un estado final...", ventana)
		return
	}

	//miro por donde me puedo ir siguiendo la cadena
	for i := 0; i < len(g.listaAristas); i++ {
		if g.listaAristas[i].GetOrigen().GetDato() == estadoActual.GetDato() {
			// print(g.listaAristas[i].GetOrigen().dato, " -> ", g.listaAristas[i].GetDestino().dato, " = ", g.listaAristas[i].GetPeso(), "\n")
			if g.listaAristas[i].GetPeso() == cadenaArrInt[posicionCadena] {
				print("Pasó por ", g.listaAristas[i].GetDestino().GetDato(), "\n")
				g.CambiarEstado(g.listaAristas[i].GetDestino(), posicionCadena+1, cadenaArrInt, ventana)
			}
		}
	}
}

func (g *Grafo) GetArista(Origen string, Destino string, Peso string) *Arista {
	for i := 0; i < len(g.listaAristas); i++ {
		if Origen == g.listaAristas[i].Origen.GetDato() && Destino == g.listaAristas[i].Destino.GetDato() && Peso == g.listaAristas[i].Peso {
			return g.listaAristas[i]
		}
	}
	return nil
}

func (g *Grafo) VerificarExisteArista(arista *Arista) bool {
	for i := 0; i < len(g.listaAristas); i++ {
		if arista.Origen.GetDato() == g.listaAristas[i].Origen.GetDato() && arista.Destino.GetDato() == g.listaAristas[i].Destino.GetDato() {
			// print("Ya existe la arista\n")
			// print("Origen: ")
			// print(arista.Origen.GetDato())
			// print(" Destino: ")
			// print(arista.Destino.GetDato())
			return true
		}
	}
	return false
}
