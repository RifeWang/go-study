package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

type message struct {
	ID   int
	Body string
}

// ConsumerT ... 自定义消费者
type ConsumerT struct{}

// HandleMessage ... 返回 nil 自动 FIN 否则 REQ
func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))

	b := message{}
	if err := json.Unmarshal(msg.Body, &b); err != nil {
		panic(err)
	}
	fmt.Println("msg struct:", b, b.ID, b.Body)

	return nil
}

// InitConsumer ... 初始化消费者
func InitConsumer(topic string, channel string) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second          //设置重连时间
	c, err := nsq.NewConsumer(topic, channel, cfg) // 新建一个消费者
	if err != nil {
		panic(err)
	}
	c.SetLogger(nil, nsq.LogLevelInfo)
	c.AddHandler(&ConsumerT{}) // 加入自定义的消费者

	// 直连 nsqd
	// if err := c.ConnectToNSQD("127.0.0.1:4150"); err != nil {
	// 	panic(err)
	// }

	// 建议使用 ConnectToNSQLookupd 而不是直连 nsqd
	if err := c.ConnectToNSQLookupds([]string{"127.0.0.1:4161"}); err != nil {
		panic(err)
	}
}

// 主函数
func main() {
	InitConsumer("test", "test-channel")
	
	// handle 在另外的 goroutine 中执行，不得不保持主线程
	for {
		time.Sleep(time.Hour * 3600)
	}
}
