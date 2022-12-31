package main

import (
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/katsuragawaa/btc-billionaire/config"
	"github.com/katsuragawaa/btc-billionaire/internal/server"
	"github.com/katsuragawaa/btc-billionaire/pkg/db/postgres"
	"github.com/katsuragawaa/btc-billionaire/pkg/logger"
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

	appLogger := logger.NewAPILogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Env)

	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
	} else {
		appLogger.Infof("Postgres connected, Status: %#v", psqlDB.Stats())
	}
	defer func(psqlDB *sqlx.DB) {
		err := psqlDB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(psqlDB)

	s := server.NewServer(cfg, appLogger)
	if err = s.Run(); err != nil {
		log.Panic(err)
	}
}
