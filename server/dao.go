package main

import (
	"database/sql"
	// Using the blank identifier in order to solely
	// provide the side-effects of the package.
	// Eseentially the side effect is calling the `init()`
	// method of `lib/pq`:
	//  func init () {  sql.Register("postgres", &Driver{} }
	// which you can see at `github.com/lib/pq/conn.go`
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

var db *sql.DB

const (
	dbhost = "localhost"
	dbport = "5432"
	dbuser = "postgres"
	dbpass = ""
	dbname = "nairastock"
)

func dbConfig() map[string]string {
	conf := make(map[string]string)
	// host, ok := os.LookupEnv(dbhost)
	// if !ok {
	//     panic("DBHOST environment variable required but not set")
	// }
	// port, ok := os.LookupEnv(dbport)
	// if !ok {
	//     panic("DBPORT environment variable required but not set")
	// }
	// user, ok := os.LookupEnv(dbuser)
	// if !ok {
	//     panic("DBUSER environment variable required but not set")
	// }
	// password, ok := os.LookupEnv(dbpass)
	// if !ok {
	//     panic("DBPASS environment variable required but not set")
	// }
	// name, ok := os.LookupEnv(dbname)
	// if !ok {
	//     panic("DBNAME environment variable required but not set")
	// }
	conf[dbhost] = dbhost
	conf[dbport] = dbport
	conf[dbuser] = dbuser
	conf[dbpass] = dbpass
	conf[dbname] = dbname
	return conf
}

func initDb() *sql.DB {
	config := dbConfig()
	var err error
	//var db *sql.DB
	/*psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])*/
	/* postgres: //myname:password@localhost:5432/five_three_one_development?sslmode=disable */
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config[dbuser], config[dbpass], config[dbhost], config[dbport],
		config[dbname])

	db, err = sql.Open("postgres", psqlInfo)
	checkForDBError(err)

	err = db.Ping()
	checkForDBError(err)

	fmt.Println("Successfully connected to DB!")
	return db
}

func checkForDBError(err error) string {
	errMsg := ""
	if err == sql.ErrNoRows {
		errMsg = "No Records Found"
	} else if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return errMsg
}

func queryForLastDate(date *time.Time) {
	//var date time.Time
	dateQuery := `SELECT MAX(date) from stocks`
	row := db.QueryRow(dateQuery)
	if err := row.Scan(date); err != nil {
		panic(err)
	}
	//return date
}

func queryForPriceList(date time.Time) PriceList {
	stocks := make([]Stock, 0)
	var pricelist PriceList
	/*var date time.Time
	if len(dateStr) > 0 {
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			//fmt.Fprintln("error converting %s to date", err)
			queryForLastDate(&date)
		}
		//fmt.Println(date)
	} else {
		queryForLastDate(&date)
	}*/
	if date.IsZero() {
		queryForLastDate(&date)
	}
	pricelist.Date = date

	//queryStmt, err := db.Prepare(query)
	//Results from query
	var rows *sql.Rows
	var err error

	queryStmt := `SELECT DISTINCT symbol, date, open, high, low, close, prev_close, movement, percent_movement, volume, value FROM stocks WHERE date=$1`
	rows, err = db.Query(queryStmt, date)

	dbMessage := checkForDBError(err)
	if dbMessage != "" {
		return pricelist
	}

	defer rows.Close()
	for rows.Next() {
		var stock Stock
		err = rows.Scan(&stock.Symbol, &stock.Date, &stock.Open, &stock.High, &stock.Low, &stock.Close, &stock.PrevClose, &stock.Movement, &stock.PercentMovement, &stock.Volume, &stock.Value)
		checkForDBError(err)
		stocks = append(stocks, stock)
		//prices[price.Symbol] = price
	}
	pricelist.Stocks = stocks
	return pricelist
}

func queryForStockHistory(symbol string, start time.Time, end time.Time) StockHistory {
	var stockHistory StockHistory
	stockHistory.Symbol = symbol

	var rows *sql.Rows
	var err error

	queryStmt := `SELECT DISTINCT date, open, high, low, close, prev_close, movement, percent_movement, volume, value FROM stocks WHERE symbol=$1 AND date >= $2 AND date < $3 ORDER BY date`

	if start.IsZero() {
		//start, end = getDateRange(time.Duration(time.Hour * 24 * 30 * 3))
		queryStmt = `SELECT date, open, high, low, close, prev_close, movement, percent_movement, volume, value FROM stocks WHERE symbol=$1 ORDER BY date`
		rows, err = db.Query(queryStmt, symbol)
	} else {
		rows, err = db.Query(queryStmt, symbol, start, end)
	}

	dbMessage := checkForDBError(err)
	if dbMessage != "" {
		return stockHistory
	}

	defer rows.Close()
	for rows.Next() {
		var date Date
		var open float64
		var high float64
		var low float64
		var close float64
		var prev_close float64
		var movement float64
		var percent_movement float64
		var volume float64
		var value float64
		err = rows.Scan(&date, &open, &high, &low, &close, &prev_close, &movement, &percent_movement, &volume, &value)
		checkForDBError(err)
		stockHistory.Dates = append(stockHistory.Dates, date)
		stockHistory.Opens = append(stockHistory.Opens, open)
		stockHistory.Highs = append(stockHistory.Highs, high)
		stockHistory.Lows = append(stockHistory.Lows, low)
		stockHistory.Closes = append(stockHistory.Closes, close)
		stockHistory.PrevCloses = append(stockHistory.PrevCloses, prev_close)
		stockHistory.Movements = append(stockHistory.Movements, movement)
		stockHistory.PercentMovements = append(stockHistory.PercentMovements, percent_movement)
		stockHistory.Volumes = append(stockHistory.Volumes, volume)
		stockHistory.Values = append(stockHistory.Values, value)
	}

	return stockHistory
}
