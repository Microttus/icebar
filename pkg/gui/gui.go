package gui

import (
	"fmt"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/microttus/icebar/pkg/config"
)

type App struct {
	Config  *config.Config
	Window  *gtk.Window
	MainBox *gtk.Box
}

func (app *App) applyColors() error {
	// Set the name of the main container to "dock" for the CSS to apply
	app.MainBox.SetName("dock")

	// Create a CSS provider
	cssProvider, err := gtk.CssProviderNew()
	if err != nil {
		return err
	}

	//Build the CSS string with background color and edge color
	css := fmt.Sprintf(`
	#dock {
	   background-color: %s;
	   border: 1px solid %s;
	}`, app.Config.Appearance.MainColor, app.Config.Appearance.EdgeColor)

	// Load the CSS data
	err = cssProvider.LoadFromData(css)
	if err != nil {
		return err
	}

	// Get the default screen
	screen, err := gdk.ScreenGetDefault()
	if err != nil {
		return err
	}

	// Apply the CSS provider to the screen
	gtk.AddProviderForScreen(screen, cssProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	return nil
}

func NewApp(cfg *config.Config) *App {
	return &App{
		Config: cfg,
	}
}

func (app *App) Run() error {
	// Initialize GTK
	gtk.Init(nil)

	// Create the main window
	var err error
	app.Window, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return err
	}

	// Create a box to hold dock items
	app.MainBox, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return err
	}

	// Add the main box to the window
	app.Window.Add(app.MainBox)

	// Apply colors after initializing GUI components
	if err := app.applyColors(); err != nil {
		return err
	}

	// Show all windows
	app.Window.ShowAll()

	// Start the GTK main loop
	gtk.Main()

	return nil
}

//func parseColor(colorStr string) (*gdk.RGBA, error) {
//	color := &gdk.RGBA{}
//	if !color.Parse(colorStr) {
//		return nil, fmt.Errorf("invalid color form
//		at: %s", colorStr)
//	}
//	return color, nil
//}
