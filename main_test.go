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
