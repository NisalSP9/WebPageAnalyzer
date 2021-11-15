package routes

import (
	"github.com/NisalSP9/WebPageAnalyzer/api"
	"github.com/gorilla/mux"
)

func UserRoutes() *mux.Router  {
	var router = mux.NewRouter()
	router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/gethtml", api.GetHTMLPage)
	return router
}
