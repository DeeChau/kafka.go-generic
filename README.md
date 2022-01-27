# kafka.go-generic
Experiments using Go 1.18beta1's Generic typings and the Segmentio kafka-go consumer client

## Get Started
1. Install go beta with generics - more found here: https://go.dev/doc/tutorial/generics
2. Install kafka.go
```bash
GOPRIVATE=github.com/wishabi go install github.com/wishabi/kafka.go/cmd/kgo@latest
```
3. Start GDC for kafka - `/gdc up`
4. Run the app