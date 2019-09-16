# gabtms
gather BTC Websocket.
This is library for aggregating trade data from some exchanges through web socket

# Install
```
$ go get github.com/mkazutaka/gabtms
```

# How to use
show simple example

```go
package main

import (
	"encoding/json"
	"log"

	"github.com/mkazutaka/gabtms"
	"github.com/mkazutaka/gabtms/bitmex"
)

func main() {
	u := bitmex.NewURL()
	bmc := bitmex.NewChannel([]string{gabtms.BitmexTopicTrade})

	client := gabtms.NewClient(u, bmc)
	_ = client.Connect()
	ch, _ := client.Subscribe()
	for {
		select {
		case v := <- ch:
			trade := bitmex.ResponseTrade{}
			err := json.Unmarshal(v, &trade)
			if err != nil {
				panic(err)
			}
			for _, v := range trade.Data {
				m := v.Generalize()
				log.Printf("bitmex: %+v", m)
			}
		}
	}
}
```

# License
This project is licensed under the MIT License - see the LICENSE.md file for details
