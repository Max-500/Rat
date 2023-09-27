package models

import (
	"fmt"
	"image/color"
	"math/rand"
	"myFirstGame/views"
	"sync"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var mutex = sync.Mutex{}
type Rat struct {
    Button       *widget.Button
    Window       fyne.Window
    life         bool
    hitsReceived uint
    counterRat  *canvas.Text
    counterFinalRat *canvas.Text
    initTimer chan bool
}

func NewRatModel(window fyne.Window, initTimer chan bool) *Rat {
    r := &Rat{
        Button:       widget.NewButton("", nil),
        Window:       window,
        life:         true,
        hitsReceived: 0,
        counterRat: views.NewCanvaText("Hits: 0", color.Black, 20, true, fyne.NewSize(100, 20), fyne.NewPos(0, 25)),
        counterFinalRat: views.NewCanvaText("", color.Black, 30, true, fyne.NewSize(100, 20), fyne.NewPos(300, 250)),
        initTimer: initTimer,
    }

    r.counterRat.Hide()
    r.Button.OnTapped = r.IncrementHits
    r.Button.Resize(fyne.NewSize(50, 40))
    return r
}

// CreateContainer crea un contenedor con el botón y otros componentes si es necesario
func (r *Rat) CreateContainer(components ...fyne.CanvasObject) (*fyne.Container, *fyne.Container, *fyne.Container)  {
    content := container.NewWithoutLayout(r.Button)
	content.Hide()
    content2 := container.NewWithoutLayout(r.counterRat)
    contentFinal := container.NewWithoutLayout(r.counterFinalRat)
    
    for _, component := range components {
        content.Add(component)
    }
    
    content.Resize(fyne.NewSize(50, 40))
    content.Move(fyne.NewPos(350, 260))

    return content, content2, contentFinal
}

func (r *Rat) PreStart(e chan <- bool) {
    if <- r.initTimer {
        r.Button.Enable()
        r.counterRat.Show()
        e <- true
        close(e)
    }
}

func (r *Rat) StartMove (container *fyne.Container) {
	for r.life {
		newX := float32(rand.Intn(750))
        newY := float32(rand.Intn(550))
        container.Move(fyne.NewPos(newX, newY))
        time.Sleep(1 * time.Second)
	}
}

func (r *Rat) IncrementHits() {
    mutex.Lock()
    r.hitsReceived++
    r.counterRat.Text = fmt.Sprintf("Hits: %d", r.hitsReceived)
    mutex.Unlock()
    r.counterRat.Refresh()
}