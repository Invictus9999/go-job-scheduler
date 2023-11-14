package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Invictus9999/go-job-scheduler/proto/gen/scheduler/v1/v1connect"
	"go.uber.org/dig"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	c := dig.New()
	c.Provide(NewHTTPServer)
	c.Provide(NewServeMux)
	c.Provide(NewSchedulerServer)
	c.Provide(NewDBPool)
	c.Provide(NewRedisClient)

	c.Invoke(func(srv *http.Server) {
		idleConnsClosed := make(chan struct{})
		go listenForShutdownSignal(srv, idleConnsClosed)

		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener:
			log.Printf("HTTP server ListenAndServe: %v", err)
		}

		<-idleConnsClosed
	})
}

func NewHTTPServer(mux *http.ServeMux) (*http.Server, error) {
	// Use h2c so we can serve HTTP/2 without TLS.
	srv := &http.Server{Addr: ":8080", Handler: h2c.NewHandler(mux, &http2.Server{})}
	return srv, nil
}

func NewServeMux(schedulerServer *SchedulerServer) *http.ServeMux {
	path, handler := v1connect.NewSchedulerServiceHandler(schedulerServer)

	mux := http.NewServeMux()
	mux.Handle(path, handler)
	return mux
}

func listenForShutdownSignal(srv *http.Server, idleConnsClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)

	// interrupt signal sent from terminal
	signal.Notify(sigint, os.Interrupt)
	// sigterm signal sent from kubernetes
	signal.Notify(sigint, syscall.SIGTERM)

	<-sigint

	// We received an interrupt signal, shut down.
	if err := srv.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("HTTP server Shutdown: %v", err)
	}

	close(idleConnsClosed)
}
