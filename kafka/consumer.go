package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

// 处理 topic/partition claims(声明)，提供 session 生命周期的钩子
type consumerGroupHandler struct{}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (h consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		key := string(msg.Key)
		value := string(msg.Value)
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		fmt.Println("message key:", key, "message value:", value, "timestamp:", msg.Timestamp, "BlockTimestamp:", msg.BlockTimestamp)
		session.MarkMessage(msg, "") // 确认消息完成消费
	}
	return nil
}

func main() {
	hosts := []string{
		"192.168.64.10:9092",
		"192.168.64.11:9092",
		"192.168.64.12:9092",
		"192.168.64.13:9092",
		"192.168.64.14:9092",
		"192.168.64.15:9092",
		"192.168.64.16:9092",
		"192.168.64.17:9092",
		"192.168.64.18:9092",
		"192.168.64.19:9092",
	}
	// topic := "uplog_yupoo"
	topic := "uplog_es"
	groupID := "yupoo_impress"

	// 初始化配置
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_2_1                                // kafka 版本
	config.ClientID = "yupoo_wy"                                     // 自定义客户端名称
	config.Consumer.Return.Errors = true                             // 消费时产生的错误都将返回到 Errors channel，默认关闭
	config.Consumer.Offsets.CommitInterval = 1000 * time.Millisecond // 自动提交 offset 的间隔时间，默认 1s
	config.Consumer.Offsets.Retry.Max = 3                            // 提交 offset 失败的最大重试次数，默认 3 次
	config.Consumer.Offsets.Initial = sarama.OffsetNewest            // 默认 OffsetNewest 从最新消费，OffsetOldest 从头开始消费

	// 创建 client
	client, err := sarama.NewClient(hosts, config)
	if err != nil {
		log.Fatal("create kafka client error:", err)
	}
	defer func() { _ = client.Close() }()     // client 必须关闭
	partitions, _ := client.Partitions(topic) // 获取 topic 下的 partitions
	fmt.Println("partitions:", partitions)

	// 创建 consumer group
	group, err := sarama.NewConsumerGroupFromClient(groupID, client)
	if err != nil {
		log.Fatal("create client group error:", err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR:", err)
		}
	}()

	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		handler := consumerGroupHandler{} // 必须使用 consumerGroupHandler

		err := group.Consume(ctx, []string{topic}, handler)
		if err != nil {
			log.Fatal("group consume error:", err)
		}
	}
}

// 消费者加入 group 和进行消费都通过 ConsumerGroupHandler 来完成。
//
// 消费者会话的生命周期如下:
//
// 1. consumers 加入 group 并被分配 partitions, 这个过程也叫 claims 。
// 2. 开始消费前, ConsumerGroupHandler 的 Setup() 钩子被触发。
// 3. ConsumeClaim() 方法被调用，进行消费，需要注意保证并发安全性。
// 4. session 会话一直存在，直到某个 ConsumeClaim() 方法退出. 这可能会由 parent context 取消造成，或者服务器 rebalance .
// 5. 一旦所有 ConsumeClaim() 循环退出, Cleanup() 方法被触发，用户可以在此处进行一些任务，先于 rebalance.
// 6. 最后 offsets 被提交。
//
// Please note, that once a rebalance is triggered, sessions must be completed within
// Config.Consumer.Group.Rebalance.Timeout. This means that ConsumeClaim() functions must exit
// as quickly as possible to allow time for Cleanup() and the final offset commit. If the timeout
// is exceeded, the consumer will be removed from the group by Kafka, which will cause offset
// commit failures.
