package main

import (
	"fmt"

	"github.com/Moonlington/cardcastgo"
)

func main() {
	cc, err := cardcastgo.New("x-auth-token")
	if err != nil {
		fmt.Println("error,", err)
	}
	fmt.Println(cc.Calls("Deck ID"))
}
