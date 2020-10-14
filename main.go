package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/suotas/chihuahua/adapter"
	"github.com/suotas/chihuahua/domain"
	"github.com/suotas/chihuahua/infra"
	"github.com/suotas/chihuahua/usecase"
)

// main function
func main() {
	// loading envconfig
	godotenv.Load(".env")
	var config domain.Config
	envconfig.Process("", &config)
	
	// scheduler start
	client, _ := infra.NewClient(config.BITBANK_API_KEY, config.BITBANK_SECRET)
	usecase := usecase.NewSMAUseCase(client)
	scheduler := adapter.NewScheduler(usecase)
	scheduler.Execute(config)
}
