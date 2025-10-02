package main

import "fmt"

type Mac struct{}

func (m *Mac) InsertIntoLightningPort() string {
	fmt.Println("Lightning connector is plugged into mac machine.")
	return "Lightning"
}
