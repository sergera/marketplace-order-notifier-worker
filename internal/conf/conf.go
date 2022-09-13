package conf

import (
	"log"
	"sync"

	"github.com/gurkankaymak/hocon"
)

var once sync.Once
var instance *conf

type conf struct {
	hocon              *hocon.Config
	Port               string
	KafkaHost          string
	KafkaPort          string
	MarketplaceAPIHost string
	MarketplaceAPIPort string
}

func GetConf() *conf {
	once.Do(func() {
		var c *conf = &conf{}
		c.setup()
		instance = c
	})
	return instance
}

func (c *conf) setup() {
	c.parseHOCONConfigFile()
	c.setKafkaHost()
	c.setKafkaPort()
	c.setMarketplaceAPIHost()
	c.setMarketplaceAPIPort()
}

func (c *conf) parseHOCONConfigFile() {
	hocon, err := hocon.ParseResource("application.conf")
	if err != nil {
		log.Panic("error while parsing configuration file: ", err)
	}

	log.Printf("configurations: %+v", *hocon)

	c.hocon = hocon
}

func (c *conf) setKafkaHost() {
	kafkaHost := c.hocon.GetString("kafka.host")
	if len(kafkaHost) == 0 {
		log.Panic("kafka host environment variable not found")
	}

	c.KafkaHost = kafkaHost
}

func (c *conf) setKafkaPort() {
	kafkaPort := c.hocon.GetString("kafka.port")
	if len(kafkaPort) == 0 {
		log.Panic("kafka port environment variable not found")
	}

	c.KafkaPort = kafkaPort
}

func (c *conf) setMarketplaceAPIHost() {
	marketplaceAPIHost := c.hocon.GetString("marketplace-api.host")
	if len(marketplaceAPIHost) == 0 {
		log.Panic("marketplace api host environment variable not found")
	}

	c.MarketplaceAPIHost = marketplaceAPIHost
}

func (c *conf) setMarketplaceAPIPort() {
	marketplaceAPIPort := c.hocon.GetString("marketplace-api.port")
	if len(marketplaceAPIPort) == 0 {
		log.Panic("marketplace api port environment variable not found")
	}

	c.MarketplaceAPIPort = marketplaceAPIPort
}
