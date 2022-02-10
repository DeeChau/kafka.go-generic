# kafka.go-generic
Experiments using Go 1.18beta1's Generic typings and the Segmentio kafka-go consumer client

## Progress
### 01-28-2022
* Successfully adapted generated code from `kafka.go` to use Generics
  * Created `generic_schema.go` and `generic_consumer.go` which aims to unify the duplication of code and streamline the process for creating new consumers.
* Created demo-code in `main.go`

### 02-19-2022
* Extract demo-code from `main.go` to subfolders dedicated for examples and unit-tests (`./internal/demo`)
* Updated documentation for if code where generics are applied
* Added unit tests for using the Schema
* Created `generic_producer.go` and added it to demo-code

## Next Steps
### 02-09-2022
* Add unit tests for Consumer and Producer - use `segmentio/kafka-go` or `wishabi/kafka-etl` as reference. Work with Go Guild to build.
* Look into alternative patterns to utilize a consumer, e.g. having it poll Kafka for messages (on a loop) -> See https://github.com/wishabi/kafka-etl/blob/main/stream/processor.go for an example.
  * Can also look at how phobos/deimos does kafka consumption for inspiration, would like a more seamless interaction w./ kafka & libraries
  * Look at other PubSub interfaces and how they're exposed.
  * Possibly function pointers if that's an interesting paradigm?
* Create feature-branch for adding generics to `kafka.go` and updating generators accordingly

## Get Started
1. Install go beta with generics - more found here: https://go.dev/doc/tutorial/generics and here https://go.dev/blog/go1.18beta2
2. Install kafka.go
```bash
GOPRIVATE=github.com/wishabi go install github.com/wishabi/kafka.go/cmd/kgo@latest
```
3. Start GDC for kafka - `/gdc up`
4. Tidy up app
```bash
GOPRIVATE=github.com/wishabi go mod tidy
GOPRIVATE=github.com/DeeChau go mod tidy
```
5. Demo - Run the app
```bash
# Without using generics
go run .
# With using generics
go run . generics
```

## Setting up Generic Producers/Consumers
### Pre-requisites
Before getting started with using Generic producers/consumers in your app, make sure you install the latest version of Go 1.18. Currently, 1.18 is in Beta (`1.18beta2`) but will be released sometime in Feb 2022.

Once done, you should use `kgo` to generate schema packages that are the types for your avro schemas keys and values.

After that is done, you should be good to start with using generics!
### Generic Consumers
```go
// Initialize Schema Registry
schemaRegistry := avro.NewRegistry(env.SchemaRegistry, &http.Client{})
// Setup Dialer
dialer, dialerErr := kafkaUtil.Dialer()

if dialerErr != nil {
  log.Error().Msg("Error getting kafka dialer.")
  return false, dialerErr
}

// fsa Consumer code
fsaConsumer := genericconsumers.NewAvroConsumer[schema.FsaKey, schema.Fsa](
  kafka.ReaderConfig{
    Topic:   "Fsa",
    GroupID: "TestConsumerGenerics",
    Dialer:  dialer,
    Brokers: env.BrokersList,
  }, schemaRegistry)

ctx := context.Background()
fsaMsg, fsaErr := fsaConsumer.AutoCommitConsume(ctx)
if fsaErr != nil {
  log.Error().Err(fsaErr).Msg("Kafka message could not be received. Error consuming message.")
  return false, fsaErr
}

fsaKey := fsaMsg.Key
fsaValue := fsaMsg.Value
```

### Generic Producers
```go
// Initialize Schema Registry
schemaRegistry := avro.NewRegistry(env.SchemaRegistry, &http.Client{})
// Setup Dialer
dialer, dialerErr := kafkaUtil.Dialer()
if dialerErr != nil {
  log.Error().Msg("Error getting kafka dialer.")
  return false, dialerErr
}

fsa := "M6G"
lat := float32(88.88)
long := float32(99.99)
key := schema.FsaKey{Label: fsa}
payload := schema.Fsa{
  Label:      fsa,
  Latitude:   lat,
  Longitude:  long,
  Message_id: fmt.Sprintf("FSA-%v", timestamp),
  Created_at: timestamp.Unix(),
  Updated_at: timestamp.Unix(),
  Timestamp:  timestamp.Format(time.UnixDate),
}

// Fsa producer code
fsaProducer := genericproducers.NewAvroProducer[schema.FsaKey, schema.Fsa](
  kafka.WriterConfig{
    Topic:   "Fsa",
    Dialer:  dialer,
    Brokers: env.BrokersList,
  }, schemaRegistry
)

err := fsaProducer.Produce(
  context.Background(),
  &key,
  &payload)
if err != nil {
  log.Error().Err(err).Msg("Kafka message could not be sent. Error producing message.")
  return false, err
}
```

## Additional setup for using VS Code for development
https://github.com/golang/vscode-go/blob/master/docs/advanced.md#using-go118

## Additional Resources
- https://go.dev/doc/tutorial/generics
- https://go.dev/blog/go1.18beta2
- https://stackoverflow.com/questions/69573113/how-can-i-instantiate-a-new-pointer-of-type-argument-with-generic-go
- https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#pointer-method-example

## `kgo` commands used
#### Schemas
```bash
kgo schema avro \
  -i schemas/Fsa.avsc \
  -o internal/schema

kgo schema avro \
  -i schemas/State.avsc \
  -o internal/schema

kgo schema avro \
  -i schemas/FsaKey.avsc \
  -o internal/schema

kgo schema avro \
  -i schemas/StateKey.avsc \
  -o internal/schema
```

#### Consumers
```bash
kgo consumer avro \
  -k FsaKey \
  -v Fsa \
  -s github.com/DeeChau/kafka.go-generic/internal/schema \
  -n Fsa \
  -o internal/consumer/fsa.go

kgo consumer avro \
  -k StateKey \
  -v State \
  -s github.com/DeeChau/kafka.go-generic/internal/schema \
  -n State \
  -o internal/consumer/state.go
```

#### Producers
```bash
kgo producer avro \
  -k FsaKey \
  -v Fsa \
  -s github.com/DeeChau/kafka.go-generic/internal/schema \
  -n Fsa \
  -o internal/producer/fsa.go

kgo producer avro \
  -k StateKey \
  -v State \
  -s github.com/DeeChau/kafka.go-generic/internal/schema \
  -n State \
  -o internal/producer/state.go
```