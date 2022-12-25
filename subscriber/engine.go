package subscriber

import (
	"FoodDelivery/common"
	"FoodDelivery/components/appcontext"
	"FoodDelivery/components/asyncjob"
	"FoodDelivery/pubsub"
	"context"
	"log"
)

type consumerJob struct {
	Title string
	Hdl   func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx appcontext.AppContext
}

func NewEngine(appCtx appcontext.AppContext) *consumerEngine {
	return &consumerEngine{appCtx: appCtx}
}

func (engine *consumerEngine) Start() error {
	engine.startSubTopic(
		common.TopicUserLikeRestaurant,
		true,
		IncreaseLikeCountAfterUserLikeRestaurant(engine.appCtx),
	)

	engine.startSubTopic(
		common.TopicUserDisLikeRestaurant,
		true,
		DecreaseLikeCountAfterUserLikeRestaurant(engine.appCtx),
		PushNotificationWhenUserDisLikeRestaurant(engine.appCtx),
	)

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(
	topic pubsub.Topic,
	isConcurrent bool,
	consumerJobs ...consumerJob) error {

	c, _ := engine.appCtx.GetPubSub().Subscribe(context.Background(), topic)

	for _, item := range consumerJobs {
		log.Println("Setup consumer for: ", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("running job for: ", job.Title, ". Value: ", message.Data())
			return job.Hdl(ctx, message)
		}
	}

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c

			jobHdlArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], msg)
				jobHdlArr[i] = asyncjob.NewJob(jobHdl)
			}

			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}