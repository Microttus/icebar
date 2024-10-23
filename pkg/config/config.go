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
	AppNameOnHover      string  `toml:"app_name_on_hover"`
}

type AppearanceSettings struct {
	MainColor       string `toml:"main_color"`
	EdgeColor       string `toml:"edge_color"`
	BorderThickness int    `toml:"border_thickness"`
	BlockStyle      string `toml:"block_style"`
	DockMargins     int    `toml:"dock_margins"`
}

type DockSettings struct {
	Applications []Application `toml:"applications"`
}

type Application struct {
	Name string `toml:"name"`
	Exec string `toml:"exec"`
	Icon string `toml:"icon"`
}
