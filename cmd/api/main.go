package main

import (
	"log"

	"github.com/katsuragawaa/btc-billionaire/config"
	"github.com/katsuragawaa/btc-billionaire/internal/server"
)

// @title BTC Billionaire
// @version 1.0
// @description App that allows users to track the balance of a bitcoin wallet over time
// @contact.name Andre Katsuragawa
// @contact.url https://github.com/katsuragawaa
// @contact.email andre.katsuragawa@gmail.com
// @BasePath /api/v1
func main() {
	cfgFile, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	s := server.NewServer(cfg)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
