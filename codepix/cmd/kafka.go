/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/alexdevelp/code-pix/application/kafka"
	"github.com/alexdevelp/code-pix/infrastructure/db"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start cosuming transactions using Apache Kafka",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("produzindo mensagem")

		producer := kafka.NewKafkaProducer()
		database := db.ConnectDB(os.Getenv("env"))
		deliveryChan := make(chan ckafka.Event)

		//kafka.Publish("Olá Kafka - consumer", "Teste", producer, deliveryChan)
		// rodando em processo paralelo - outra thread
		go kafka.DeliveryReport(deliveryChan)

		//Criando um Processor kafka e consumindo
		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)
}
