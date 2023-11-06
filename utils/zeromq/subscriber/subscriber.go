package subscriber

import (
	tacticalfigure "be-tactical-figure/app/controller/rest/tactical-figure"
	"be-tactical-figure/app/db"
	"be-tactical-figure/app/models"
	"encoding/json"
	"fmt"
	"log"
	"os"

	zmq "github.com/pebbe/zmq4"
)

func SubscriberZeroMQ() error {
	port := os.Getenv("ZMQ_PORT")
	mockID := os.Getenv("MOCK_ID")
	//  Prepare our subscriber PULL == SUBS
	subscriber, _ := zmq.NewSocket(zmq.PULL)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:" + port)
	topics := []string{"Point", "Single", "Multi"}
	for _, topic := range topics {
		subscriber.SetSubscribe(topic)
	}
	fmt.Printf("ZeroMQ Subscriber Already Connect at Port %v \n", port)

	for {
		msg, err := subscriber.RecvMessage(0)
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			continue
		}
		if string(msg[2]) == mockID {
			if string(msg[0]) == "Point" {
				fmt.Println("Data Receive", string(msg[1]))
				if string(msg[3]) == "true" {
					var point *models.Point
					if err := json.Unmarshal([]byte(string(msg[1])), &point); err != nil {
						fmt.Println("Error:", err)
						return err
					}
					db.InsertDBPoint(point)
					fmt.Println("Save To DB")
				} else {
					fmt.Println("Didnt Save To DB")
				}
				tacticalfigure.SseChannel <- string(msg[1])

			} else if string(msg[0]) == "Single" {
				fmt.Println("Data Receive", string(msg[1]))
				if string(msg[3]) == "true" {
					var single *models.SingleLine
					if err := json.Unmarshal([]byte(string(msg[1])), &single); err != nil {
						fmt.Println("Error:", err)
						return err
					}
					db.InsertDBSingle(single)
					fmt.Println("Save To DB")
				} else {
					fmt.Println("Didnt Save To DB")
				}
				tacticalfigure.SseChannel <- string(msg[1])

			} else if string(msg[0]) == "Multi" {
				fmt.Println("Data Receive", string(msg[1]))
				if string(msg[3]) == "true" {
					var multi *models.MultiLine
					if err := json.Unmarshal([]byte(string(msg[1])), &multi); err != nil {
						fmt.Println("Error:", err)
						return err
					}
					db.InsertDBMulti(multi)
					fmt.Println("Save To DB")
				} else {
					fmt.Println("Didnt Save To DB")
				}
				tacticalfigure.SseChannel <- string(msg[1])
			}
		} else {
			fmt.Println("Data Receive But Same Mock")
		}
	}
}
