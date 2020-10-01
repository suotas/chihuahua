package main

import (
	"context"
	"fmt"
	"os"
	"time"
	bitbank "github.com/jjjjpppp/bitbank-go-client/v1"
	"github.com/joho/godotenv"
)

const (
	ohlcvIdx = 0
	candlestickOpenPriceIdx = 0
	candlestickClosePriceIdx = 3
)

func getCandlesticks(client *bitbank.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	candlesticks, _ := client.GetCandlesticks(ctx, "btc_jpy", "1day", "2020")
	ohlcv := candlesticks.Data.Candlesticks[ohlcvIdx].Ohlcv

	var temp = 0
	if (len(ohlcv) > 75) {
		calcTemp := ohlcv[len(ohlcv)-5:]
		for _, v := range calcTemp {
			closePriceInt, _ := v[candlestickClosePriceIdx].Int64()
			temp += int(closePriceInt)
		}
		temp /= 5
		fmt.Println(temp)
	} else {
		fmt.Println("Not enough data.");
	}
}

func main() {
	godotenv.Load(".env")
	client, _ := bitbank.NewClient(os.Getenv("BITBANK_API_KEY"), os.Getenv("BITBANK_SECRET"), nil)
	getCandlesticks(client)
}