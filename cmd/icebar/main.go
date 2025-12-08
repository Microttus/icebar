package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/microttus/icebar/pkg/config"
	"github.com/microttus/icebar/pkg/gui"
)

func main() {
	// --- CLI flags ---
	help := flag.Bool("h", false, "Show help")
	helpLong := flag.Bool("help", false, "Show help")
	configPath := flag.String("config", "", "Path to custom config directory")

	flag.Parse()

	// Handle help flags first
	if *help || *helpLong {
		fmt.Println("Usage: icebar [--config <path>]")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("  -h, --help           Show this help message")
		fmt.Println("      --config <path>  Use a custom config directory")
		os.Exit(0)
	}

	// Load configuration
	var cfg *config.Config
	var err error

	if *configPath != "" {
		cfg, err = config.Load(*configPath)
	} else {
		cfg, err = config.Load()
	}
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
