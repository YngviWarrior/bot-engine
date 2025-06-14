package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/YngviWarrior/bot-engine/job"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(".env file is missing")
	}

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	switch os.Getenv("ENVIROMENT") {
	case "local":
		log.SetOutput(os.Stdout)
		log.SetOutput(file)

	case "testnet", "demo":
		log.SetOutput(os.Stdout)

	case "server":
		loc, _ := time.LoadLocation("America/Sao_Paulo")
		time.Local = loc

		log.SetOutput(file)
	}

	if err != nil {
		log.Printf("%v", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	job.NewJobs().InitJobs()

	wg.Wait()
}
