package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
	cs  ClipboardService
}

// NewApp creates a new App application struct
func NewApp(clipboardService ClipboardService) *App {
	return &App{
		cs: clipboardService,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetClipboardHistory() []string {
	return a.cs.GetClipboardHistory()
}

func (a *App) SetValueToClipboard(val string) {
	a.cs.SetValueToClipboard(val)
}
