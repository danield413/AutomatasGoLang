// create the clas View
package view

import (
	"fmt"
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"bufio"
	"fyne.io/fyne/v2/dialog"
	"strconv"
	models "myapp/models"
)

type View struct {
	app   fyne.App
	grafo *models.Grafo
}

func NewView() *View {
	return &View{
		app:   app.New(),
		grafo: models.NewGrafo(),
	}
}

//* VENTANA PRINCIPAL DE LA APLICACIÓN *//
func (v *View) CargarVentanaPrincipal() {

	// v.app.Settings().SetTheme(theme.LightTheme())

	// v.grafo.IngresarVertices("A")
	// v.grafo.IngresarVertices("B")

	// v.grafo.GetVertice("B").SetEstadoFinal(true)

	// v.grafo.IngresarArista("A", "A", "1")
	// v.grafo.IngresarArista("A", "B", "0")
	// v.grafo.IngresarArista("B", "B", "0")
	// v.grafo.IngresarArista("B", "A", "1")

	// v.grafo.MostrarVertices()

	// v.grafo.ConvertirAutomataACompleto()

	myWindow := v.app.NewWindow("Autómatas")
	myWindow.Resize(fyne.NewSize(600, 300))

	//* BOTÓN PARA CREAR EL AUTOMATA MANUALMENTE *//
	button := widget.NewButton("Crear autómata manualmente", func() {
		ventanaCargar := v.app.NewWindow("Cargar")
		ventanaCargar.Resize(fyne.NewSize(400, 180))
		ventanaCargar.SetContent(widget.NewLabel("Cargar"))
		ventanaCargar.Show()

		input := widget.NewEntry()
		input.SetPlaceHolder("Cuantos estados tiene el automata?")

		widgets := []*widget.Entry{}

		input.OnSubmitted = func(text string) {
			//string to integer
			num, err := strconv.Atoi(text)
			if err == nil {

				//create a new window
				ventanaAgregar := v.app.NewWindow("Cargar")
				ventanaAgregar.Resize(fyne.NewSize(400, 500))
				ventanaAgregar.SetContent(widget.NewLabel("Cargar"))

				//create a new container
				container := container.New(layout.NewVBoxLayout())

				form := &widget.Form{
					Items: []*widget.FormItem{},
					OnSubmit: func() {
						// Obtener los valores de los widgets
						for _, widget := range widgets {
							valor := widget.Text
							// Agregar el estado al grafo
							v.grafo.IngresarVertices(valor)
						}

						v.CargarVentanaEstadosFinales()

						v.CargarVentanaTransiciones()

						ventanaAgregar.Close()
						ventanaCargar.Close()
					},
				}

				for i := 0; i < num; i++ {
					entry := widget.NewEntry()
					entry.SetPlaceHolder("Nombre del estado")
					formItem := widget.NewFormItem("Estado "+strconv.Itoa(i+1), entry)
					form.Items = append(form.Items, formItem)
					widgets = append(widgets, entry)
				}

				container.Add(form)

				ventanaAgregar.SetContent(container)
				ventanaAgregar.Show()
			}
		}
		ventanaCargar.SetContent(container.NewVBox(input))
	})

	//* BOTÓN PARA CARGAR UN ARCHIVO CON LA CADENA A LEER *//
	selectFileButton := widget.NewButton("Cargar cadena", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err == nil && reader != nil {
				fmt.Println("Archivo seleccionado: ", reader.URI().String())
				//leer el archivo seleccionado
				scanner := bufio.NewScanner(reader)

				for scanner.Scan() {
					texto := scanner.Text()
					fmt.Println("Cadena a leer: ", texto)

					//* Convertir el automata a completo *//
					v.grafo.ConvertirAutomataACompleto()

					//* Luego lo recorremos
					v.grafo.RecorrerAutomata(texto, myWindow)

				}

				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}
			}
		}, myWindow)
	})

	//* BOTÓN PARA VER EL AUTOMATA *//
	buttonVer := widget.NewButton("Ver autómata", func() {
		v.CargarVentanaMostrarAutomata()
	})

	//* Agregamos el botón a un contenedor y lo mostramos
	content := container.New(layout.NewVBoxLayout(), button, selectFileButton, buttonVer)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

//* VENTANA PARA INGRESAR LAS TRANSICIONES DEL AUTOMATA *//
func (v *View) CargarVentanaTransiciones() {

	if len(v.grafo.GetListaVertices()) == 0 {
		fmt.Println("No hay estados para crear transiciones, crea el automata")
		return
	}

	ventanaTransiciones := v.app.NewWindow("Transiciones")
	ventanaTransiciones.Resize(fyne.NewSize(400, 500))

	estados := v.grafo.GetNombreVertices()

	label1 := widget.NewLabel("Estado inicial de la transición")
	radioGroup1 := widget.NewRadioGroup(estados, func(string) {})
	label2 := widget.NewLabel("Estado final de la transición")
	radioGroup2 := widget.NewRadioGroup(estados, func(string) {})
	label3 := widget.NewLabel("Valor de la transición")
	radioGroup3 := widget.NewRadioGroup([]string{"0", "1"}, func(string) {})

	container := container.New(layout.NewVBoxLayout())
	container.Add(widget.NewButton("Crear transición", func() {
		//* Obtener los valores de los inputs
		inicio := radioGroup1.Selected
		fin := radioGroup2.Selected
		valor := radioGroup3.Selected

		fmt.Println(inicio)
		fmt.Println(fin)
		fmt.Println(valor)

		//* Agregamos la transición al grafo
		v.grafo.IngresarArista(inicio, fin, valor)
		v.grafo.MostrarVertices()
		dialog.ShowInformation("Mensaje", "Nueva transición creada!", ventanaTransiciones)
	}))
	container.Add(widget.NewButton("Regresar", func() {
		ventanaTransiciones.Close()
	}))

	container.Add(label1)
	container.Add(radioGroup1)
	container.Add(label2)
	container.Add(radioGroup2)
	container.Add(label3)
	container.Add(radioGroup3)

	ventanaTransiciones.SetContent(container)
	ventanaTransiciones.Show()

}

//* VENTANA PARA MOSTAR EL AUTOMATA EN LA INTERFAZ GRAFICA *//
func (v *View) CargarVentanaMostrarAutomata() {

	ventanaMostrarAutomata := v.app.NewWindow("Ver autómata")
	ventanaMostrarAutomata.Resize(fyne.NewSize(800, 500))
	contenedor := container.NewWithoutLayout()

	if len(v.grafo.GetListaVertices()) == 0 {
		dialog.ShowInformation("Mensaje", "No hay estados para mostrar, crea el automata", ventanaMostrarAutomata)
		ventanaMostrarAutomata.Show()
		fmt.Println("No hay estados para mostrar, crea el automata")
		return
	}

	//* Convertir el automata a completo *//
	v.grafo.ConvertirAutomataACompleto()

	//* Si el automata tiene 4 o menos estados, lo mostramos en la interfaz grafica *//
	if len(v.grafo.GetListaVertices()) <= 4 {

		inicial := v.grafo.GetListaVertices()[0]
		inicial.SetPosicionX(100)
		inicial.SetPosicionY(100)
		estadoInicial := canvas.NewCircle(color.RGBA{R: 0, G: 255, B: 0, A: 255})
		estadoInicial.Move(fyne.NewPos(float32(inicial.GetPosicionX()), float32(inicial.GetPosicionY())))
		estadoInicial.Resize(fyne.NewSize(50, 50)) // 50x50 pixels
		contenedor.Add(estadoInicial)

		//* Creamos los estados del automata (circulos)
		for i := 1; i < len(v.grafo.GetListaVertices()); i++ {

			if v.grafo.GetListaVertices()[i].GetDato() == "Sumidero" {
				circulo := canvas.NewCircle(color.RGBA{R: 255, G: 0, B: 0, A: 255})
				circulo.Move(fyne.NewPos(200, 200))
				v.grafo.GetListaVertices()[i].SetPosicionX(200)
				v.grafo.GetListaVertices()[i].SetPosicionY(200)
				circulo.Resize(fyne.NewSize(50, 50)) // 50x50 pixels
				contenedor.Add(circulo)
			} else {
				estadoAnterior := v.grafo.GetListaVertices()[i-1]

				estadoActual := v.grafo.GetListaVertices()[i]
				estadoActual.SetPosicionX(estadoAnterior.GetPosicionX() + 200)
				estadoActual.SetPosicionY(estadoAnterior.GetPosicionY())
				circulo := canvas.NewCircle(color.RGBA{R: 0, G: 255, B: 0, A: 255})
	
				if estadoActual.GetEstadoFinal() {
					circulo = canvas.NewCircle(color.RGBA{R: 0, G: 0, B: 255, A: 255})
				}
	
				circulo.Resize(fyne.NewSize(50, 50)) // 50x50 pixels
				circulo.Move(fyne.NewPos(float32(estadoActual.GetPosicionX()), float32(estadoActual.GetPosicionY())))
				contenedor.Add(circulo)
			}
		}

		//* Creamos las transiciones del automata (lineas)
		for i := 0; i < len(v.grafo.GetListaAristas()); i++ {

			origen := v.grafo.GetListaAristas()[i].GetOrigen()
			destino := v.grafo.GetListaAristas()[i].GetDestino()
			peso := v.grafo.GetListaAristas()[i].GetPeso()

			if origen == destino {

				if peso == "0" {
					Text := canvas.NewText(peso, color.RGBA{R: 255, G: 255, B: 0, A: 255})
					Text.TextStyle = fyne.TextStyle{Bold: true}
					Text.TextSize = 20
					Text.Move(fyne.NewPos(float32(origen.GetPosicionX()+20), float32(origen.GetPosicionY()-30)))
					Line := canvas.NewLine(color.RGBA{R: 0, G: 0, B: 0, A: 255})
					Line.StrokeColor = color.NRGBA{R: 255, G: 255, B: 0, A: 255}
					Line.StrokeWidth = 6
	
					Line.Position1 = fyne.NewPos(float32(origen.GetPosicionX()+20), float32(origen.GetPosicionY()-5))
					Line.Position2 = fyne.NewPos(float32(destino.GetPosicionX()+40), float32(destino.GetPosicionY()-5))
	
					contenedor.Add(Text)
					contenedor.Add(Line)
				} else {
					Text := canvas.NewText(peso, color.RGBA{R: 255, G: 255, B: 0, A: 255})
					Text.TextStyle = fyne.TextStyle{Bold: true}
					Text.TextSize = 20
					Text.Move(fyne.NewPos(float32(origen.GetPosicionX()+20), float32(origen.GetPosicionY()+65)))
					Line := canvas.NewLine(color.RGBA{R: 0, G: 0, B: 0, A: 255})
					Line.StrokeColor = color.NRGBA{R: 255, G: 255, B: 0, A: 255}
					Line.StrokeWidth = 6
	
					Line.Position1 = fyne.NewPos(float32(origen.GetPosicionX()+20), float32(origen.GetPosicionY()+60))
					Line.Position2 = fyne.NewPos(float32(destino.GetPosicionX()+40), float32(destino.GetPosicionY()+60))
	
					contenedor.Add(Text)
					contenedor.Add(Line)
				}

			} else {

				Line := canvas.NewLine(color.RGBA{R: 255, G: 255, B: 0, A: 255})
				Line.StrokeColor = color.NRGBA{R: 255, G: 255, B: 0, A: 255}
				Line.StrokeWidth = 6

				Text := canvas.NewText(peso, color.RGBA{R: 255, G: 255, B: 0, A: 255})
				Text.TextStyle = fyne.TextStyle{Bold: true}
				Text.TextSize = 20
				Text.Move(fyne.NewPos(float32(origen.GetPosicionX()+15), float32(origen.GetPosicionY()-5)))

				Line.Position1 = fyne.NewPos(float32(destino.GetPosicionX()+25), float32(destino.GetPosicionY()+20))
				Line.Position2 = fyne.NewPos(float32(origen.GetPosicionX()+25), float32(origen.GetPosicionY()+20))

				contenedor.Add(Text)
				contenedor.Add(Line)

			}

		}
		println("------------------------------------")

		ventanaMostrarAutomata.SetContent(
			contenedor,
		)

		ventanaMostrarAutomata.Show()

	} else {
		dialog.ShowInformation("Warning", "No se puede mostrar ese autómata, tiene más de 4 estados", ventanaMostrarAutomata)
		fmt.Println("No se puede mostrar ese autómata, tiene más de 4 estados")
	}
}

//* VENTANA PARA INGRESAR LOS ESTADOS FINALES DEL AUTOMATA *//
func (v *View) CargarVentanaEstadosFinales() {

	ventanaEstadosFinales := v.app.NewWindow("Estados finales")
	ventanaEstadosFinales.Resize(fyne.NewSize(400, 500))

	estados := v.grafo.GetNombreVertices()

	label1 := widget.NewLabel("Seleccione los estados finales:")
	radioGroup1 := widget.NewRadioGroup(estados, func(string) {})

	container := container.New(layout.NewVBoxLayout())
	container.Add(widget.NewButton("Crear estado final", func() {
		// Obtener los valores de los widgets
		estadoFinal := radioGroup1.Selected

		fmt.Println(estadoFinal)

		v.grafo.GetVertice(estadoFinal).SetEstadoFinal(true)
		dialog.ShowInformation("Mensaje", "Nuevo estado final creado!", ventanaEstadosFinales)
	}))
	container.Add(widget.NewButton("Regresar", func() {
		ventanaEstadosFinales.Close()
	}))

	container.Add(label1)
	container.Add(radioGroup1)

	ventanaEstadosFinales.SetContent(container)
	ventanaEstadosFinales.Show()

}
