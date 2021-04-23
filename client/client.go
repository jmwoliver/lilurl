package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/jmwoliver/lilurl/url_service"
	"google.golang.org/grpc"
)

var (
	url = flag.String("url", "", "The URL to be shortened.")
)

func getShortenURL(client pb.URLServiceClient, link *pb.Link) *pb.Link {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newLink, err := client.GetShortenURL(ctx, link)
	if err != nil {
		log.Fatalf("%v.GetShortenURL(_) = _, %v: ", client, err)
	}
	return newLink
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	serverAddr := "localhost:9000"
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewURLServiceClient(conn)

	// Passing in URL as a flag is required
	if *url == "" {
		log.Fatal("'url' flag is required.\n")
	}

	log.Printf("Shortening URL for %s. Sending to server.\n", *url)

	// Shorten a URL passed in
	newLink := getShortenURL(client, &pb.Link{Url: *url})

	log.Printf("Got response back from server. Shortened URL is %s\n", newLink.Url)
}
