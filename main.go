package main

import (
	"myFirstGame/scenes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Atrapa la rata!!!")
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(800, 600))

	game := scenes.NewIntroScene(myWindow)
	game.LoadGame()
	myWindow.ShowAndRun()
}