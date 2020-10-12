package calculator

import "github.com/suotas/chihuahua/domain/model"

type ICalculator interface {
	GetSMA(ohlcv *model.Candlesticks) (int64)
	GetESMA(ohlcv *model.Candlesticks) (int64)
}