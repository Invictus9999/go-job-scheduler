package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	dbScheduler "github.com/Invictus9999/go-job-scheduler/db/sqlc/scheduler"
	schedulerv1 "github.com/Invictus9999/go-job-scheduler/proto/gen/scheduler/v1"
)

type SchedulerServer struct {
	dbpool *pgxpool.Pool
	rdb    *redis.Client
}

func (s *SchedulerServer) Schedule(
	ctx context.Context,
	req *connect.Request[schedulerv1.ScheduleRequest],
) (*connect.Response[schedulerv1.ScheduleResponse], error) {
	job, err := createJob(ctx, req.Msg, s.dbpool)

	if err != nil {
		return nil, err
	}

	jobId, err := storeAndScheduleInRedis(ctx, s.rdb, job)

	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&schedulerv1.ScheduleResponse{
		JobId: jobId.String(),
	})

	log.Printf("Job created id: %s", jobId.String())
	res.Header().Set("Scheduler-Version", "v1")
	return res, nil
}

type RedisJob struct {
	JobType      dbScheduler.Jobtype
	Payload      []byte
	Frequency    int32
	NextRun      time.Time
	TimeoutAfter int32
}

func storeAndScheduleInRedis(ctx context.Context, rdb *redis.Client, job *dbScheduler.Job) (*uuid.UUID, error) {
	jobId, err := uuid.FromBytes(job.ID.Bytes[:])
	if err != nil {
		return nil, err
	}

	rJob := RedisJob{
		JobType:      job.JobType,
		Payload:      job.Payload,
		Frequency:    job.Frequency.Int32,
		NextRun:      job.NextRun.Time,
		TimeoutAfter: job.TimeoutAfter,
	}
	bytes, _ := json.Marshal(rJob)
	value := string(bytes)
	log.Println(value)
	rdb.Set(ctx, jobId.String(), value, 0)
	rdb.ZAdd(ctx, "ScheduledJobs", redis.Z{Score: float64(job.NextRun.Time.Unix()), Member: jobId.String()})

	return &jobId, nil
}

func (s *SchedulerServer) Status(
	ctx context.Context,
	req *connect.Request[schedulerv1.StatusRequest],
) (*connect.Response[schedulerv1.StatusResponse], error) {
	job, err := getJobById(ctx, req.Msg, s.dbpool)

	if err != nil {
		return nil, err
	}

	jobId, err := uuid.FromBytes(job.ID.Bytes[:])

	if err != nil {
		return nil, err
	}

	// Unmarshall Payload to the correct type
	var simpleJobPayload schedulerv1.SimpleJobPayload
	proto.Unmarshal(job.Payload, &simpleJobPayload)

	// handle err from unmarshalling

	res := connect.NewResponse(&schedulerv1.StatusResponse{
		JobDetails: &schedulerv1.JobDetails{
			Id:      jobId.String(),
			EmailId: job.EmailID,
			NextRun: timestamppb.New(job.NextRun.Time),
			Payload: &schedulerv1.JobDetails_SimpleJobPayload{
				SimpleJobPayload: &simpleJobPayload,
			},
		},
	})

	res.Header().Set("Scheduler-Version", "v1")
	return res, nil
}

func createJob(ctx context.Context, scheduleReq *schedulerv1.ScheduleRequest, dbpool *pgxpool.Pool) (*dbScheduler.Job, error) {
	queries := dbScheduler.New(dbpool)

	id := uuid.New()

	// Serialize the Person message to a byte slice
	data, err := proto.Marshal(scheduleReq.GetSimpleJobPayload())
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// create job
	job, err := queries.CreateJob(ctx, dbScheduler.CreateJobParams{
		ID:        pgtype.UUID{Bytes: id, Valid: true},
		EmailID:   scheduleReq.EmailId,
		JobType:   "Simple",
		Payload:   data,
		Recurring: scheduleReq.Recurring,
		Frequency: pgtype.Int4{
			Int32: int32(scheduleReq.Frequency.Seconds),
			Valid: true,
		},
		NextRun: pgtype.Timestamp{
			Time:  scheduleReq.Start.AsTime(),
			Valid: true,
		},
		TimeoutAfter: int32(scheduleReq.TimeoutAfter.Seconds),
		JobStatus:    dbScheduler.JobstatusScheduled,
	})

	if err != nil {
		return nil, err
	}
	log.Println(job)

	return &job, nil
}

func getJobById(ctx context.Context, statusReq *schedulerv1.StatusRequest, dbpool *pgxpool.Pool) (*dbScheduler.Job, error) {
	queries := dbScheduler.New(dbpool)

	id, err := uuid.Parse(statusReq.JobId)

	if err != nil {
		return nil, err
	}

	// get job by id
	job, err := queries.GetJobById(ctx, pgtype.UUID{Bytes: id, Valid: true})

	if err != nil {
		return nil, err
	}

	log.Println(job)

	return &job, nil
}

func NewSchedulerServer(dbpool *pgxpool.Pool, rdb *redis.Client) *SchedulerServer {
	return &SchedulerServer{
		dbpool: dbpool,
		rdb:    rdb,
	}
}

func NewDBPool() (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), os.Getenv("POSTGRESQL_URL"))
}

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
