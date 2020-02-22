package main

import (
	"context"
	"github.com/judwhite/go-svc/svc"
	"log"
	"math/rand"
	"time"
)

type Program struct {
	producer1 *nsqProducer
	producer2 *nsqProducer
}

func (p *Program) Init(env svc.Environment) error {
	if env.IsWindowsService() {

	} else {

	}
	rand.Seed(time.Now().Unix())
	return nil
}

// Start is called after Init. This method must be non-blocking.
func (p *Program) Start() error {
	var err error
	p.producer1, err = NewProducer("127.0.0.1:4150")
	if err != nil {
		log.Println("init producer1 error:", err)
		return err
	}

	p.producer2, err = NewProducer("127.0.0.1:4152")
	if err != nil {
		log.Println("init producer2 error:", err)
		return err
	}

	go runProducer(context.Background(), p.producer1, "test", "test")
	go runProducer(context.Background(), p.producer2, "test2", "test2")

	return nil
}

// Stop is called in response to syscall.SIGINT, syscall.SIGTERM, or when a
// Windows Service is stopped.
func (p *Program) Stop() error {
	p.producer1.Stop()
	p.producer2.Stop()
	return nil
}

func main() {
	pro := &Program{}
	if err := svc.Run(pro); err != nil {
		log.Println(err)
	}
}
