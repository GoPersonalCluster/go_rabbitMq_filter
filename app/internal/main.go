package main

import (
	"log"
	"net/http"

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
	svc.Start()
	log.Println("[main] service.Start() retornou (não bloqueante) — subindo servidor de health check")

	// Endpoint de health check exigido pelas probes do Kubernetes.
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Println("[main] servidor HTTP ouvindo na porta 8080")
	// ListenAndServe é bloqueante — mantém o processo vivo até dar erro fatal.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("[main] erro no servidor HTTP: %v", err)
	}
}
