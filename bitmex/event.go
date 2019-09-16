package bitmex

import "encoding/json"

type ConnectEvent struct {
	Info      string          `json:"info"`
	Version   string          `json:"version"`
	Timestamp string          `json:"timestamp"`
	Docs      string          `json:"docs"`
	Limit     json.RawMessage `json:"limit"`
}
