package main

import (
	"context"
	"fmt"
	"grpc_lesson/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v¥n", err)
	}
	defer conn.Close()
	
	client := pb.NewFileServiceClient(conn)

	callListFiles(client)
}

func callListFiles(client pb.FileServiceClient) {
	
	res, err := client.ListFiles(context.Background(), &pb.ListFilesRequest{})
	if err != nil {
		log.Fatalf("error while calling ListFiles RPC: %v¥n", err)
	}

	fmt.Println(res.GetFilenames())
}