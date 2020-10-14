package converter

import "github.com/suotas/chihuahua/domain/model"

type BitBankDataConverter struct {}

func (b BitBankDataConverter) GetOhlcv(candlesticks *model.Candlesticks, days, offset int) (*model.Candlesticks) {
	baseOhlcv := candlesticks.Data
	startIdxMinus := days + offset
	startIdx := len(baseOhlcv) - startIdxMinus -1
	endIdx := len(baseOhlcv) - offset -1
	result := model.Candlesticks{Data: baseOhlcv[startIdx:endIdx]}
	return &result
}
