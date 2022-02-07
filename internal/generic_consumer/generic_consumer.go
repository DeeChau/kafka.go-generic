// Modified from code generated by github.com/wishabi/kafka.go.
package genericconsumers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/rs/zerolog/log"

	"github.com/segmentio/kafka-go"
	"github.com/wishabi/kafka.go/avro"

	genericschema "github.com/DeeChau/kafka.go-generic/internal/generic_schema"
)

// This type is sharable with others.
type AvroConsumer[K, V genericschema.AvroSchemaConstraint, PTK genericschema.AvroSchemaStruct[K], PTV genericschema.AvroSchemaStruct[V]] struct {
	reader   *kafka.Reader
	registry *avro.Registry
	topic    string
}

type KafkaMessage[K, V genericschema.AvroSchemaConstraint] struct {
	Topic     string
	Partition int
	Offset    int64
	Key       *K
	Value     *V
	Headers   []kafka.Header
	Time      time.Time
}

// TODO: -> Add unit tests!
// Create helpers
func parseKafkaMessage[K, V genericschema.AvroSchemaConstraint,
					   PTK genericschema.AvroSchemaStruct[K], PTV genericschema.AvroSchemaStruct[V]](
		message kafka.Message, registry *avro.Registry) (*KafkaMessage[K, V], error) {

	valueData, valueSchemaID, err := avro.Parse(message.Value)
	if err != nil && !errors.Is(err, avro.ErrEmpty) {
		return nil, fmt.Errorf("can't parse Kafka Message: %w", err)
	}

	log.Printf("Consumer parsed avro message: %s", valueData)

	var value *V

	log.Printf("Consumer parsed value: %v - %s - %T", value, value, value)

	// if value is tombstone message, the error will be avro.ErrEmpty
	// if it's not tombdstone message, then we need to parse the value
	if !errors.Is(err, avro.ErrEmpty) {
		valueSchema, err := registry.Schema(valueSchemaID)


		log.Printf("Registry found schema: %s", valueSchema)
		value, err = genericschema.DeserializeFromSchema[V, PTV](bytes.NewReader(valueData), valueSchema)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize Kafka  message: %w", err)
		}
	}

	log.Printf("Consumer parsed value: %s", value)

	keyData, keySchemaID, err := avro.Parse(message.Key)
	if err != nil && !errors.Is(err, avro.ErrEmpty) {
		return nil, fmt.Errorf("can't parse Kafka Key: %w", err)
	}

	var key *K

	// if key is tombstone message, the error will be avro.ErrEmpty
	// if it's not tombstone message, then we need to parse the key
	// NOTE: if the key is empty, there must be an error happened in publisher side
	// whoever receiving this message should consider adding null check at key
	if !errors.Is(err, avro.ErrEmpty) {
		keySchema, err := registry.Schema(keySchemaID)

		// Flag
		key, err = genericschema.DeserializeFromSchema[K, PTK](bytes.NewReader(keyData), keySchema)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize Kafka Key message: %w", err)
		}		
	}

	log.Printf("Consumer parsed key: %s", key)

	return &KafkaMessage[K, V]{
		Topic:     message.Topic,
		Partition: message.Partition,
		Offset:    message.Offset,
		Key:       key,
		Value:     value,
		Headers:   message.Headers,
		Time:      message.Time,
	}, nil
}

// TODO: Implement Batch consume too!
// Unit Tests -> How do we mock out the Kafka Consumer/Producer? Does such a library even exist?
// Consume consumes the next message and returns KafkaMessage, commit function and an error
// The caller of this call is responsible to call the commit.
// Note: this method is a blocking call and can be cancled by passing timeout or deadline context
func (c *AvroConsumer[K, V, PTK, PTV]) Consume(ctx context.Context)(*KafkaMessage[K, V], func(context.Context) error, error) {
	message, err := c.reader.FetchMessage(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read Kafka message: %w", err)
	}

	commit := func(ctx context.Context) error {
		return c.reader.CommitMessages(ctx, message)
	}

	msg, err := parseKafkaMessage[K, V, PTK, PTV](message, c.registry)
	if err != nil {
		return nil, commit, err
	}

	log.Printf("Consumer parsed messages: %s", msg)

	return msg, commit, nil
}

// // AutoCommitConsume consumes the next message and returns KafkaMessage and automatically calls commit
// // by passing a timeout context, ones can control timeout of consuming and commiting a message
func (c *AvroConsumer[K, V, PTK, PTV]) AutoCommitConsume(ctx context.Context) (*KafkaMessage[K, V], error) {
	msg, commit, err := c.Consume(ctx)
	if err != nil {
		return nil, err
	}

	return msg, commit(ctx)
}

// This can likely be extracted into another file... Will need to ask Go Experts on how the package layout should/could be
// NewAvroKafka Consumer creates a consumer which can consume and returns Kafka Message messages
func NewAvroConsumer[K, V genericschema.AvroSchemaConstraint, PTK genericschema.AvroSchemaStruct[K], PTV genericschema.AvroSchemaStruct[V]](
		config kafka.ReaderConfig, registry *avro.Registry) *AvroConsumer[K, V, PTK, PTV] {
	return &AvroConsumer[K, V, PTK, PTV]{
		reader:   kafka.NewReader(config),
		registry: registry,
		topic:    config.Topic,
	}
}
