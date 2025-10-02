package main

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	result := com.InsertIntoLightningPort()
	fmt.Printf("Connection result: %s\n", result)
}
