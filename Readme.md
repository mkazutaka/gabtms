# gabtms
gather BTC Websocket.
This is library for aggregating trade data from some exchanges through web socket

![demo](https://user-images.githubusercontent.com/4601360/64967843-282db680-d8dc-11e9-8454-301f19e54f8c.gif)

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

	"github.com/mkazutaka/gabtms/bitmex"
)

func main() {
	ch, _, err := bitmex.Subscribe([]string{bitmex.TopicTrade})
	if err != nil {
		log.Fatal(err)
	}

	for v := range ch {
		trade := bitmex.ResponseTrade{}
		err := json.Unmarshal(v, &trade)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range trade.Data {
			m := v.Generalize()
			log.Printf("bitmex: %+v", m)
		}
	}
}
```

# License
This project is licensed under the MIT License - see the LICENSE.md file for details
