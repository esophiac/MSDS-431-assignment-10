package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dixonwille/wmenu/v5"
	_ "modernc.org/sqlite"
)

// create a check error function to use whereever I need it
func checkError(err error) {

	if err != nil {
		log.Fatal(err)
	}
}

// different options for interacting with the sqlite database
func userAct(database *sql.DB, opts []wmenu.Opt) {

	switch opts[0].Value {

	case 0:
		fmt.Println("Send a SQLite query to the database.")
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
