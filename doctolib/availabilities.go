package doctolib

import (
	"fmt"
	"strings"
	"time"

	"github.com/alixq/rdv-sante/app"
)

func (hh *DoctolibHealthHub) fetchCenterAvailabilities(center healthCenter) ([]app.Availability, error) {
	today := time.Now().Format("2006-01-02")

	agendaStrings := []string{}
	for _, agendaId := range center.Agendas {
		agendaStrings = append(agendaStrings, fmt.Sprint(agendaId))
	}

	agendaIds := strings.Join(agendaStrings, "-")

	availabilityUrl := fmt.Sprintf(
		"/availabilities.json?start_date=%s&visit_motive_ids=%d&agenda_ids=%s&insurance_sector=public&destroy_temporary=true&limit=3",
		today,
		center.MotiveId,
		agendaIds,
	)

	var ca centerAvailabilityResponse
	err := hh.getJSON(availabilityUrl, &ca)
	if err != nil {
		return nil, err
	}

	return hh.buildAvailabilities(center, &ca), nil
}

func (hh *DoctolibHealthHub) buildAvailabilities(center healthCenter, ca *centerAvailabilityResponse) (av []app.Availability) {
	for _, date := range ca.Availabilities {
		for _, a := range date.Slots {
			a.healthCenter = center
			av = append(av, &a)
		}
	}
	return
}
