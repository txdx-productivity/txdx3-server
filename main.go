package main

import (
	"log"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

import _ "txdx3-server/migrations"

func main() {
	app := pocketbase.New()

	// If we are running with `go run`, we enable automigrate
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	if isGoRun {
		log.Println("Assuming `go run`, enabling automigrate")
	}

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: isGoRun,
	})

	// Serve static files
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*",
			apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	// Start the app
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
