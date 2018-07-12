package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
	"os"
)

const (
	STATIC_DIR = "/public/"
)

func main() {
	db := initDb()
	defer db.Close()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter() //.StrictSlash(false)
	router.HandleFunc("/pricelist", PriceListHandler).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":"+port, n)
}

func PriceListHandler(rw http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	dateStr := queryValues.Get("date")
	//fmt.Print(dateStr)
	pricelist := queryForPriceList(dateStr)

	out, err := json.Marshal(pricelist)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	fmt.Fprintf(rw, string(out))
}
