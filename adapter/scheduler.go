package adapter

import (
	"time"

	"github.com/suotas/chihuahua/domain"
	"github.com/suotas/chihuahua/usecase"
)

// Execute is scheduler function
func Execute(config domain.Config) {
	for true {
		usecase.SMA(config)
		time.Sleep(time.Second * time.Duration(config.INTERVAL_TIME))
	}
}