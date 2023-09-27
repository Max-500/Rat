package views

import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func NewCanvaText(text string, color color.Gray16, sizeWord uint, bold bool, size fyne.Size, position fyne.Position) *canvas.Text {
	customLabel := canvas.NewText(text, color)
	customLabel.TextSize = float32(sizeWord)
	customLabel.TextStyle.Bold = bold
	customLabel.Resize(size)
	customLabel.Move(position)
	return customLabel
}

func AddContainerTitle(object *canvas.Text, size fyne.Size, position fyne.Position) *fyne.Container {
	container := container.New(layout.NewHBoxLayout(), object)
	container.Resize(size)
	container.Move(position)
	return container
}

func SetImage(image string, size fyne.Size) *canvas.Image {
	img := canvas.NewImageFromURI(storage.NewFileURI(image))
	img.Resize(size)
	return img
}

func CreateButton(label string, clickFunc func(), size fyne.Size, pos fyne.Position, val bool) *widget.Button {
	btn := widget.NewButton(label, func ()  {
		clickFunc()
	})

	btn.Resize(size)
	btn.Move(pos)
	btn.Hidden = val
	return btn
}