package bitmex

import (
	"encoding/json"
	"log"
)

// see: https://www.bitmex.com/app/wsAPI#Subscriptions
const (
	TopicAnnouncement        = "announcement"
	TopicChat                = "chat"
	TopicConnected           = "connected"
	TopicFunding             = "funding"
	TopicInstrument          = "instrument"
	TopicInsurance           = "insurance"
	TopicLiquidation         = "liquidation"
	TopicOrderBookL225       = "orderBookL2_25"
	TopicOrderBookL2         = "orderBookL2"
	TopicOrderBook10         = "orderBook10"
	TopicPublicNotifications = "publicNotifications"
	TopicQuote               = "quote"
	TopicQuoteBin1m          = "quoteBin1m"
	TopicQuoteBin5m          = "quoteBin5m"
	TopicQuoteBin1h          = "quoteBin1h"
	TopicQuoteBin1d          = "quoteBin1d"
	TopicSettlement          = "settlement"
	TopicTrade               = "trade"
	TopicTradeBin1m          = "tradeBin1m"
	TopicTradeBin5m          = "tradeBin5m"
	TopicTradeBin1h          = "tradeBin1h"
	TopicTradeBin1d          = "tradeBin1d"
)

// subscribe=instrument,orderBook:XBTUSD
type Channel struct {
	Op   string   `json:"op"`
	Args []string `json:"args"`
}

func NewChannel(topics []string) json.RawMessage {
	channel := &Channel{
		Op:   "subscribe",
		Args: topics,
	}
	msg, err := json.Marshal(channel)
	if err != nil {
		log.Fatal("failed encode", err)
	}
	return msg
}
