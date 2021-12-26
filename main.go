package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/riemerio/pellets-price-scraper/pellets"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("No local .env file found")
	}

	fmt.Println("post code:", os.Getenv("SCRAPING_POSTCODE"))

	fmt.Println("main started")

	s1 := gocron.NewScheduler(time.UTC)

	s1.Every(10).Minutes().Do(doRequests)
	s1.StartBlocking()
}

func doRequests() {
	go func() {
		result := pellets.FetchPriceForLoosePellets(6000)
		fmt.Println(result)
	}()

	go func() {
		result := pellets.FetchPriceForSackedPellets(6)
		fmt.Println(result)
	}()
}
