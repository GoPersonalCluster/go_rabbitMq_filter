package filter

import (
	"errors"

	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/config"
)

type FilterCommand struct {
}

func (c *FilterCommand) GetQueue(event *consumer.IntegrationEvent) (consumer.IntegrationEvent, error) {

	switch event.EventName {
	case "PII":
		return c.GetPIIQueue(event)
	default:
		return c.GetDefaultErrorResponse(event)
	}

}
func (c *FilterCommand) GetDefaultErrorResponse(event *consumer.IntegrationEvent) (consumer.IntegrationEvent, error) {
	event.CreateMetaHeader(config.GetHostName(), "ErrorMatchingEvent")
	return *event, errors.New(event.EventName + "event not found")
}


func (c *FilterCommand) GetPIIQueue(event *consumer.IntegrationEvent) (consumer.IntegrationEvent, error) {
	mh := event.CreateMetaHeader(config.GetHostName(), "ErrorMatchingEvent")
	mh.Args = append(mh.Args, mh.CreateArgs("NextQueue", "PII_Queue"))
	event.MetaHeader = append(event.MetaHeader, mh)
	return *event, nil
}
