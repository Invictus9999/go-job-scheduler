package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Invictus9999/go-job-scheduler/platform/poller"
	"go.uber.org/dig"
)

func main() {
	c := dig.New()
	c.Provide(poller.NewPoller)
	c.Provide(NewPollerConfig)
	c.Provide(NewPollWork)
	c.Provide(NewJetStreamConsumer)
	c.Provide(NewWorkerPool)
	c.Provide(NewRedisClient)

	log.Println("List of Providers and their dependencies")
	log.Println(c.String())

	err := c.Invoke(func(poller poller.Poller) {
		idleConnsClosed := make(chan struct{})
		go listenForShutdownSignal(poller, idleConnsClosed)

		poller.Start()
		<-idleConnsClosed
	})

	if err != nil {
		log.Println(err)
	}

	log.Println("Bye bye from main")

}

func listenForShutdownSignal(poller poller.Poller, idleConnsClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)

	// interrupt signal sent from terminal
	signal.Notify(sigint, os.Interrupt)
	// sigterm signal sent from kubernetes
	signal.Notify(sigint, syscall.SIGTERM)

	<-sigint

	// We received an interrupt signal, shut down.
	poller.Stop()

	close(idleConnsClosed)
}
