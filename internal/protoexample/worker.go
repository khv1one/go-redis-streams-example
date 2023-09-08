package protoexample

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/khv1one/goxstreams"
)

type Worker[E interface{ Event }] struct {
	Name string
}

func NewWorker[E interface{ Event }](name string) Worker[E] {
	return Worker[E]{Name: name}
}

func (w Worker[E]) Process(event goxstreams.RedisMessage[E]) error {
	time.Sleep(1000 * time.Millisecond)

	a := rand.Intn(20)
	if a == 0 {
		return errors.New("rand error")
	} else {
		fmt.Printf("read event from %s: id: %s, retry: %d, body: %v, worker: %v\n",
			"mystream", event.ID, event.RetryCount, event.Body, w.Name)
	}

	return nil
}

func (w Worker[E]) ProcessBroken(broken goxstreams.RedisBrokenMessage) error {
	fmt.Printf("read broken event from %s: id: %s, retry: %d, body: %v, worker: %v, err: %s\n",
		"mystream", broken.ID, broken.RetryCount, broken.Body, w.Name, broken.Error.Error())

	return nil
}

func (w Worker[E]) ProcessDead(dead goxstreams.RedisMessage[E]) error {
	fmt.Printf("read from %s is dead!!! id: %s, retry: %d, body: %v, worker: %v\n",
		"mystream", dead.ID, dead.RetryCount, dead.Body, w.Name)

	return nil
}
