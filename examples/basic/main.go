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
	bmCh, _, err := bitmex.Subscribe([]string{bitmex.TopicTrade})
	if err != nil {
		log.Fatal(err)
	}

	liqCh, liqChErr, err := liquid.Subscribe(liquid.ChannelCash, liquid.PairBTCJPY)
	if err != nil {
		log.Fatal(err)
	}

	bfCh, _, err := bitflyer.Subscribe(bitflyer.ChannelExecutionsFXBTCJPY)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case v := <-bmCh:
			trade := bitmex.ResponseTrade{}
			err := json.Unmarshal(v, &trade)
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range trade.Data {
				m := v.Generalize()
				log.Printf("bitmex: %+v", m)
			}
		case v := <-bfCh:
			trade := bitflyer.ResponseExecutions{}
			err := json.Unmarshal(v, &trade)
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range trade.Params.Message {
				m := v.Generalize()
				log.Printf("bitflyer: %+v", m)
			}
		case v := <-liqCh:
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
		case <-liqChErr:
			log.Fatal(err)
		}
	}
}
