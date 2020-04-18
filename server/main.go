package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	STATIC_DIR = "/public/"
)

var DURATION = map[string]int{
	"1mnths": 24 * 30,
	"3mnths": 24 * 3 * 30,
	"6mnths": 24 * 6 * 30,
	"1yr":    24 * 12 * 30,
	"2yr":    24 * 24 * 30,
	"3yr":    24 * 36 * 30,
	"5yr":    24 * 60 * 30,
}

func main() {
	db := initDb()
	defer db.Close()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter() //.StrictSlash(false)
	router.HandleFunc("/pricelist", PriceListHandler).Methods("GET")
	router.HandleFunc("/stockhistory/{symbol}", StockHistoryHandler).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":"+port, n)
}

func getDateRange(duration time.Duration) (time.Time, time.Time) {
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	//Go back in time by 30 days
	start := date.Add(-1 * time.Duration(duration))
	end := date
	fmt.Println("start:", start, ", end:", end)
	return start, end
}

func PriceListHandler(rw http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	dateStr := queryValues.Get("date")
	//fmt.Print(dateStr)
	var date time.Time
	//var err error
	if len(dateStr) > 0 {
		date, _ = time.Parse("2006-01-02", dateStr)
	}
	pricelist := queryForPriceList(date)

	out, err := json.Marshal(pricelist)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	fmt.Fprintf(rw, string(out))
}

func StockHistoryHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	symbol := vars["symbol"]
	/*startDateStr := vars["start"]
	endDateStr := vars["end"]*/

	queryValues := r.URL.Query()
	durationStr := queryValues.Get("duration")
	startDateStr := queryValues.Get("start")
	endDateStr := queryValues.Get("end")
	duration := DURATION[durationStr]
	var startDate time.Time
	var endDate time.Time

	if duration > 0 {
		startDate, endDate = getDateRange(time.Duration(duration))
	} else {
		if len(startDateStr) > 0 {
			startDate, _ = time.Parse("2006-01-02", startDateStr)
		}
		if len(endDateStr) > 0 {
			endDate, _ = time.Parse("2006-01-02", endDateStr)
		}
	}
	symbol = strings.ToUpper(symbol)
	//fmt.Print(dateStr)
	stockHistory := queryForStockHistory(symbol, startDate, endDate)

	out, err := json.Marshal(stockHistory)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	fmt.Fprintf(rw, string(out))
}
