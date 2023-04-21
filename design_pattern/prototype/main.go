package main

import (
	"golang.org/x/text/language"
	"log"
	"tutorial/design_pattern/draft/prototype"
	"tutorial/design_pattern/draft/publisher"
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
