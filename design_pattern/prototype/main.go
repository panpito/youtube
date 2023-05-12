package main

import (
	"github.com/panpito/youtube/design_pattern/prototype/prototype"
	"github.com/panpito/youtube/design_pattern/prototype/publisher"
	"golang.org/x/text/language"
	"log"
)

func main() {
	publisher := prototype_publisher.NewMyPublisher(language.Turkish)
	run(publisher, "example of i")

	registry := prototype.NewRegistry()
	registry.Add("tur", publisher)

	turkishPublisher, err := registry.Get("eng")
	if err != nil {
		log.Print(err)
	}
	run(turkishPublisher.(prototype_publisher.IPublisher), "example of i 2")

	publisher2 := publisher.Clone().(prototype_publisher.IPublisher)
	run(publisher2, "i live")
}

func run(publisher prototype_publisher.IPublisher, message string) {
	publisher.Publish(message)
}
