package dock

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"github.com/microttus/icebar/pkg/config"
	"github.com/microttus/icebar/pkg/gui"
	"github.com/microttus/icebar/pkg/launcher"
	"log"
)

type Dock struct {
	Item []Item
}

func NewDock(apps []config.Application) *Dock {
	// Convert config.Application to Dock Item
	return &Dock{
		Item: createItems(apps),
	}
}

func createItems(apps []config.Application) []Item {
	// Create dock items from applications
	return nil
}

func AddApplicationButton(app *gui.App, application config.Application) error {
	iconSize := app.Config.General.IconSize
	//magnificationEnabled := app.Config.Behavior.Magnification
	//magnificationFactor := app.Config.Behavior.MagnificationFactor

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
	//iconTheme, err := gtk.IconThemeGetDefault()
	if err != nil {
		//log.Printf("Unable to load icon for %s: %v", application.Name, err)
		return fmt.Errorf("Unable to load icon for %s: %v", application.Name, err)
	}

	// Scale icon
	//icon, err := iconTheme.LoadIcon(application.Icon, iconSize, gtk.ICON_LOOKUP_FORCE_SIZE)
	img.SetSizeRequest(iconSize, iconSize)

	// Set initial icon size and set img for button
	//img.SetFromPixbuf(icon)
	button.Add(img)

	// Set tooltip with application name
	button.SetTooltipText(application.Name)

	// Clicked to launch
	button.Connect("clicked", func() {
		err := launcher.Launch(app, application.Name, application.Exec)
		if err != nil {
			return
		}
	})

	if app.Config.Behavior.Magnification {
		// Connect the "enter-notify-event" and "leave-notify-event" for magnification
		button.Connect("enter-notify-event", func() {
			log.Printf("Hovering over: %s", application.Name)
			img.SetPixelSize(int(float64(app.Config.General.IconSize) * app.Config.Behavior.MagnificationFactor))
		})
		button.Connect("leave-notify-event", func() {
			img.SetPixelSize(app.Config.General.IconSize)
		})
	}

	// Add button to the MainBox
	app.MainBox.PackStart(button, false, false, 0)

	return nil
}
