package main

// rpc so that an external computer can call to interface will call interface remotely
// functions need to have 2 agruments to be rpc
// make sure a server can call pieces of our code
import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// crud application
type Item struct { // created a struct to store data in a group is non descript
	Route string
	Body  string
}
type API int        // make functions methods
var database []Item // database passed through reply pointer and passed through database
func (a *API) GetDB(empty string, reply *[]Item) error { // grabs information from database
	*reply = database //passes to reply pointer
	return nil
}

//reply is a pointer
//Get by name get items by their name
func (a *API) GetByName(route string, reply *Item) error {
	var getItem Item
	for _, val := range database { // iterate through data
		if val.Route == route { // if title is equal to title
			getItem = val
		}
	}

	*reply = getItem
	return nil // statisfy error type return type
}


func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)

	*reply = item
	return nil 

func (a *API) UpdateItem(item Item, reply *Item) error {
	for i, val := range database {
		if val.Route == item.Route {
			database[i] = item
		}
	}

	*reply = item
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	for i, val := range database {
		if val.Route == item.Route {
			database = append(database[:i], database[i+1:]...)
		}
	}

	*reply = item
	return nil
}

// where server is made
func main() {
	api := new(API)
	err := rpc.Register(api) // register type to call methods remotely

	if err != nil { // this is where I want to put !!
		log.Fatal("error registering API", err)
	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040") // listen on port 4040

	if err != nil {
		log.Fatal("Listner error ", err)
	}

	log.Printf("Serving RPC server on port %d", 4040)

	err = http.Serve(listener, nil)

	if err != nil {

		log.Fatal("Error serving:", err)

	}

}
