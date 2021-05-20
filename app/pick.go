package app

import (
	"fmt"

	"github.com/alixq/rdv-sante/utils"
)

type picker struct {
	centers []HealthCenter
	picks   []int
}

func (app *App) FilterCenters(centers []HealthCenter) []int {
	fmt.Println("\nNous avons récupéré toutes les structures proposant des rendez-vous. Merci de sélectionner celles qui vous conviennent pour un rendez-vous.")
	fmt.Println("Pour répondre, taper o pour oui, n pour non, s pour arrêter de sélectionner des structures, puis entrée.")

	p := &picker{}
	for i := 0; i < len(centers); i++ {
		keepGoing := p.PickCenter(i, centers[i])
		if !keepGoing {
			break
		}
	}

	p.DisplayPicks()
	return p.picks
}

func (p *picker) PickCenter(i int, center HealthCenter) bool {
	fmt.Printf(
		"\n%v | %v\n%v\nSélectionner ? ",
		center.GetName(),
		center.GetFormattedAddress(),
		center.GetBookingUrl(),
	)

	answer := utils.ReadCharFromStdin()
	if answer == "s" {
		return false
	}

	if answer == "o" {
		p.centers = append(p.centers, center)
		p.picks = append(p.picks, i)
	}

	return true
}

func (p *picker) DisplayPicks() {
	fmt.Println("\nLes centres suivants ont été sélectionnés")
	for _, center := range p.centers {
		fmt.Printf("- %v | %v\n", center.GetName(), center.GetFormattedAddress())
	}
	fmt.Println()
}
