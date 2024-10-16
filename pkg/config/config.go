package config

type Config struct {
	General    GeneralSettings    `toml:"general"`
	Behavior   BehaviorSettings   `toml:"behavior"`
	Appearance AppearanceSettings `toml:"appearance"`
	Dock       DockSettings       `toml:"dock"`
}

type GeneralSettings struct {
	Position string `toml:"position"`
	AutoHide bool   `toml:"auto_hide"`
	IconSize int    `toml:"icon_size"`
}

type BehaviorSettings struct {
	Magnification       bool    `toml:"magnification"`
	MagnificationFactor float64 `toml:"magnification_factor"`
}

type AppearanceSettings struct {
	MainColor  string `toml:"main_color"`
	EdgeColor  string `toml:"edge_color"`
	BlockStyle string `toml:"block_style"`
}

type DockSettings struct {
	Applications []Application `toml:"applications"`
}

type Application struct {
	Name string `toml:"name"`
	Exec string `toml:"exec"`
	Icon string `toml:"icon"`
}
