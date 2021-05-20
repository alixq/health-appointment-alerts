package doctolib

import (
	"fmt"
	"time"

	"github.com/alixq/rdv-sante/app"
)

type centerInfoResponse struct {
	CenterInfo healthCenter `json:"search_result"`
}

type availability struct {
	healthCenter healthCenter
	StartDate    time.Time `json:"start_date"`
}

type availabilityDate struct {
	StartDate string         `json:"date"`
	Slots     []availability `json:"slots"`
}

type centerAvailabilityResponse struct {
	Total          int                `json:"total"`
	Availabilities []availabilityDate `json:"availabilities"`
}

type healthCenter struct {
	Address    string `json:"address"`
	ZipCode    string `json:"zipcode"`
	OrgType    string `json:"organization_status"`
	City       string `json:"city"`
	Name       string `json:"name_with_title"`
	Agendas    []int  `json:"agenda_ids"`
	MotiveId   int    `json:"visit_motive_id"`
	BookingUrl string `json:"url"`
}

func (c healthCenter) GetName() string {
	return c.Name
}

func (c healthCenter) GetFormattedAddress() string {
	return fmt.Sprintf("%v, %v %v", c.Address, c.ZipCode, c.City)
}

func (c healthCenter) GetBookingUrl() string {
	return c.BookingUrl
}

func (a availability) GetHealthCenter() app.HealthCenter {
	return a.healthCenter
}

func (a availability) GetDate() time.Time {
	return a.StartDate
}

func healthCentersConversion(centers []healthCenter) []app.HealthCenter {
	newCenters := make([]app.HealthCenter, len(centers))
	for i, v := range centers {
		newCenters[i] = app.HealthCenter(v)
	}
	return newCenters
}
