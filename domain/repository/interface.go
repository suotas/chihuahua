package repository

import (
	"github.com/suotas/chihuahua/domain/model"
)

type IApiClient interface {
	GetAssets() (*model.Assets, error)
	GetDepth(pair string) (*model.Depth, error)
	GetCandlesticks(pair, candleType, yaer string) (*model.Candlesticks, error)
}

type IDataConverter interface {
	GetOhlcv(candlesticks *model.Candlesticks, days, offset int) (*model.Candlesticks)
}
type ICalculator interface {
	GetSMA(ohlcv *model.Candlesticks) (int64)
	GetESMA(ohlcv *model.Candlesticks) (int64)
}