/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/imduchuyyy/golang-patterns/generator"
	"github.com/spf13/cobra"
)

// generatorCmd represents the generator command
var generatorCmd = &cobra.Command{
	Use:   "generator",
	Short: "Generator Pattern",
	Long:  `The generator pattern is a way to create a function that gives us a series of values. In other programming languages, people might use things called enumerators or iterators. But in Go, the best tool for this job is a channel.`,
	Run: func(cmd *cobra.Command, args []string) {
		generator := generator.NewGenerator()
		generator.Run()
	},
}

func init() {
	rootCmd.AddCommand(generatorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
