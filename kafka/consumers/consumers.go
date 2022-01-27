package consumers

import (
	consumers "github.com/DeeChau/kafka.go-generic/internal/consumer"
	"github.com/segmentio/kafka-go"
	"github.com/wishabi/kafka.go/avro"
)

// Should not have to provide the schema reg, kafka broker url
// Or topic even. Should be able to get this via. helper function
// Need a better way to store/do configs per topic
func GetFsaConsumer(schemaRegistry *avro.Registry, dialer *kafka.Dialer, kafkaBrokerUrl string) *consumers.AvroFsaConsumer {
	var fsaConsumer = consumers.NewAvroFsaConsumer(kafka.ReaderConfig{
		Topic:       "Fsa",
		GroupID:     "TestFsaConsumer",
		Dialer:      dialer,
		StartOffset: kafka.LastOffset,
		Brokers: []string{
			kafkaBrokerUrl,
		},
	}, schemaRegistry)
	return fsaConsumer
}
