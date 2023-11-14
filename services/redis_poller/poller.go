package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/Invictus9999/go-job-scheduler/platform/poller"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

const (
	GetAndRemoveLuaScriptString = `
	local key = KEYS[1]
	local min = ARGV[1]
	local max = ARGV[2]

	local value = redis.call("ZRANGE", key, min, max, "BYSCORE")
	redis.call("ZREMRANGEBYSCORE", key, min, max)

	return value
	`
	SortedSetKey      = "ScheduledJobs"
	SortedSetRangeMin = "-inf"
	StreamName        = "JOBS"
	SubjectName       = "JOBS.scheduled"
)

type redisPollWork struct {
	rdb                   *redis.Client
	getAndRemoveLuaScript *redis.Script
	js                    jetstream.JetStream
}

func (p *redisPollWork) Execute() error {
	ctx := context.Background()

	min := SortedSetRangeMin
	max := strconv.FormatFloat(float64(time.Now().Unix()), 'f', -1, 64)

	keys := []string{SortedSetKey}
	values := []interface{}{min, max}
	val, err := p.getAndRemoveLuaScript.Run(ctx, p.rdb, keys, values...).StringSlice()

	if err != nil {
		return err
	}

	log.Println(val)
	return p.publishToStream(ctx, val)
}

func (p *redisPollWork) publishToStream(ctx context.Context, jobsIDs []string) error {
	for _, jobID := range jobsIDs {
		ack, err := p.js.Publish(ctx, SubjectName, []byte(jobID))

		if err != nil {
			log.Printf("Error while publishing %s", err)
		} else {
			log.Printf("Publish Ack %s", ack)
		}
	}

	// TODO: return partial error
	return nil
}

func NewPollerConfig(pollWork poller.PollWork) *poller.PollerConfig {
	return &poller.PollerConfig{
		Freq:     2 * time.Second,
		PollWork: pollWork,
	}
}

func NewPollWork(rdb *redis.Client, js jetstream.JetStream) (poller.PollWork, error) {
	getAndRemoveLuaScript := redis.NewScript(GetAndRemoveLuaScriptString)

	val, err := loadLuaScript(getAndRemoveLuaScript, rdb)
	log.Printf("Script hash is %s", val)

	if err != nil {
		return nil, err
	}

	return &redisPollWork{
		rdb:                   rdb,
		getAndRemoveLuaScript: getAndRemoveLuaScript,
		js:                    js,
	}, nil
}

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}

func NewJetStreamClient() (jetstream.JetStream, error) {
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

	_, err = js.CreateOrUpdateStream(context.Background(), jetstream.StreamConfig{
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
		log.Println(err)
		return nil, err
	}

	return js, nil
}

func loadLuaScript(getAndRemoveLuaScript *redis.Script, rdb *redis.Client) (string, error) {
	return getAndRemoveLuaScript.Load(context.Background(), rdb).Result()
}
