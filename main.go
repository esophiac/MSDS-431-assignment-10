package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	"reflect"
	"strings"

	"github.com/dixonwille/wmenu/v5"
	_ "modernc.org/sqlite"
)

// interface to hold responses from the database
type DBresponse struct {
	id       int64
	name     string
	year     string
	rank     any
	movie_id any
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

	columns, err := rows.Columns()
	checkError(err)
	colNum := len(columns)

	var values = make([]interface{}, colNum)
	for i, _ := range values {
		var ii interface{}
		values[i] = &ii
	}

	for rows.Next() {
		err := rows.Scan(values...)
		fmt.Println(err)
		for i, colName := range columns {
			var raw_value = *(values[i].(*interface{}))
			var raw_type = reflect.TypeOf(raw_value)

			fmt.Println(colName, raw_type, raw_value)
		}
	}

	return err
}

// accept and run a user SQL statement to update a row in the database or delete a row
func updateDelete(database *sql.DB, userUpdate string) int64 {
	result, err := database.Exec(userUpdate)
	checkError(err)

	rowNum, _ := result.RowsAffected()

	return rowNum
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
		// add rows to the database

		// get query from user
		sqlQuery := infoRequest("SQLite table addition")

		rowsChanged := updateDelete(database, sqlQuery)
		upResp := fmt.Sprintf("Changed %v rows.", rowsChanged)
		fmt.Println(upResp)

		break

	case 2:
		// update rows in the database

		// get query from user
		sqlQuery := infoRequest("SQLite table update")

		rowsChanged := updateDelete(database, sqlQuery)
		upResp := fmt.Sprintf("Changed %v rows.", rowsChanged)
		fmt.Println(upResp)

		break

	case 3:
		// delete records from the database

		// get query from user
		sqlQuery := infoRequest("SQLite deletion query")

		rowsChanged := updateDelete(database, sqlQuery)
		upResp := fmt.Sprintf("Deleted %v rows.", rowsChanged)
		fmt.Println(upResp)

		break

	case 4:
		fmt.Println("Run a demonstration of CRUD operations.")
	case 5:
		// quit the application
		fmt.Println("Application closing")
		os.Exit(3)
	}
}

func main() {

	database, err := sql.Open("sqlite", "movieDB")
	checkError(err)
	// defer close
	defer database.Close()

	menu := wmenu.NewMenu("Welcome to the movie database. Please make a selection.")

	menu.Action(func(opts []wmenu.Opt) error { userAct(database, opts); return nil })

	menu.Option("Send SQLite query to the database", 0, true, nil)
	menu.Option("Add data to the database", 1, false, nil)
	menu.Option("Update existing records in the databse", 2, false, nil)
	menu.Option("Delete records from the database", 3, false, nil)
	menu.Option("Run a demonstration of CRUD operations.", 4, false, nil)
	menu.Option("Quit the applicaton", 5, false, nil)

	runErr := menu.Run()

	if runErr != nil {
		log.Fatal(runErr)
	}

}
