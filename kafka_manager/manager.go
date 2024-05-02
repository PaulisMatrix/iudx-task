package kafkamanager

import (
	"bytes"
	"context"
	"errors"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var logger = logrus.StandardLogger()

// Manager Settings
type Settings struct {
	Brokers         []string // List of brokers to connect. currently it would be localhost:9092
	ProducerSetting *ProducerSettings
}

// Topic specific settings
type TopicConfig struct {
	Topic             string
	TopicPartitions   int
	ReplicationFactor int
}

// Writer specific settings
type ProducerSettings struct {
}

// Callback func to read the messages
type Callback func(*kafka.Reader, *sync.WaitGroup)

// Reader specific settings
type ConsumerSettings struct {
	Topic          string
	Partition      int
	GroupID        string
	Callback       Callback
	Offset         int64
	ReadBackoffMin time.Duration
	ReadBackoffMax time.Duration
	MinBytes       int
	MaxBytes       int
}

type KafkaManager struct {
	sync.Mutex
	wg              *sync.WaitGroup
	producer        *kafka.Writer
	consumers       map[string]*kafka.Reader // Topic to consumer map. Supports multiple consumers
	ConsumerSetting []*ConsumerSettings
	settings        *Settings
}

func DefaultSettings() *Settings {
	return &Settings{
		Brokers:         []string{"localhost:9092"},
		ProducerSetting: DefaultProducerSettings(),
	}
}

// Exeriment latter with different producer settings
func DefaultProducerSettings() *ProducerSettings {
	return &ProducerSettings{}
}

func DefaultConsumerSettings(topic string, partition int, startOffset int64, callback Callback) *ConsumerSettings {
	return &ConsumerSettings{
		Topic:          topic,
		Callback:       callback,
		Offset:         startOffset,
		ReadBackoffMin: 100 * time.Millisecond,
		ReadBackoffMax: time.Second,
		MinBytes:       1,
		MaxBytes:       1e6, // 1 MB
	}
}

func NewKafkaManager(s *Settings) (*KafkaManager, error) {
	manager := &KafkaManager{
		Mutex:           sync.Mutex{},
		wg:              &sync.WaitGroup{},
		consumers:       make(map[string]*kafka.Reader),
		ConsumerSetting: make([]*ConsumerSettings, 0),
		settings:        s,
	}
	manager.producer = manager.getNewWriter()

	return manager, nil
}

func (m *KafkaManager) getNewWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(m.settings.Brokers...),
		Balancer:     kafka.CRC32Balancer{},
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Async:        false,
	}
}

func (m *KafkaManager) getNewReader(c *ConsumerSettings) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Topic:             c.Topic,
		GroupID:           c.GroupID,
		Brokers:           m.settings.Brokers,
		StartOffset:       c.Offset,
		ReadBackoffMin:    c.ReadBackoffMin,
		ReadBackoffMax:    c.ReadBackoffMax,
		MinBytes:          c.MinBytes,
		MaxBytes:          c.MaxBytes,
		HeartbeatInterval: 4 * time.Second,
	})
}
func (m *KafkaManager) Produce(ctx context.Context, topic, key string, msgList [][]byte) (bool, error) {
	// Incase produce is called from multiple places
	m.Lock()

	var err error
	{
		messages := []kafka.Message{}

		for _, msg := range msgList {
			messages = append(messages, kafka.Message{
				Topic: topic,
				Key:   bytes.NewBufferString(key).Bytes(),
				Value: msg,
			})
		}

		err = m.producer.WriteMessages(ctx, messages...)
	}

	m.Unlock()

	if err != nil {
		logger.WithError(err).Error("failed to write to the broker.")
		return false, err
	}

	return true, nil
}

func (m *KafkaManager) AddConsumer(c *ConsumerSettings) error {
	m.ConsumerSetting = append(m.ConsumerSetting, c)

	_, ok := m.consumers[c.Topic]
	if ok {
		return errors.New("consumer already initialised")
	}

	m.consumers[c.Topic] = m.getNewReader(c)
	return nil
}

// Consume starts consumers
// It blocks until all the consumers are done consuming message
func (m *KafkaManager) Consume() {
	for _, c := range m.ConsumerSetting {
		reader, ok := m.consumers[c.Topic]
		if ok {
			m.wg.Add(1)
			go c.Callback(reader, m.wg)
		}
	}

	m.wg.Wait()
}

// Close closes all the open connections from producer & consumers
func (m *KafkaManager) Close() {

	if err := m.producer.Close(); err != nil {
		logger.WithError(err).Error("failed to close producer")
	}

	for _, c := range m.consumers {
		if err := c.Close(); err != nil {
			logger.WithError(err).WithField("topic", c.Config().Topic).Error("failed to close consumer")
		}
	}

}
