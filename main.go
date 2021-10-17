package main

import (
	_ "embed"

	"github.com/wailsapp/wails"
)

//go:embed frontend/public/build/bundle.js
var js string

//go:embed frontend/public/build/bundle.css
var css string

func main() {
	connectDB()
	defer close()

	go clipboardManagerInit()

	app := wails.CreateApp(&wails.AppConfig{
		Width:  400,
		Height: 400,
		Title:  "cloppy",
		JS:     js,
		CSS:    css,
	})

	app.Bind(getClipboardHistory)
	app.Run()
}
