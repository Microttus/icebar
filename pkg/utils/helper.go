package utils

import (
	"github.com/gotk3/gotk3/gtk"
)

func ShowErrorDialog(parent *gtk.Window, message string) {
	dialog := gtk.MessageDialogNew(parent, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_CLOSE, message)
	dialog.Run()
	dialog.Destroy()
}
