package filter

import (
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer",
	"errors"
)

type FilterCommand struct {
}

func (c *FilterCommand) GetQueue(event *consumer.IntegrationEvent) (consumer.IntegrationEvent, error) {

	switch event.EventName {
		case "PII": return "PII_Queue", nil
		default : return "err" , errors.New(event.EventName + "event not found")
	}

}


