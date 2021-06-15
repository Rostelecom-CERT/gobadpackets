package main

import (
	"flag"
	"fmt"
	"log"
	"time"

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

	// Format data from string to Time type
	timeStart,err := time.Parse(time.RFC3339,"2021-05-19T07:50:00Z")
	if err!= nil {
		log.Fatal(err)
	}
	timeEnd,err := time.Parse(time.RFC3339,"2021-05-19T08:15:00Z")
	if err!= nil {
		log.Fatal(err)
	}

	// Request data between 2 datetime
	data, err = conf.Query(&gobadpackets.Request{
		FirstSeenAfter: timeStart.Add(-time.Minute * 15),
		LastSeenBefore: timeEnd.Add(time.Minute * 15),
		SourceIPAddress: "192.168.1.1",
	})
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println(data.Count)
}
