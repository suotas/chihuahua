package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
	bitbank "github.com/jjjjpppp/bitbank-go-client/v1"
	"github.com/jjjjpppp/bitbank-go-client/v1/models"
	"github.com/joho/godotenv"
)

const (
	ohlcvIdx = 0
	candlestickOpenPriceIdx = 0
	candlestickClosePriceIdx = 3

	pairBtcJpy = "btc_jpy"
	candleType1Day = "1day"
)

func getCandlesticks() (*models.Candlesticks, error) {
	client, _ := bitbank.NewClient(os.Getenv("BITBANK_API_KEY"), os.Getenv("BITBANK_SECRET"), nil)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	return client.GetCandlesticks(ctx, pairBtcJpy, candleType1Day, "2020")
}

func getOhlcv(candlesticks *models.Candlesticks, days, offset int) ([][]json.Number) {
	baseOhlcv := candlesticks.Data.Candlesticks[ohlcvIdx].Ohlcv
	startIdxMinus := days + offset
	startIdx := len(baseOhlcv) - startIdxMinus -1
	endIdx := len(baseOhlcv) - offset -1
	
	return baseOhlcv[startIdx:endIdx]
}

func getMovingAverage(ohlcv [][]json.Number) (int) {
	temp := 0
	for _, v := range ohlcv {
		closePriceInt, _ := v[candlestickClosePriceIdx].Int64()
		temp += int(closePriceInt)
	}
	return temp / len(ohlcv)
}

func main() {
	godotenv.Load(".env")
	candlesticks, _ := getCandlesticks()

	var shortOhlcv, middleOhlcv, longOhlcv [][][]json.Number
	var shortMa, middleMa, longMa []int

	offsets := [...] int{0,1,2,3,4}

	for _, v := range offsets {
		shortOhlcv = append(shortOhlcv, getOhlcv(candlesticks, 7, v))
		middleOhlcv = append(middleOhlcv, getOhlcv(candlesticks, 28, v))
		longOhlcv = append(longOhlcv, getOhlcv(candlesticks, 74, v))

		shortMa = append(shortMa, getMovingAverage(shortOhlcv[v]))
		middleMa = append(middleMa, getMovingAverage(middleOhlcv[v]))
		longMa = append(longMa, getMovingAverage(longOhlcv[v]))
	}

	fmt.Printf("short moving average:\t%d\n", shortMa)
	fmt.Printf("middle moving average:\t%d\n", middleMa)
	fmt.Printf("long moving average:\t%d\n", longMa)
}