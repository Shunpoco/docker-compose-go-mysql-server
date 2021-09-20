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

	http.HandleFunc("/", handler)
	server := &http.Server{
		Addr:    ":5000",
		Handler: nil,
	}
	log.Fatal(server.ListenAndServe())
}

func handler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dataSource)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("gogo!")

	var newUser *User
	newUser, err = SelectUser(db)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(newUser)

	json.NewEncoder(w).Encode(newUser)
}

func SelectUser(db *sql.DB) (user *User, err error) {
	user = new(User)

	query := "SELECT * FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	if rows.Next() == false {
		return
	}

	err = rows.Scan(&user.ID, &user.Name)
	if err != nil {
		return
	}

	return
}
