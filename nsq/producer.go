package main

import (
	"encoding/json"
	"log"

	"github.com/nsqio/go-nsq"
)

// NSQProducer ... 高级特性，只能用于TCP端口，自定义TCP协议
var NSQProducer *nsq.Producer = producer()

func producer() *nsq.Producer {
	NSQProducer, err := nsq.NewProducer("localhost:4150", nsq.NewConfig())
	if err != nil {
		log.Fatal("failed to connect NSQ", err, NSQProducer)
	}
	log.Println("========111111===========", NSQProducer)
	return NSQProducer
}

type message struct {
	ID   int
	Body string
}

func main() {
	log.Println("===================", NSQProducer)
	topic := "test"

	body := message{
		ID:   123458888,
		Body: "/safdfdsfafa/fsdfdfsfs",
	}
	b, err := json.Marshal(body)
	if err != nil {
		log.Println("json marshal error:", err)
	}
	log.Println("msg body:", string(b))

	err = NSQProducer.Publish(topic, b)
	if err != nil {
		log.Println("publish error:", err)
	}
}
