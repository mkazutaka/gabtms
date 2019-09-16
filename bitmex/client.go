package bitmex

import (
	"encoding/json"
	"log"

	"github.com/mkazutaka/gabtms"
)

func Subscribe(topics []string) (chan json.RawMessage, chan error, error) {
	u := NewURL()
	bmc := NewChannel(topics)

	option := func(c *gabtms.Client) error {
		c.OnConnectEvent = &ConnectEvent{}
		return nil
	}

	client, err := gabtms.NewClient(u, bmc, option)
	if err != nil {
		return nil, nil, err
	}

	err = client.Connect()
	if err != nil {
		log.Fatal(err)
	}

	ch, chErr, err := client.Subscribe()
	if err != nil {
		log.Fatal(err)
	}

	return ch, chErr, err
}
