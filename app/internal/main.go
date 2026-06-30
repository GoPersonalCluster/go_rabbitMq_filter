package main

import (
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service"
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/filter"
)

func main() {
	service := service.RabbitMQConfigComposite{}
	service.ConfigureConnection()

	filterCommand := filter.FilterCommand{}

	filterConsumer := consumer.FilterConsumer{}
	config := consumer.FilterConfig{}
	config.AbstractFactory = &filterCommand
	config.Durable = false
	config.Exclusive = false
	config.AutoDelete = false
	config.NoWait = true
	config.QueueName = "filter_queue"
	config.Args= nil 

	filterConsumer.SetConfiguration(&config)

	service.AddConsumer("test_queue", &filterConsumer)
	service.Start()
	service.TestPublish()

}
