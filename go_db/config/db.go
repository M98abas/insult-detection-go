package config

import (
	"log"
	"os"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB{
	opts := &pg.Options{
		User:"mahmoud-dark",
		Password: "MahmoudDark123",
		Addr: "localhost:5432",
		Database: "blogdb",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect !!!!")
		os.Exit(200)
	}
	log.Printf("Everything Ok")
	return db
}