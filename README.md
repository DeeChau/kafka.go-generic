# kafka.go-generic
Experiments using Go 1.18beta1's Generic typings and the Segmentio kafka-go consumer client

## Progress
### 01-28-2022
* Successfully adapted generated code from `kafka.go` to use Generics
  * Created `generic_schema.go` and `generic_consumer.go` which aims to unify the duplication of code and streamline the process for creating new consumers.
* Created demo-code in `main.go`
## Next Steps
### 02-07-2022
* Extrace demo-code from `main.go` to a subfolder dedicated for examples (faux unit-tests)
* Better document steps in code where generics are applied
* Add unit tests where available
  * Schema, Avro, Producer & Consumer
* Create `generic_producer.go` which applies `generics` to the producer interface
* Look into alternative patterns to utilize a consumer, e.g. having it poll Kafka for messages (on a loop) -> See https://github.com/wishabi/kafka-etl/blob/main/stream/processor.go for an example.
  * Can also look at how phobos/deimos does kafka consumption for inspiration, would like a more seamless interaction w./ kafka & libraries
  * Look at other PubSub interfaces and how they're exposed.
  * Maybe function pointers if that's an interesting paradigm?

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

## VS Code setup
https://github.com/golang/vscode-go/blob/master/docs/advanced.md#using-go118

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