/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/alexdevelp/code-pix/application/grpc"
	"github.com/alexdevelp/code-pix/application/kafka"
	"github.com/alexdevelp/code-pix/infrastructure/db"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
)

var (
	gRPCPortNumber int
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run gRPC and Apache Kafka",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("all called")

		// Iniciando servidor gRPC em outra thread
		database := db.ConnectDB(os.Getenv("env"))
		go grpc.StartGrpcServer(database, gRPCPortNumber)

		producer := kafka.NewKafkaProducer()
		deliveryChan := make(chan ckafka.Event)

		// rodando em processo paralelo - outra thread
		go kafka.DeliveryReport(deliveryChan)

		//Criando um Processor kafka e consumindo
		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		kafkaProcessor.Consume()

	},
}

func init() {
	rootCmd.AddCommand(allCmd)
	allCmd.Flags().IntVarP(&gRPCPortNumber, "grpc-port", "p", 50051, "gRPC Port")
}
