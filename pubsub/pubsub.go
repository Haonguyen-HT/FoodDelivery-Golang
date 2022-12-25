package pubsub

import "context"

type Topic string

type Pubsub interface {
	Publish(ctx context.Context, chanel Topic, data *Message) error
	Subscribe(ctx context.Context, chanel Topic) (ch <-chan *Message, close func())
}