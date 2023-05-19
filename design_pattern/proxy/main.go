package main

import (
	"github.com/panpito/youtube/design_pattern/proxy/consumer"
	"github.com/panpito/youtube/design_pattern/proxy/database"
)

func main() {
	consumer := proxy_consumer.Consumer{}

	db := proxy_database.NewDatabase("../design_pattern/proxy/db.json")
	cache := proxy_database.NewCache(*db)

	consumer.Process(cache)
}
