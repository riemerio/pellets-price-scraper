package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/riemerio/pellets-price-scraper/db"
	"github.com/riemerio/pellets-price-scraper/pellets"
	"github.com/riemerio/pellets-price-scraper/prices"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	connection := db.Connect("localhost", 27017)
	fmt.Println(connection.NumberSessionsInProgress())
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("No local .env file found")
	}

	fmt.Println("post code:", os.Getenv("SCRAPING_POSTCODE"))

	fmt.Println("main started")

	c := cron.New()

	c.AddFunc(os.Getenv("SCRAPING_CRON"), doRequests)
	c.Start()

	//sig := make(chan os.Signal)
	//signal.Notify(sig, os.Interrupt, os.Kill)
	//<-sig

	select {}
}

func doRequests() {
	go func() {
		result := pellets.FetchPriceForLoosePellets(6000)
		entry, err := db.Collection.InsertOne(db.CTX, &prices.Price{
			ID:        primitive.NewObjectID(),
			Price:     result.NettoEinzelpreis,
			Timestamp: time.Now(),
			RecordId:  pellets.Loose,
		})
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("loose pellets: %s with id %s", result.NettoEinzelpreis, entry.InsertedID)
	}()

	go func() {
		result := pellets.FetchPriceForSackedPellets(6)
		entry, err := db.Collection.InsertOne(db.CTX, &prices.Price{
			ID:        primitive.NewObjectID(),
			Price:     result.NettoEinzelpreis,
			Timestamp: time.Now(),
			RecordId:  pellets.Sacked,
		})
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("sacked pellets: %s with id %s", result.NettoEinzelpreis, entry.InsertedID)
	}()
}
