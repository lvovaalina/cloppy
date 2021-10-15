package main

import (
	_ "embed"

	"github.com/wailsapp/wails"
)

const serviceName = "cloppy"
const serviceDescription = "Clipboard Manager"

//go:embed frontend/public/build/bundle.js
var js string

//go:embed frontend/public/build/bundle.css
var css string

func main() {

	go clipboardManagerInit()

	app := wails.CreateApp(&wails.AppConfig{
		Width:  400,
		Height: 400,
		Title:  "cloppy",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(getClipboardHistory)
	app.Run()
}
