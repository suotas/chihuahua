package model

type Candlesticks struct {
	Data		[]*Ohlcv
}

type Ohlcv struct {
	Open				int64
	High				int64
	Low					int64
	Close				int64
	Volume				int64
	Timestamp			int64
}
