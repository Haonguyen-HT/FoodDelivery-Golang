package localPb

import (
	"FoodDelivery/common"
	"FoodDelivery/pubsub"
	"context"
	"log"
	"sync"
)

type localPubsub struct {
	messageQueue chan *pubsub.Message
	mapChanel    map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewPubSub() *localPubsub {
	pb := &localPubsub{
		messageQueue: make(chan *pubsub.Message, 10000),
		mapChanel:    make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}

	pb.run()

	return pb
}

func (ps *localPubsub) Publish(
	ctx context.Context,
	topic pubsub.Topic,
	data *pubsub.Message) error {
	data.SetChanel(topic)

	go func() {
		defer common.AppRecover()
		ps.messageQueue <- data
		log.Println("New event published:", data.String(), "with data", data.Data())
	}()

	return nil
}

func (ps *localPubsub) Subscribe(
	ctx context.Context,
	topic pubsub.Topic) (
	ch <-chan *pubsub.Message,
	close func()) {

	c := make(chan *pubsub.Message)

	ps.locker.Lock()

	if val, ok := ps.mapChanel[topic]; ok {
		val = append(ps.mapChanel[topic], c)
		ps.mapChanel[topic] = val
	} else {
		ps.mapChanel[topic] = []chan *pubsub.Message{c}
	}

	ps.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")

		if chans, ok := ps.mapChanel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.mapChanel[topic] = chans
					ps.locker.Unlock()
					break
				}
			}
		}
	}
}

func (ps *localPubsub) run() error {
	log.Println("Pubsub started")

	go func() {
		defer common.AppRecover()

		for {
			mess := <-ps.messageQueue
			log.Println("Message dequeue", mess.String())

			if subs, ok := ps.mapChanel[mess.Chanel()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						defer common.AppRecover()
						c <- mess
					}(subs[i])
				}
			}
		}

	}()

	return nil
}