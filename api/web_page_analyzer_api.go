package api

import (
	"encoding/json"
	"github.com/NisalSP9/WebPageAnalyzer/controller"
	"github.com/NisalSP9/WebPageAnalyzer/models"
	"log"
	"net/http"
)

func GetHTMLPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(r.Body)
	var requURL models.REQUURL
	err := decoder.Decode(&requURL)
	if err != nil {
		log.Println(err.Error())
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(err)
		return
	}
	res, err := controller.GetHTMLPage(requURL.URL)
	if err != nil {
		log.Println(err.Error())
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode("Please check entered URL")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err.Error())
	}
	return

}
