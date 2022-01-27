# kafka.go-generic
Experiments using Go 1.18beta1's Generic typings and the Segmentio kafka-go consumer client

## Get Started
1. Install go beta with generics - more found here: https://go.dev/doc/tutorial/generics
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
5. Run the app


# kgo commands
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