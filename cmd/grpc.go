/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/alexdevelp/code-pix/application/grpc"
	"github.com/alexdevelp/code-pix/infrastructure/db"
	"github.com/spf13/cobra"
)

var portNumber int
var testAlexandre string

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Use codepix software to intermediate bank transaction with apache kafka and grpc",
	// Long: `A longer description that spans multiple lines and likely contains examples
	// 	and usage of using your command. For example:

	// 	Cobra is a CLI library for Go that empowers applications.
	// 	This application is a tool to generate the needed files
	// 	to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		grpc.StartGrpcServer(database, portNumber)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().IntVarP(&portNumber, "port", "p", portNumber, "gRPC server port")
	grpcCmd.Flags().StringVarP(&testAlexandre, "author", "a", "Alexandre", "test author name default")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
