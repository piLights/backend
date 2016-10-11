package main

import (
	"context"
	"io"
	"log"

	"gitlab.com/piLights/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:13337", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := LighterGRPC.NewLighterClient(conn)

	setColor(client, &LighterGRPC.ColorMessage{
		true,
		255,
		0,
		0,
		100,
		"Penis",
		"",
	})

	loadServerLog(client, &LighterGRPC.LogRequest{
		1,
		300,
		"",
	})
}

func loadServerLog(client LighterGRPC.LighterClient, logRequest *LighterGRPC.LogRequest) {
	stream, err := client.LoadServerLog(context.Background(), logRequest)
	if err != nil {
		log.Fatal(err)
	}

	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			grpclog.Fatalf("%v.LoadServerLog(_) = _, %v", client, err)
		}
		grpclog.Println(feature)
	}

	log.Println(stream)
}

// printFeatures lists all the features within the given bounding Rectangle.
func setColor(client LighterGRPC.LighterClient, colorMessage *LighterGRPC.ColorMessage) {
	stream, err := client.SetColor(context.Background(), colorMessage)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(stream)
}
