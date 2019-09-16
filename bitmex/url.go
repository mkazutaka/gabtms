package bitmex

import "net/url"

const (
	host = "www.bitmex.com"
	path = "realtime"
)

func NewURL() url.URL {
	return url.URL{Scheme: "wss", Host: host, Path: path}
}
