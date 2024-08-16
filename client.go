package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn       *websocket.Conn
	PlayerID   string
	WSURL      string
	APIURL     string
	HttpClient *http.Client
}
type apiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func newClient(wsURL string, apiUrl string, playerID string) *Client {
	return &Client{
		WSURL:      wsURL,
		APIURL:     apiUrl,
		PlayerID:   playerID,
		HttpClient: &http.Client{},
	}
}

func (c *Client) registerTestNet(id string) error {

	req, err := http.NewRequest("POST", c.APIURL+"/register_testnet", nil)
	if err != nil {
		return fmt.Errorf("Couldn't make post request")
	}
	id = fmt.Sprintf("{%s}", id)
	req.Header.Set("Playerid", id)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("Couldn't send register request ")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Couldn't send register request ")
	}
	var respStruct apiResponse
	json.Unmarshal(body, &respStruct)
	if respStruct.Status == "SUCCESS" {
		fmt.Printf("Connected to TestNet: %s", respStruct.Message)
	} else {
		return fmt.Errorf(respStruct.Message)
	}
	return nil
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

func (c *Client) ListenToMessages(messageChan chan<- []byte) {
	for {
		_, payload, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Printf("An error Occured %v", err)
		} else {
			// will call in main loop with routine and channels.
			messageChan <- payload
		}
	}
}

func (c *Client) PlaceOrder(order *Order) error {

	orderJson, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("can't marshall order")
	}

	req, err := http.NewRequest("POST", c.APIURL+"/order", bytes.NewBuffer(orderJson))
	if err != nil {
		return fmt.Errorf("falied to create order")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HttpClient.Do(req)

	if err != nil {
		return fmt.Errorf("couldn't send order")

	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("failed to read response body")
	}
	var respStruct apiResponse
	err = json.Unmarshal(body, &respStruct)
	if err != nil {
		return fmt.Errorf("couldn't unmarshall error")
	}
	if respStruct.Status == "SUCCESS" {
		fmt.Println("Order placed successfully:", respStruct.Message)

	} else {
		return fmt.Errorf(respStruct.Status)
	}

	return nil
}
