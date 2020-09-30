package main

import (
	"context"
	"fmt"
	"os"
	"time"
	bitbank "github.com/jjjjpppp/bitbank-go-client/v1"
	"github.com/joho/godotenv"
)

func getCandlesticks(client *bitbank.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	candlesticks, _ := client.GetCandlesticks(ctx, "btc_jpy", "1day", "2020")
	candlesticksData := candlesticks.Data.Candlesticks
	for i, v := range candlesticksData {
		fmt.Println(i, v)
	}
}

func main() {
	godotenv.Load(".env")
	client, _ := bitbank.NewClient(os.Getenv("BITBANK_API_KEY"), os.Getenv("BITBANK_SECRET"), nil)
	getTicker(client)
	time.Sleep(time.Second * 1)
}