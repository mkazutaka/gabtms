package liquid

import (
	"encoding/json"
	"strconv"

	"github.com/mkazutaka/gabtms"
)

// execution: https://developers.liquid.com/#get-executions
type ResponseDetailExecution struct {
	ID          uint64  `json:"id"`
	Price       float64 `json:"price"`
	Quantity    string  `json:"quantity"`
	SellOrderID uint64  `json:"sell_order_id"`
	BuyOrderID  uint64  `json:"buy_order_id"`
	CreatedAt   int64   `json:"created_at"`
	TakerSide   string  `json:"taker_side"`
}

func (e *ResponseDetailExecution) Generalize() *gabtms.Execution {
	size, _ := strconv.ParseFloat(e.Quantity, 64)
	return &gabtms.Execution{
		Side:  e.TakerSide,
		Size:  size,
		Price: e.Price,
		Tick:  e.CreatedAt,
	}
}

// {"created_at":1568464769,"id":193293288,"price":1111921.22557,"quantity":0.06,"taker_side":"buy"}

type ResponseExecution struct {
	Channel string          `json:"channel"`
	Data    json.RawMessage `json:"data"`
	Event   string          `json:"event"`
}

type Execution struct {
	ID        uint64  `json:"id"`
	Price     float64 `json:"price"`
	Quantity  float64 `json:"quantity"`
	CreatedAt int64   `json:"created_at"`
	TakerSide string  `json:"taker_side"`
}

func (e *Execution) Generalize() *gabtms.Execution {
	return &gabtms.Execution{
		Side:  e.TakerSide,
		Size:  e.Quantity,
		Price: e.Price,
		Tick:  e.CreatedAt,
	}
}
