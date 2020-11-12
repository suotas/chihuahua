package usecase

import (
	"github.com/suotas/chihuahua/domain"
	"github.com/suotas/chihuahua/domain/model"
)

type IIndicatorUseCase interface {
	Execute(config domain.Config) (string)
}

type IOrderUseCase interface {
	AllBuy(config domain.Config) (model.Order)
	AllSell(config domain.Config) (model.Order)
}