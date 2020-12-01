package main

import (
	"log"

	"github.com/tsuika98/GRPC-Testing/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	// create a connection to TCP port 7777
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	// create a new client
	c := api.NewPingClient(conn)

	// send a ping to the server and read its response
	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "Ping"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Greeting)

}
