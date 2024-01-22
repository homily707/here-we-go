package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()
	version, err := sarama.ParseKafkaVersion("3.4.0")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	producer.SendMessage(&sarama.ProducerMessage{
		Topic: "",
	})
}

