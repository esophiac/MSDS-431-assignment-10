# Week 10 Assginment - Building a Personal Movie Database with Go
For this assignment, we created a database with SQLite to house open-source movie data, and then used Go to interface with the database.

Describe how you might add to this database by including a table showing the movies that you have in your personal collection, where those movies are located, and perhaps your personal ratings of the movies.
<br>
text
<br>
Describe plans for drawing on the personal movie database. What purpose would it serve?
<br>
Text
<br>
Describe possible user interactions with the database (beyond what can be obtained from SQL queries alone).  In other words, what would a useful Go movie application look like? What would be the advantages of this application over IMDb alone?
<br>
text
<br>
Describe possible database enhancements and further application development.
<br>
Text
<br>
(Optional) Consider possibilities for adding movie review information to the database. What might be the possibilities for building a personal movie recommendation system based on your personal movie reviews?
<br>
Text
<br>

## Background
The database for this assignment was made with [SQLite](https://www.sqlite.org/index.html), which is a C-language library that implements a small, fast, self-contained, high-reliability, full-featured, SQL database engine. Additionally, the database was managed using [SQLiteStudio](https://sqlitestudio.pl/).

In order to interface with the database, I used the [database/sql](https://pkg.go.dev/database/sql) package from the Go standard library and then pure-Go port of SQLite from [https://pkg.go.dev/modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite).

## Roles of Programs and Data
This program was completed in Go, and the database was created in SQLite. 

### Data
The data in this assignment comes from the Internet Movie Database (IMDb, imdb.com). The data is currently available from the Northwestern library at [https://arch.library.northwestern.edu/concern/datasets/3484zh40n?locale=en](https://arch.library.northwestern.edu/concern/datasets/3484zh40n?locale=en). Only the movies and movie genre tables were used for this assignment.

### Programs
These are the programs in the repository:
- go.mod: defines the module's properties
- go.sum: record of the libraries the project depends on
- main_test.go: tests and benchmarks the fuctions in the main.go file
- main.go: run a menu that allows the user to interact with the movieDB database. You can select menu option 5 to get a demonstration of CRUD with the movies database.
- movieDB: a database of moves from IMBD. It has two tables: movies and generes.
- README.md: the readme file for the repository
- assignment10.exe: the final build for the assignment

## Application
An executable for this project was created using Windows. To create your own executable, run **go build** in the same directory as the go program. For more information, see the Gopher documentation on creating an executable [here](https://go.dev/doc/tutorial/compile-install).

## Use of AI
AI was not used for this assignment.