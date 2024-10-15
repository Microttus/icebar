package gui

import (
	"github.com/microttus/icebar/pkg/config"
)

type App struct {
	Config *config.Config
}

func NewApp(cfg *config.Config) *App {
	return &App{
		Config: cfg,
	}
}

func (app *App) Run() error {

	return nil
}
