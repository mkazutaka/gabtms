package liquid

import (
	"encoding/json"
	"log"

	"github.com/mkazutaka/gabtms"
)

func Subscribe(channelFormat string, pair string) (chan json.RawMessage, chan error, error) {
	u := NewURL()
	c := NewChannel(channelFormat, pair)
	option := func(c *gabtms.Client) {
		c.OnConnectEvent = &Event{}
		c.OnSubscribeEvent = &Event{}
	}

	client := gabtms.NewClient(u, c, option)
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
