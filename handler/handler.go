package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Websites struct {
	SitesName []string `json:"websites"`
}

var (
	Siteslist = Websites{}

	Webstatus = map[string]string{}

	QParam string
)

type StatusChecker interface {
	Check(ctx context.Context, name string) (status bool, err error)
}
type httpChecker struct {
}

func (h httpChecker) Check(ctx context.Context, name string) (status bool, err error) {
	_, err = http.Get("http://" + name[4:])
	if err != nil {
		return false, err

	}
	return true, err

}

var httpCheck = httpChecker{}

func Getstatus(w http.ResponseWriter) {
	for _, website := range Siteslist.SitesName {
		status, _ := httpCheck.Check(context.Background(), website)
		if status == false {
			Webstatus[website] = "DOWN"
			continue
		}

		Webstatus[website] = "UP"

	}
	//json.NewEncoder(w).Encode(webstatus)
	jsonResp, err := json.Marshal(Webstatus)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func Postall(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&Siteslist)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "200 OK")
	fmt.Println("POST received")
	fmt.Println(Siteslist)
}

func Postquerry(w http.ResponseWriter, qparam string) {
	QParam = qparam
	if searchSite() {
		var qMap = map[string]string{}
		_, err := http.Get("http://" + QParam[4:])
		if err != nil {
			qMap[QParam] = "DOWN"

		} else {

			qMap[QParam] = "UP"
		}
		json.NewEncoder(w).Encode(qMap)

	} else {
		fmt.Fprintf(w, "Required website is not availabe in server. Please use POST request to add that site.")
	}
}

func searchSite() bool {
	for _, site := range Siteslist.SitesName {
		if site == QParam {
			return true
		}

	}
	return false
}
