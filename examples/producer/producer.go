package main

import (
	"context"
	"time"

	kafkamgr "datakaveri/kafka_manager"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	mgr, err := kafkamgr.NewKafkaManager(kafkamgr.DefaultSettings())
	if err != nil {
		panic(err)
	}

	for i := 0; i < 3; i++ {
		ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
		_, err = mgr.Produce(ctx, "testing-kafka", "datakaveri", [][]byte{[]byte("hello")})
		cancel()

		if err != nil {
			logrus.WithError(err).Error("failed to produce")
			continue
		}

		logrus.Info("produced a msg to the given topic testing-kafka")
	}

	// Shutdown all the producers and consumers
	mgr.Close()
}
