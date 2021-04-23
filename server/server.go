package main

import (
	"context"
	"fmt"
	"hash/fnv"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/jmwoliver/lilurl/url_service"
)

const (
	baseURL = "lilurl.cc/"
)

type urlServiceServer struct {
	pb.UnimplementedURLServiceServer
	link *pb.Link
}

func generateHash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// GetShortenURL returns a shortened URL.
func (s *urlServiceServer) GetShortenURL(ctx context.Context, link *pb.Link) (*pb.Link, error) {
	hash := generateHash(link.Url)
	newUrl := baseURL + fmt.Sprint(hash)
	log.Printf("Got request from client. Shortening URL to %s\n", newUrl)

	// Connect to the database and put it in if it doesn't already exist

	return &pb.Link{Url: newUrl}, nil
}

func newServer() *urlServiceServer {
	s := &urlServiceServer{}
	return s
}

func main() {
	serviceAddr := "localhost:9000"
	lis, err := net.Listen("tcp", serviceAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterURLServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
