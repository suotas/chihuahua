package model

type Orders struct {
	Data []*Order
}

type Order struct {
	OrderID int				`json:"order_id,omitempty"`
	Pair string				`json:"pair,omitempty"`
	Side string				`json:"side,omitempty"`
	Type string				`json:"type,omitempty"`
	StartAmount string		`json:"start_amount,omitempty"`
	RemainingAmount string	`json:"remaining_amount,omitempty"`
	ExecutedAmount string	`json:"executed_amount,omitempty"`
	Price string			`json:"price,omitempty"`
	AveragePrice string		`json:"average_price,omitempty"`
	OrderedAt int64			`json:"ordered_at,omitempty"`
	Status string			`json:"status,omitempty"`
}