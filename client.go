package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	PlayerID string
	WSURL    string
	APIURL   string
}

func newClient(wsURL string, apiUrl string, playerID string) *Client {
	return &Client{
		WSURL:    wsURL,
		APIURL:   apiUrl,
		PlayerID: playerID,
	}
}

func (c *Client) ConnectWebSocket() error {
	u, err := url.Parse(c.WSURL)
	if err != nil {
		return err
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), http.Header{})
	if err != nil {
		return err
	}

	c.Conn = conn
	fmt.Printf("Connected to %s", wsURL)
	return nil
}

func (c *Client) CloseWebSocket() {
	if c.Conn != nil {
		c.Conn.Close()
	}
}

func (c *Client) readMessages() {
	for {
		_, payload, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Printf("An error Occured %v", err)
		} else {
			handleMessage(payload)
		}
	}
}

func handleMessage(payload []byte) {
	var m Message
	err := json.Unmarshal(payload, &m)
	if err != nil {
		fmt.Printf("Couldn't unmarshall paylaod %v", err)
	}
	switch m.Kind {
	case dealing:

	}

}

func (c *Client) PlaceOrder() error {
	// Implement REST API order placement logic
	// This will involve making HTTP POST requests to c.HTTPURL + "/order"
	return nil
}
