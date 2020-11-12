package usecase

import (
	"fmt"

	"github.com/suotas/chihuahua/domain"
	"github.com/suotas/chihuahua/domain/api"
	"github.com/suotas/chihuahua/domain/calculator"
	"github.com/suotas/chihuahua/domain/converter"
	"github.com/suotas/chihuahua/domain/model"
)

type ESMAUseCase struct {
	client api.IApiClient
}

func NewESMAUseCase (client api.IApiClient) (IIndicatorUseCase){
	usecase := ESMAUseCase{client: client}
	return &usecase
}

// Execute get ohlcv and calculate ESMA function.
func (u *ESMAUseCase) Execute(config domain.Config) (string) {
	candlesticks, _ := u.client.GetCandlesticks(config.TRADE_PAIR, config.USE_CANDLE_TYPE, "2020")

	var shortOhlcv, middleOhlcv, longOhlcv []*model.Candlesticks
	var shortESMA, middleESMA, longESMA []int64

	var conv converter.IDataConverter
	conv = new(converter.BitBankDataConverter)
	var calc calculator.ICalculator
	calc = new(calculator.Calculator)
	offsets := [...] int{0,1,2,3,4}

	for _, v := range offsets {
		shortOhlcv = append(shortOhlcv, conv.GetOhlcv(candlesticks, 7, v))
		middleOhlcv = append(middleOhlcv, conv.GetOhlcv(candlesticks, 28, v))
		longOhlcv = append(longOhlcv, conv.GetOhlcv(candlesticks, 74, v))

		shortESMA = append(shortESMA, calc.GetESMA(shortOhlcv[v]))
		middleESMA = append(middleESMA, calc.GetESMA(shortOhlcv[v]))
		longESMA = append(longESMA, calc.GetESMA(shortOhlcv[v]))
	}

	fmt.Printf("short ESMA:\t%d\n", shortESMA)
	fmt.Printf("middle ESMA:\t%d\n", middleESMA)
	fmt.Printf("long ESMA:\t%d\n", longESMA)

	return ""
}