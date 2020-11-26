package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jjjjpppp/bitbank-go-client/v1"
	"github.com/jjjjpppp/bitbank-go-client/v1/request"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/suotas/chihuahua/domain"
)

func main() {
	// loading envconfig
	godotenv.Load("../.env")
	var config domain.Config
	envconfig.Process("", &config)
			
	client, err := bitbank.NewClient(config.BITBANK_API_KEY, config.BITBANK_SECRET, nil)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	params := request.GetActiveOrdersParams{}
	params.Pair = "btc_jpy"
	result, err := client.GetActiveOrders(ctx, params)
	if err != nil {
		panic(-1)
	}

	fmt.Println(result.Data.Orders)
}