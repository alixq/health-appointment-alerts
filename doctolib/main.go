package doctolib

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/alixq/rdv-sante/app"
	"github.com/alixq/rdv-sante/utils"
)

const resultsPerPage = 10

type DoctolibHealthHub struct {
	Persistence  app.PersistenceService
	domain       string
	search       string
	specialityId string
	maxCenters   int
	centers      []healthCenter
}

var HealthHub = DoctolibHealthHub{
	domain: "https://www.doctolib.fr",
}

func (hh *DoctolibHealthHub) PromptForSearch(maxCenters int) {
	hh.maxCenters = maxCenters

	fmt.Println("Pour faire fonctionner l'application, il faut insérer l'URL de votre recherche Doctolib.")
	fmt.Println("Une fenêtre va s'ouvrir, taper votre recherche et copiez-coller l'URL une fois arrivé.e sur la page de résultats")
	fmt.Print("Ouvrir Doctolib? ")
	_ = utils.ReadFromStdin()
	exec.Command("open", hh.domain).Start()

	fmt.Print("\nCollez votre URL ici: ")
	searchUrl := utils.ReadFromStdin()
	hh.search = strings.Replace(searchUrl, hh.domain, "", 1)
}

func (hh *DoctolibHealthHub) FetchAllCenters() []app.HealthCenter {
	centerIds := hh.fetchCenterIds()
	hh.fetchAllCenterInfos(centerIds)

	return healthCentersConversion(hh.centers)
}

func (hh *DoctolibHealthHub) FetchAllAvailabilities(c chan app.Availability) error {
	if len(hh.centers) == 0 {
		return app.ErrNoCenters
	}

	for i, center := range hh.centers {
		go func(i int, center healthCenter) {
			avs, err := hh.fetchCenterAvailabilities(center)
			if err != nil {
				log.Println(err)
				if strings.HasPrefix(err.Error(), "404") {
					picked := []int{}
					for j := 0; j < len(hh.centers); j++ {
						if i != j {
							picked = append(picked, j)
						}
					}
					hh.PickCenters(picked)
				}
			}
			for _, av := range avs {
				c <- av
			}
		}(i, center)
	}

	return nil
}

func (hh *DoctolibHealthHub) LoadCenters() ([]app.HealthCenter, error) {
	var centers []healthCenter
	if err := hh.Persistence.Retrieve(&centers); err != nil {
		return nil, err
	}

	hh.centers = centers
	return healthCentersConversion(centers), nil
}

func (hh *DoctolibHealthHub) ClearCenters() error {
	if err := hh.Persistence.Save([]healthCenter{}); err != nil {
		return err
	}

	hh.centers = []healthCenter{}

	return nil
}

func (hh *DoctolibHealthHub) PickCenters(indexes []int) error {
	newCenters := []healthCenter{}
	for i, center := range hh.centers {
		contained := false
		for _, j := range indexes {
			if i == j {
				contained = true
				break
			}
		}
		if contained {
			newCenters = append(newCenters, center)
		}
	}

	if err := hh.Persistence.Save(newCenters); err != nil {
		return err
	}

	hh.centers = newCenters
	return nil
}
