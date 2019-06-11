package pubsub

import (
	"context"
	"fmt"
)

type Subscription struct {
	topic    string
	ch       chan *Message
	cancelCh chan<- *Subscription
	err      error
}

func (sub *Subscription) Topic() string {
	return sub.topic
}

func (sub *Subscription) Next(ctx context.Context) (*Message, error) {
	select {
	case msg, ok := <-sub.ch:
		if !ok {
			fmt.Println("not ok")
			return msg, sub.err
		}
		fmt.Println("ok")
		return msg, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (sub *Subscription) Cancel() {
	sub.cancelCh <- sub
}
