package bitmex

import (
	"strings"
	"time"

	"github.com/mkazutaka/btcwsc"
)

type ResponseTrade struct {
	Table       string            `json:"table"`
	Keys        []string          `json:"keys"`
	Types       map[string]string `json:"types"`
	ForeignKeys map[string]string `json:"foreignKeys"`
	Attributes  map[string]string `json:"attributes"`
	Action      string            `json:"action"`
	Data        []TradeData       `json:"data"`
}

type TradeData struct {
	ForeignNotional float64 `json:"foreignNotional"`
	GrossValue      float64 `json:"grossValue"`
	HomeNational    float64 `json:"homeNotional"`
	Price           float64 `json:"price"`
	Side            string  `json:"side"`
	Size            float64 `json:"size"`
	Symbol          string  `json:"symbol"`
	TickDirection   string  `json:"tickDirection"`
	Timestamp       string  `json:"timestamp"`
	TryMatchID      string  `json:"trdMatchID"`
}

func (e TradeData) Generalize() *btcwsc.Execution {
	t, _ := time.Parse(time.RFC3339, e.Timestamp)
	return &btcwsc.Execution{
		Side:  strings.ToLower(e.Side),
		Size:  e.Size,
		Price: e.Price,
		Tick:  t.Unix(),
	}
}
