package main

import (
	"github.com/alixq/rdv-sante/app"
	"github.com/alixq/rdv-sante/doctolib"
	"github.com/alixq/rdv-sante/utils"
)

func main() {
	var hub = doctolib.HealthHub
	hub.Persistence = &utils.PersistenceService

	var App = app.App{
		HealthHub: &hub,
	}

	App.Run()
}
