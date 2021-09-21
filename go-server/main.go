package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbDriver   = "mysql"
	dataSource string
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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

	var newUsers []*User
	newUsers, err = SelectUser(db)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(newUsers)

	json.NewEncoder(w).Encode(newUsers)
}

func SelectUser(db *sql.DB) (users []*User, err error) {

	query := "SELECT * FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	if rows.Next() == false {
		return
	}

	for rows.Next() {
		var user = new(User)
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return
		}
		users = append(users, user)
	}

	return
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dataSource)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	query := "INSERT INTO users (username) VALUES ('George');"
	_, err = db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
}
