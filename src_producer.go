package main

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"datakaveri/events"
	kafkamgr "datakaveri/kafka_manager"

	"github.com/sirupsen/logrus"
)

var myEvents = []events.EventType{events.PageViewEvent, events.ButtonClickEvent, events.SearchEvent, events.PurchaseEvent, events.AdvertEvent, events.ScrollEvent}

func main() {
	// Tickers to randomly keep generating the events
	interval1 := float64(1000)
	interval2 := float64(500)
	interval3 := float64(1500)
	interval4 := float64(2000)
	interval5 := float64(2500)

	ticker1 := time.NewTicker(time.Duration(interval1) * time.Millisecond)
	ticker2 := time.NewTicker(time.Duration(interval2) * time.Millisecond)
	ticker3 := time.NewTicker(time.Duration(interval3) * time.Millisecond)
	ticker4 := time.NewTicker(time.Duration(interval4) * time.Millisecond)
	ticker5 := time.NewTicker(time.Duration(interval5) * time.Millisecond)

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// init kafka manager
	parentCtx := context.Background()
	manager, err := kafkamgr.NewKafkaManager(kafkamgr.DefaultSettings())
	if err != nil {
		panic(err)
	}

	bufArray := make([][]byte, 0)
	dataChan := make(chan []byte)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(manager *kafkamgr.KafkaManager) {

		for d := range dataChan {
			if len(bufArray) == 5 {
				// Drain the buf and produce the messages
				ctx, cancel := context.WithTimeout(parentCtx, 2*time.Second)
				_, err := manager.Produce(ctx, "data-kaveri", "", bufArray)
				if err != nil {
					logrus.WithError(err).Error("error in producing the event")
				}
				cancel()

				bufArray = nil
			}

			bufArray = append(bufArray, d)
		}

		manager.Close()
		wg.Done()

	}(manager)

	for {
		select {
		case <-stop:
			ticker1.Stop()
			ticker2.Stop()
			ticker3.Stop()
			ticker4.Stop()
			ticker5.Stop()

			close(dataChan)

			// Wait for the manager to close all producers and consumers
			wg.Wait()

		case <-ticker1.C:
			source := rand.NewSource(time.Now().Unix())
			r := rand.New(source)
			data := FakerProduce(myEvents[r.Intn(len(myEvents))])

			dataChan <- data

		case <-ticker2.C:
			source := rand.NewSource(time.Now().Unix())
			r := rand.New(source)
			data := FakerProduce(myEvents[r.Intn(len(myEvents))])

			dataChan <- data

		case <-ticker3.C:
			source := rand.NewSource(time.Now().Unix())
			r := rand.New(source)
			data := FakerProduce(myEvents[r.Intn(len(myEvents))])

			dataChan <- data

		case <-ticker4.C:
			source := rand.NewSource(time.Now().Unix())
			r := rand.New(source)
			data := FakerProduce(myEvents[r.Intn(len(myEvents))])

			dataChan <- data

		case <-ticker5.C:
			source := rand.NewSource(time.Now().Unix())
			r := rand.New(source)
			data := FakerProduce(myEvents[r.Intn(len(myEvents))])

			dataChan <- data
		}
	}
}
