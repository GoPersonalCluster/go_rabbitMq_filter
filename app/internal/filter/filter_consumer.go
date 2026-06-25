package filter

import (
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
)
type FilterCommand struct {

}
	
func (c *FilterCommand)CreateStrategy(event *consumer.IntegrationEvent) (consumer.StrategyHandler, error){
	switch event.EventName {
	case "PII" : 
	}
}

func (d *FilterCommand)Command(mh *consumer.MetaHeader, eN string)  (consumer.StrategyHandler , error){


}