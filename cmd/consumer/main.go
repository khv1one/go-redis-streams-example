package main

import (
	"context"
	"fmt"
	"time"

	"go-redis-streams-example/internal/app"

	"github.com/khv1one/goxstreams"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	consumerCtx, _ := context.WithCancel(ctx)

	consumerInit().Run(consumerCtx)
	fmt.Println("Consumer Started")

	<-ctx.Done()
	// timer1 := time.NewTimer(30 * time.Second)
	// <-timer1.C
	// cancel()
	// fmt.Println("shutdown....")
	//
	// timer := time.NewTimer(30 * time.Second)
	// <-timer.C
}

func consumerInit() goxstreams.Consumer[app.Event] {
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

	myConsumer := goxstreams.NewConsumer[app.Event](
		redisClient,
		app.NewConverter[app.Event](),
		app.NewWorker[app.Event]("foo"),
		config,
	)

	return myConsumer
}
