package main

import (
	"encoding/json"
	"log"

	"github.com/mkazutaka/gabtms"
	"github.com/mkazutaka/gabtms/bitflyer"
	"github.com/mkazutaka/gabtms/bitmex"
	"github.com/mkazutaka/gabtms/liquid"
)

func main() {
	liqch := createLiquid()
	bmch := createBitmex()
	bfch := createBitflyer()
	for {
		select {
		case v := <-bmch:
			trade := bitmex.ResponseTrade{}
			err := json.Unmarshal(v, &trade)
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range trade.Data {
				m := v.Generalize()
				log.Printf("bitmex: %+v", m)
			}
		case v := <-bfch:
			trade := bitflyer.ResponseExecutions{}
			err := json.Unmarshal(v, &trade)
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range trade.Params.Message {
				m := v.Generalize()
				log.Printf("bitflyer: %+v", m)
			}
		case v := <-liqch:
			event := &liquid.ResponseExecution{}
			err := json.Unmarshal(v, &event)
			if err != nil {
				log.Fatal(err)
			}

			e := &liquid.Execution{}
			err = gabtms.UnmarshalDataString(event.Data, &e)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("liquid: %+v", e.Generalize())
		}
	}
}

func createBitmex() chan json.RawMessage {
	u := bitmex.NewURL()
	bmc := bitmex.NewChannel([]string{bitmex.TopicTrade})
	option := func(c *gabtms.Client) {
		c.OnConnectEvent = &bitmex.ConnectEvent{}
	}

	client := gabtms.NewClient(u, bmc, option)
	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
	ch, err := client.Subscribe()
	if err != nil {
		log.Fatal(err)
	}
	return ch
}

func createBitflyer() chan json.RawMessage {
	u := bitflyer.NewURL()
	c := bitflyer.NewChannel(bitflyer.ChannelExecutionsFXBTCJPY)

	client := gabtms.NewClient(u, c)
	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
	ch, err := client.Subscribe()
	if err != nil {
		log.Fatal(err)
	}
	return ch
}

func createLiquid() chan json.RawMessage {
	u := liquid.NewURL()
	c := liquid.NewChannel(liquid.ChannelCash, liquid.PairBTCJPY)
	option := func(c *gabtms.Client) {
		c.OnConnectEvent = &liquid.Event{}
		c.OnSubscribeEvent = &liquid.Event{}
	}

	client := gabtms.NewClient(u, c, option)
	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
	ch, err := client.Subscribe()
	if err != nil {
		log.Fatal(err)
	}
	return ch
}
