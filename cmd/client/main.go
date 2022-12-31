package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/katsuragawaa/btc-billionaire/internal/models"
)

func main() {
	fmt.Println("simple client to keep sending BTC")

	client := resty.New()

	for {
		loc, err := time.LoadLocation(pickRandomLocal())
		if err != nil {
			log.Fatal("error loading location")
		}
		dt := time.Now().In(loc).Format("2006-01-02T15:04:05-07:00")

		t := models.Transaction{}
		res, err := client.R().SetBody(
			map[string]interface{}{
				"datetime": dt,
				"amount":   getRandomAmount(),
			},
		).SetResult(&t).Post("http://localhost:8080/api/v1/transactions")
		if err != nil {
			log.Fatal("failure to send transaction")
		}

		log.Printf("status %s | sent a transaction [ID: %s]", res.Status(), t.ID)
		time.Sleep(time.Second)
	}
}

func pickRandomLocal() string {
	locals := [10]string{
		"Etc/GMT",
		"Etc/GMT+2",
		"Etc/GMT+4",
		"Etc/GMT+8",
		"Etc/GMT+12",
		"Etc/GMT-2",
		"Etc/GMT-4",
		"Etc/GMT-8",
		"Etc/GMT-12",
	}

	randomIndex := rand.Intn(len(locals))

	return locals[randomIndex]
}

func getRandomAmount() float64 {
	var min = 0.01
	var max float64 = 10

	return min + rand.Float64()*(max-min)
}
