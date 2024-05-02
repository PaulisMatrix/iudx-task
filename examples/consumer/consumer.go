package main

import (
	"context"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"

	kafkamgr "datakaveri/kafka_manager"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func main() {
	// Kafka cluster with one broker.
	mgr, err := kafkamgr.NewKafkaManager(kafkamgr.DefaultSettings())
	if err != nil {
		panic(err)
	}

	consumer := kafkamgr.DefaultConsumerSettings("testing-kafka", 0, kafka.FirstOffset, kafkamgr.Callback(handleMsg))
	consumer.GroupID = "test-consumer-group"
	if err := mgr.AddConsumer(consumer); err != nil {
		logrus.WithError(err).Fatalln("failed to add the consumer")
	}

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-stop
		mgr.Close()
	}()

	// Blocking call
	mgr.Consume()
}

// Callback Handler responsbile for reading the messages
func handleMsg(reader *kafka.Reader, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			if err == io.EOF {
				logrus.WithError(err).Error("done reading messages")
				return
			}
			logrus.WithError(err).Error("failed to fetch message")
			return
		}

		logrus.WithField("msg", string(msg.Value)).Infof("received message from topic[%s] with partition[%d] & offset[%d]", msg.Topic, msg.Partition, msg.Offset)

	}
}
