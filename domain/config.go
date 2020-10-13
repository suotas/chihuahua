package domain

type Config struct {
	BITBANK_API_KEY string `required:"true"`
	BITBANK_SECRET string `required:"true"`
	TRADE_PAIR string `default:"btc_jpy"`
	USE_CANDLE_TYPE string `default:"1day"`
	INTERVAL_TIME int `default:"86400"`
	SHORT_OHLCV_LENGTH int `default:"5"`
	MIDDLE_OHLCV_LENGTH int `default:"25"`
	LONG_OHLCV_LENGTH int `default:"75"`
}