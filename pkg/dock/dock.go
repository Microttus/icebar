package dock

import (
	"fmt"
	"github.com/gotk3/gotk3/gdk"
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
	button.SetMarginTop(5)
	button.SetMarginBottom(5)

	originalPixbuf, err := gdk.PixbufNewFromFileAtScale(application.Icon, iconSize, iconSize, false)

	// Create an image for each application
	//img, err := gtk.ImageNewFromFile(application.Icon)
	img, err := gtk.ImageNewFromPixbuf(originalPixbuf)
	if err != nil {
		//log.Printf("Unable to load icon for %s: %v", application.Name, err)
		return nil, fmt.Errorf("Unable to load icon for %s: %v", application.Name, err)
	}

	// Scale icon
	//icon, err := iconTheme.LoadIcon(application.Icon, iconSize, gtk.ICON_LOOKUP_FORCE_SIZE)
	//img.SetSizeRequest(iconSize, iconSize)

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

/*
func LoadAndScaleImage(cfg *config.Config, application *config.Application) (*gtk.Widget, error) {
	// Open the image file
	file, err := os.Open(application.Icon)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %v", err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	// Create a new RGBA image with the desired size
	dst := image.NewRGBA(image.Rect(0, 0, cfg.General.IconSize, cfg.General.IconSize))

	// Scale the original image into the new RGBA image
	draw.Draw(dst, dst.Bounds(), img, img.Bounds().Min, draw.Over)
	// Create a GdkPixbuf from the scaled image
	pixbuf, err := gdk.PixbufNewFromFileAtScale()
}

*/
