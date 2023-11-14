package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Invictus9999/go-job-scheduler/platform/poller"
	"github.com/jackc/puddle/v2"
	"github.com/redis/go-redis/v9"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

const (
	MaxWorkerPoolSize = 5
	StreamName        = "JOBS"
	ConsumerName      = "WORKER"
	SubjectFilter     = "JOBS.scheduled"
)

type streamPollWork struct {
	jc         jetstream.Consumer
	workerPool *puddle.Pool[Handler]
}

func (s *streamPollWork) Execute() error {
	log.Println("Worker Polling")
	ctx := context.Background()

	msgs, err := s.jc.FetchNoWait(int(MaxWorkerPoolSize))

	if err != nil {
		return err
	}

	for msg := range msgs.Messages() {
		log.Println(string(msg.Data()))
		handler, err := s.workerPool.Acquire(context.Background())

		if err != nil {
			msg.Nak()
			continue
		}

		go func(msg jetstream.Msg, handler *puddle.Resource[Handler]) {
			// Release the acquired worker
			defer handler.Release()

			err = handler.Value().Process(string(msg.Data()))

			if err != nil {
				msg.Nak()
			}

			err = msg.DoubleAck(ctx)

			// Optimistically acking once more. Can still fail
			if err != nil {
				msg.Ack()
			}
		}(msg, handler)
	}

	return nil

}

func NewPollerConfig(pollWork poller.PollWork) *poller.PollerConfig {
	return &poller.PollerConfig{
		Freq:     2 * time.Second,
		PollWork: pollWork,
	}
}

func NewPollWork(workerPool *puddle.Pool[Handler], jc jetstream.Consumer) (poller.PollWork, error) {
	return &streamPollWork{
		jc:         jc,
		workerPool: workerPool,
	}, nil
}

func NewJetStreamConsumer() (jetstream.Consumer, error) {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)

	if err != nil {
		return nil, err
	}
	// TODO: Figure out how to shutdown safely
	// defer nc.Drain()

	js, err := jetstream.New(nc)

	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	stream, err := js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:       StreamName,
		Subjects:   []string{"JOBS.*"},
		NoAck:      false,
		MaxMsgs:    -1,
		MaxBytes:   -1,
		MaxAge:     time.Hour * 24 * 365,
		Storage:    jetstream.FileStorage,
		Retention:  jetstream.LimitsPolicy,
		MaxMsgSize: -1,
		Discard:    jetstream.DiscardOld,
	})

	if err != nil {
		return nil, err
	}

	cons, _ := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:       ConsumerName,
		FilterSubject: SubjectFilter,
		AckPolicy:     jetstream.AckExplicitPolicy,
		DeliverPolicy: jetstream.DeliverAllPolicy,
	})

	return cons, nil
}

func NewWorkerPool(rdb *redis.Client) *puddle.Pool[Handler] {
	constructor := func(context.Context) (Handler, error) {
		return &handler{rdb: rdb}, nil
	}
	destructor := func(value Handler) {
		return
	}
	maxPoolSize := int32(MaxWorkerPoolSize)

	pool, err := puddle.NewPool(&puddle.Config[Handler]{Constructor: constructor, Destructor: destructor, MaxSize: maxPoolSize})
	if err != nil {
		log.Fatal(err)
	}

	return pool
}
