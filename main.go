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

func getDepth() (*models.Depth, error) {
	client, _ := bitbank.NewClient(os.Getenv("BITBANK_API_KEY"), os.Getenv("BITBANK_SECRET"), nil)
	ctx, err := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		// fatal error
	}

	return client.GetDepth(ctx, pairBtcJpy)
}

func getCandlesticks() (*models.Candlesticks, error) {
	client, _ := bitbank.NewClient(os.Getenv("BITBANK_API_KEY"), os.Getenv("BITBANK_SECRET"), nil)
	ctx, err := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		// fatal error
	}

	return client.GetCandlesticks(ctx, pairBtcJpy, candleType1Day, "2020")
}

func getOhlcv(candlesticks *models.Candlesticks, days, offset int) ([][]json.Number) {
	baseOhlcv := candlesticks.Data.Candlesticks[ohlcvIdx].Ohlcv
	startIdxMinus := days + offset
	startIdx := len(baseOhlcv) - startIdxMinus -1
	endIdx := len(baseOhlcv) - offset -1
	
	return baseOhlcv[startIdx:endIdx]
}

func getSMA(ohlcv [][]json.Number) (int) {
	temp := 0
	for _, v := range ohlcv {
		closePriceInt, _ := v[candlestickClosePriceIdx].Int64()
		temp += int(closePriceInt)
	}

	return temp / len(ohlcv)
}

func getESMA(ohlcv [][]json.Number) (int) {
	temp := 0
	divisor := 0
	for i, v := range ohlcv {
		closePriceInt, _ := v[candlestickClosePriceIdx].Int64()
		temp += int(closePriceInt) * (i + 1)
		divisor += (i + 1)
	}

	return temp / divisor
}

func main() {
	godotenv.Load(".env")
	candlesticks, _ := getCandlesticks()

	var shortOhlcv, middleOhlcv, longOhlcv [][][]json.Number
	var shortSMA, middleSMA, longSMA []int
	var shortESMA, middleESMA, longESMA []int

	offsets := [...] int{0,1,2,3,4}

	for _, v := range offsets {
		shortOhlcv = append(shortOhlcv, getOhlcv(candlesticks, 7, v))
		middleOhlcv = append(middleOhlcv, getOhlcv(candlesticks, 28, v))
		longOhlcv = append(longOhlcv, getOhlcv(candlesticks, 74, v))

		shortSMA = append(shortSMA, getSMA(shortOhlcv[v]))
		shortESMA = append(shortESMA, getESMA(shortOhlcv[v]))

		middleSMA = append(middleSMA, getSMA(middleOhlcv[v]))
		middleESMA = append(middleESMA, getESMA(shortOhlcv[v]))
		
		longSMA = append(longSMA, getSMA(longOhlcv[v]))
		longESMA = append(longESMA, getESMA(shortOhlcv[v]))
	}

	fmt.Printf("short SMA:\t%d\n", shortSMA)
	fmt.Printf("short ESMA:\t%d\n", shortESMA)

	fmt.Printf("middle SMA:\t%d\n", middleSMA)
	fmt.Printf("middle ESMA:\t%d\n", middleESMA)

	fmt.Printf("long SMA:\t%d\n", longSMA)
	fmt.Printf("long ESMA:\t%d\n", longESMA)

	depth, _ := getDepth()

	for _, v := range depth.Data.Asks {
		price, _ := v[0].Int64()
		volume, _ := v[1].Float64()
		fmt.Printf("depth ask price: \t%d\n", price)
		fmt.Printf("depth ask volume: \t%g\n", volume)
	}

	for _, v := range depth.Data.Bids {
		price, _ := v[0].Int64()
		volume, _ := v[1].Float64()
		fmt.Printf("depth bid price: \t%d\n", price)
		fmt.Printf("depth bid volume: \t%g\n", volume)
	}
}
