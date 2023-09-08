package main

import (
	"context"
	"fmt"
	"time"

	"go-redis-streams-example/internal/protoexample"

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

func consumerInit() goxstreams.Consumer[protoexample.Event] {
	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	config := goxstreams.ConsumerConfig{
		Stream:         "mystream",
		Group:          "mygroup",
		ConsumerName:   "consumer",
		BatchSize:      100,
		MaxConcurrency: 200,
		MaxRetries:     3,
	}

	myConsumer := goxstreams.NewConsumerWithConverter[protoexample.Event](
		redisClient,
		protoexample.NewWorker("foo"),
		protoexample.ConvertTo,
		config,
	)

	return myConsumer
}
