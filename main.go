package main

import (
	"fmt"

	demo "github.com/DeeChau/kafka.go-generic/internal/demo"
)

func main() {
	fmt.Println("---Begin Hackathon for Generic Kafka experimentation---")
	// zerolog.SetGlobalLevel(zerolog.InfoLevel)
	demo.ConsumeWithGenerics()

	// useGenerics := len(os.Args) >= 2 && (os.Args[1] == "generic")

	// if useGenerics {
	// 	// demo.ExperimentWithGenerics()

	// } else {
	// 	demo.PublishAndConsumeNoGenerics()
	// }
}
