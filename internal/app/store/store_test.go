package store_test

import (
	"os"
	"testing"
)

var (
	dataBaseURL string
)

func TestMain(m *testing.M) {
	dataBaseURL = os.Getenv("DATABASE_URL")
	if dataBaseURL == "" {
		dataBaseURL = "host=localhost dbname=restapi_test user=postgres password=admin sslmode=disable"
	}
	os.Exit(m.Run())
}
