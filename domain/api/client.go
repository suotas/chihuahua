package api

import (
	"github.com/suotas/chihuahua/domain/model"
)

type IApiClient interface {
	GetAssets() (*model.Assets, error)
	GetDepth(pair string) (*model.Depth, error)
	GetCandlesticks(pair, candleType, yaer string) (*model.Candlesticks, error)
}
