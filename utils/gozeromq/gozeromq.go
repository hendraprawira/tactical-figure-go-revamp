package gozeromq

import (
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
	port := "tcp://localhost:" + os.Getenv("ZMQ_PORT")
	//  Prepare our subscriber
	sub := zmq4.NewSub(context.Background())
	defer sub.Close()

	// Subscribe to multiple topics
	topics := []string{"Point", "Single", "Multi"}
	for _, topic := range topics {
		err := sub.SetOption(zmq4.OptionSubscribe, topic)
		if err != nil {
			log.Fatalf("could not subscribe: %v", err)
			return err
		}
	}

	// Function to handle reconnection to the publisher
	reconnect := func() {
		for {
			err := sub.Dial(port)
			if err != nil {
				log.Printf("could not dial: %v", err)
			} else {
				log.Println("Connected to publisher")
				return // Exit the function when the connection is established
			}

			// Sleep for some time before trying to connect again
			time.Sleep(500 * time.Millisecond)
		}
	}

	// Start a goroutine to handle initial connection and reconnection
	go reconnect()

	go func() {
		for {
			// Read envelope
			msg, err := sub.Recv()
			if err != nil {
				log.Printf("could not receive message: %v", err)
				// Sleep for some time before trying to receive again
				time.Sleep(500 * time.Millisecond)

				// If the connection is lost, attempt to reconnect
				log.Println("Connection lost, attempting to reconnect...")
				reconnect()
				continue
			}
			if string(msg.Frames[0]) == "Point" {
				var point *models.Point
				if err := json.Unmarshal([]byte(string(msg.Frames[1])), &point); err != nil {
					fmt.Println("Error:", err)
					return
				}
				db.InsertDBPoint(point)
				fmt.Println("Data Receive")
			} else if string(msg.Frames[0]) == "Single" {
				var single *models.SingleLine
				if err := json.Unmarshal([]byte(string(msg.Frames[1])), &single); err != nil {
					fmt.Println("Error:", err)
					return
				}
				db.InsertDBSingle(single)
				fmt.Println("Data Receive")
			} else if string(msg.Frames[0]) == "Multi" {
				var multi *models.MultiLine
				if err := json.Unmarshal([]byte(string(msg.Frames[1])), &multi); err != nil {
					fmt.Println("Error:", err)
					return
				}
				db.InsertDBMulti(multi)
				fmt.Println("Data Receive")
			}
		}
	}()

	// Keep the main goroutine running to handle other tasks if needed
	select {}
}
