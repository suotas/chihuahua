package infra

import (
	"context"
	"strconv"
	"time"

	bitbank "github.com/jjjjpppp/bitbank-go-client/v1"
	"github.com/jjjjpppp/bitbank-go-client/v1/request"
	"github.com/suotas/chihuahua/domain/model"
)

// BitBankAPI bitbank API specific client struct.
type BitBankAPI struct {
	client *bitbank.Client
}

// NewClient is constructor for BitBankAPI.
func NewClient(apiKey, secret string) (*BitBankAPI, error) {
	client, err := bitbank.NewClient(apiKey, secret, nil)
	if err != nil {
		return nil, err
	}
	return &BitBankAPI{client: client}, nil
}

// GetAssets get assets info from bitbank API.
func (b *BitBankAPI) GetAssets() (*model.Assets, error) {
	result := new(model.Assets)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	bitbankAssets, err := b.client.GetAssets(ctx)
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

// GetDepth get depth info from bitbank API.
func (b *BitBankAPI) GetDepth(pair string) (*model.Depth, error) {
	result := model.NewDepth()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	bitbankDepth, err := b.client.GetDepth(ctx, pair)
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

// GetCandlesticks get candlestick data from bitbank API.
func (b *BitBankAPI) GetCandlesticks(pair, candleType, date string) (*model.Candlesticks, error) {
	result := new(model.Candlesticks)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	bitbankCandlesticks, err := b.client.GetCandlesticks(ctx, pair, candleType, date)
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

// GetActiveOrders get active sent orders data from bitbank API.
func (b *BitBankAPI) GetActiveOrders(params request.GetActiveOrdersParams) (*model.Orders, error){
	result := new(model.Orders)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	bitbankActiveOrders, err := b.client.GetActiveOrders(ctx, params)
	if err != nil {
		return nil, err
	}

	for _, v := range bitbankActiveOrders.Data.Orders {
		orderId := v.OrderID
		pair := v.Pair
		side := v.Side
		_type := v.Type
		startAmount := v.StartAmount
		remainingAmount := v.RemainingAmount
		executedAmount := v.ExecutedAmount
		price := v.Price
		averagePrice := v.AveragePrice
		orderedAt := v.OrderedAt
		status := v.Status
		orderData := model.Order{OrderID: orderId, Pair: pair, Side: side, Type: _type, StartAmount: startAmount, RemainingAmount: remainingAmount, ExecutedAmount: executedAmount, Price: price, AveragePrice: averagePrice, OrderedAt: orderedAt, Status: status}
		result.Data = append(result.Data, &orderData)
	}

	return result, nil
}

// CreateOrder send buy or sell order to bitbank API.
func (b *BitBankAPI) CreateOrder(params request.CreateOrderParams) (*model.Order, error) {
	result := new(model.Order)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	bitbankOrder, err := b.client.CreateOrder(ctx, params)
	if err != nil {
		return nil, err
	}

	orderId := bitbankOrder.Data.OrderID
	pair := bitbankOrder.Data.Pair
	side := bitbankOrder.Data.Side
	_type := bitbankOrder.Data.Type
	startAmount := bitbankOrder.Data.StartAmount
	remainingAmount := bitbankOrder.Data.RemainingAmount
	executedAmount := bitbankOrder.Data.ExecutedAmount
	price := bitbankOrder.Data.Price
	averagePrice := bitbankOrder.Data.AveragePrice
	orderedAt := bitbankOrder.Data.OrderedAt
	status := bitbankOrder.Data.Status
	result = &model.Order{OrderID: orderId, Pair: pair, Side: side, Type: _type, StartAmount: startAmount, RemainingAmount: remainingAmount, ExecutedAmount: executedAmount, Price: price, AveragePrice: averagePrice, OrderedAt: orderedAt, Status: status}

	return result, nil
}
