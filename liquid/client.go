package liquid

import (
	"encoding/json"
	"log"

	"github.com/mkazutaka/gabtms"
)

func Subscribe(channelFormat string, pair string) (chan json.RawMessage, chan error, error) {
	u := NewURL()
	c := NewChannel(channelFormat, pair)
	option := func(c *gabtms.Client) error {
		c.OnConnectEvent = &Event{}
		c.OnSubscribeEvent = &Event{}
		return nil
	}

	client, err := gabtms.NewClient(u, c, option)
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
