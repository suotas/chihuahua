package usecase

import (
	"fmt"

	"github.com/suotas/chihuahua/domain"
	"github.com/suotas/chihuahua/domain/api"
	"github.com/suotas/chihuahua/domain/calculator"
	"github.com/suotas/chihuahua/domain/converter"
	"github.com/suotas/chihuahua/domain/model"
)

type SMAUseCase struct {
	client api.IApiClient
}

func NewSMAUseCase (client api.IApiClient) (IIndicatorUseCase){
	usecase := SMAUseCase{client: client}
	return &usecase
}

// Execute get ohlcv and calculate SMA function.
func (u *SMAUseCase) Execute(config domain.Config) (string) {
	candlesticks, _ := u.client.GetCandlesticks(config.TRADE_PAIR, config.USE_CANDLE_TYPE, "2020")

	var shortOhlcv, middleOhlcv, longOhlcv []*model.Candlesticks
	var shortSMA, middleSMA, longSMA []int64

	var conv converter.IDataConverter
	conv = new(converter.BitBankDataConverter)
	var calc calculator.ICalculator
	calc = new(calculator.Calculator)
	offsets := [...] int{0,1,2,3,4}

	for _, v := range offsets {
		shortOhlcv = append(shortOhlcv, conv.GetOhlcv(candlesticks, 7, v))
		middleOhlcv = append(middleOhlcv, conv.GetOhlcv(candlesticks, 28, v))
		longOhlcv = append(longOhlcv, conv.GetOhlcv(candlesticks, 74, v))

		shortSMA = append(shortSMA, calc.GetSMA(shortOhlcv[v]))
		middleSMA = append(middleSMA, calc.GetSMA(middleOhlcv[v]))
		longSMA = append(longSMA, calc.GetSMA(longOhlcv[v]))
	}

	fmt.Printf("short SMA:\t%d\n", shortSMA)
	fmt.Printf("middle SMA:\t%d\n", middleSMA)
	fmt.Printf("long SMA:\t%d\n", longSMA)

	return ""
}
