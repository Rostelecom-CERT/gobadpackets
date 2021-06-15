package gobadpackets

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var APIKey string
var URL string

func init() {
	flag.StringVar(&APIKey, "api", "", "API key badpackets")
	flag.StringVar(&URL, "url", "", "URL badpackets")
	flag.Parse()
	if APIKey == "" {
		fmt.Println("API key is required to run the tests against badpackets")
		os.Exit(1)
	}

}

func TestNew(t *testing.T) {
	conf, err := New(APIKey, URL)
	if err != nil {
		t.Fatal(err)
	}
	if conf == nil {
		t.Error("connection to api.badpackets.net isn't established")
		return
	}
}

func TestPing(t *testing.T) {
	conf, err := New(APIKey, URL) // Initiate new connection to API
	if err != nil {
		t.Fatal(err)
	}
	stat := conf.Ping()
	if !stat {
		t.Error("access forbidden ", err)
		return
	}
}

func TestQuery(t *testing.T) {
	conf, err := New(APIKey, URL) // Initiate new connection to API
	if err != nil {
		t.Fatal(err)
	}
	data, err := conf.Query(&Request{Protocol: "tcp", TargetPort: 443})
	if err != nil {
		t.Fatal(err)
	}
	if data == nil {
		t.Error("service return nil data")
		return
	}
}
