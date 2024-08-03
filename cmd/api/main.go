package main

import (
	"log"
	"os"

	"github.com/Zeo-dev3/fashionmart/config"
	"github.com/Zeo-dev3/fashionmart/internal/server"
	"github.com/Zeo-dev3/fashionmart/pkg/db"
	"github.com/Zeo-dev3/fashionmart/pkg/utils"
)

func main() {
	log.Println("Starting api server")

	cfgPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("load config : %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("parse config : %v", err)
	}

	db, err := db.NewDB(cfg)
	if err != nil {
		log.Fatalf("cannot connect to database : %v", err)
	}

	s := server.NewServer(cfg, db)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
