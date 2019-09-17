package gabtms

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	URL     url.URL
	Channel []byte

	OnConnectEvent   interface{}
	OnSubscribeEvent interface{}

	PingDeadline time.Duration
	PingPeriod   time.Duration

	conn *websocket.Conn
}

type Option func(wsc *Client) error

func NewClient(url url.URL, channel []byte, options ...Option) (*Client, error) {
	client := &Client{
		URL:               url,
		Channel:           channel,
		OnConnectEvent:    nil,
		OnSubscribeEvent:  nil,
		PingPeriod:        5 * time.Second,
		PingDeadline:      10 * time.Second,
	}
	for _, option := range options {
		err := option(client)
		if err != nil {
			return nil, err
		}
	}
	return client, nil
}

func (c *Client) Connect() error {
	var err error
	c.conn, _, err = websocket.DefaultDialer.Dial(c.URL.String(), nil)
	if err != nil {
		return err
	}

	if c.OnConnectEvent != nil {
		err = c.conn.ReadJSON(c.OnConnectEvent)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) Subscribe() (chan json.RawMessage, chan error, error) {
	err := c.conn.WriteMessage(websocket.TextMessage, c.Channel)
	if err != nil {
		return nil, nil, err
	}

	if c.OnSubscribeEvent != nil {
		err = c.conn.ReadJSON(c.OnConnectEvent)
		if err != nil {
			return nil, nil, err
		}
	}

	chErr := make(chan error)
	go c.KeepAlive(chErr)

	ch := make(chan json.RawMessage)
	go func() {
		defer close(ch)
		for {
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				chErr <- err
				return
			}
			ch <- message
		}
	}()

	return ch, chErr, nil
}

func (c *Client) KeepAlive(errChan chan error) {
	ticker := time.NewTicker(c.PingPeriod)
	defer ticker.Stop()

	for range ticker.C {
		if err := c.Ping(); err != nil {
			errChan <- err
		}
	}
}

func (c *Client) Ping() error {
	err := c.conn.SetWriteDeadline(time.Now().Add(c.PingDeadline))
	if err != nil {
		return err
	}
	return c.conn.WriteMessage(websocket.PingMessage, []byte{})
}

func (c *Client) Disconnect() error {
	return c.conn.Close()
}
