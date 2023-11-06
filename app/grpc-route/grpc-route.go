package grpcroute

import (
	"fmt"
	"log"
	"net"
	"os"

	grpc_tacticalfigure "be-tactical-figure/app/controller/grpc/tactical-figure"
	pbPoint "be-tactical-figure/app/generated-protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartgRPC() {
	port := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", "localhost"+":"+port)
	if err != nil {
		log.Fatalf("gRPC failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pbPoint.RegisterTodoServiceServer(s, &grpc_tacticalfigure.PointServer{})
	reflection.Register(s)
	fmt.Printf("gRPC Server Listening at Port %v \n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("gRPC failed to serve: %v", err)
	}
}
