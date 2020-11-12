package api

import (
	"github.com/suotas/chihuahua/domain/model"
	"github.com/jjjjpppp/bitbank-go-client/v1/request"
)

type IApiClient interface {
	GetAssets() (*model.Assets, error)
	GetDepth(pair string) (*model.Depth, error)
	GetCandlesticks(pair, candleType, yaer string) (*model.Candlesticks, error)
	GetActiveOrders(params request.GetActiveOrdersParams) (*model.Orders, error)
	CreateOrder(params request.CreateOrderParams) (*model.Order, error)
}
