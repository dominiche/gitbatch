package main

import (
	"fmt"
	"gitbatch/cmd/checkout"
	"gitbatch/cmd/clone"
	"gitbatch/cmd/fetch"
	"gitbatch/cmd/pull"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gitbatch",
	Short: "git batch tool",
	Long:  `A git batch tool. Support gitlab only for now.`,
}

func init() {
	rootCmd.AddCommand(clone.Cmd())
	rootCmd.AddCommand(checkout.Cmd())
	rootCmd.AddCommand(pull.Cmd())
	rootCmd.AddCommand(fetch.Cmd())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
