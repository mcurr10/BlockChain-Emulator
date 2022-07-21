package main

import (
	"fmt"
	"log" // calls ox exit 1
	"net/rpc"
)

// Call AddItem from main.go and pass in a Item struct
// Call GetDB from main.go and pass in an empty string
// Call EditItem from main.go and pass in a Item struct
// Call DeleteItem from main.go and pass in a Item struct

type Item struct { // item is record for data being called
	Route string
	Body  string
}

func main() {
	var reply Item
	var db []Item
	// Creates a tcp connection to local host on port 4040

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error:", err)
	}

	buyAlias := Item{"/v1/register/buy/alias/", "buyAlias response"}
	sellAlias := Item{"/v1/register/sell/alias/", "sellAlias response"}
	c := Item{"Third", "A third item"}

	client.Call("API.AddItem", buyAlias, &reply)
	client.Call("API.AddItem", sellAlias, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("Initial Database: ", db) //see whats inside database

	client.Call("API.EditItem", Item{"Third ", " A new second item"}, &reply)
	client.Call("API.DeleteItem,", c, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("Updated Database: ", db)

	client.Call("API.GetByName", "/v1/register/buy/alias/", &reply)
	fmt.Println("BuyAlias: ", reply)

}
