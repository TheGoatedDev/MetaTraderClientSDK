package broker

import (
	"github.com/TheGoatedDev/MetaTraderClientSDK/internal/domain/broker"
	"github.com/TheGoatedDev/MetaTraderClientSDK/internal/infrastructure/metatrader4"
	"github.com/TheGoatedDev/MetaTraderClientSDK/internal/infrastructure/metatrader5"
)

type Broker struct{}

func NewBroker() *Broker {
	return &Broker{}
}

func (b *Broker) SearchMT4(company string) ([]broker.Company, error) {
	return metatrader4.Search(company)
}

func (b *Broker) SearchMT5(company string) ([]broker.Company, error) {
	return metatrader5.Search(company)
}
