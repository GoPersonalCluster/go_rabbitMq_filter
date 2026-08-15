package main

import (
	"log"

	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service"
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/filter"
)

func main() {
	log.Println("[main] iniciando aplicação...")

	svc := service.FilterRabbitMQConfigComposite{}

	svc.ConfigureConnection()

	filterCommand := filter.FilterFactory{}

	filterConsumer := consumer.FilterConsumer{}
	config := consumer.ConsumerConfig{}

	config.AbstractFactory = &filterCommand
	config.Durable = true
	config.Exclusive = false
	config.AutoDelete = false
	config.NoWait = true
	config.QueueName = "filter_queue"
	config.Args = nil

	filterConsumer.SetConfiguration(&config)

	svc.AddConsumer("filter_queue", &filterConsumer)

	// Mantém a aplicação em execução.
	svc.Start()
}
