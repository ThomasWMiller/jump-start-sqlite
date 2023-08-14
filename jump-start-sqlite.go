package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	_ "github.com/glebarez/go-sqlite"
)

func main() {

	if err := createDatabase(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := listDatabase(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func readData() [][]string {
	csvFileName := "QandA.csv"
	csvInput, err := os.Open(csvFileName)
	if err != nil {
		log.Fatalf("open %s failed: %s", csvFileName, err)
	}
	defer csvInput.Close()

	csvReader := csv.NewReader(csvInput)
	csvReader.FieldsPerRecord = -1

	fmt.Printf("Reading %s\n", csvFileName)
	allRecords, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("csvReader.ReadAll failed: %s", err)
	}

	fmt.Printf("%d records read from csv file %s:\n",
		len(allRecords), csvFileName)
	fmt.Println("--------------------------------------------------")

	for i := 0; i < len(allRecords); i++ {
		record := allRecords[i]
		fmt.Printf("question: %s answer: %s \n", record[0], record[1])
	}
	fmt.Printf("Finished reading file %s \n", csvFileName)
	return allRecords
}

func createDatabase() error {
	dbDirName := "tutorDB"
	dbFileName := "QandA.db"
	// Check if the tutorDB directory already exists
	if _, err := os.Stat(dbDirName); os.IsNotExist(err) {
		// The directory tutorDB does not exist, so create it

		fmt.Println("--------------------------------------------------")
		fmt.Printf("Creating database directory %s and database file %s \n",
			dbDirName, dbFileName)

		allRecords := readData()
		fmt.Printf("Input data allRecords has %d question/answer pairs\n", len(allRecords))

		err := os.Mkdir(dbDirName, 0755) // read and write permissions
		if err != nil {
			log.Fatalf("os.Mkdir failed: %s", err)
		}

		fn := filepath.Join(dbDirName, dbFileName)

		db, err := sql.Open("sqlite", fn)
		if err != nil {
			log.Fatalf("sql.Open failed: %s", err)
		}

		defer db.Close()

		// let qat be the name of the question-and-answer table
		stmt, err := db.Prepare(`create table if not exists qat(id integer, question text, answer text)`)
		if err != nil {
			log.Fatalf("dbPrepare create table failed: %s", err)
		}

		if _, err = stmt.Exec(); err != nil {
			log.Fatalf("stmt.Exec create table failed: %s", err)
		}

		for id := 0; id < len(allRecords); id++ {
			record := allRecords[id]
			question := record[0]
			answer := record[1]

			fmt.Printf("defining database id %d question %s and answer %s\n",
				id, question, answer)

			stmt, err = db.Prepare("insert into qat(id, question, answer) values(?, ?, ?)")
			if err != nil {
				log.Fatalf("db.Prepare insert statement failed: %s", err)
			}

			_, err = stmt.Exec(strconv.Itoa(id), question, answer)
			if err != nil {
				log.Fatalf("db.Exec populate table failed: %s", err)
			}
		}

		fmt.Printf("Finished creating database file %s \n", dbFileName)
		return nil
	} else {
		fmt.Printf("Database directory %s already exists\n", dbDirName)
		return nil
	}
}

func listDatabase() error {
	dbDirName := "tutorDB"
	dbFileName := "QandA.db"

	fn := filepath.Join(dbDirName, dbFileName)

	db, err := sql.Open("sqlite", fn)
	if err != nil {
		log.Fatalf("sql.Open failed: %s", err)
	}

	defer db.Close()

	theQuery := "select * from qat"
	// fmt.Println(theQuery)

	rows, err := db.Query(theQuery)
	if err != nil {
		log.Fatalf("db.Query failed: %s", err)
	}

	defer rows.Close()

	fmt.Println("--------------------------------------------------")
	fmt.Println("Listing all database questions and answers:")
	fmt.Println("--------------------------------------------------")
	for rows.Next() {
		var id int
		var question string
		var answer string
		rows.Scan(&id, &question, &answer)
		fmt.Printf("The answer to %s is %s\n", question, answer)
	}
	fmt.Println("--------------------------------------------------")
	return nil
}
