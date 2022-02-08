package demo

import (
	"context"
	"fmt"
	"net/http"
	"time"

	genericconsumers "github.com/DeeChau/kafka.go-generic/internal/generic_consumer"
	genericschema "github.com/DeeChau/kafka.go-generic/internal/generic_schema"
	"github.com/DeeChau/kafka.go-generic/internal/schema"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/wishabi/kafka.go/avro"
	env "github.com/wishabi/pkg/kafka"
	kafkaUtil "github.com/wishabi/pkg/kafka"
)

// Using generics
func ConsumeWithGenerics() (bool, error) {
	fmt.Println("---Using Generic Kafka Consumers---")
	// Begin Kafka code here
	dialer, dialerErr := kafkaUtil.Dialer()

	if dialerErr != nil {
		log.Error().Msg("Error getting kafka dialer.")
		return false, dialerErr
	}

	schemaRegistry := avro.NewRegistry(env.SchemaRegistry, &http.Client{})

	// fsa Consumer code
	fmt.Println("---Using Generic Kafka Consumers for Fsas---")
	fsaConsumer := genericconsumers.NewAvroConsumer[schema.FsaKey, schema.Fsa](
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
	stateConsumer := genericconsumers.NewAvroConsumer[schema.StateKey, schema.State](
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

// With this pattern, we can use this as a building block to define constraints on what a Avro Key or Value should be
// - and have it generated to redude duplication of code.
// "Hack"(?) that allows us to call methods
func printAvroKey[K any, PT genericschema.AvroSchemaStruct[K]](key K) {
	fmt.Printf("%T Key: %v\n", key, key)
	fmt.Printf("Key Schema: %s\n\n", PT(&key).Schema())
}

func printAvroValue[V any, PT genericschema.AvroSchemaStruct[V]](value V) {
	fmt.Printf("%T Value %v\n", value, value)
	fmt.Printf("Value Schema: %s\n\n", PT(&value).Schema())
}
