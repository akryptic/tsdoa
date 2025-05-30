package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "Todo",
		DisableResize: true,
		Frameless:     true,
		MaxWidth:      1280,
		Width:         1280,
		MinWidth:      1280,
		Height:        720,
		MinHeight:     720,
		MaxHeight:     720,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app.taskService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
