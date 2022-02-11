package demo

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/DeeChau/kafka.go-generic/internal/schema"
	generic_avro "github.com/wishabi/kafka.go/generic_avro"
	generic_consumer "github.com/wishabi/kafka.go/generic_consumer"
	generic_producer "github.com/wishabi/kafka.go/generic_producer"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/wishabi/kafka.go/avro"
	env "github.com/wishabi/pkg/kafka"
	kafkaUtil "github.com/wishabi/pkg/kafka"
)

// Using generics
func ConsumeWithGenerics() (bool, error) {
	fmt.Println("---Using Generic Kafka Consumers---")
	dialer, dialerErr := kafkaUtil.Dialer()

	if dialerErr != nil {
		log.Error().Msg("Error getting kafka dialer.")
		return false, dialerErr
	}

	schemaRegistry := avro.NewRegistry(env.SchemaRegistry, &http.Client{})

	// fsa Consumer code
	fmt.Println("---Using Generic Kafka Consumers for Fsas---")
	fsaConsumer := generic_consumer.NewAvroConsumer[schema.FsaKey, schema.Fsa](
		kafka.ReaderConfig{
			Topic:   "Fsa",
			GroupID: "TestConsumerGenerics",
			Dialer:  dialer,
			Brokers: env.BrokersList,
		}, schemaRegistry)

	log.Info().Msgf("Fsa Consumer Initialized: %s\n", fsaConsumer)

	ctx := context.Background()
	fsaMsg, fsaErr := fsaConsumer.AutoCommitConsume(ctx)
	if fsaErr != nil {
		log.Error().Err(fsaErr).Msg("Kafka message could not be received. Error consuming message.")
		return false, fsaErr
	}

	fsaKey := fsaMsg.Key
	fsaValue := fsaMsg.Value
	log.Info().Msgf("Fsa -- Kafka: Received Key %v, Message %v from", fsaKey, fsaValue)

	// State Consumer code
	fmt.Println("---Using Generic Consumers for States---")
	stateConsumer := generic_consumer.NewAvroConsumer[schema.StateKey, schema.State](
		kafka.ReaderConfig{
			Topic:   "State",
			GroupID: "TestConsumerGenerics",
			Dialer:  dialer,
			Brokers: env.BrokersList,
		}, schemaRegistry)

	log.Info().Msgf("State Consumer Initialized: %s\n", stateConsumer)

	ctx = context.Background()
	stateMsg, stateErr := stateConsumer.AutoCommitConsume(ctx)
	if stateErr != nil {
		log.Error().Err(stateErr).Msg("Kafka message could not be received. Error consuming message.")
		return false, stateErr
	}

	stateKey := stateMsg.Key
	stateValue := stateMsg.Value
	log.Info().Msgf("State -- Kafka: Received Key %v, Message %v from", stateKey, stateValue)

	return true, nil
}

func ProduceWithGenerics() (bool, error) {
	fmt.Println("---Using Generic Kafka Producers---")

	// fsa producer code
	dialer, dialerErr := kafkaUtil.Dialer()
	if dialerErr != nil {
		log.Error().Msg("Error getting kafka dialer.")
		return false, dialerErr
	}

	fsa := "M6G"
	lat := float32(88.88)
	long := float32(99.99)
	key := generateFsaKey(fsa)
	payload := generateFsa(fsa, lat, long)

	log.Info().Msgf("Key to send %v, Message to send %v", key, payload)
	log.Printf("Schema Registry location %v, Kafka Broker location %v", env.SchemaRegistry, env.BrokersList)

	schemaRegistry := avro.NewRegistry(env.SchemaRegistry, &http.Client{})

	fsaProducer := generic_producer.NewAvroProducer[schema.FsaKey, schema.Fsa](kafka.WriterConfig{
		Topic:   "Fsa",
		Dialer:  dialer,
		Brokers: env.BrokersList,
	}, schemaRegistry)

	err := fsaProducer.Produce(
		context.Background(),
		&key,
		&payload)
	if err != nil {
		log.Error().Err(err).Msg("Kafka message could not be sent. Error producing message.")
		return false, err
	}

	log.Info().Msg("FSA kafka message sent successfully.")
	return true, nil
}

// Experimental work with generics
func ExperimentWithGenerics() {
	fmt.Println("---Experimenting with Generic Types and Methods---")
	timestamp := time.Now()

	fsa_key := schema.FsaKey{Label: "M5V"}
	fsa_value := schema.Fsa{
		Label:      "M5V",
		Latitude:   12.34,
		Longitude:  56.78,
		Message_id: timestamp.Format(time.UnixDate),
		Created_at: timestamp.Unix(),
		Updated_at: timestamp.Unix(),
		Timestamp:  timestamp.Format(time.UnixDate),
	}
	state_key := schema.StateKey{Abbreviation: "ON"}
	state_value := schema.State{
		Abbreviation: "ON",
		Name:         "Ontario",
		Message_id:   timestamp.Format(time.UnixDate),
		Timestamp:    timestamp.Format(time.UnixDate),
	}

	printAvroKey(fsa_key)
	printAvroKey(state_key)
	printAvroValue(fsa_value)
	printAvroValue(state_value)
}

// With this pattern, we can use this as a building block to define constraint types for our Avro Key or Values
// and reduce duplication of code. Not sure if this is a "Hack" that allows us to call methods.
func printAvroKey[K any, PT generic_avro.AvroSchemaStruct[K]](key K) {
	fmt.Printf("%T Key: %v\n", key, key)
	fmt.Printf("Key Schema: %s\n\n", PT(&key).Schema())
}

func printAvroValue[V any, PT generic_avro.AvroSchemaStruct[V]](value V) {
	fmt.Printf("%T Value %v\n", value, value)
	fmt.Printf("Value Schema: %s\n\n", PT(&value).Schema())
}
