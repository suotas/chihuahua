package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/suotas/chihuahua/domain/calculator"
	"github.com/suotas/chihuahua/domain/converter"
	"github.com/suotas/chihuahua/domain/model"
	"github.com/suotas/chihuahua/domain/repository"
	"github.com/suotas/chihuahua/infra"
)

type Config struct {
	BITBANK_API_KEY string `required:"true"`
	BITBANK_SECRET string `required:"true"`
	TRADE_PAIR string `default:"btc_jpy"`
	USE_CANDLE_TYPE string `default:"1day"`
	SHORT_OHLCV_LENGTH int `default:5`
	MIDDLE_OHLCV_LENGTH int `default:25`
	LONG_OHLCV_LENGTH int `default:75`
}

func main() {
	godotenv.Load(".env")
	var config Config
	envconfig.Process("", &config)

	var api repository.IApiClient
	api, _ = infra.NewClient(config.BITBANK_API_KEY, config.BITBANK_SECRET)
	candlesticks, _ := api.GetCandlesticks(config.TRADE_PAIR, config.USE_CANDLE_TYPE, "2020")

	var shortOhlcv, middleOhlcv, longOhlcv []*model.Candlesticks
	var shortSMA, middleSMA, longSMA []int64
	var shortESMA, middleESMA, longESMA []int64

	var conv repository.IDataConverter
	conv = new(converter.BitBankDataConverter)
	var calc repository.ICalculator
	calc = new(calculator.Calculator)
	offsets := [...] int{0,1,2,3,4}

	for _, v := range offsets {
		shortOhlcv = append(shortOhlcv, conv.GetOhlcv(candlesticks, 7, v))
		middleOhlcv = append(middleOhlcv, conv.GetOhlcv(candlesticks, 28, v))
		longOhlcv = append(longOhlcv, conv.GetOhlcv(candlesticks, 74, v))

		shortSMA = append(shortSMA, calc.GetSMA(shortOhlcv[v]))
		shortESMA = append(shortESMA, calc.GetESMA(shortOhlcv[v]))

		middleSMA = append(middleSMA, calc.GetSMA(middleOhlcv[v]))
		middleESMA = append(middleESMA, calc.GetESMA(shortOhlcv[v]))
		
		longSMA = append(longSMA, calc.GetSMA(longOhlcv[v]))
		longESMA = append(longESMA, calc.GetESMA(shortOhlcv[v]))
	}

	fmt.Printf("short SMA:\t%d\n", shortSMA)
	fmt.Printf("short ESMA:\t%d\n", shortESMA)

	fmt.Printf("middle SMA:\t%d\n", middleSMA)
	fmt.Printf("middle ESMA:\t%d\n", middleESMA)

	fmt.Printf("long SMA:\t%d\n", longSMA)
	fmt.Printf("long ESMA:\t%d\n", longESMA)

	assets, _ := api.GetAssets()

	fmt.Printf("assets type:\t%v\n", assets.Data[0].Asset)
	fmt.Printf("free amount:\t%f\n", assets.Data[0].FreeAmount)
	fmt.Printf("assets type:\t%v\n", assets.Data[1].Asset)
	fmt.Printf("free amount:\t%f\n", assets.Data[1].FreeAmount)

	// depth, _ := api.GetDepth(pairBtcJpy)

	// for _, v := range depth.Data.Asks {
	// 	price := v.Price
	// 	volume := v.Volume
	// 	fmt.Printf("depth ask price: \t%d\n", price)
	// 	fmt.Printf("depth ask volume: \t%g\n", volume)
	// }

	// for _, v := range depth.Data.Bids {
	// 	price := v.Price
	// 	volume := v.Volume
	// 	fmt.Printf("depth bid price: \t%d\n", price)
	// 	fmt.Printf("depth bid volume: \t%g\n", volume)
	// }
}
