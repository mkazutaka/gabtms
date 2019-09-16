package gabtms

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewClient(t *testing.T) {
	u := url.URL{Scheme: "wss", Host: "www.hoge.com", Path: "test"}
	value := map[string]string{
		"key": "value",
	}
	msg, err := json.Marshal(value)
	if err != nil {
		t.Error(err)
	}

	option := func(c *Client) error {
		value := map[string]string{
			"subscribeKey": "subscribeValue",
		}
		c.OnSubscribeEvent = value
		return nil
	}

	client, err := NewClient(u, msg, option)
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(client.URL, u); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	if diff := cmp.Diff(client.Channel, msg); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	if client.OnConnectEvent != nil {
		t.Errorf("OnConnect Event should be null")
	}
	if client.OnSubscribeEvent == nil {
		t.Errorf("OnSubscribeEvent Event should not be null")
	}
}
