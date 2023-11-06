package publisher

import (
	"fmt"
	"log"

	zmq "github.com/pebbe/zmq4"
)

var GlobalPublisher *zmq.Socket

func PublisherZeroMQ() *zmq.Socket {
	// Create a PUSH socket and bind to the interface
	pushSocket, err := zmq.NewSocket(zmq.PUSH)
	if err != nil {
		fmt.Println("Error creating socket:", err)
	}
	err = pushSocket.Bind("tcp://*:5555")
	if err != nil {
		fmt.Println("Error binding socket:", err)
	}
	GlobalPublisher = pushSocket

	log.Print("Pubsliher Connected")
	return pushSocket
}
