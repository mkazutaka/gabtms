package bitflyer

import (
	"strings"
	"time"

	"github.com/mkazutaka/btcwsc"
)

type ResponseExecutions struct {
	Version string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  ExecutionsParam `json:"params"`
	Result  interface{}     `json:"result"`
	Id      *int            `json:"id"`
}

type ExecutionsParam struct {
	Channel string      `json:"channel"`
	Message []Execution `json:"message"`
}

type Execution struct {
	ID                         int     `json:"id"`
	Side                       string  `json:"side"`
	Size                       float64 `json:"size"`
	Price                      float64 `json:"price"`
	ExecDate                   string  `json:"exec_date"`
	BuyChildOrderAcceptanceId  string  `json:"buy_child_order_acceptance_id"`
	SellChildOrderAcceptanceId string  `json:"sell_child_order_acceptance_id"`
}

func (e Execution) Generalize() *btcwsc.Execution {
	t, _ := time.Parse(time.RFC3339, e.ExecDate)
	return &btcwsc.Execution{
		ID:    e.ID,
		Side:  strings.ToLower(e.Side),
		Size:  e.Size,
		Price: e.Price,
		Tick:  t.Unix(),
	}
}
