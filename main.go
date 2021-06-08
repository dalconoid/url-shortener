package main

import (
	"flag"
	"github.com/dalconoid/url-shortener/server"
	"github.com/dalconoid/url-shortener/storage"
	"github.com/dalconoid/url-shortener/utils"
	log "github.com/sirupsen/logrus"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to application config file")
	flag.Parse()
	config, err := utils.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	utils.CreatePGSQLConnString(config)

	s := server.New()
	db := storage.Database{ConnString: config.DBConnectionString}
	if err = db.Open(); err != nil {
		log.Fatal(err)
	}
	s.ConfigureRouter(&db)

	log.Infof("Starting sever on %s", config.ServerAddress)
	log.Fatal(s.Start(config.ServerAddress))
}