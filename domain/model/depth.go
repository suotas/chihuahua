package model

type Depth struct {
	Data *DepthData
}

type DepthData struct {
	Asks []DepthPair
	Bids []DepthPair
}

type DepthPair struct {
	Price int64
	Volume float64
}

func NewDepth() (*Depth) {
	depth := new(Depth)
	var asks []DepthPair
	var bids []DepthPair
	depthData := DepthData{Asks: asks, Bids: bids}
	depth.Data = &depthData

	return depth
}