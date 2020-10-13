package usecase

import "github.com/suotas/chihuahua/domain"

type IIndicatorUseCase interface {
	Execute(config domain.Config)
}