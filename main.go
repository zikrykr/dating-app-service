package main

import (
	"github.com/dating-app-service/cmd/rest"
	appSetup "github.com/dating-app-service/cmd/setup"
	"github.com/dating-app-service/config"
)

func main() {
	// config init
	config.InitConfig()

	// app setup init
	setup := appSetup.InitSetup()

	// starting REST server
	rest.StartServer(setup)
}
