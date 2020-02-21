package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/nsqio/go-nsq"
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

func createConsumer(addr, topic, channel string) error {
	config := nsq.NewConfig()
	config.HeartbeatInterval = time.Second * time.Duration(10)
	config.LookupdPollInterval = 100 * time.Millisecond
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Println("init Consumer NewConsumer error:", err)
		return err
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
	err = consumer.ConnectToNSQLookupd(addr)
	if err != nil {
		log.Println("Consumer ConnectToNSQLookupd error:", err)
		return err
	}

	return nil
}

var (
	cfg = []struct {
		addr    string
		topic   string
		channel string
	}{
		{"127.0.0.1:4161", "test1", "test-channel1"},
		{"127.0.0.1:4161", "test2", "test-channel2"},
		{"127.0.0.1:4161", "test", "test-channel2"},
	}
)

func main() {
	var err error
	for _, c := range cfg {
		err = createConsumer(c.addr, c.topic, c.channel)
		if err != nil {
			log.Fatal("init consumer error")
		}
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	log.Println(<-c)
}
