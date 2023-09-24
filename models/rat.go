package models

import (
	"fmt"
	"image/color"
	"math/rand"
	"myFirstGame/views"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// AntModel es un modelo para el botón y la lógica relacionada
type Rat struct {
    Button       *widget.Button
    Window       fyne.Window
    life         bool
    hitsReceived uint
    timer        uint
    counterText  *canvas.Text
    counterFinal *canvas.Text
}

// NeeRatModel crea una nueva instancia de AntModel con un botón y asigna la ventana
func NewRatModel(window fyne.Window) *Rat {
    r := &Rat{
        Button:       widget.NewButton("", nil),
        Window:       window,
        life:         true,
        hitsReceived: 0,
        timer:        10,
        counterText: views.NewCanvaText("Hits: 0", color.Black, 20, true, fyne.NewSize(100, 20), fyne.NewPos(0, 25)),
        counterFinal: views.NewCanvaText("", color.Black, 30, true, fyne.NewSize(100, 20), fyne.NewPos(300, 250)),
    }

    r.counterText.Hide()

    r.Button.OnTapped = r.IncrementHits
    r.Button.Resize(fyne.NewSize(50, 40))

    return r
}

// CreateContainer crea un contenedor con el botón y otros componentes si es necesario
func (r *Rat) CreateContainer(components ...fyne.CanvasObject) (*fyne.Container, *fyne.Container, *fyne.Container)  {
    content := container.NewWithoutLayout(r.Button)
	content.Hide()

    content2 := container.NewWithoutLayout(r.counterText)

    contentFinal := container.NewWithoutLayout(r.counterFinal)
    
    for _, component := range components {
        content.Add(component)
    }
    
    content.Resize(fyne.NewSize(50, 40))
    content.Move(fyne.NewPos(350, 260))

    return content, content2, contentFinal
}

func (r *Rat) PreStart() {
	r.Button.Enable()
    r.counterText.Show()
}

func (r *Rat) StartMove (container *fyne.Container) {
	for r.life {
		newX := float32(rand.Intn(750))
        newY := float32(rand.Intn(550))
        container.Move(fyne.NewPos(newX, newY))
        time.Sleep(1 * time.Second)
	}
}

func (r *Rat) StartTimer (timer *canvas.Text) {
	for r.timer > 0 {
        time.Sleep(time.Second * 1)
        r.timer--
        cadena := "Tiempo Restante: " + strconv.FormatUint(uint64(r.timer), 10) + " segundos"
        timer.Text = cadena
        timer.Refresh()
    }
    timer.Text = "Tu tiempo se agoto!!!"
    r.life = false
    r.Button.Disable()
    r.counterFinal.Text = fmt.Sprintf("Puntuación Final: %d", r.hitsReceived)
    r.counterFinal.Show()
    r.counterFinal.Refresh()
}

func (r *Rat) IncrementHits() {
    r.hitsReceived++
    r.counterText.Text = fmt.Sprintf("Hits: %d", r.hitsReceived)
    r.counterText.Refresh()
}