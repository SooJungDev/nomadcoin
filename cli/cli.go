package cli

import (
	"flag"
	"fmt"
	"github.com/SooJungDev/nomadcoin/explorer"
	"github.com/SooJungDev/nomadcoin/rest"
	"os"
)

func usage() {
	fmt.Printf("Welcome to 노마드코인\n")
	fmt.Printf("Please use the following flag:\n\n")
	fmt.Printf("-port=4000:  Set the PORT of the server\n")
	fmt.Printf("-mode=rest:  Start the REST API(recommended) \n")
	os.Exit(0)
}

func Start() {

	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()
	switch *mode {
	case "rest":
		// start rest api
		rest.Start(*port)
	case "html":
		// start html explorer
		explorer.Start(*port)
	default:
		usage()
	}
	fmt.Println(*port, *mode)
}
