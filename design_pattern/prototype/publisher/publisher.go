package prototype_publisher

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/panpito/youtube/design_pattern/prototype/prototype"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
)

type IPublisher interface {
	prototype.IPrototype
	Publish(message string)
}

type MyPublisher struct {
	id       string
	language language.Tag
}

func (publisher MyPublisher) Clone() interface{} {
	return MyPublisher{
		id:       publisher.id,
		language: publisher.language,
	}
}

func NewMyPublisher(language language.Tag) *MyPublisher {
	return &MyPublisher{
		id:       uuid.New().String(),
		language: language,
	}
}

func (publisher MyPublisher) Publish(message string) {
	fullMessage := fmt.Sprint("Publisher ", publisher.id, " > ", message)

	translatedMessage := cases.Upper(publisher.language).String(fullMessage)

	log.Print(translatedMessage)
}
