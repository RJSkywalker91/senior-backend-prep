package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "matchmaking/cmd/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {

	username := flag.String("username", "", "Username for the player")
	password := flag.String("password", "", "Password for the player")
	flag.Parse()

	if *username == "" || *password == "" {
		log.Fatalf("❗ All parameters are required.\nUsage: go run main.go -username=<name> -password=<pass>")
	}

	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPlayerServiceClient(conn)

	// Prepare the request
	req := &pb.PlayerLoginRequest{
		Username: *username,
		Password: *password,
	}

	// Add a timeout to the call
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call the Create RPC
	resp, err := client.Login(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		switch st.Code() {
		case codes.AlreadyExists:
			fmt.Printf("❌ Username '%s' already taken", req.Username)
		case codes.Unavailable:
			fmt.Printf("🚫 A problem occurred. Please try again later.")
		default:
			fmt.Printf("⚠️  An unknown problem occurred: %s", st)
		}
	} else {
		fmt.Printf("✅ Successfully logged in. Please use token:%s\n", resp.Token)
	}
}
