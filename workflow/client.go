package workflow

import (
	"log"

	"go.temporal.io/sdk/client"
)

type Client struct {
	client.Client
}

func NewClient() *Client {
	// Create the client object just once per process
	c, err := client.NewLazyClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
		panic(err)
	}

	return &Client{
		c,
	}
}

func (c *Client) Close() {
	c.Client.Close()
}
