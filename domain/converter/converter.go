package converter

import "github.com/suotas/chihuahua/domain/model"

type IDataConverter interface {
	GetOhlcv(candlesticks *model.Candlesticks, days, offset int) (*model.Candlesticks)
}