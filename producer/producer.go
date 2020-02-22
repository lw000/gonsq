package main

import (
	"context"
	"github.com/nsqio/go-nsq"
	"github.com/satori/go.uuid"
	"log"
	"os"
	"time"
)

type nsqProducer struct {
	producer *nsq.Producer
}

func NewProducer(addr string) (*nsqProducer, error) {
	config := nsq.NewConfig()
	config.HeartbeatInterval = time.Second * time.Duration(10)
	producer, err := nsq.NewProducer(addr, config)
	if err != nil {
		return nil, err
	}
	producer.SetLogger(log.New(os.Stderr, "", log.Flags()), nsq.LogLevelError)
	return &nsqProducer{producer}, nil

}

func (np *nsqProducer) Publish(topic, message string) error {
	err := np.producer.Publish(topic, []byte(message))
	if err != nil {
		log.Println("nsq public error:", err)
		return nil
	}
	return nil
}

func (np *nsqProducer) Stop() {
	np.producer.Stop()
}

func (np *nsqProducer) Message() string {
	u1 := uuid.NewV4()
	return u1.String()
}

func runProducer(ctx context.Context, producer *nsqProducer, topic string, message string) {
	t := time.NewTicker(time.Millisecond * time.Duration(100))
	for {
		select {
		case <-t.C:
			err := producer.Publish(topic, message)
			if err != nil {
				log.Println("Publish error:", err)
			}
		case <-ctx.Done():
			return
		}
	}
}
