package grpcserver

import (
	"be-tactical-figure/app/db"
	pbPoint "be-tactical-figure/app/generated-protos"
	"be-tactical-figure/app/models"
	"be-tactical-figure/utils/zeromq/publisher"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type pointServer struct {
	pbPoint.UnimplementedTodoServiceServer
}

func (ts *pointServer) CreatePoint(ctx context.Context, in *pbPoint.CreatePointRequest) (*pbPoint.CreatePointResponse, error) {
	point := in.GetPoint()
	mockID := os.Getenv("MOCK_ID")
	newPoint := models.Point{
		Coordinates:    point.GetCoordinates(),
		Color:          point.Color,
		Amplifications: point.Amplifications,
		Opacity:        int(point.Opacity),
		Altitude:       point.Altitude,
		IdUnique:       point.IdUnique,
		SaveDB:         point.SaveDb,
	}

	newCoord := fmt.Sprintf("%v", newPoint.Coordinates)
	newSaveDb := fmt.Sprintf("%v", newPoint.SaveDB)

	if newPoint.SaveDB {
		result := db.DB.Create(&models.TacticalFigureInput{
			FigureType:     "Poinst",
			Coordinates:    newCoord,
			Color:          newPoint.Color,
			Amplifications: newPoint.Amplifications,
			Opacity:        newPoint.Opacity,
			Altitude:       newPoint.Altitude,
			UpdatedAt:      time.Now(),
			IdUnique:       newPoint.IdUnique,
			IsDeleted:      false,
		})

		if result.Error != nil {
			log.Panic(result.Error)
			return nil, result.Error
		}
	}

	// Encode the Point to a JSON []byte
	jsonData, err := json.Marshal(newPoint)
	if err != nil {
		fmt.Println("Error encoding Point to JSON:", err)
		return nil, err
	}
	messageBuffer := make([][]string, 0)
	datas := []string{"Point", string(jsonData), mockID, newSaveDb}
	messageBuffer = append(messageBuffer, datas)
	log.Print(messageBuffer[0])
	// Publish Message
	for len(messageBuffer) > 0 {
		_, err := publisher.GlobalPublisher.SendMessageDontwait(messageBuffer[0])
		if err != nil {
			log.Print(err)
			// If message could not be sent, break the loop and try again later
			break
		} else {
			fmt.Println("Sent")
			messageBuffer = messageBuffer[1:]
		}
	}

	return &pbPoint.CreatePointResponse{Point: &pbPoint.TacticPoint{
		Coordinates:    newPoint.Coordinates,
		Color:          newPoint.Color,
		Amplifications: newPoint.Amplifications,
		Opacity:        int32(newPoint.Opacity),
		Altitude:       newPoint.Altitude,
		IdUnique:       newPoint.IdUnique,
		SaveDb:         newPoint.SaveDB,
	}}, nil

}

func (ts *pointServer) GetPoint(ctx context.Context, in *pbPoint.GetPointRequest) (*pbPoint.GetPointResponse, error) {
	var a []float64
	return &pbPoint.GetPointResponse{Point: &pbPoint.TacticPoint{
		Coordinates:    a,
		Color:          "newPoint.Color",
		Amplifications: "newPoint.Amplifications",
		Opacity:        2,
		Altitude:       2,
		IdUnique:       "2",
		SaveDb:         true,
	}}, nil
}

func StartPoint() {
	lis, err := net.Listen("tcp", "localhost"+":8881")
	if err != nil {
		log.Fatalf("grpc-todo failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pbPoint.RegisterTodoServiceServer(s, &pointServer{})
	reflection.Register(s)
	log.Printf("grpc-todo server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("grpc-todo failed to serve: %v", err)
	}
}
