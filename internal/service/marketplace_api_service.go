package service

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sergera/marketplace-order-notifier-worker/internal/conf"
	"github.com/sergera/marketplace-order-notifier-worker/internal/domain"
)

type MarketplaceAPIService struct {
	host        string
	port        string
	contentType string
	client      *http.Client
}

func NewMarketplaceAPIService() *MarketplaceAPIService {
	conf := conf.GetConf()
	return &MarketplaceAPIService{
		conf.MarketplaceAPIHost,
		conf.MarketplaceAPIPort,
		"application/json; charset=UTF-8",
		&http.Client{},
	}
}

func (mkt MarketplaceAPIService) Post(route string, jsonData []byte) error {
	request, err := http.NewRequest("POST", mkt.host+":"+mkt.port+"/"+route, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("failed to create post request: " + err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, err := mkt.client.Do(request)
	if err != nil {
		log.Println("failed to perform marketplace api post request: " + err.Error())
		return err
	}

	defer response.Body.Close()
	return nil
}

func (mkt MarketplaceAPIService) Put(route string, jsonData []byte) error {
	request, err := http.NewRequest("PUT", mkt.host+":"+mkt.port+"/"+route, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("failed to create put request: " + err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, err := mkt.client.Do(request)
	if err != nil {
		log.Println("failed to perform marketplace api put request: " + err.Error())
		return err
	}

	defer response.Body.Close()
	return nil
}

func (mkt MarketplaceAPIService) UpdateOrderStatus(o domain.OrderModel) error {
	m, err := json.Marshal(o)
	if err != nil {
		log.Println("failed to marshal order model into json")
		return err
	}

	err = mkt.Post("update-order", m)
	if err != nil {
		return err
	}

	return nil
}
