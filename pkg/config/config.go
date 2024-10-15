package config

type Config struct {
	General    GeneralSettings
	Behavior   BehaviorSettings
	Appearance AppearanceSettings
	Dock       DockSettings
}

type GeneralSettings struct {
	Position string
	AutoHide bool
	IconSize int
}

type BehaviorSettings struct {
	Magnification       bool
	MagnificationFactor float64
}

type AppearanceSettings struct {
	MainColor  string
	EdgeColor  string
	BlockStyle string
}

type DockSettings struct {
	Applications []Application
}

type Application struct {
	Name string
	Exec string
	Icon string
}
