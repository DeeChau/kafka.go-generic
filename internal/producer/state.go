// Code generated by github.com/wishabi/kafka.go. DO NOT EDIT.
package producers

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/wishabi/kafka.go/avro"

	schema "github.com/DeeChau/kafka.go-generic/internal/schema"
)

// AvroStateProducer kafka producer for the Schema Type State
type AvroStateProducer struct {
	writer   *kafka.Writer
	registry *avro.Registry
	topic    string
}

// Produce publish a message using key and value
func (p *AvroStateProducer) Produce(ctx context.Context, key *schema.StateKey, value *schema.State) error {
	keySchemaID, err := p.registry.SchemaID(p.topic+"-key", key.Schema())
	if err != nil {
		return err
	}

	keyBuffer := avro.Build(keySchemaID)
	err = key.Serialize(keyBuffer)
	if err != nil {
		return fmt.Errorf("failed to write StateKey to buffer: %w", err)
	}

	valueSchemaID, err := p.registry.SchemaID(p.topic+"-value", value.Schema())
	if err != nil {
		return err
	}

	valueBuffer := avro.Build(valueSchemaID)
	err = value.Serialize(valueBuffer)
	if err != nil {
		return fmt.Errorf("failed to write State to buffer: %w", err)
	}

	err = p.writer.WriteMessages(ctx, kafka.Message{
		Key:   keyBuffer.Bytes(),
		Value: valueBuffer.Bytes(),
	})

	if err != nil {
		return fmt.Errorf("failed to publish State to kafka: %w", err)
	}

	return nil
}

// NewAvroStateProducer creates a producer for sending the Schema Type State message
func NewAvroStateProducer(config kafka.WriterConfig, registry *avro.Registry) *AvroStateProducer {
	return &AvroStateProducer{
		writer:   kafka.NewWriter(config),
		registry: registry,
		topic:    config.Topic,
	}
}
