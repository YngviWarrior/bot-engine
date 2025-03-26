package jobtest_test

import (
	"log"
	"os"
	"testing"

	"github.com/YngviWarrior/bot-engine/job"
	"github.com/joho/godotenv"
)

var jobs job.JobInterface

func TestMain(m *testing.M) {
	err := godotenv.Load(`.env`)

	if err != nil {
		log.Fatal(".env file is missing")
	}

	jobs = job.NewJobs()

	code := m.Run()
	os.Exit(code)
}
