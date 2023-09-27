package models

import (
	"fmt"
	"myFirstGame/views"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Timer struct {
	timer uint
    initTimer chan bool
}

func NewTimerModel(initTimer chan bool) *Timer {
    return &Timer{
        timer: 10,
        initTimer: initTimer,
    }
}

func (t *Timer) StartTimer(timer *canvas.Text, r *Rat, imgContainter *fyne.Container, bntFinish *widget.Button) {
    timerStarted := false

    for t.timer > 0 {
        if !timerStarted {
            timerStarted = true
            go func() {
                time.Sleep(time.Second * 1)
                t.initTimer <- true // Envía una señal después del primer segundo
                close(t.initTimer)
            }()
        }

        time.Sleep(time.Second * 1)
        t.timer--
        cadena := "Tiempo Restante: " + strconv.FormatUint(uint64(t.timer), 10) + " segundos"
        timer.Text = cadena
        timer.Refresh()
    }

    // Resto de la lógica del temporizador después de que termine el bucle
    timer.Text = "Tu tiempo se agotó!!!"
    r.life = false
    r.Button.Disable()
    mutex.Lock()
    r.counterFinalRat.Text = fmt.Sprintf("Puntuación Final: %d", r.hitsReceived)
    mutex.Unlock()
    r.counterFinalRat.Show()
    r.counterFinalRat.Refresh()
    img := views.SetImage("./assets/ratF.png", fyne.NewSize(50,40))
    imgContainter.Add(img)
    bntFinish.Show()
}