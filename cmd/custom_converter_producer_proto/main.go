package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go-redis-streams-example/internal/protoexample"

	"github.com/khv1one/goxstreams"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	producer := goxstreams.NewProducerWithConverter[protoexample.Event](
		redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
		protoexample.ConvertFrom,
	)

	go write(producer, ctx)

	fmt.Println("Producer started")
	<-ctx.Done()
}

func write(producer goxstreams.Producer[protoexample.Event], ctx context.Context) {
	for {
		event := protoexample.Event{
		Message: "message", Name: "name", Foo: int64(rand.Intn(1000)), Bar: int64(rand.Intn(1000)),
		SubEvent: &protoexample.SubEvent{
			Barbar: "1234",
			Foofoo: &protoexample.SubSubEvent{Foofoofoo: 777},
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
