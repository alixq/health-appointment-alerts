package app

import (
	"fmt"
	"time"

	"github.com/alixq/rdv-sante/utils"
)

type HealthCenter interface {
	GetName() string
	GetFormattedAddress() string
	GetBookingUrl() string
}

type Availability interface {
	GetHealthCenter() HealthCenter
	GetDate() time.Time
}

type HealthHub interface {
	PromptForSearch(maxCenters int)
	FetchAllCenters() []HealthCenter
	FetchAllAvailabilities(chan Availability) error
	LoadCenters() ([]HealthCenter, error)
	PickCenters([]int) error
	ClearCenters() error
}

type PersistenceService interface {
	Save(value interface{}) error
	Retrieve(ref interface{}) error
}

type App struct {
	HealthHub HealthHub
}

func (app *App) Run() {
	app.start()
	app.checkAllAvailabilities()
}

func (app *App) start() {
	centers, err := app.HealthHub.LoadCenters()
	if err != nil {
		utils.CrashWithError(err.Error())
	}

	if len(centers) == 0 {
		app.firstTime()
	} else {
		app.alreadyUsed()
	}

	fmt.Println("En attente de rendez-vous")
	fmt.Println()
}
