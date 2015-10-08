package gui

import (
	"fmt"

	"github.com/twstrike/go-gtk/gtk"
)

func verifyFingerprintDialog(account *Account, uid string) *gtk.Dialog {
	dialog := gtk.NewDialog()
	dialog.SetTitle("Fingerprint verification")
	dialog.SetPosition(gtk.WIN_POS_CENTER)
	vbox := dialog.GetVBox()

	conversation := account.GetConversationWith(uid)
	fpr := conversation.GetTheirKey().DefaultFingerprint()

	// message copied from libpurple
	message := fmt.Sprintf(`
	Fingerprint for you (%s): %x

	Purported fingerprint for %s: %x

	Is this the verifiably correct fingerprint for %s?
	`, account.Account, account.Session.PrivateKey.DefaultFingerprint(), uid, fpr, uid)
	vbox.Add(gtk.NewLabel(message))

	button := gtk.NewButtonWithLabel("Verify")
	vbox.Add(button)

	button.Connect("clicked", func() {
		account.AuthorizeFingerprint(uid, string(fpr))
		//SAVE ??
		dialog.Destroy()
	})

	return dialog
}