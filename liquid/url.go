package liquid

import "net/url"

const (
	host = "tap.liquid.com"
	path = "app/Liquid"
	query = "protocol=7"
)

func NewURL() url.URL {
	return url.URL{Scheme: "wss", Host: host, Path: path, RawQuery: query}
}
