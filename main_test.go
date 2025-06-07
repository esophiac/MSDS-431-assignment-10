package main

import (
	"database/sql"
	"testing"
)

// test the infoRequest test
func TestInfoRequest(t *testing.T) {

	// user has to input test for this to pass
	expected := ""

	out := infoRequest("test")

	if expected != out {
		t.Errorf("Expected %v, got %v", expected, out)
	}
}

// test sending a query to a database
func TestUserRequest(t *testing.T) {

	database, _ := sql.Open("sqlite", "movieDB")
	defer database.Close()

	in := "SELECT * FROM movies WHERE id=3"

	out := userRequest(database, in)

	if out != nil {
		t.Errorf("Expected nil, got %v", out)
	}
}

// test the updateDelete function
func TestUpdateDelete(t *testing.T) {

	database, _ := sql.Open("sqlite", "movieDB")
	defer database.Close()

	in := "INSERT INTO movies (year, rank) VALUES(20000 , 10)"

	out := updateDelete(database, in)

	expected := int64(1)

	if out != expected {
		t.Errorf("Expected nil, got %v", out)
	}

	//delete the test data from the database
	in2 := "DELETE FROM movies WHERE year = 20000"
	updateDelete(database, in2)
}
