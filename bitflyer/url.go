package bitflyer

import "net/url"

const (
	host = "ws.lightstream.bitflyer.com"
	path = "json-rpc"
)

func NewURL() url.URL {
	return url.URL{Scheme: "wss", Host: host, Path: path}
}
