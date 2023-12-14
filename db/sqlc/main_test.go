package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	defer conn.Close(context.Background())
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
