package dock

import (
	"fmt"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/microttus/icebar/pkg/config"
	"github.com/microttus/icebar/pkg/launcher"
)

/*
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
*/

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
	dockMargin := cfg.Appearance.DockMargins
	if dockMargin < 0 {
		dockMargin = 0
	}
	button.SetMarginStart(dockMargin)
	button.SetMarginEnd(dockMargin)
	button.SetMarginTop(dockMargin)
	button.SetMarginBottom(dockMargin)

	originalPixbuf, err := gdk.PixbufNewFromFileAtScale(application.Icon, iconSize, iconSize, false)

	// Create an image for each application
	img, err := gtk.ImageNewFromPixbuf(originalPixbuf)
	if err != nil {
		//log.Printf("Unable to load icon for %s: %v", application.Name, err)
		return nil, fmt.Errorf("Unable to load icon for %s: %v", application.Name, err)
	}

	// Set initial icon size and set img for button
	button.Add(img)

	// Set tooltip with application name
	if cfg.Behavior.AppNameOnHover {
		button.SetTooltipText(application.Name)
	}

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
			//log.Printf("Hovering over: %s", application.Name)
			img.SetPixelSize(int(float64(cfg.General.IconSize) * cfg.Behavior.MagnificationFactor))
		})
		button.Connect("leave-notify-event", func() {
			img.SetPixelSize(cfg.General.IconSize)
		})
	}

	return button, nil
}
