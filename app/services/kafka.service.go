package services

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"src/helpers"
)

var Producer sarama.SyncProducer

func CreateKafkaProducer() sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// SASL/PLAIN
	//config.Net.SASL.Enable = true
	//config.Net.SASL.User = helpers.GetEnv("KAFKA_USERNAME", "root")
	//config.Net.SASL.Password = helpers.GetEnv("KAFKA_PASSWORD", "secret")
	//config.Net.SASL.Handshake = true
	//config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	//
	//tlsConfig := tls.Config{}
	//config.Net.TLS.Enable = true
	//config.Net.TLS.Config = &tlsConfig

	host := helpers.GetEnv("KAFKA_HOST", "beego_kafka_local")
	port := helpers.GetEnv("KAFKA_PORT", "9092")
	brokers := []string{fmt.Sprintf("%s:%s", host, port)}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Printf("Failed to create Kafka producer: %v", err)
	}

	Producer = producer

	return Producer
}
