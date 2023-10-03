package main

import (
	"be-tactical-figure/app/router"
	zeromqGo "be-tactical-figure/utils/gozeromq"
	"fmt"
	"log"
	"os"

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

	// Start the Gin server
	if err := router.Routes().Run(port); err != nil {
		log.Fatalln(err)
	}
	select {}

}
