package main

import (
	grpcserver "be-tactical-figure/app/grpc-server"
	"be-tactical-figure/app/router"
	"be-tactical-figure/utils/zeromq/publisher"
	zeroMQ "be-tactical-figure/utils/zeromq/subscriber"
	"fmt"
	"log"
	"os"

	// zmq4 "github.com/go-zeromq/zmq4"
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

	go grpcserver.StartPoint()

	publisher.PublisherZeroMQ()

	// Start the Gin server
	if err := router.Routes().Run(port); err != nil {
		log.Fatalln(err)
	}

	select {}
}
