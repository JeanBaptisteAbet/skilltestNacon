package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "skilltestnacon/api/grpcserver/liveevents"
	"skilltestnacon/constant"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	err := godotenv.Load(".env") // WARNING, change this path if needed
	if err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}

	conn, err := grpc.NewClient(
		fmt.Sprintf(":%v", os.Getenv(constant.API_PORT)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()

	c := pb.NewLiveEventsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	md := metadata.MD{}
	md.Set("Authorization", os.Getenv(constant.GRPC_API_KEY))

	ctx = metadata.NewOutgoingContext(ctx, md)

	r, err := c.ListEvents(ctx, &pb.ListEventsRequest{})
	if err != nil {
		log.Fatalf("error calling function ListEvents: %v", err)
	}

	log.Printf("Response from gRPC server's ListEvents function: %s", r.LiveEvents)

	r2, err := c.CreateEvent(ctx, &pb.CreateEventRequest{Title: "eventTitle", Description: "description", StartTime: time.Now().Unix(), Rewards: "rewards"})
	if err != nil {
		log.Fatalf("error calling function CreateEvent: %v", err)
	}

	log.Printf("Response from gRPC server's CreateEvent function: %v", r2.Id)
}
