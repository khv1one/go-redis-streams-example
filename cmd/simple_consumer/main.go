package main

import (
	"context"
	"fmt"
	"time"

	"go-redis-streams-example/internal/jsonexample"

	"github.com/khv1one/goxstreams"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	consumerCtx, _ := context.WithCancel(ctx)

	consumerInit().Run(consumerCtx)
	fmt.Println("Consumer Started")

	<-ctx.Done()
}

func consumerInit() goxstreams.Consumer[jsonexample.Event] {
	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	config := goxstreams.ConsumerConfig{
		Stream:         "mystream",
		Group:          "mygroup",
		ConsumerName:   "consumer",
		BatchSize:      100,
		MaxConcurrency: 5000,
		NoAck:          false,
		MaxRetries:     3,
		CleaneUp:       false,
		FailReadTime:   1000 * time.Millisecond,
		FailIdle:       5000 * time.Millisecond,
	}

	myConsumer := goxstreams.NewConsumer[jsonexample.Event](
		redisClient,
		jsonexample.NewWorker("foo"),
		config,
	)

	return myConsumer
}
