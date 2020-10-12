package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/suotas/chihuahua/adapter"
	"github.com/suotas/chihuahua/domain"
)

// main function
func main() {
	// loading envconfig
	godotenv.Load(".env")
	var config domain.Config
	envconfig.Process("", &config)
	
	// scheduler start
	adapter.Execute(config)
}
