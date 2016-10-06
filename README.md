CardcastGo - CardCast API for Go
===================

## Setup
1. `go get github.com/Moonlington/cardcastgo`
2. Use the lib бе\_(е─)_/бе

## Example

```go
package main

import (
	"fmt"

	"github.com/Moonlington/cardcastgo"
)

func main() {
	cc, err := cardcastgo.New("Token")
	if err != nil {
		fmt.Println("error,", err)
	}
	fmt.Println(cc.GetDeck("DeckID"))
	fmt.Println(cc.PostCall("DeckID", "When _ happens, I do _."))
	fmt.Println(cc.PostResponse("DeckID", "Test"))
}
```