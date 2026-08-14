package main

import (
	"log"

	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service"
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/filter"
)

func main() {
	log.Println("[main] iniciando aplicação...")

	svc := service.RabbitMQConfigComposite{}

	log.Println("[main] configurando conexão com RabbitMQ...")
	svc.ConfigureConnection()
	log.Println("[main] conexão configurada com sucesso")

	filterCommand := filter.FilterFactory{}

	filterConsumer := consumer.FilterConsumer{}
	config := consumer.ConsumerConfig{}

	config.AbstractFactory = &filterCommand
	config.Durable = false
	config.Exclusive = false
	config.AutoDelete = false
	config.NoWait = true
	config.QueueName = "filter_queue"
	config.Args = nil

	log.Println("[main] configurando consumer...")
	filterConsumer.SetConfiguration(&config)

	log.Println("[main] registrando consumer 'filter_queue'...")
	svc.AddConsumer("filter_queue", &filterConsumer)

	log.Println("[main] chamando service.Start()...")
	go svc.Start()
	log.Println("[main] service.Start() retornou (não bloqueante) — mantendo processo vivo")

	select {}
}
