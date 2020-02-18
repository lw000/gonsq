package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

type nsqProducer struct {
	*nsq.Producer
}

func initProducer(addr string) (*nsqProducer, error) {
	log.Println("init producer address:", addr)
	producer, err := nsq.NewProducer(addr, nsq.NewConfig())
	if err != nil {
		return nil, err
	}
	return &nsqProducer{producer}, nil

}

func (np *nsqProducer) public(topic, message string) error {
	err := np.Publish(topic, []byte(message))
	if err != nil {
		log.Println("nsq public error:", err)
		return nil
	}
	return nil
}

func UUID() string {
	u1, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	return u1.String()
}

func main() {
	strIP1 := "127.0.0.1:4150"
	strIP2 := "127.0.0.1:4152"
	producer1, err := initProducer(strIP1)
	if err != nil {
		log.Fatal("init producer1 error:", err)
	}
	producer2, err := initProducer(strIP2)
	if err != nil {
		log.Fatal("init producer1 error:", err)
	}

	defer producer1.Stop()
	defer producer2.Stop()

	// 读取控制台输入
	// reader := bufio.NewReader(os.Stdin)

	count := 0
	for {
		// fmt.Print("please say:")
		// data, _, _ := reader.ReadLine()
		// command := string(data)
		command := UUID()
		log.Println(command)
		if command == "stop" {
			fmt.Println("stop producer!")
			return
		}
		if count%2 == 0 {
			err := producer1.public("test1", command)
			if err != nil {
				log.Fatal("producer1 public error:", err)
			}
		} else {
			err := producer2.public("test2", command)
			if err != nil {
				log.Fatal("producer1 public error:", err)
			}
		}
		count++

		time.Sleep(time.Millisecond * time.Duration(10))
	}
}
