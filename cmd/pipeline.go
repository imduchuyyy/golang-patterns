/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/imduchuyyy/golang-patterns/pipeline"
	"github.com/spf13/cobra"
)

// pipelineCmd represents the pipeline command
var pipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Pipeline pattern",
	Long:  `The pipeline pattern is a way of splitting tasks where each step does a particular job. The output of one step becomes the input for the next and this technique is good for tasks that can be divided into smaller, separate jobs.`,
	Run: func(cmd *cobra.Command, args []string) {
		pipeline := pipeline.NewPipeline()
		pipeline.Run()
	},
}

func init() {
	rootCmd.AddCommand(pipelineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pipelineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pipelineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
