package main

import (
	"time"
)

type PriceList struct {
	Date   time.Time `json:"Date"`
	Stocks []Stock   `json:"stocks"`
}

type Stock struct {
	Symbol          string  `json:"symbol"`
	Date            Date    `json:"date"`
	Open            float64 `json:"open"`
	High            float64 `json:"high"`
	Low             float64 `json:"low"`
	Close           float64 `json:"close"`
	PrevClose       float64 `json:"prev_close"`
	Movement        float64 `json:"movement"`
	PercentMovement float64 `json:"percent_movement"`
	Volume          float64 `json:"volume"`
	Value           float64 `json:"value"`
}

type StockHistory struct {
	Symbol           string    `json:"symbol"`
	Dates            []Date    `json:"dates"`
	Opens            []float64 `json:"opens"`
	Highs            []float64 `json:"highs"`
	Lows             []float64 `json:"lows"`
	Closes           []float64 `json:"closes"`
	PrevCloses       []float64 `json:"prev_closes"`
	Movements        []float64 `json:"movements"`
	PercentMovements []float64 `json:"percent_movements"`
	Volumes          []float64 `json:"volumes"`
	Values           []float64 `json:"values"`
}

type Date time.Time

func (date Date) MarshalJSON() ([]byte, error) {
	t := time.Time(date)
	return []byte(t.Format("\"2006-01-02\"")), nil
}
