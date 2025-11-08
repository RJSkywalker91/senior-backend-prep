package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	q "matchmaking/cmd/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	jwt := flag.String("jwt", "", "Authenticated JWT Token")
	port := flag.String("port", "8080", "Port to run client on")
	flag.Parse()

	target := fmt.Sprintf("localhost:%s", *port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := q.NewMatchmakingServiceClient(conn)
	md := metadata.Pairs("authorization", "Bearer "+*jwt)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	stream, err := client.Queue(ctx, &q.QueueRequest{Region: "NA"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		ev, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		switch ev.Type {
		case q.QueueEvent_QUEUED:
			log.Print("Queued up for match...")
		case q.QueueEvent_MATCH_FOUND:
			log.Print("Match found!")
			log.Printf("Match ID: %s | Opponent Id: %s", ev.MatchId, ev.OpponentId)
			return
		}
	}
}
