package dao

import (
	"os"
	"testing"

	"github.com/content-services/content-sources-backend/pkg/db"
	"github.com/rs/zerolog/log"
)

func TestMain(m *testing.M) {
	//open database connection
	var err = db.Connect()
	if err != nil {
		log.Fatal().Err(err)
	}

	// run tests
	exitCode := m.Run()

	if err := db.Close(); err != nil {
		log.Fatal().Err(err)
	}
	os.Exit(exitCode)
}