package gozeromq

import (
	tacticalfigure "be-tactical-figure/app/controller/tactical-figure"
	"be-tactical-figure/app/db"
	"be-tactical-figure/app/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-zeromq/zmq4"
)

func StartZMQSubs() error {
	mockID := os.Getenv("MOCK_ID")

	// Define a list of IP addresses to connect to (multiple)
	zmqPorts := []string{
		"tcp://192.168.160.132:" + os.Getenv("ZMQ_PORT"),
		"tcp://192.168.160.133:" + os.Getenv("ZMQ_PORT"),
		// Add more addresses as needed
	}

	// Create a map to track the connection state for each host
	connectionState := make(map[string]bool)

	// Prepare a subscriber for each IP address
	subs := make([]zmq4.Socket, 0)
	for _, port := range zmqPorts {
		sub := zmq4.NewSub(context.Background())
		subs = append(subs, sub)

		// Subscribe to the desired topics
		topics := []string{"Point", "Single", "Multi"}
		for _, topic := range topics {
			err := sub.SetOption(zmq4.OptionSubscribe, topic)
			if err != nil {
				log.Fatalf("could not subscribe: %v", err)
				return err
			}
		}

		// Initialize the connection state for this host
		connectionState[port] = true

		// Function to handle reconnection to the publisher for this address
		reconnect := func(port string, sub zmq4.Socket) {
			for {
				err := sub.Dial(port)
				if err != nil {
					log.Printf("could not dial %s: %v", port, err)
					// Update the connection state for this host
					connectionState[port] = false
				} else {
					log.Printf("Connected to publisher at %s", port)
					// Update the connection state for this host
					connectionState[port] = true
					return
				}
				time.Sleep(10 * time.Second) // Adjust the sleep time as needed
			}
		}

		// Start a goroutine to handle initial connection and reconnection for this address
		go reconnect(port, sub)

		go func(port string, sub zmq4.Socket) {
			for {
				// Read envelope and process messages (similar to your existing code)
				msg, err := sub.Recv()
				if err != nil {
					log.Printf("could not receive message: %v", err)
					time.Sleep(10 * time.Second)
					log.Println("Connection lost, attempting to reconnect...")
					reconnect(port, sub)
					continue
				}

				// Process the received message (similar to your existing code)
				if string(msg.Frames[2]) != mockID {
					if string(msg.Frames[0]) == "Point" {
						fmt.Println("Data Receive", string(msg.Frames[1]))
						if string(msg.Frames[2]) == "true" {
							var point *models.Point
							if err := json.Unmarshal([]byte(string(msg.Frames[1])), &point); err != nil {
								fmt.Println("Error:", err)
								return
							}
							db.InsertDBPoint(point)
							fmt.Println("Save To DB")
						}
						tacticalfigure.SseChannel <- string(msg.Frames[1])
						fmt.Println("Didnt Save To DB")

					} else if string(msg.Frames[0]) == "Single" {
						fmt.Println("Data Receive", string(msg.Frames[1]))
						if string(msg.Frames[2]) == "true" {
							var single *models.SingleLine
							if err := json.Unmarshal([]byte(string(msg.Frames[1])), &single); err != nil {
								fmt.Println("Error:", err)
								return
							}
							db.InsertDBSingle(single)
							fmt.Println("Save To DB")
						}
						tacticalfigure.SseChannel <- string(msg.Frames[1])
						fmt.Println("Didnt Save To DB")

					} else if string(msg.Frames[0]) == "Multi" {
						fmt.Println("Data Receive", string(msg.Frames[1]))
						if string(msg.Frames[2]) == "true" {
							var multi *models.MultiLine
							if err := json.Unmarshal([]byte(string(msg.Frames[1])), &multi); err != nil {
								fmt.Println("Error:", err)
								return
							}
							db.InsertDBMulti(multi)
							fmt.Println("Save To DB")
						}
						tacticalfigure.SseChannel <- string(msg.Frames[1])
						fmt.Println("Didnt Save To DB")
					}
				} else {
					fmt.Println("Data Receive But Same Mock")
				}

			}
		}(port, sub)
	}

	// Keep the main goroutine running to handle other tasks if needed
	select {}
}
