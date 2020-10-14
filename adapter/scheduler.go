package adapter

import (
	"time"

	"github.com/suotas/chihuahua/domain"
	"github.com/suotas/chihuahua/usecase"
)

type IScheduler interface {
	Execute(config domain.Config)
}

type scheduler struct {
	usecase usecase.IIndicatorUseCase
}

func NewScheduler(usecase usecase.IIndicatorUseCase) (IScheduler) {
	scheduler := scheduler{usecase: usecase}
	return &scheduler
}

// Execute is scheduler function
func (s *scheduler) Execute(config domain.Config) {
	for true {
		s.usecase.Execute(config)
		time.Sleep(time.Second * time.Duration(config.INTERVAL_TIME))
	}
}