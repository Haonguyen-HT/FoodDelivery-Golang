package main

import (
	"FoodDelivery/common"
	"FoodDelivery/pubsub"
	"FoodDelivery/pubsub/localPb"
	"context"
	"log"
	"time"
)

func main() {
	var localPs pubsub.Pubsub = localPb.NewPubSub()

	var topic pubsub.Topic = "OrderCreate"

	sub1, close1 := localPs.Subscribe(context.Background(), topic)
	sub2, _ := localPs.Subscribe(context.Background(), topic)

	localPs.Publish(context.Background(), topic, pubsub.NewMessage(1))
	localPs.Publish(context.Background(), topic, pubsub.NewMessage(2))

	go func() {
		defer common.AppRecover()
		for {
			log.Println("Sub1", (<-sub1).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()
	go func() {
		defer common.AppRecover()
		for {
			log.Println("Sub2", (<-sub2).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	time.Sleep(time.Second * 3)
	close1()

	localPs.Publish(context.Background(), topic, pubsub.NewMessage(3))
	time.Sleep(time.Second * 3)
}