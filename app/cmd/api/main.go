package main

import (
	"log"
	"net/http"

	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service"
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/filter"
	"github.com/gin-gonic/gin"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/routes"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/docs"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strings"

)
func CaseInsensitiveRouter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.URL.Path = strings.ToLower(c.Request.URL.Path)
		c.Next()
	}
}

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
	
		// Routes
	routes.Setup(
		router,
		userHandler,
	)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router := gin.Default()
	router.Use(CaseInsensitiveRouter())

	api := router.Group("/api/v1")
	{
		api.GET("/healthCheck", handlers.HealthCheck)
	}

	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)
	
	docs.Init()
	router.Run(":8080")
}
