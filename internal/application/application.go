package application

import "blog-go-api/internal/config"

type Application struct {
	Cfg *config.Config
}

func Get() (*Application, error) {
	cfg := config.Get()

	return &Application{
		Cfg: cfg,
	}, nil

}
