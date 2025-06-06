package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dixonwille/wmenu/v5"
	_ "modernc.org/sqlite"
)

// interface to hold responses from the database
type dbResponse struct {
	id       int
	name     string
	year     string
	rank     float64
	movie_id int
	genre    string
}

// create a check error function to use whereever I need it
func checkError(err error) {

	if err != nil {
		log.Fatal(err)
	}
}

// prompts the user for information
func infoRequest(info string) string {

	reader := bufio.NewReader(os.Stdin)

	promptString := fmt.Sprintf("Please enter %v", info)
	fmt.Println(promptString)

	result, _ := reader.ReadString('\n')
	if result != "\n" {
		result = strings.TrimSuffix(result, "\n")
	}

	return result

}

// send a string as a query to the sqlite movie database
func userRequest(database *sql.DB, userIn string) (err error) {

	// send request
	rows, err := database.Query(userIn)
	checkError(err)
	defer rows.Close()

	// process the returns
	for rows.Next() {
		r := &dbResponse{}
		err := rows.Scan(&r.id, &r.name, &r.year, &r.rank, &r.movie_id, &r.genre)
		checkError(err)
		fmt.Println(*r)
	}

	return err
}

// different options for interacting with the sqlite database
func userAct(database *sql.DB, opts []wmenu.Opt) {

	switch opts[0].Value {

	case 0:
		// send a SQLite query to the database
		fmt.Println("Send a SQLite query to the database.")
		// get the input from the user
		input := infoRequest("SQL query")

		// send the query to the database
		sentQ := userRequest(database, input)
		checkError(sentQ)

		break
	case 1:
		fmt.Println("Add data to the database.")
	case 2:
		fmt.Println("Update existing records in the database.")
	case 3:
		fmt.Println("Delete records in the database.")
	case 4:
		fmt.Println("Run a demonstration of CRUD operations.")
	case 5:
		fmt.Println("Quit the application.")
	}
}

func main() {

	database, err := sql.Open("sqlite", "./movieDB.db")
	checkError(err)
	// defer close
	defer database.Close()

	menu := wmenu.NewMenu("Welcome to the movie database. Please make a selection.")

	menu.Action(func(opts []wmenu.Opt) error { userAct(database, opts); return nil })

	menu.Option("Send SQLite query to the database", 0, true, nil)
	menu.Option("Add data to the database", 1, false, nil)
	menu.Option("Update existing records in the databse", 2, false, nil)
	menu.Option("Run a demonstration of CRUD operations.", 3, false, nil)
	menu.Option("Quit the applicaton", 4, false, nil)

	runErr := menu.Run()

	if runErr != nil {
		log.Fatal(runErr)
	}

}
