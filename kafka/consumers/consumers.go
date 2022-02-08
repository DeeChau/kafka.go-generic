package consumers

import (
	consumers "github.com/DeeChau/kafka.go-generic/internal/consumer"
	"github.com/segmentio/kafka-go"
	"github.com/wishabi/kafka.go/avro"
)

func GetFsaConsumer(schemaRegistry *avro.Registry, dialer *kafka.Dialer, kafkaBrokerUrl string) *consumers.AvroFsaConsumer {
	var fsaConsumer = consumers.NewAvroFsaConsumer(kafka.ReaderConfig{
		Topic:       "Fsa",
		GroupID:     "TestFsaConsumer",
		Dialer:      dialer,
		StartOffset: kafka.FirstOffset,
		Brokers: []string{
			kafkaBrokerUrl,
		},
	}, schemaRegistry)
	return fsaConsumer
}

func GetStateConsumer(schemaRegistry *avro.Registry, dialer *kafka.Dialer, kafkaBrokerUrl string) *consumers.AvroStateConsumer {
	var stateConsumer = consumers.NewAvroStateConsumer(kafka.ReaderConfig{
		Topic:       "Fsa",
		GroupID:     "TestStateConsumer",
		Dialer:      dialer,
		StartOffset: kafka.FirstOffset,
		Brokers: []string{
			kafkaBrokerUrl,
		},
	}, schemaRegistry)
	return stateConsumer
}
