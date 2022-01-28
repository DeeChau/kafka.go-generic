package producers

import (
	producers "github.com/DeeChau/kafka.go-generic/internal/producer"
	"github.com/segmentio/kafka-go"
	"github.com/wishabi/kafka.go/avro"
)

/*
	Duplicated code per producer AND consumer ... this is not ideal
	Also Should not have to provide the schema reg, kafka broker url
	Or even a topic.
	Should be able to get this via. helper function
	Need a better way to store/do configs per topic!
*/
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

func GetStateProducer(schemaRegistry *avro.Registry, dialer *kafka.Dialer, kafkaBrokerUrl string) *producers.AvroStateProducer {
	var stateProducer = producers.NewAvroStateProducer(kafka.WriterConfig{
		Topic:  "State",
		Dialer: dialer, // just add the dialer here
		Brokers: []string{
			kafkaBrokerUrl,
		},
	}, schemaRegistry)
	return stateProducer
}
