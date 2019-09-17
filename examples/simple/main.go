package main

import (
	"encoding/json"
	"log"

	"github.com/mkazutaka/gabtms/bitmex"
)

func main() {
	ch, _, err := bitmex.Subscribe([]string{bitmex.TopicTrade})
	if err != nil {
		log.Fatal(err)
	}

	for v := range ch {
		trade := bitmex.ResponseTrade{}
		err := json.Unmarshal(v, &trade)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range trade.Data {
			m := v.Generalize()
			log.Printf("bitmex: %+v", m)
		}
	}
}
