package main

import (
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	ctx := context.Background()
	loadEnv()

	repo := NewRepo()
	repo.ConnectDB()

	cs := NewClipboardService(repo)
	go cs.StartClipboardService(ctx)

	// Create an instance of the app structure
	app := NewApp(cs)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "cloppy",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
