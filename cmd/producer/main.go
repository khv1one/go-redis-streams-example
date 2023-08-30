package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go-redis-streams-example/internal/app"

	streams "github.com/khv1one/goxstreams/pkg/goxstreams/client"
	"github.com/khv1one/goxstreams/pkg/goxstreams/producer"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	streamClient := streamClientInit()
	converter := app.Converter[app.Event]{}
	producer := producer.NewProducer[app.Event](streamClient, converter)
	go write(producer, ctx)

	fmt.Printf("Producer started")
	<-ctx.Done()
}

func streamClientInit() streams.StreamClient {
	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	clientParams := streams.Params{Stream: "mystream"}

	streamClient := streams.NewClient(redisClient, clientParams)

	return streamClient
}

func write(producer producer.Producer[app.Event], ctx context.Context) {
	for {
		event := app.Event{Message: "message", Name: "name", Foo: rand.Intn(1000), Bar: rand.Intn(1000)}

		err := producer.Produce(ctx, event, "mystream")
		fmt.Printf("produced %v\n", event)
		if err != nil {
			log.Printf("write error %v\n", err)
			time.Sleep(time.Second)
			continue
		}

		time.Sleep(10 * time.Millisecond)
	}
}
