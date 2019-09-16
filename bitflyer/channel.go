package bitflyer

import (
	"encoding/json"
	"log"

	"github.com/gorilla/rpc/v2/json2"
)

const (
	ChannelExecutionsBTCJPY   = "lightning_executions_BTC_JPY"
	ChannelExecutionsFXBTCJPY = "lightning_executions_FX_BTC_JPY"
	ChannelExecutionsETHJPY   = "lightning_executions_ETH_JPY"
)

type Channel struct {
	Channel string `json:"channel"`
}

func NewChannel(name string) json.RawMessage {
	channel := &Channel{
		Channel: name,
	}
	msg, err := json2.EncodeClientRequest("subscribe", channel)
	if err != nil {
		log.Fatal("failed encode", err)
	}
	return msg
}
