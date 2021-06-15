[![GoDoc](https://godoc.org/github.com/Rostelecom-CERT/gobadpackets?status.svg)](http://godoc.org/github.com/Rostelecom-CERT/gobadpackets)
[![Go Report Card](https://goreportcard.com/badge/github.com/Rostelecom-CERT/gobadpackets)](https://goreportcard.com/report/github.com/Rostelecom-CERT/gobadpackets)

BadPackets REST API client library 
---------------------
BadPackets is TI IoT service provider with data about botnets and other threats. 

Link to BadPackets: 
* [official site](https://badpackets.net/)
* [twitter](https://twitter.com/bad_packets)

Usage example
------------------------------------------------

```sh
go get -u github.com/Rostelecom-CERT/gobadpackets
```

```sh
go test -api APIKEY -url URL
```

Simple example using library in cmd/gobadpackets/main.go

```go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Rostelecom-CERT/gobadpackets"
)

func main () {
	APIKeyPtr := flag.String("api", "", "API key Badpackets")
	URLPtr := flag.String("url", "", "URL Badpackets")
	flag.Parse()
	var err error
	var conf *gobadpackets.Client

	// init function
	conf,err = gobadpackets.New(*APIKeyPtr,*URLPtr)
	if err != nil {
		log.Fatalln(err)
	}

	// example ping function
	status := conf.Ping()
	fmt.Println(status)

	// example Query function
	data, err := conf.Query(&gobadpackets.Request{Country: "RU", Tags: "Mirai"})
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println(data.Count)

	// print all tags description
	for _, v := range data.Results {
		for _, tv := range v.Tags {
			fmt.Println(tv.Description)
		}
	}

	// Format data from string to Time type
	timeTest,err := time.Parse(time.RFC3339,"2018-12-31T09:04:22Z")
	if err!= nil {
		log.Fatal(err)
	}

	// Request data with time parameter
	data, err = conf.Query(&gobadpackets.Request{LastSeenBefore: timeTest})
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println(data.Count)
}
```