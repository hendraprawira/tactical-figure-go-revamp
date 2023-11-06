package main

import (
	pbPoint "be-tactical-figure/app/generated-protos"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

const userServiceAddress = "localhost:8881"

func main() {
	// Create connection to gRPC server
	conn, err := grpc.Dial(userServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to service: %v", err)
		return
	}
	defer conn.Close()

	// Create new userService client
	userServiceClient := pbPoint.NewTodoServiceClient(conn)

	// Create connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	a := []float64{-13611388.13050057, 6015418.379246738}

	// Record the start time
	startTime := time.Now()

	response, err := userServiceClient.CreatePoint(ctx, &pbPoint.CreatePointRequest{
		Point: &pbPoint.TacticPoint{
			Coordinates:    a,
			Color:          "rgba(0, 0, 0, 1)",
			Amplifications: "Search",
			Opacity:        100,
			Altitude:       2,
			IdUnique:       "1697788077309-1121",
			SaveDb:         true,
		},
	})
	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}

	// Calculate and print the response time
	elapsedTime := time.Since(startTime)
	fmt.Printf("Response received in %v\n", elapsedTime)

	// Show response
	fmt.Println(response)
}
