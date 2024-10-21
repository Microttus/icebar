package gui

import (
	"fmt"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/microttus/icebar/pkg/config"
	"log"
	"os/exec"
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
		// Create button for each application
		button, err := gtk.ButtonNew()
		if err != nil {
			return fmt.Errorf("unable to create button: %v", err)
		}
		button.SetName("dock-button")
		button.SetRelief(gtk.RELIEF_NONE)
		button.SetCanFocus(false)

		// Set margins for spacing between buttons
		button.SetMarginStart(5)
		button.SetMarginEnd(5)

		// Create an image for each application
		img, err := gtk.ImageNewFromFile(application.Icon)
		if err != nil {
			log.Printf("Unable to load icon for %s: %v", application.Name, err)
			continue // Skip this application
		}

		// Set initial icon size and set img for button
		img.SetPixelSize(app.Config.General.IconSize)
		button.Add(img)

		// Set tooltip with application name
		button.SetTooltipText(application.Name)

		// Clicked to launch
		appName := application.Name
		execPath := application.Exec
		button.Connect("clicked", func() {
			log.Printf("Launching application: %s (%s)", appName, execPath)
			cmd := exec.Command(execPath)
			err := cmd.Start()
			if err != nil {
				log.Printf("Failed to launch %s: %v", appName, err)
				showErrorDialog(app.Window, fmt.Sprintf("Failed to launch %s:\n%v", appName, err))

			}
		})

		if app.Config.Behavior.Magnification {
			// Connect the "enter-notify-event" and "leave-notify-event" for magnification
			button.Connect("enter-notify-event", func() {
				log.Printf("Hovering over: %s", appName)
				img.SetPixelSize(int(float64(app.Config.General.IconSize) * app.Config.Behavior.MagnificationFactor))
			})
			button.Connect("leave-notify-event", func() {
				img.SetPixelSize(app.Config.General.IconSize)
			})
		}

		// Add button to the MainBox
		app.MainBox.PackStart(button, false, false, 0)
	}
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

	app.Window.SetTitle("icebar")
	app.Window.SetResizable(false)
	app.Window.SetDecorated(false)
	app.Window.SetSkipTaskbarHint(true)
	app.Window.SetKeepAbove(true)
	app.Window.Connect("destroy", func() {
		log.Println("Destroy signal received. Quitting GTK main loop.")
		gtk.MainQuit()
	})

	// Create a box to hold dock items
	app.MainBox, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
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
		return fmt.Errorf("Unable to add applications: %v", err)
	}

	// Show all windows
	app.Window.ShowAll()

	// Position window
	app.positionWindow()

	// Start the GTK main loop
	gtk.Main()

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

	// Determine position based on config
	switch app.Config.General.Position {
	case "bottom":
		posX = (screenWidth - winWidth) / 2
		posY = screenHeight - winHeight
	case "top":
		posX = (screenWidth - winWidth) / 2
		posY = 0
	case "left":
		posX = 0
		posY = (screenHeight - winHeight) / 2
	case "right":
		posX = screenWidth - winWidth
		posY = (screenHeight - winHeight) / 2
	default:
		posX = (screenWidth - winWidth) / 2
		posY = screenHeight - winHeight
	}

	// Move the window to pos
	app.Window.Move(posX, posY)
}

func showErrorDialog(parent *gtk.Window, message string) {
	dialog := gtk.MessageDialogNew(parent, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_CLOSE, message)
	dialog.Run()
	dialog.Destroy()
}
