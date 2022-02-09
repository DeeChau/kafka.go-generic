package main

import (
	"fmt"
	"os"

	demo "github.com/DeeChau/kafka.go-generic/internal/demo"
	"github.com/rs/zerolog"
)

func main() {
	fmt.Println("---Begin Hackathon for Generic Kafka experimentation---")
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	useGenerics := len(os.Args) >= 2 && (os.Args[1] == "generic")

	if useGenerics {
		demo.ProduceWithGenerics()
		demo.ConsumeWithGenerics()
	} else {
		demo.PublishAndConsumeNoGenerics()
	}
}
