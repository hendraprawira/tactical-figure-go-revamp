package publisher

import (
	"fmt"
	"os"

	zmq "github.com/pebbe/zmq4"
)

var GlobalPublisher *zmq.Socket

func PublisherZeroMQ() *zmq.Socket {
	// Create a PUSH socket and bind to the interface
	port := os.Getenv("ZMQ_PORT")
	pushSocket, err := zmq.NewSocket(zmq.PUSH)
	if err != nil {
		fmt.Println("Error creating socket:", err)
	}
	err = pushSocket.Bind("tcp://*:" + port)
	if err != nil {
		fmt.Println("Error binding socket:", err)
	}
	GlobalPublisher = pushSocket

	fmt.Printf("Pubsliher Connected %v \n", port)
	return pushSocket
}
