package asyncjob

import (
	"context"
	"time"
)

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDurations(times []time.Duration)
}

const (
	defaultMaxTimeout = time.Second * 10
)

var (
	defaultRetryTime = []time.Duration{time.Second, time.Second * 2, time.Second * 4}
)

type JobHandler func(ctx context.Context) error
type JobState int

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

func (j JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[j]
}

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(handler JobHandler) *job {
	j := job{
		config: jobConfig{
			MaxTimeout: defaultMaxTimeout,
			Retries:    defaultRetryTime,
		},
		handler:    handler,
		state:      StateInit,
		retryIndex: -1,
		stopChan:   make(chan bool),
	}

	return &j
}

func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning

	var err error
	err = j.handler(ctx)
	if err != nil {
		j.state = StateFailed
		return err
	}

	j.state = StateCompleted
	return nil

}
func (j *job) Retry(ctx context.Context) error {
	j.retryIndex += 1

	time.Sleep(j.config.Retries[j.retryIndex])

	err := j.Execute(ctx)

	if err == nil {
		j.state = StateCompleted
		return nil
	}

	if j.retryIndex == len(j.config.Retries)-1 {
		j.state = StateRetryFailed
		return err
	}

	j.state = StateFailed
	return err

}

func (j *job) State() JobState {
	return j.state
}

func (j *job) RetryIndex() int {
	return j.retryIndex
}

func (j *job) SetRetryDurations(times []time.Duration) {
	if len(times) == 0 {
		return
	}

	j.config.Retries = times
}