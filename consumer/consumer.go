package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"time"
)

type testHandler struct {
	messagesReceived int
}

type test1Handler struct {
	messagesReceived int
}

type test2Handler struct {
	messagesReceived int
}

// 处理消息
func (nh *testHandler) HandleMessage(msg *nsq.Message) error {
	nh.messagesReceived++
	body := string(msg.Body)
	log.Printf("%d >> receive ID:%s, addr:%s, message:%s\n", nh.messagesReceived, msg.ID, msg.NSQDAddress, body)
	return nil
}

// 处理消息
func (nh *test1Handler) HandleMessage(msg *nsq.Message) error {
	nh.messagesReceived++
	body := string(msg.Body)
	log.Printf("%d >> receive ID:%s, addr:%s, message:%s\n", nh.messagesReceived, msg.ID, msg.NSQDAddress, body)
	return nil
}

// 处理消息
func (nh *test2Handler) HandleMessage(msg *nsq.Message) error {
	nh.messagesReceived++
	body := string(msg.Body)
	log.Printf("%d >> receive ID:%s, addr:%s, message:%s\n", nh.messagesReceived, msg.ID, msg.NSQDAddress, body)
	return nil
}

func createConsumerNSQD(addr, topic, channel string) (*nsq.Consumer, error) {
	config := nsq.NewConfig()
	config.HeartbeatInterval = time.Second * time.Duration(10)
	config.LookupdPollInterval = 1000 * time.Millisecond
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Println("init Consumer NewConsumer error:", err)
		return nil, err
	}
	consumer.SetLogger(log.New(os.Stderr, "", log.Flags()), nsq.LogLevelError)

	var handler nsq.Handler
	switch topic {
	case "test":
		handler = &testHandler{}
	case "test1":
		handler = &test1Handler{}
	case "test2":
		handler = &test2Handler{}
	}
	consumer.AddHandler(handler)
	// err = consumer.ConnectToNSQLookupd(addr)
	err = consumer.ConnectToNSQD(addr)
	if err != nil {
		log.Println("Consumer ConnectToNSQLookupd error:", err)
		return nil, err
	}

	return consumer, nil
}
