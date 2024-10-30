/*
Copyright © 2024 Simplex <kontakt.simplex@pm.me>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/SimplexDE/goTasks/cmd/goTasks"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add \"Task\"",
	Aliases: []string{
		"a",
		"create",
		"new",
		"n",
	},
	Short: "Add Task to the list",
	Long: `Add Task to the list`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := &gotasks.Task{}
		err := task.Add(args[0])
		
		if err != nil {
			panic(err)
		}

		fmt.Println("Task added!")
}}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
