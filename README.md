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
	data, err := conf.Query()
	fmt.Println(data.Count)

	for _, v:=range data.Results {
		fmt.Println(v.PostData)
	}
}
```