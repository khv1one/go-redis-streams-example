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

	producer := goxstreams.NewProducer[app.Event](redis.NewClient(&redis.Options{Addr: "localhost:6379"}))
	go write(producer, ctx)

	fmt.Println("Producer started")
	<-ctx.Done()
}

func write(producer goxstreams.Producer[app.Event], ctx context.Context) {
	for {
		event := app.Event{
			Message: "message", Name: "name", Foo: rand.Intn(1000), Bar: rand.Intn(1000),
			SubEvent: app.SubEvent{
				BarBar: "1234",
				FooFoo: app.SubSubEvent{FooFooFoo: 777},
			},
		}

		err := producer.Produce(ctx, event, "mystream")
		if err != nil {
			fmt.Printf("write error %v\n", err)
			time.Sleep(time.Second)
			continue
		}

		fmt.Printf("produced %v\n", event)

		time.Sleep(100 * time.Millisecond)
	}
}
