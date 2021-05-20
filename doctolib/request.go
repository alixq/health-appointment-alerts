package doctolib

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/alixq/rdv-sante/utils"
)

func (hh *DoctolibHealthHub) getRequest(uri string) *http.Response {
	res, err := http.Get(hh.domain + uri)
	if err != nil {
		utils.CrashWithError(err.Error())
	}

	return res
}

func (hh *DoctolibHealthHub) getHTML(uri string) *goquery.Document {
	res := hh.getRequest(uri)
	utils.CrashIfErrorStatus(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		utils.CrashWithError(err.Error())
	}

	return doc
}

func (hh *DoctolibHealthHub) getJSON(uri string, ref interface{}) (err error) {
	res := hh.getRequest(uri)

	if res.StatusCode > 299 {
		return fmt.Errorf(res.Status)
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&ref)
	return
}
