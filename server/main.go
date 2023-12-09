package main

import (
	"bufio"
	"context"
	"fmt"
	"grpc_lesson/pb"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFileServiceServer
}
func (*server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	fmt.Println("ListFiles invoked")

	dir := "/Users/takuya/go/src/workspace/grpc_lesson/storage"

	paths, err := os.ReadDir(dir);
	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, path := range paths {
		if !path.IsDir() {
			filenames = append(filenames, path.Name())
		}  
	}

	res := &pb.ListFilesResponse{
		Filenames: filenames,
	}

	return res, nil
}

func (*server) Download(req *pb.DownloadRequest, stream pb.FileService_DownloadServer) error {
	fmt.Println("Download invoked")

	filename := req.GetFilename()
	path := "/Users/takuya/go/src/workspace/grpc_lesson/storage" + filename

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	buf := make([]byte, 5)
	for {
		n, err := file.Read(buf)
		if n == 0 || err = io.EOF {
			break
		}
		if err != nil {
			return err
		}

		res := &pb.DownloadResponse{Data: buf[:n]}
		sendErr := stream.Send(res)
		if sendErr != nil {
			return sendErr
		}
		time.Spleep(1 * time.Second)
	}
	return nil

}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v¥n", err)
	}

	s := grpc.NewServer()
	pb.RegisterFileServiceServer(s, &server{})

	fmt.Println("start server on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v¥n", err)
	} 

}