package liquid

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	PairBTCJPY = "btcjpy"
	PairBTCUSD = "btcusd"

	ChannelCash        = "executions_cash_%s"
	ChannelDetailsCash = "execution_details_cash_%s"
)

type Channel struct {
	Channel     string `json:"channel"`
}

func NewChannel(channelFormat string, pair string) json.RawMessage {
	channel := &Channel{
		Channel: fmt.Sprintf(channelFormat, pair),
	}
	msg, err := json.Marshal(channel)
	if err != nil {
		log.Fatal("failed encode", err)
	}

	e := &Event{
		Event: "pusher:subscribe",
		Data: msg,
		Channel: channelFormat,
	}
	msg, err = json.Marshal(e)
	if err != nil {
		log.Fatal("failed encode", err)
	}

	return msg
}
