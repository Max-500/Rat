package main

import (
	"fmt"
	"runtime"
	"myFirstGame/scenes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Juego Pelota")
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(800, 600))

	game := scenes.NewIntroScene(myWindow)
	game.LoadGame()
	myWindow.ShowAndRun()

	fmt.Printf("Número de subprocesos activos: %d\n", runtime.NumGoroutine())
}