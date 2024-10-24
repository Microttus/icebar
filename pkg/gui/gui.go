package gui

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/microttus/icebar/pkg/dock"

	"fmt"
	"github.com/microttus/icebar/pkg/config"
	"log"
)

type App struct {
	Config        *config.Config
	Window        *gtk.Window
	MainBox       *gtk.Box
	HotZoneWindow *gtk.Window
}

func (app *App) applyColors() error {
	// Set the name of the main container to "dock" for the CSS to apply
	app.MainBox.SetName("dock")

	// Create a CSS provider
	cssProvider, err := gtk.CssProviderNew()
	if err != nil {
		return err
	}

	// Ensure no negative boarder thickness
	borderThickness := app.Config.Appearance.BorderThickness
	if borderThickness < 0 {
		borderThickness = 0
	}

	//Build the CSS string with options
	css := fmt.Sprintf(`
	#dock {
	   	background-color: %s;
	   	border: %dpx solid %s;
		border-radius: 5px;
        padding: 5px;
    }
    .dock-button {
        border: none;
        background: transparent;
    }
    .dock-button:hover {
        background-color: rgba(255, 255, 255, 0.1);
	}`, app.Config.Appearance.MainColor, borderThickness, app.Config.Appearance.EdgeColor)

	// Load the CSS data
	err = cssProvider.LoadFromData(css)
	if err != nil {
		return err
	}

	// Get the default screen
	screen := app.Window.GetScreen()
	if screen == nil {
		return err
	}

	// Apply the CSS provider to the screen
	gtk.AddProviderForScreen(screen, cssProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	return nil
}

func (app *App) addApplications() error {
	for _, application := range app.Config.Dock.Applications {
		button, err := dock.AddApplicationButton(app.Config, application)
		if err != nil {
			log.Printf("Failed to add application button for %s: %v", application.Name, err)
			continue
		}
		if button == nil {
			continue
		}

		// Add button to the MainBox
		app.MainBox.PackStart(button, false, false, 0)
	}
	return nil
}

func (app *App) positionWindow() {
	// Get the default display
	display, err := gdk.DisplayGetDefault()
	if err != nil || display == nil {
		log.Println("Unable to get default display")
		return
	}

	// Get the primary monitor
	monitor, _ := display.GetPrimaryMonitor()
	if monitor == nil {
		log.Println("Unable to get primary monitor")
		return
	}

	// Get monitor geometry
	screen := monitor.GetGeometry()

	screenWidth := screen.GetWidth()
	screenHeight := screen.GetHeight()

	// Get the windows dimensions
	winWidth, winHeight := app.Window.GetSize()

	// Calculate position
	var posX, posY int

	var offsetX, offsetY int = 5, 5

	// Determine position based on config
	switch app.Config.General.Position {
	case "bottom":
		posX = (screenWidth - winWidth) / 2
		posY = screenHeight - winHeight - offsetY
	case "top":
		posX = (screenWidth - winWidth) / 2
		posY = offsetY
	case "left":
		posX = offsetX
		posY = (screenHeight - winHeight) / 2
	case "right":
		posX = screenWidth - winWidth
		posY = (screenHeight - winHeight) / 2
	default:
		posX = (screenWidth - winWidth) / 2
		posY = screenHeight - winHeight - offsetY
	}

	// Move the window to pos
	app.Window.Move(posX, posY)
}

func (app *App) positionHotZone() {
	// Get the default display
	display, err := gdk.DisplayGetDefault()
	if err != nil || display == nil {
		log.Println("Unable to get default display")
		return
	}

	// Get the primary monitor
	monitor, _ := display.GetPrimaryMonitor()
	if monitor == nil {
		log.Println("Unable to get primary monitor")
		return
	}

	// Get monitor geometry
	screen := monitor.GetGeometry()

	screenWidth := screen.GetWidth()
	screenHeight := screen.GetHeight()

	// Define hot zone dimensions
	hotZoneThickness := 5 // Thickness of the hot zone in pixels

	// Variables for position and size
	var posX, posY, width, height int

	// Determine position based on config
	switch app.Config.General.Position {
	case "bottom":
		posX = 0
		posY = screenHeight - hotZoneThickness
		width = screenWidth
		height = hotZoneThickness
	case "top":
		posX = 0
		posY = 0
		width = screenWidth
		height = hotZoneThickness
	case "left":
		posX = 0
		posY = 0
		width = hotZoneThickness
		height = screenHeight
	case "right":
		posX = screenWidth - hotZoneThickness
		posY = 0
		width = hotZoneThickness
		height = screenHeight
	default:
		// Default to bottom
		posX = 0
		posY = screenHeight - hotZoneThickness
		width = screenWidth
		height = hotZoneThickness
	}

	// Move and resize the hot zone window
	app.HotZoneWindow.Move(posX, posY)
	app.HotZoneWindow.SetDefaultSize(width, height)
}

func (app *App) CreateHotZone() error {
	var err error
	app.HotZoneWindow, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return fmt.Errorf("unable to create hot zone window: %v", err)
	}

	app.HotZoneWindow.SetTypeHint(gdk.WINDOW_TYPE_HINT_DOCK)
	app.HotZoneWindow.SetDecorated(false)
	app.HotZoneWindow.SetSkipTaskbarHint(true)
	app.HotZoneWindow.SetAcceptFocus(false)
	app.HotZoneWindow.SetKeepAbove(true)
	app.HotZoneWindow.SetOpacity(0) // make window invisible

	// Position of the hotzone
	app.positionHotZone()

	// Connect to enter-notify-event to show the dock
	app.HotZoneWindow.Connect("enter-notify-event", func(widget *gtk.Widget, event *gdk.Event) {
		app.ShowDock()
	})

	app.HotZoneWindow.ShowAll()

	return nil
}

func (app *App) HideDock() {
	app.Window.Hide()
}

func (app *App) ShowDock() {
	app.Window.ShowAll()
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

	app.Window.SetTitle("icebar")
	app.Window.SetResizable(false)
	app.Window.SetDecorated(false)
	app.Window.SetSkipTaskbarHint(false)
	app.Window.SetKeepAbove(true)
	app.Window.Connect("destroy", func() {
		log.Println("Destroy signal received. Quitting GTK main loop.")
		gtk.MainQuit()
	})

	// Auto hide
	app.Window.AddEvents(int(gdk.EVENT_ENTER_NOTIFY | gdk.EVENT_LEAVE_NOTIFY)) // Add event pointer to notify
	app.Window.Connect("leave-notify-event", func(widget *gtk.Window, event *gdk.Event) bool {
		app.HideDock()
		return false
	})
	app.Window.Connect("enter-notify-event", func(widget *gtk.Window, event *gdk.Event) bool {
		app.ShowDock()
		return false
	})

	// Find orientation
	var orientation gtk.Orientation
	switch app.Config.General.Position {
	case "left", "right":
		orientation = gtk.ORIENTATION_VERTICAL
	default:
		orientation = gtk.ORIENTATION_HORIZONTAL
	}

	// Create a box to hold dock items
	app.MainBox, err = gtk.BoxNew(orientation, 0)
	if err != nil {
		return err
	}

	// Apply colors after initializing GUI components
	if err := app.applyColors(); err != nil {
		return err
	}

	// Add the main box to the window
	app.Window.Add(app.MainBox)

	if err := app.addApplications(); err != nil {
		return fmt.Errorf("unable to add applications: %v", err)
	}

	// Show all windows
	app.Window.ShowAll()

	// Position window
	app.positionWindow()

	//Create the hot zone
	if err := app.CreateHotZone(); err != nil {
		return fmt.Errorf("unable to create hot zone: %v", err)
	}

	// Start the GTK main loop
	gtk.Main()

	return nil
}
