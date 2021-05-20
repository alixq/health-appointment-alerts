package app

import (
	"fmt"

	"github.com/alixq/rdv-sante/utils"
)

const maxCenters = 50

func (app *App) firstTime() {
	fmt.Println("Bonjour. Cette app vous permet d'être notifiés des nouveaux rendez-vous vaccins.")
	app.loadAllCenters()
}

func (app *App) loadAllCenters() {
	app.HealthHub.PromptForSearch(maxCenters)

	fmt.Println("\nChargement...")
	centers := app.HealthHub.FetchAllCenters()

	picks := app.FilterCenters(centers)
	if err := app.HealthHub.PickCenters(picks); err != nil {
		utils.CrashWithError(fmt.Sprintf("Could not persist: %v", err.Error()))
	}
}

func (app *App) alreadyUsed() {
	fmt.Println("Bonjour. Vous avez déjà une recherche enregistrée. Souhaitez-vous la lancer ?")
	fmt.Println(
		"Taper o pour utiliser la recherche enregistrée, n pour en créer une nouvelle (la précédente sera perdue).",
	)
	fmt.Print("Votre choix ? ")
	answer := utils.ReadFromStdin()
	fmt.Println()

	if answer == "n" {
		app.HealthHub.ClearCenters()
		app.loadAllCenters()
	}
}
