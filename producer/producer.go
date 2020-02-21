package main

import (
	"github.com/nsqio/go-nsq"
	"github.com/satori/go.uuid"
	"log"
	"math/rand"
	"os"
	"time"
)

type nsqProducer struct {
	producer *nsq.Producer
}

func New(addr string) (*nsqProducer, error) {
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

func UUID() string {
	u1 := uuid.NewV4()
	return u1.String()
}

func main() {
	strIP1 := "127.0.0.1:4150"
	strIP2 := "127.0.0.1:4152"
	producer1, err := New(strIP1)
	if err != nil {
		log.Fatal("init producer1 error:", err)
	}
	defer producer1.Stop()

	producer2, err := New(strIP2)
	if err != nil {
		log.Fatal("init producer2 error:", err)
	}
	defer producer2.Stop()

	// 读取控制台输入
	// reader := bufio.NewReader(os.Stdin)

	rand.Seed(time.Now().Unix())

	for {
		// log.Println("please say:")
		// data, _, _ := reader.ReadLine()
		// msg := string(data)
		// if msg == "stop" {
		// 	fmt.Println("stop producer!")
		// 	return
		// }

		index := rand.Intn(3)
		// msg := UUID()
		switch index {
		case 0:
			err := producer1.Publish("test", "test")
			if err != nil {
				log.Fatal("producer1 Publish error:", err)
			}
		case 1:
			err := producer2.Publish("test1", "test1")
			if err != nil {
				log.Fatal("producer1 Publish error:", err)
			}
		default:
			err := producer2.Publish("test2", "test2")
			if err != nil {
				log.Fatal("producer1 Publish error:", err)
			}
		}

		time.Sleep(time.Millisecond * time.Duration(100))
	}
}
