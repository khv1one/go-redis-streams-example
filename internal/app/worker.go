package app

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Worker[E any] struct {
	Name string
}

func NewWorker[E any](name string) Worker[E] {
	return Worker[E]{Name: name}
}

func (w Worker[E]) Process(event Event) error {
	time.Sleep(1000 * time.Millisecond)

	a := rand.Intn(20)
	if a == 0 {
		return errors.New("rand error")
	} else {
		fmt.Printf("read event from %v: %v, worker: %v\n", "mystream", event, w.Name)
	}

	return nil
}

func (w Worker[E]) ProcessBroken(broken map[string]interface{}) error {
	fmt.Printf("read broken event from %v: %v, worker: %v\n", "mystream", broken, w.Name)

	return nil
}

func (w Worker[E]) ProcessDead(dead Event) error {
	fmt.Printf("event %v from stream %v is dead!, worker: %v\n", dead.RedisID, "mystream", w.Name)

	return nil
}
