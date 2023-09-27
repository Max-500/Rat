package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"myFirstGame/models"
	"myFirstGame/views"
	"time"
)

type MainScene struct {
	window fyne.Window
}

func NewIntroScene(window fyne.Window) *MainScene {
	return &MainScene{window: window}
}

var startButton *widget.Button

func (s *MainScene) LoadGame() {
	end := make(chan bool)
	timerInit := make(chan bool)

	background := views.SetImage("./assets/fondo.jpeg", fyne.NewSize(800, 600))

	text := views.NewCanvaText("Rata En Accion!!!", color.Black, 35, true, fyne.NewSize(300, 200), fyne.NewPos(200, 50))

	title := views.AddContainerTitle(text, fyne.NewSize(300, 200), fyne.NewPos(200, 25))
	timer := views.NewCanvaText("Tiempo Restante: 30 segunfos", color.Black, 20, true, fyne.NewSize(100, 25), fyne.NewPos(0, 0))
	timer.Hide()

	img := views.SetImage("./assets/rat.png", fyne.NewSize(50, 40))

	ratModel := models.NewRatModel(s.window, timerInit)
	timerModel := models.NewTimerModel(timerInit)
	rat, ratPoints, ratFinalPoints := ratModel.CreateContainer(img)

	finishButton := views.CreateButton("Finish Game", s.window.Close, fyne.NewSize(150, 30), fyne.NewPos(300, 350), true)

	startButton = views.CreateButton("Start Game", func() {
		timer.Show()
		rat.Show()
		title.Hide()
		startButton.Hide()
		go ratModel.PreStart(end)
		go timerModel.StartTimer(timer, ratModel, rat, finishButton) // Hilo de logica
		go ratModel.StartMove(rat) // Hilo de logica
	}, fyne.NewSize(150, 30), fyne.NewPos(300, 300), false)

	go toggleVisibility(title, startButton, end) // Hilo decorador

	s.window.SetContent(container.NewWithoutLayout(background, title, startButton, rat, timer, ratPoints, ratFinalPoints, finishButton))
}

func toggleVisibility(label *fyne.Container, button *widget.Button, s <-chan bool) {
	for {
		select {
		case ok := <-s:
			if ok {
				label.Hide()
				button.Hide()
			}
			return
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
