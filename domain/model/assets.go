package model

type Assets struct {
	Data			[]*Asset
}

type Asset struct {
	Asset			string
	FreeAmount		float64
}
