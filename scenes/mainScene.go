package scenes

import (
	"image/color"
	"myFirstGame/models"
	"myFirstGame/views"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type MainScene struct {
	window fyne.Window
}

func NewIntroScene(window fyne.Window) *MainScene {
	return &MainScene{window: window}
}

var startButton *widget.Button
func (s *MainScene) LoadGame () {
	end := make(chan bool)
	background := canvas.NewImageFromURI(storage.NewFileURI("./assets/fondo.jpeg"))
	background.Resize(fyne.NewSize(800,600))

	text := views.NewCanvaText("Rata En Accion!!!", color.Black, 35, true, fyne.NewSize(300, 200), fyne.NewPos(200, 50))

	title := views.AddContainerTitle(text, fyne.NewSize(300, 200), fyne.NewPos(200, 25))
	timer := views.NewCanvaText("Tiempo Restante: 60", color.Black, 20, true, fyne.NewSize(100, 25), fyne.NewPos(0, 0))
	timer.Hide()

	img := canvas.NewImageFromURI(storage.NewFileURI("./assets/pelota.png"))
	img.Resize(fyne.NewSize(50,40))

	ratModel := models.NewRatModel(s.window)
    rat, ratPoints, ratFinalPoints := ratModel.CreateContainer(img)

	startButton = widget.NewButton("Start Game", func() {
		timer.Show()
		rat.Show()
		title.Hide()
		startButton.Hide()
		ratModel.PreStart(end)
		go ratModel.StartTimer(timer)
		go ratModel.StartMove(rat)
	})
	
	startButton.Resize(fyne.NewSize(150,30))
	startButton.Move(fyne.NewPos(300,300))

	go toggleVisibility(title, startButton, end)

	s.window.SetContent(container.NewWithoutLayout(background, title, startButton, rat, timer, ratPoints, ratFinalPoints))
}

func toggleVisibility(label *fyne.Container, button *widget.Button, s <-chan bool) {
    for {
        select {
		case ok := <-s:
        	if ok {
				label.Hide()
				button.Hide()
			}
			return;
		default:
			time.Sleep(time.Second * 2)
        	label.Hide()
        	button.Hide()
        	time.Sleep(time.Second * 1)
        	label.Show()
       	    button.Show()
		}
    }
}