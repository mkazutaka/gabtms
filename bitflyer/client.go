package bitflyer

import (
	"encoding/json"
	"log"

	"github.com/mkazutaka/gabtms"
)

func Subscribe(channelName string) (chan json.RawMessage, chan error, error) {
	u := NewURL()
	c := NewChannel(channelName)

	client := gabtms.NewClient(u, c)
	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
	ch, chErr, err := client.Subscribe()
	if err != nil {
		log.Fatal(err)
	}

	return ch, chErr, err
}
