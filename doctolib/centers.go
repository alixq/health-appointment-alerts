package doctolib

import (
	"fmt"
	"log"
	"math"
	"net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/alixq/rdv-sante/utils"
)

func (hh *DoctolibHealthHub) fetchCenterIds() (centerIds []string) {
	pages := int(math.Floor(float64(hh.maxCenters) / resultsPerPage))

	var wg sync.WaitGroup
	c := make(chan []string)

	for i := 1; i <= pages; i++ {
		wg.Add(1)
		go (func(page int) {
			defer wg.Done()
			c <- hh.fetchPageCenterIds(page)
		})(i)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for ids := range c {
		centerIds = append(centerIds, ids...)
	}

	return centerIds
}

func (hh *DoctolibHealthHub) fetchPageCenterIds(page int) (centerIds []string) {
	searchResultsUrl := fmt.Sprintf("%s&page=%d", hh.search, page)

	doc := hh.getHTML(searchResultsUrl)

	sel := doc.Find("[data-speciality-id]").First()
	id, exists := sel.Attr("data-speciality-id")
	if !exists {
		utils.CrashWithError("Error parsing document")
	}
	hh.specialityId = id

	doc.Find(".dl-search-result").Each(func(i int, s *goquery.Selection) {
		wholeId, exists := s.Attr("id")
		if !exists {
			utils.CrashWithError("Error parsing document")
		}

		id := strings.Replace(wholeId, "search-result-", "", 1)
		centerIds = append(centerIds, id)
	})

	return
}

func (hh *DoctolibHealthHub) fetchCenterInfo(centerId string) (*healthCenter, error) {
	url, err := url.Parse(hh.search)
	if err != nil {
		utils.CrashWithError("problem parsing url: " + err.Error())
	}

	centerInfoUrl := fmt.Sprintf(
		"/search_results/%s.json?%v&speciality_id=%s&search_result_format=json",
		centerId,
		url.Query().Encode(),
		hh.specialityId,
	)

	var infoRes centerInfoResponse
	if err := hh.getJSON(centerInfoUrl, &infoRes); err != nil {
		return nil, err
	}

	return &infoRes.CenterInfo, nil
}

func (hh *DoctolibHealthHub) fetchAllCenterInfos(centerIds []string) {
	c := make(chan healthCenter)
	var wg sync.WaitGroup

	for _, centerId := range centerIds {
		wg.Add(1)
		go (func(centerId string) {
			defer wg.Done()
			info, err := hh.fetchCenterInfo(centerId)
			if err != nil {
				log.Print(err.Error())
				return
			}
			c <- *info
		})(centerId)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for center := range c {
		center.BookingUrl = hh.domain + center.BookingUrl
		if center.MotiveId != 0 && len(center.Agendas) > 0 {
			hh.centers = append(hh.centers, center)
		}
	}
}
