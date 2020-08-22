package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Rostelecom-CERT/gobadpackets"
)

func main() {
	APIKeyPtr := flag.String("api", "", "API key Badpackets")
	URLPtr := flag.String("url", "", "URL Badpackets")
	flag.Parse()
	var err error
	var conf *gobadpackets.Client

	// init function
	conf, err = gobadpackets.New(*APIKeyPtr, *URLPtr)
	if err != nil {
		log.Fatalln(err)
	}

	// example ping function
	status := conf.Ping()
	fmt.Println(status)

	// example Query function
	data, err := conf.Query()
	fmt.Println(data.Count)

	for _, v := range data.Results {
		fmt.Println(v.PostData)
	}
}
