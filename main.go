package main

import (
	"github.com/NisalSP9/WebPageAnalyzer/routes"
	"log"
	"net/http"
)

func main() {

	//Starting the API server
	router := routes.UserRoutes()
	http.Handle("/api/", router)

	//Starting the web server
	fs := http.FileServer(http.Dir("server/webapps"))
	http.Handle("/", fs)

	log.Println("Listening... port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}