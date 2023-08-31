package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go-redis-streams-example/internal/app"

	"github.com/khv1one/goxstreams"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	converter := app.Converter[app.Event]{}
	producer := goxstreams.NewProducer[app.Event](redis.NewClient(&redis.Options{Addr: "localhost:6379"}), converter)
	go write(producer, ctx)

	fmt.Println("Producer started")
	<-ctx.Done()
}

func write(producer goxstreams.Producer[app.Event], ctx context.Context) {
	for {
		event := app.Event{Message: "message", Name: "name", Foo: rand.Intn(1000), Bar: rand.Intn(1000)}

		err := producer.Produce(ctx, event, "mystream")
		fmt.Printf("produced %v\n", event)
		if err != nil {
			fmt.Printf("write error %v\n", err)
			time.Sleep(time.Second)
			continue
		}

		time.Sleep(100 * time.Millisecond)
	}
}
