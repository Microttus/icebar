package main

import (
	"log"

	"github.com/microttus/icebar/pkg/config"
	"github.com/microttus/icebar/pkg/gui"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the GUI application
	app := gui.NewApp(cfg)

	// Run the application
	if err := app.Run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
