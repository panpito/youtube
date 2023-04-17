package main

import (
	"go.uber.org/fx"
	"log"
)

func main() {
	// Traditional dependency injection by hand
	//t := Title("goodbye")
	//p := NewPublisher(&t)
	//m := NewMainService(p)
	//m.Run()

	// FX framework
	fx.New(
		fx.Provide(NewMainService),
		fx.Provide(
			fx.Annotate(
				NewPublisher,
				fx.As(new(IPublisher)),
				fx.ParamTags(`group:"titles"`),
			),
		),
		fx.Provide(
			titleComponent("hello"),
		),
		fx.Provide(
			titleComponent("bonjour"),
		),
		fx.Provide(
			titleComponent("goodbye"),
		),
		fx.Invoke(func(service *MainService) {
			service.Run()
		}),
	).Run()
}

// Refactoring for component in the titles group
func titleComponent(title string) any {
	return fx.Annotate(
		func() *Title {
			t := Title(title)
			return &t
		},
		fx.ResultTags(`group:"titles"`),
	)
}

// Main service
type MainService struct {
	publisher IPublisher
}

func NewMainService(publisher IPublisher) *MainService {
	return &MainService{publisher: publisher}
}

func (service *MainService) Run() {
	service.publisher.Publish()
	log.Print("main program")
}

// Dependency
type IPublisher interface {
	Publish()
}

type Publisher struct {
	titles []*Title
}

func NewPublisher(titles ...*Title) *Publisher {
	return &Publisher{titles: titles}
}

func (publisher *Publisher) Publish() {
	for _, title := range publisher.titles {
		log.Print("publisher: ", *title)
	}
}

// Dependency of publisher
type Title string
