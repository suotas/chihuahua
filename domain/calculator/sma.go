package calculator

import (
	"github.com/suotas/chihuahua/domain/model"
)

type Calculator struct {}


func (c *Calculator) GetSMA(ohlcv *model.Candlesticks) (int64) {
	temp := int64(0)
	for _, v := range ohlcv.Data {
		closePrice := v.Close
		temp += closePrice
	}

	return temp / int64(len(ohlcv.Data))
}

func (c *Calculator) GetESMA(ohlcv *model.Candlesticks) (int64) {
	temp := int64(0)
	divisor := int64(0)
	for i, v := range ohlcv.Data {
		closePrice := v.Close
		temp += closePrice * int64(i + 1)
		divisor += int64(i + 1)
	}

	return temp / divisor
}
