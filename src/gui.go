package main

import (
	"log"
	"os"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func main() {
	// Initialize astilectron
	var a, _ = astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName:            "Notenrechner",
		AppIconDefaultPath: "", // If path is relative, it must be relative to the data directory
		AppIconDarwinPath:  "", // Same here
		BaseDirectoryPath:  "Deps",
		VersionAstilectron: "0.55.0",
		VersionElectron:    "18.2.3",
	})
	defer a.Close()

	// Start astilectron
	a.Start()

	// Create a new window
	var w, _ = a.NewWindow("http://127.0.0.1:4000", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(600),
		Width:  astikit.IntPtr(600),
	})
	w.Create()

	// Blocking pattern
	a.Wait()
}
