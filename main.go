package main

import (
	"log"
	"net/http"

	"www.github.com/satyendra/bootcamp/handler"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			handler.Getstatus(w)
		}

	case "POST":
		{
			qparam := r.FormValue("name")
			if handler.QParam == "" {
				handler.Postall(w, r)
			} else {
				handler.Postquerry(w, qparam)
			}
		}

	}
}

func main() {

	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))

}
