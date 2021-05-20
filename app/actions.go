package app

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/alixq/rdv-sante/utils"
)

var ErrNoCenters = fmt.Errorf("NO CENTERS")
var lock = WinnerLock{}

type WinnerLock struct {
	mu sync.Mutex
}

func (w *WinnerLock) Win(info HealthCenter) {
	w.mu.Lock()
	defer w.mu.Unlock()

	err := exec.Command("open", info.GetBookingUrl()).Start()
	if err != nil {
		utils.CrashWithError(fmt.Sprintf("Could not launch browser: %v\n", err.Error()))
	}

	notification := fmt.Sprintf("'Rendez-vous trouvé à %s'", info.GetName())
	err = exec.Command("say", "-v", "Thomas", notification).Start()
	if err != nil {
		utils.CrashWithError(fmt.Sprintf("Could not launch browser: %v\n", err.Error()))
	}

	os.Exit(0)
}

func (app *App) checkAvailability(av Availability) {
	tomorrow := utils.GetStartOfDay(time.Now().Add(time.Hour * 48))

	if av.GetDate().Before(tomorrow) {
		lock.Win(av.GetHealthCenter())
	}
}

func (app *App) checkAllAvailabilities() {
	c := make(chan Availability)

	go (func() {
		for {
			err := app.HealthHub.FetchAllAvailabilities(c)
			if err == ErrNoCenters {
				fmt.Printf("You didn't select any center. Exiting...")
				break
			}

			fmt.Print(".")
			time.Sleep(time.Second)
		}
		close(c)
	})()

	for av := range c {
		app.checkAvailability(av)
	}
}
