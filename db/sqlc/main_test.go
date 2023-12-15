package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	testDB, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	defer testDB.Close()
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
