package main

import (
	"desent/src/bootstrap"
	"desent/src/pkg/customvalidator"
	"desent/src/routes"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	app := bootstrap.NewApplication()

	e := echo.New()
	e.Validator = customvalidator.NewCustomValidator()

	routes.InitRoutes(e, app)

	wg := sync.WaitGroup{}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Initiate gracefully shutdown with exit signal")
		waitTimeout(&wg, 10*time.Second)
		log.Println("Gracefully shutting down...")
		e.Close()
	}()

	if err := e.Start(":9999"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("ERR Echo = %+v", err)
	}
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()
	select {
	case <-done:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
