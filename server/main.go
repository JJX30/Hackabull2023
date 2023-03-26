// This is the main entry point for the CampUs API application.

package main

import "github.com/JJX30/Rizzerator-App/server/api/app"

func main() {
	// Initializes a new instance of the App struct
	app := &app.App{}

	// Opens a connection to the SQLite database and sets up the router
	// with the appropriate handlers.
	app.Initialize()

	// Starts the HTTP server and listens for incoming requests on port 3000
	app.Run(":4000")
}
