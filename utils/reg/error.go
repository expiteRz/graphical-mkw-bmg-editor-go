package reg

import (
	"github.com/sqweek/dialog"
	"os"
)

func errorDialog(err error) {
	dialog.Message("%v\n\nPress OK to exit the application", err).Title("Fatal error").Info()
	os.Exit(1)
}
