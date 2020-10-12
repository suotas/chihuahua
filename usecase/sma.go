package usecase

import (
	"fmt"

	"github.com/suotas/chihuahua/domain"
	"github.com/suotas/chihuahua/domain/api"
	"github.com/suotas/chihuahua/domain/calculator"
	"github.com/suotas/chihuahua/domain/converter"
	"github.com/suotas/chihuahua/domain/model"
	"github.com/suotas/chihuahua/infra"
)

// SMA get ohlcv and calculate SMA function.
func SMA(config domain.Config) {
	var api api.IApiClient
	api, _ = infra.NewClient(config.BITBANK_API_KEY, config.BITBANK_SECRET)
	candlesticks, _ := api.GetCandlesticks(config.TRADE_PAIR, config.USE_CANDLE_TYPE, "2020")

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
}

// ESMA get ohlcv and calculate ESMA function.
func ESMA(config domain.Config) {
	var api api.IApiClient
	api, _ = infra.NewClient(config.BITBANK_API_KEY, config.BITBANK_SECRET)
	candlesticks, _ := api.GetCandlesticks(config.TRADE_PAIR, config.USE_CANDLE_TYPE, "2020")

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
}