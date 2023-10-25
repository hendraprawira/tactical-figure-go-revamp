package main

import (
	"be-tactical-figure/app/router"
	zeromqGo "be-tactical-figure/utils/gozeromq"
	"context"
	"fmt"
	"log"
	"os"

	zmq4 "github.com/go-zeromq/zmq4"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	port := ":" + os.Getenv("ACTIVE_PORT")

	//start Kafka
	go func() {
		err := zeromqGo.StartZMQSubs()
		if err != nil {
			log.Fatalf("Subscriber error: %v", err)
		}
	}()

	pub := zmq4.NewPub(context.Background())
	defer pub.Close()

	errs := pub.Listen("tcp://*:5563")
	if errs != nil {
		log.Fatalf("could not listen: %v", errs)
	}

	// Start the Gin server
	if err := router.Routes(pub).Run(port); err != nil {
		log.Fatalln(err)
	}
	select {}

}
