package gabtms

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type Client struct {
	URL     url.URL
	Channel json.RawMessage

	OnConnectEvent   interface{}
	OnSubscribeEvent interface{}

	conn *websocket.Conn
}

type Option func(wsc *Client)

func NewClient(url url.URL, channel json.RawMessage, options ...Option) *Client {
	client := &Client{
		URL:              url,
		Channel:          channel,
		OnConnectEvent:   nil,
		OnSubscribeEvent: nil,
	}
	for _, option := range options {
		option(client)
	}
	return client
}

func (c *Client) Connect() error {
	var err error
	c.conn, _, err = websocket.DefaultDialer.Dial(c.URL.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	if c.OnConnectEvent != nil {
		err = c.conn.ReadJSON(c.OnConnectEvent)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) Subscribe() (chan json.RawMessage, error) {
	err := c.conn.WriteMessage(websocket.TextMessage, c.Channel)
	if err != nil {
		log.Fatal("failed write message:", err)
	}

	if c.OnSubscribeEvent != nil {
		err = c.conn.ReadJSON(c.OnConnectEvent)
		if err != nil {
			return nil, err
		}
	}

	ch := make(chan json.RawMessage)
	go func() {
		defer close(ch)
		for {
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			ch <- message
		}
	}()

	return ch, nil
}
