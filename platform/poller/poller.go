package poller

import (
	"log"
	"time"
)

type poller struct {
	ticker     *time.Ticker
	tickerDone chan bool
	pollerDone chan bool
	pollWork   PollWork
}

type Poller interface {
	Start()
	Stop()
}

type PollWork interface {
	Execute() error
}

type PollerConfig struct {
	Freq     time.Duration
	PollWork PollWork
}

func (r *poller) Start() {
	log.Println("Starting Poller")

	go func() {
		for {
			select {
			case <-r.tickerDone:
				log.Println("Ticker Stopped")
				r.pollerDone <- true
				return
			case <-r.ticker.C:
				//log.Println("Tick at", t)
				r.poll()
			}
		}
	}()
}

func (r *poller) Stop() {
	r.ticker.Stop()
	r.tickerDone <- true
	<-r.pollerDone

	log.Println("Poller Stopped")
}

func (r *poller) poll() {
	err := r.pollWork.Execute()

	if err != nil {
		log.Println(err)
	}
}

func NewPoller(pollConfig *PollerConfig) (Poller, error) {
	return &poller{
		ticker:     time.NewTicker(pollConfig.Freq),
		tickerDone: make(chan bool),
		pollerDone: make(chan bool),
		pollWork:   pollConfig.PollWork,
	}, nil
}
