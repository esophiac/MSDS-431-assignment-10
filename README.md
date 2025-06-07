# Week 10 Assginment - Building a Personal Movie Database with Go
For this assignment, we created a database with SQLite to house open-source movie data, and then used Go to interface with the database.

*Describe how you might add to this database by including a table showing the movies that you have in your personal collection, where those movies are located, and perhaps your personal ratings of the movies.*
<br>
I would add to the database by creating a table for each concept (one for personal collection, one for movie location, etc.) and then connecting to the other tables in the database based on the movie id. I would also include the movie name because its hard to navigate based on movie id alone, but there is enough repitition in movie titles that it would be better to connect everything with the movie id. The movie id would be the foreign key in each of the tables.
<br>
<br>
*Describe plans for drawing on the personal movie database. What purpose would it serve?*
<br>
I would mostly use it to keep track of movies that I liked or disliked, but that would also require a lot more tables.
<br>
<br>
*Describe possible user interactions with the database (beyond what can be obtained from SQL queries alone).  In other words, what would a useful Go movie application look like? What would be the advantages of this application over IMDb alone?*
<br>
I think the advantages would be to run more strategic level analysis of movies and the entertainment industry. IMBDb has a user interface that (to me) is kind of difficult to interact with. Using Go would give people more freedom to interact with the information in a way that makes sense to them.
<br>
<br>
*Describe possible database enhancements and further application development.*
<br>
The main thing would be GUI. The menu option is a good start, but is also rigid. I also think that I would implement some method of cleaning up the SQLite queries so that its easier to work with. SQLite was just different enough from regular SQL that it was odd to use when building the application. I would also include more robust genre definitions as a table, include other people affiliated with movies that just actors (directors, cinematoraphers, producers, etc.), and other long-form text information about the movies (reviews, summaries, etc.). Maybe a controversies section would be fun?
<br>
<br>
*(Optional) Consider possibilities for adding movie review information to the database. What might be the possibilities for building a personal movie recommendation system based on your personal movie reviews?*
<br>
I think that this is where I would want to bring in machine learning as a recommendation method. It would be fun to take the movies that I have already watched, rate them myself, and then create a prediction table where it predicts if I would enjoy the next movies or not.
<br>
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