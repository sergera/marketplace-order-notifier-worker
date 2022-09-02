package evt

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sergera/marketplace-order-notifier-worker/internal/domain"
	"github.com/sergera/marketplace-order-notifier-worker/internal/service"
)

var Topics map[domain.Status]string = map[domain.Status]string{
	domain.Unconfirmed: "orders__unconfirmed",
	domain.InProgress:  "orders__in_progress",
	domain.Ready:       "orders__ready",
	domain.InTransit:   "orders__in_transit",
	domain.Delivered:   "orders__delivered",
}

type OrderListener struct {
	service      *service.MarketplaceAPIService
	eventHandler *EventHandler
}

func NewOrderListener() *OrderListener {
	return &OrderListener{service.NewMarketplaceAPIService(), NewEventHandler()}
}

func (l *OrderListener) Listen() {
	l.eventHandler.Consume([]string{
		Topics[domain.InProgress],
		Topics[domain.Ready],
		Topics[domain.InTransit],
		Topics[domain.Delivered]},
		func(msg kafka.Message) {
			updatedOrder := domain.OrderModel{}
			if err := json.Unmarshal(msg.Value, &updatedOrder); err != nil {
				log.Println("error unmarshalling kafka message: ", err.Error())
				return
			}
			l.NotifyMarketplace(updatedOrder)
		})
}

func (l *OrderListener) NotifyMarketplace(o domain.OrderModel) {
	l.service.UpdateOrderStatus(o)
}
