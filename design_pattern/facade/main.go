package main

import (
	"github.com/panpito/youtube/design_pattern/facade/phone_number"
)

func main() {
	consumer := phone_number.NewConsumer()

	consumer.Process("0044445645")
}
