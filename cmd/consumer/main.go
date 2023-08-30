package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"go-redis-streams-example/internal/app"

	streams "github.com/khv1one/goxstreams/pkg/goxstreams/client"
	"github.com/khv1one/goxstreams/pkg/goxstreams/consumer"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	workerCtx, _ := context.WithCancel(context.Background())

	streamClient := streamClientInit()

	worker := Worker[app.Event]{Name: "foo"}
	converter := app.Converter[app.Event]{}

	consumer := consumer.NewConsumer[app.Event](streamClient, converter, worker, 3)
	consumer.Run(workerCtx)

	fmt.Printf("Redis started %s", "localhost:6379")
	<-ctx.Done()
}

func streamClientInit() streams.StreamClient {
	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	streamParams := streams.Params{
		Stream:   "mystream",
		Group:    "mygroup",
		Consumer: "consumer",
		Batch:    50,
	}

	streamClient := streams.NewClient(redisClient, streamParams)

	return streamClient
}

type Worker[E any] struct {
	Name string
}

func (w Worker[E]) Process(event app.Event) error {
	fmt.Printf("read event from %v: %v, worker: %v\n", "mystream", event, w.Name)

	a := rand.Intn(5)
	if a == 0 {
		return errors.New("rand error")
	}

	time.Sleep(100 * time.Millisecond)
	return nil
}

func (w Worker[E]) ProcessBroken(broken map[string]interface{}) error {
	fmt.Printf("read broken event from %v: %v, worker: %v\n", "mystream", broken, w.Name)

	return nil
}

func (w Worker[E]) ProcessDead(dead app.Event) error {
	fmt.Printf("event %v from stream %v is dead!, worker: %v\n", dead.RedisID, "mystream", w.Name)

	return nil
}
