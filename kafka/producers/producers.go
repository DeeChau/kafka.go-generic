package producers

import (
	producers "github.com/DeeChau/kafka.go-generic/internal/producer"
	"github.com/segmentio/kafka-go"
	"github.com/wishabi/kafka.go/avro"
)

// Should not have to provide the schema reg, kafka broker url
// Or topic even. Should be able to get this via. helper function
// Need a better way to store/do configs per topic
func GetFsaProducer(schemaRegistry *avro.Registry, dialer *kafka.Dialer, kafkaBrokerUrl string) *producers.AvroFsaProducer {
	var fsaProducer = producers.NewAvroFsaProducer(kafka.WriterConfig{
		Topic:  "Fsa",
		Dialer: dialer, // just add the dialer here
		Brokers: []string{
			kafkaBrokerUrl,
		},
	}, schemaRegistry)
	return fsaProducer
}
