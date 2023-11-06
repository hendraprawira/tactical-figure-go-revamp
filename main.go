package main

import (
	grpcroute "be-tactical-figure/app/grpc-route"
	"be-tactical-figure/app/router"
	"be-tactical-figure/utils/zeromq/publisher"
	zeroMQ "be-tactical-figure/utils/zeromq/subscriber"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	zmq "github.com/pebbe/zmq4"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	port := ":" + os.Getenv("ACTIVE_PORT")

	context, err := zmq.NewContext()
	if err != nil {
		fmt.Println("Error creating context:", err)
		return
	}
	defer context.Term()

	go func() {
		err := zeroMQ.SubscriberZeroMQ()
		if err != nil {
			log.Fatalf("Subscriber error: %v", err)
		}
	}()

	go grpcroute.StartgRPC()

	publisher.PublisherZeroMQ()

	// Start the Gin server
	if err := router.Routes().Run(port); err != nil {
		log.Fatalln(err)
	}

	select {}
}
