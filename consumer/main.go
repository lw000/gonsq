package main

import (
	"github.com/judwhite/go-svc/svc"
	"github.com/nsqio/go-nsq"
	"log"
)

type ConsumerConfig struct {
	addr     string
	topic    string
	channel  string
	consumer *nsq.Consumer
}

type Program struct {
	consumerConfigs []*ConsumerConfig
}

func (p *Program) Init(env svc.Environment) error {
	if env.IsWindowsService() {

	} else {

	}
	// 127.0.0.1:4161
	p.consumerConfigs = append(p.consumerConfigs, &ConsumerConfig{addr: "127.0.0.1:4150", topic: "test1", channel: "test-channel1"})
	p.consumerConfigs = append(p.consumerConfigs, &ConsumerConfig{addr: "127.0.0.1:4152", topic: "test2", channel: "test-channel2"})
	// p.consumerConfigs = append(p.consumerConfigs, &ConsumerConfig{addr: "127.0.0.1:4153", topic: "test", channel: "test-channel2"})

	return nil
}

// Start is called after Init. This method must be non-blocking.
func (p *Program) Start() error {
	var err error
	for _, c := range p.consumerConfigs {
		c.consumer, err = createConsumerNSQD(c.addr, c.topic, c.channel)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

// Stop is called in response to syscall.SIGINT, syscall.SIGTERM, or when a
// Windows Service is stopped.
func (p *Program) Stop() error {
	var err error
	for _, cfg := range p.consumerConfigs {
		err = cfg.consumer.DisconnectFromNSQD(cfg.addr)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

func main() {
	pro := &Program{}
	if err := svc.Run(pro); err != nil {
		log.Println(err)
	}
}
