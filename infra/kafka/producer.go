package kafka_code_bank

import "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaProducer struct {
	Producer *kafka.Producer
}

func NewKafkaProducer() KafkaProducer {
	return KafkaProducer{}
}

func (k *KafkaProducer) SetupProducer(bs string) error {
	config := &kafka.ConfigMap{
		"bootstrap.servers": bs,
	}

	producer, err := kafka.NewProducer(config)
	k.Producer = producer

	if err != nil {
		return err
	}

	return nil
}

func (k *KafkaProducer) Publish(payload string, topic string) error {
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(payload),
	}

	err := k.Producer.Produce(message, nil)

	if err != nil {
		return err
	}

	return nil
}
