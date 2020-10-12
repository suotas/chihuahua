package infra

import (
	"context"
	"strconv"
	"time"

	bitbank "github.com/jjjjpppp/bitbank-go-client/v1"
	"github.com/suotas/chihuahua/domain/model"
)

type BitBankApi struct {
	client *bitbank.Client
	ctx *context.Context
}

func NewClient(apiKey, secret string) (*BitBankApi, error) {
	client, err := bitbank.NewClient(apiKey, secret, nil)
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	return &BitBankApi{client: client, ctx: &ctx}, nil
}

func (b *BitBankApi) GetAssets() (*model.Assets, error) {
	result := new(model.Assets)
	bitbankAssets, err := b.client.GetAssets(*b.ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range bitbankAssets.Data.Assets {
		asset := v.Asset
		freeAmount, _ := strconv.ParseFloat(v.FreeAmount, 64)
		assetData := model.Asset{Asset: asset, FreeAmount: freeAmount}
		result.Data = append(result.Data, &assetData)
	}
	return result, nil
}

func (b *BitBankApi) GetDepth(pair string) (*model.Depth, error) {
	result := model.NewDepth()
	bitbankDepth, err := b.client.GetDepth(*b.ctx, pair)
	if err != nil {
		return nil, err
	}
	for _, v := range bitbankDepth.Data.Asks {
		price, _ := v[0].Int64()
		volume, _ := v[1].Float64()
		depthPair := model.DepthPair{Price: price, Volume: volume}
		result.Data.Asks = append(result.Data.Asks, depthPair)
	}
	for _, v := range bitbankDepth.Data.Bids {
		price, _ := v[0].Int64()
		volume, _ := v[1].Float64()
		depthPair := model.DepthPair{Price: price, Volume: volume}
		result.Data.Bids = append(result.Data.Bids, depthPair)
	}
	return result, nil
}

func (b *BitBankApi) GetCandlesticks(pair, candleType, date string) (*model.Candlesticks, error) {
	result := new(model.Candlesticks)
	bitbankCandlesticks, err := b.client.GetCandlesticks(*b.ctx, pair, candleType, date)
	if err != nil {
		return nil, err
	}
	for _, v := range bitbankCandlesticks.Data.Candlesticks[0].Ohlcv {
		open, _ := v[0].Int64()
		high, _ := v[1].Int64()
		low, _ := v[2].Int64()
		close, _ := v[3].Int64()
		volume, _ := v[4].Int64()
		timestamp, _ := v[5].Int64()
		assetData := model.Ohlcv{Open: open, High: high, Low: low, Close: close, Volume: volume, Timestamp: timestamp}
		result.Data = append(result.Data, &assetData)
	}
	return result, nil
}
