package models

import (
	"fmt"
	"math/rand"
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
    initTimer chan bool
}

func NewRatModel(window fyne.Window, initTimer chan bool, counterRat *canvas.Text) *Rat {
    r := &Rat{
        Button:       widget.NewButton("", nil),
        Window:       window,
        life:         true,
        hitsReceived: 0,
        initTimer: initTimer,
    }

    r.Button.OnTapped = func() {
        r.IncrementHits(counterRat)
    }

    r.Button.Resize(fyne.NewSize(50, 40))
    return r
}

// CreateContainer crea un contenedor con el botón y otros componentes si es necesario
func (r *Rat) CreateContainer(counterRat, counterFinalRat *canvas.Text, components ...fyne.CanvasObject) (*fyne.Container, *fyne.Container, *fyne.Container)  {
    content := container.NewWithoutLayout(r.Button)
	content.Hide()
    content2 := container.NewWithoutLayout(counterRat)
    contentFinal := container.NewWithoutLayout(counterFinalRat)
    
    for _, component := range components {
        content.Add(component)
    }
    
    content.Resize(fyne.NewSize(50, 40))
    content.Move(fyne.NewPos(350, 260))
    return content, content2, contentFinal
}

func (r *Rat) PreStart(e chan <- bool, counterRat *canvas.Text) {
    if <- r.initTimer {
        r.Button.Enable()
        counterRat.Show()
        e <- true
        close(e)
    }
}

func (r *Rat) StartMove (container *fyne.Container) {
	for r.life {
		newX := float32(rand.Intn(750))
        newY := float32(rand.Intn(550))
        container.Move(fyne.NewPos(newX, newY))
        time.Sleep(750 * time.Millisecond)
	}
}

func (r *Rat) IncrementHits(counterRat *canvas.Text) {
    mutex.Lock()
    r.hitsReceived++
    counterRat.Text = fmt.Sprintf("Hits: %d", r.hitsReceived)
    mutex.Unlock()
    counterRat.Refresh()
}