package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/DeeChau/kafka.go-generic/internal/schema"
	"github.com/DeeChau/kafka.go-generic/internal/generic_consumer"
	genericschema "github.com/DeeChau/kafka.go-generic/internal/generic_schema"

	"github.com/DeeChau/kafka.go-generic/kafka/consumers"
	"github.com/DeeChau/kafka.go-generic/kafka/producers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/wishabi/kafka.go/avro"
	env "github.com/wishabi/pkg/kafka"
	kafkaUtil "github.com/wishabi/pkg/kafka"
)

func main() {
	fmt.Println("---Begin Hackathon for Generic Kafka experimentation---")
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	useGenerics := len(os.Args) >= 2 && (os.Args[1] == "generic")

	if useGenerics {
		// experimentWithGenerics()
		consumeWithGenerics()
	} else {
		publishAndConsumeNoGenerics()
	}
}

// Using generics
func consumeWithGenerics() (bool, error) {
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
			Topic:       "Fsa",
			GroupID:     "TestConsumerGenerics",
			Dialer:      dialer,
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
			Topic:       "State",
			GroupID:     "TestConsumerGenerics",
			Dialer:      dialer,
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

func experimentWithGenerics() {
	fmt.Println("---Experimenting with Generic Types and Methods---")
	timestamp := time.Now()

	fsa_key := schema.FsaKey{Label: "M5V"}
	fsa_value := schema.Fsa{
		Label: "M5V",
		Latitude: 12.34,
		Longitude: 56.78,
		Message_id: timestamp.Format(time.UnixDate),
		Created_at: timestamp.Unix(),
		Updated_at: timestamp.Unix(),
		Timestamp:  timestamp.Format(time.UnixDate),
	}
	state_key := schema.StateKey{Abbreviation: "ON"}
	state_value := schema.State{
		Abbreviation: "ON",
		Name: "Ontario",
		Message_id: timestamp.Format(time.UnixDate),
		Timestamp:  timestamp.Format(time.UnixDate),
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

// Not using generics
func publishAndConsumeNoGenerics() {
	fmt.Println("---Not using Generic Kafka Consumers/Producers---")
	_, err := publishFsa("M5V", 43.6450946, -79.3929024)
	if err != nil {
		log.Info().Msg("Some error occurred while producing to Kafka")
	} else {
		log.Info().Msg("Producers successfully finished")
	}

	_, err = publishState("ON", "Ontario")
	if err != nil {
		log.Info().Msg("Some error occurred while producing to Kafka")
	} else {
		log.Info().Msg("Producers successfully finished")
	}

	_, err = consumeFsas()
	if err != nil {
		log.Info().Msg("Some error occurred while consuming from Kafka")
	} else {
		log.Info().Msg("Consumer successfully finished")
	}
}

// Consumer code.
// Currently not being run with too much - in the future we need to
// investigate a platform for this to run within a GoRoutine.
func consumeFsas() (bool, error) {
	log.Info().Msg("Create consumers")

	dialer, dialerErr := kafkaUtil.Dialer()

	if dialerErr != nil {
		log.Error().Msg("Error getting kafka dialer.")
		return false, dialerErr
	}

	schemaRegistry := avro.NewRegistry(env.SchemaRegistry, &http.Client{})

	consumer := consumers.GetFsaConsumer(schemaRegistry, dialer, env.BrokersList[0])

	ctx := context.Background()
	msg, err := consumer.AutoCommitConsume(ctx)
	if err != nil {
		return false, err
	}

	key := msg.Key
	value := msg.Value
	log.Info().Msgf("Kafka: Received Key %v, Message %v", key, value)

	return true, nil
}

// Producer code.
func publishFsa(fsa string, lat, long float32) (bool, error) {
	log.Info().Msg("Create Producers")

	dialer, dialerErr := kafkaUtil.Dialer()

	if dialerErr != nil {
		log.Error().Msg("Error getting kafka dialer.")
		return false, dialerErr
	}

	key := generateFsaKey(fsa)
	payload := generateFsa(fsa, lat, long)

	log.Info().Msgf("Key to send %v, Message to send %v", key, payload)
	log.Printf("Schema Registry location %v, Kafka Broker location %v", env.SchemaRegistry, env.BrokersList)

	err := produceFsaMessage(dialer, &key, &payload)
	if err != nil {
		log.Error().Err(err).Msg("Kafka message could not be sent. Error producing message.")
		return false, err
	}

	log.Info().Msg("FSA kafka message sent successfully.")
	return true, nil
}

func publishState(abbreviation, name string) (bool, error) {
	log.Info().Msg("Create Producers")

	dialer, dialerErr := kafkaUtil.Dialer()

	if dialerErr != nil {
		log.Error().Msg("Error getting kafka dialer.")
		return false, dialerErr
	}

	key := generateStateKey(abbreviation)
	payload := generateState(abbreviation, name)

	log.Info().Msgf("Key to send %v, Message to send %v", key, payload)
	log.Printf("Schema Registry location %v, Kafka Broker location %v", env.SchemaRegistry, env.BrokersList)

	err := produceStateMessage(dialer, &key, &payload)
	if err != nil {
		log.Error().Err(err).Msg("Kafka message could not be sent. Error producing message.")
		return false, err
	}

	log.Info().Msg("State kafka message sent successfully.")
	return true, nil
}


/* Notes
Implementation & Use 	- These should probably be abstracted and made easier - it should be possible within the producer's library.
Linting 			 	- These keys should be standardized with Camel Case or snake case - not both.
Implementation & Use 	- Utilize Type Switch to take in a Schema Type -> Translate to the Configured Kafka Topic ID - PENDING on discussions from golang maintainers
Implementation & Use 	- Share one registry across? Is that a good idea.
Implementation & Use 	- Utilize Generics
*/
// - Begin FSA Producer helper code - //
func generateFsaKey(fsa string) schema.FsaKey {
	return schema.FsaKey{Label: fsa}
}

func generateFsa(fsa string, lat, long float32) schema.Fsa {
	
	timestamp := time.Now()
	payload := schema.Fsa{
		Label:      fsa,
		Latitude:   lat,
		Longitude:  long,
		Message_id: fmt.Sprintf("FSA-%v", timestamp),
		Created_at: timestamp.Unix(),
		Updated_at: timestamp.Unix(),
		Timestamp:  timestamp.Format(time.UnixDate),
	}

	return payload
}

func produceFsaMessage(dialer *kafka.Dialer, key *schema.FsaKey, payload *schema.Fsa) error {
	
	schemaRegistry := avro.NewRegistry(env.SchemaRegistry, &http.Client{})
	return producers.GetFsaProducer(schemaRegistry, dialer, env.BrokersList[0]).Produce(
		context.Background(),
		key,
		payload)
}

// - Begin State Producer helper code - //
func generateStateKey(abbreviation string) schema.StateKey {
	return schema.StateKey{Abbreviation: abbreviation}
}

func generateState(abbreviation, name string) schema.State {
	timestamp := time.Now()
	payload := schema.State{
		Abbreviation: abbreviation,
		Name:   name,
		Message_id: fmt.Sprintf("State-%v", timestamp),
		Timestamp:  timestamp.Format(time.UnixDate),
	}

	return payload
}

func produceStateMessage(dialer *kafka.Dialer, key *schema.StateKey, payload *schema.State) error {
	schemaRegistry := avro.NewRegistry(env.SchemaRegistry, &http.Client{})
	return producers.GetStateProducer(schemaRegistry, dialer, env.BrokersList[0]).Produce(
		context.Background(),
		key,
		payload)
}