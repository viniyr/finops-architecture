package application

import (
	"net/http"

	"aevyn/finops-arch/internal/infrastructure/rabbitmq"
	"aevyn/finops-arch/internal/interfaces/http/handlers"
	"aevyn/finops-arch/internal/interfaces/http/routes"
	"aevyn/finops-arch/internal/usecases"
)

type App struct {
	Router http.Handler
}

func New() *App {
	rabbitURL := "amqp://guest:guest@localhost:5672/"
	exchange := "x.disputes"

	conn, ch, err := rabbitmq.Dial(rabbitURL)
	if err != nil {
		panic(err)
	}

	_ = conn

	if err := rabbitmq.DeclareTopicExchange(ch, exchange); err != nil {
		panic(err)
	}

	pub := rabbitmq.Publisher{Ch: ch, Exchange: exchange}

	openUC := &usecases.OpenDisputeUseCase{Pub: pub}
	h := handlers.NewDisputeHandler(openUC)
	r := routes.SetupRoutes(h)

	return &App{
		Router: r,
	}
}
