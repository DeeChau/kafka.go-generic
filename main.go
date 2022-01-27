package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/DeeChau/kafka.go-generic/internal/schema"
	"github.com/DeeChau/kafka.go-generic/kafka/consumers"
	"github.com/DeeChau/kafka.go-generic/kafka/producers"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/wishabi/kafka.go/avro"
	env "github.com/wishabi/pkg/kafka"
	kafkaUtil "github.com/wishabi/pkg/kafka"
)

func main() {
	fmt.Println("---Begin Hackathon for Generic Kafka Consumers---")

	useGenerics := os.Args[1] == "generic"

	if useGenerics {
		return
	} else {
		publishAndConsumeNoGenerics()
	}
}

// Not using generics!!!
func publishAndConsumeNoGenerics() {
	_, err := publishFsa("M5V", 43.6450946, -79.3929024)
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
func consumeFsas() (bool, error) {
	log.Info().Msg("Create consumers")

	dialer, dialerErr := kafkaUtil.Dialer()

	if dialerErr != nil {
		log.Error().Msg("Error getting kafka dialer.")
		return false, dialerErr
	}

	// dialer, dialerErr := kafkaUtil.Dialer()
	schemaRegistry := avro.NewRegistry(env.SchemaRegistry, &http.Client{})

	consumer := consumers.GetFsaConsumer(schemaRegistry, dialer, env.BrokersList[0])

	ctx := context.Background()
	msg, err := consumer.AutoCommitConsume(ctx)
	if err != nil {
		return false, err
	}

	key := msg.Key
	value := msg.Value
	log.Printf("---> Kafka: Received Key %v, Message %v from", key, value)

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

	log.Printf("Key to send %v, Message to send %v", key, payload)
	log.Printf("Schema Registry location %v, Kafka Broker location %v", env.SchemaRegistry, env.BrokersList)

	err := produceFsaMessage(dialer, &key, &payload)
	if err != nil {
		log.Error().Err(err).Msg("Kafka message could not be sent. Error producing message.")
		return false, err
	}

	log.Info().Msg("FSA kafka message sent successfully.")
	return true, nil
}

// These should probably be abstracted and made easier, possible within the producer's library.
func generateFsaKey(fsa string) schema.FsaKey {
	return schema.FsaKey{Label: fsa}
}

func generateFsa(fsa string, lat, long float32) schema.Fsa {
	// These keys should be standardized with Camel Case or snake case - not both.
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

// Utilize Type Switch to take in a Schema Type -> Translate to the Configured Kafka Topic ID.
func produceFsaMessage(dialer *kafka.Dialer, key *schema.FsaKey, payload *schema.Fsa) error {
	// Share one registry across? Is that a good idea.
	schemaRegistry := avro.NewRegistry(env.SchemaRegistry, &http.Client{})
	return producers.GetFsaProducer(schemaRegistry, dialer, env.BrokersList[0]).Produce(
		context.Background(),
		key,
		payload)
}
