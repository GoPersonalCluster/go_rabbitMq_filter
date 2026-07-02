package filter

import (
	"errors"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/filter/strategy"
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/config"
)

type FilterFactory struct {
	event *consumer.IntegrationEvent
}

func (c *FilterFactory) CreateStrategy(event *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {

	switch event.EventName {
	case "PII":
		return &strategy.PiiQueueStrategy{event: event}, nil
	default:
		return nil, c.GetDefaultErrorResponse(event)
	}
}

func (c *FilterFactory) GetDefaultErrorResponse(event *consumer.IntegrationEvent) error {
	event.CreateMetaHeader(config.GetHostName(), "ErrorMatchingEvent")
	return errors.New(event.EventName + "event not found")
}

// func (c *FilterFactory) GetPIIQueue(event *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {
// 	mh := event.CreateMetaHeader(config.GetHostName(), "ErrorMatchingEvent")
// 	mh.Args = append(mh.Args, mh.CreateArgs("NextQueue", "PII_Queue"))
// 	event.MetaHeader = append(event.MetaHeader, mh)
// 	return *event, nil
// }
