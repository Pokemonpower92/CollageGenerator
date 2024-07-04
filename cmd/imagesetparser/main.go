package main

import (
	"log"

	"github.com/pokemonpower92/imagesetservice/config"
	"github.com/pokemonpower92/imagesetservice/internal/consumer"
	"github.com/pokemonpower92/imagesetservice/internal/jobhandler"
	"github.com/pokemonpower92/imagesetservice/internal/repository"
)

func init() {
	config.LoadEnvironmentVariables()
}

func instantiateJobHandler() *jobhandler.JobHandler {
	repo, err := repository.NewImageSetRepository()
	if err != nil {
		panic(err)
	}
	jobHandlerLogger := log.New(log.Writer(), "jobhandler: ", log.Flags())
	return jobhandler.NewJobHandler(repo, jobHandlerLogger)
}

func instantiateImageSetConsumer(jh *jobhandler.JobHandler) *consumer.ImageSetConsumer {
	consumerLogger := log.New(log.Writer(), "consumer: ", log.Flags())
	consumerConfig := config.NewConsumerConfig()
	return consumer.NewImageSetConsumer(jh, consumerLogger, consumerConfig)
}

func main() {
	jh := instantiateJobHandler()
	isc := instantiateImageSetConsumer(jh)
	isc.Consume()
}
