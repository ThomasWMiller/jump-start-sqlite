# jump-start-sqlite
This code is provided as a preliminary to the development of an AI tutoring facility. The prototype can draw on an SQLite database.

We show how to use a pure Go implementation of SQLite: [go-sqlite](https://github.com/glebarez/go-sqlite). 

Note that many online Go examples with SQLite rely on CGO and the [sqlite3 driver](https://github.com/mattn/go-sqlite3). We feel that a pure Go implementation is preferred (assuming that it tests out) because pure Go allows us to work with a code base that is more easily ported across operating systems. 

This jump-start code shows how to create an SQLite database from a comma-delimited text file.

The function readData reads the comma-delimited text file QandA.csv, which  happens to include the 25 Go keywords and their meanings.

The function createDatabase creates an SQLite database file QandA.db under the tutorDB directory if the tutorDB directory does not exist. The name "qat" is used for the question-and-answer table, the only table in the database.

Our methods are similar to an online example showing how a CSV file can be used to populate an SQLite database: [Example from Universal Glue](https://universalglue.dev/posts/csv-to-sqlite/).

The function listDatabase queries the database to produce a complete listing of the questions/keywords and answers/definitions. 

Interested in setting up tests? See [Writing Tests for Your Database Code in Go](https://markphelps.me/posts/writing-tests-for-your-database-code-in-go/).

