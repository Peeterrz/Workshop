package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"app/route"
	"app/shared/database"
	"app/shared/jsonconfig"
	"app/shared/server"
	"app/shared/session"
	"app/shared/view"
	"app/shared/view/plugin"
)

// *****************************************************************************
// Application Logic
// *****************************************************************************

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// Load the configuration file
	//jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	configPath := flag.String("config", dir+string(os.PathSeparator)+"config"+string(os.PathSeparator)+"config.json", "config filepath")
	jsonconfig.Load(*configPath, config)

	// Configure the session cookie store
	session.Configure(config.Session)

	// Connect to database
	database.Connect(config.Database)

	// Setup the views
	view.Configure(config.View)
	view.LoadTemplates(config.Template.Root, config.Template.Children)
	view.LoadPlugins(
		plugin.TagHelper(config.View),
		plugin.NoEscape(),
		plugin.PrettyTime())

	// Start the listener
	server.Run(route.LoadHTTP(), route.LoadHTTPS(), config.Server)
}

// *****************************************************************************
// Application Settings
// *****************************************************************************

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database database.Info   `json:"Database"`
	Server   server.Server   `json:"Server"`
	Session  session.Session `json:"Session"`
	Template view.Template   `json:"Template"`
	View     view.View       `json:"View"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

func XX() int {
	return 0
}
