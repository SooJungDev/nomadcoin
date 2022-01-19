package main

import (
	"github.com/SooJungDev/nomadcoin/explorer"
	"github.com/SooJungDev/nomadcoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
