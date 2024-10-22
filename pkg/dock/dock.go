package dock

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"github.com/microttus/icebar/pkg/config"
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

func AddApplicationButton(cfg *config.Config, application config.Application) (*gtk.Button, error) {
	iconSize := cfg.General.IconSize

	// Create button for each application
	button, err := gtk.ButtonNew()
	if err != nil {
		return nil, fmt.Errorf("unable to create button: %v", err)
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
		return nil, fmt.Errorf("Unable to load icon for %s: %v", application.Name, err)
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
		err := launcher.Launch(application.Name, application.Exec)
		if err != nil {
			return
		}
	})

	if cfg.Behavior.Magnification {
		// Connect the "enter-notify-event" and "leave-notify-event" for magnification
		button.Connect("enter-notify-event", func() {
			log.Printf("Hovering over: %s", application.Name)
			img.SetPixelSize(int(float64(cfg.General.IconSize) * cfg.Behavior.MagnificationFactor))
		})
		button.Connect("leave-notify-event", func() {
			img.SetPixelSize(cfg.General.IconSize)
		})
	}

	return button, nil
}
