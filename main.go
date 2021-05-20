package main

import (
	"github.com/alixq/rdv-sante/app"
	"github.com/alixq/rdv-sante/doctolib"
	"github.com/alixq/rdv-sante/utils"
)

func main() {
	var hub = doctolib.HealthHub
	var persistence = &utils.PersistenceService
	persistence.SetPath()
	hub.Persistence = &utils.PersistenceService

	var App = app.App{
		HealthHub: &hub,
	}

	App.Run()
}
