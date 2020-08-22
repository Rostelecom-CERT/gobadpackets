package gobadpackets

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Data struct with data from query
type Data struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		EventID         string `json:"event_id"`
		SourceIPAddress string `json:"source_ip_address"`
		Country         string `json:"country"`
		UserAgent       string `json:"user_agent"`
		Payload         string `json:"payload"`
		PostData        string `json:"post_data"`
		TargetPort      int    `json:"target_port"`
		Protocol        string `json:"protocol"`
		Tags            []struct {
			Cve         string `json:"cve"`
			Category    string `json:"category"`
			Description string `json:"description"`
		} `json:"tags"`
		EventCount int       `json:"event_count"`
		FirstSeen  time.Time `json:"first_seen"`
		LastSeen   time.Time `json:"last_seen"`
	} `json:"results"`
}

// Client main struct
type Client struct {
	APIKey string
	URL    string
	conn   *http.Client
}

const defaultURL = "https://api.badpackets.net/v1/"

// New constructor function
func New(APIKey string, URL string) (*Client, error) {
	client := &http.Client{}
	if URL == "" {
		URL = defaultURL
	}
	if APIKey == "" {
		return nil, errors.New("didn't find API key")
	}
	return &Client{
		conn:   client,
		URL:    URL,
		APIKey: APIKey,
	}, nil
}

func (c *Client) query(url string) ([]byte, int, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", c.APIKey)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	return body, resp.StatusCode, nil
}

// Ping return status of connection
func (c *Client) Ping() bool {
	url := fmt.Sprintf("%sping", c.URL)
	_, code, err := c.query(url)
	if err != nil {
		return false
	}
	if code == 200 {
		return true
	}
	return false
}

// Query return data from request
func (c *Client) Query() (*Data, error) {
	url := fmt.Sprintf("%squery", c.URL)
	body, code, err := c.query(url)
	if err != nil {
		return &Data{}, nil
	}
	switch code {
	case 200:
		var data Data
		err = json.Unmarshal(body, &data)
		if err != nil {
			return &Data{}, err
		}

		return &data, nil
	case 400:
		return &Data{}, errors.New("invalid query")
	case 401:
		return &Data{}, errors.New("unauthorized")
	case 403:
		return &Data{}, errors.New("forbidden")
	}
	return &Data{}, errors.New("unknown error")
}
