/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/imduchuyyy/golang-patterns/worker"
	"github.com/spf13/cobra"
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Worker Pool Pattern",
	Long:  `In the worker pool pattern, we use multiple workers to handle jobs. Each worker is a separate process, and we have a way of assigning tasks to these workers. In Go, we use goroutines for the workers and channels to give them tasks and to receive the results back.`,
	Run: func(cmd *cobra.Command, args []string) {
		worker := worker.NewWorker()
		worker.Run()
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
