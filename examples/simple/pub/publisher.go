package main

import (
	"context"
	"log"
	"os"

	"github.com/dtegapp/nexus/v3/client"
	"github.com/dtegapp/nexus/v3/wamp"
)

const (
	addr  = "ws://localhost:8080/ws"
	realm = "realm1"

	exampleTopic = "example.hello"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	cfg := client.Config{
		Realm:  realm,
		Logger: logger,
	}

	// Connect publisher session.
	publisher, err := client.ConnectNet(context.Background(), addr, cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer publisher.Close()

	// Publish to topic.
	err = publisher.Publish(exampleTopic, nil, wamp.List{"hello world"}, nil)
	if err != nil {
		logger.Fatal("publish error:", err)
	}
	logger.Println("Published", exampleTopic, "event")
}
