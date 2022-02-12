module github.com/DeeChau/kafka.go-generic

go 1.18

require (
	github.com/actgardner/gogen-avro/v7 v7.3.1
	github.com/rs/zerolog v1.20.0
	github.com/segmentio/kafka-go v0.4.27
	github.com/wishabi/kafka.go v0.0.0-20210527154254-c027742f8263
	github.com/wishabi/pkg v0.0.7
)

require (
	github.com/golang/snappy v0.0.2 // indirect
	github.com/klauspost/compress v1.9.8 // indirect
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
)

replace github.com/wishabi/kafka.go => github.com/DeeChau/kafka.go v0.2.0-alpha
