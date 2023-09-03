package kfk

import (
	"fmt"

	"github.com/IBM/sarama"
)

// KafkaProducer 发送单条
func KafkaProducer(topic, msg string) (err error) {
	producer, err := sarama.NewSyncProducerFromClient(GobalKafka)
	if err != nil {
		return
	}
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		return
	}
	fmt.Println(offset, partition)
	return
}

// KafkaProducers 发送多条，topic在messages中
func KafkaProducers(messages []*sarama.ProducerMessage) (err error) {
	producer, err := sarama.NewSyncProducerFromClient(GobalKafka)
	if err != nil {
		return
	}
	err = producer.SendMessages(messages)
	if err != nil {
		return
	}
	return
}
