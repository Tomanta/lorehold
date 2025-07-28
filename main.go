package main

import (
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"github.com/tomanta/lorehold/internal/api/scryfall"
)

//go:embed all:frontend/dist
var assets embed.FS

type Config struct {
	client *scryfall.Client
}

func main() {
	client := scryfall.NewClient()
	var config = Config{
		client: client,
	}

	result, _ := scryfall.CardSearch(config.client, "Jace")

	fmt.Printf("%d cards found:\n", result.TotalCards)
	for _, card := range result.Data {
		fmt.Printf("  - %s\n", card.Name)
	}
	// LaunchApp()
}

func LaunchApp() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "lorehold",
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
