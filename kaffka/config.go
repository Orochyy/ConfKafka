package kaffka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func StartKafka() {

	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	}

	c, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			println("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			println("Consumer error: %v (%v)\n", err, msg)
		}
	}

	//c.SubscribeTopics([]string{"myTopic"}, nil)

}
