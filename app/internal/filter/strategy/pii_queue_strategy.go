package strategy

import (
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/config"
)

type PiiQueueStrategy struct {
	event *consumer.IntegrationEvent
}

func (pQS *PiiQueueStrategy) New(iE *consumer.IntegrationEvent) (PiiQueueStrategy, error) {
	iE.EventName = "PII"
	return PiiQueueStrategy{event: iE}, nil
}

func (pQS *PiiQueueStrategy) Start() ([]byte, error) {
	mh := pQS.event.CreateMetaHeader(config.GetHostName(), "ErrorMatchingEvent")
	mh.Args = append(mh.Args, mh.CreateArgs("NextQueue", "PII_Queue"))
	pQS.event.MetaHeader = append(pQS.event.MetaHeader, mh)
	return *pQS.event , nil

}
