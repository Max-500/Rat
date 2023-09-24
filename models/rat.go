package models

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// AntModel es un modelo para el botón y la lógica relacionada
type RatModel struct {
    Button *widget.Button
    Window fyne.Window
	life bool
	hitsReceived, timer uint
}

// NewAntModel crea una nueva instancia de AntModel con un botón y asigna la ventana
func NewAntModel(window fyne.Window) *RatModel {
    antButton := widget.NewButton("", func() {
        window.SetTitle("¡Botón Transparente Clickeado!")
    })
	antButton.Disable()
    antButton.Resize(fyne.NewSize(50, 40))

    return &RatModel{
        Button: antButton,
        Window: window,
		life: true,
		hitsReceived: 0,
		timer: 60,
    }
}

// CreateContainer crea un contenedor con el botón y otros componentes si es necesario
func (a *RatModel) CreateContainer(components ...fyne.CanvasObject) *fyne.Container {
    content := container.NewWithoutLayout(a.Button)
	content.Hide()
    for _, component := range components {
        content.Add(component)
    }
    
    content.Resize(fyne.NewSize(50, 40))
    content.Move(fyne.NewPos(350, 260))

    return content
}

func (a *RatModel) PreStart() {
	a.Button.Enable()
}

func (a *RatModel) StartMove (container *fyne.Container, channel chan <- bool) {
	for a.life {
		newX := float32(rand.Intn(750))
        newY := float32(rand.Intn(550))
        container.Move(fyne.NewPos(newX, newY))
        time.Sleep(1 * time.Second)
	}
}

func (a *RatModel) StartTimer (timer *canvas.Text) {
	for a.timer > 0 {
        time.Sleep(time.Second * 1)
        a.timer--
        cadena := "Tiempo Restante: " + strconv.FormatUint(uint64(a.timer), 10) + " segundos"
        fmt.Println(a.timer)
        timer.Text = cadena
        timer.Refresh()
    }
    timer.Text = "Tu tiempo se agoto!!!"
}