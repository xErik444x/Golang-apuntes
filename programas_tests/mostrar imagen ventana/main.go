package main

import (
	"bytes"
	_ "embed"
	"image/jpeg"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

//go:embed images/pinguino.jpg
var pinguinoImage []byte

func main() {
	// Crear una nueva aplicación
	myApp := app.New()
	myWindow := myApp.NewWindow("Imagen en Fyne")
	w2 := myApp.NewWindow("Imagen en Fyne2")

	// Cargar la imagen desde el archivo embebido
	imageReader := bytes.NewReader(pinguinoImage)
	img, err := jpeg.Decode(imageReader)
	if err != nil {
		fyne.LogError("Failed to decode image", err)
		return
	}
	image := canvas.NewImageFromImage(img)

	// Ajustar el tamaño de la imagen
	image.FillMode = canvas.ImageFillContain

	// Crear un contenedor para la imagen
	content := container.NewStack(image)

	// Establecer el contenido de la ventana
	myWindow.SetContent(content)
	w2.SetContent(content)

	// Establecer el tamaño de la ventana
	myWindow.Resize(fyne.NewSize(400, 400))
	w2.Resize(fyne.NewSize(400, 400))

	// Mostrar la ventana
	myWindow.CenterOnScreen()
	w2.CenterOnScreen()

	myWindow.Show()
	w2.Show()
	myApp.Run()
}
