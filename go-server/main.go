package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbDriver   = "mysql"
	dataSource string
	tpl        = template.Must(template.ParseFiles("template/index.html"))
)

type Card struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Describe  string `json:"describe"`
	Reference string `json:"reference"`
}

func main() {
	dbhost := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	dbuser := os.Getenv("MYSQL_USER")
	dbpass := os.Getenv("MYSQL_PASSWORD")
	dbport := os.Getenv("MYSQL_PORT")

	dataSource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpass, dbhost, dbport, dbname)

	fmt.Println(dataSource)

	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/set", setHandler)
	server := &http.Server{
		Addr:    ":5000",
		Handler: nil,
	}
	log.Fatal(server.ListenAndServe())
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dataSource)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	var newCards []*Card
	newCards, err = SelectUser(db)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(newCards)
	if err = tpl.Execute(w, newCards); err != nil {
		fmt.Println(err)
	}
}

func SelectUser(db *sql.DB) (cards []*Card, err error) {

	query := "SELECT * FROM cards"
	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	if rows.Next() == false {
		return
	}

	for rows.Next() {
		var card = new(Card)
		err = rows.Scan(&card.ID, &card.Title, &card.Describe, &card.Reference)
		if err != nil {
			return
		}
		cards = append(cards, card)
	}

	return
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dataSource)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	query := "INSERT INTO cards (`title`, `describe`, `reference`) VALUES ('test-title', 'test-desc', 'test-refs');"
	_, err = db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
}
