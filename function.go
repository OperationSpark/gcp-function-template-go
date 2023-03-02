package gcpfunctemplate

import (
	"fmt"
	"net/http"
)

type (
	App struct {
		// Add any state/data/config you need on the App struct
	}
)

// New creates an new App instance.
func New() *App {
	// Set up any dependencies here
	return &App{}
}

// MainHandler is the entry point to the HTTP-triggered Cloud Function.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
