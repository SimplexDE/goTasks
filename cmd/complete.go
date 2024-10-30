/*
Copyright © 2024 Simplex <kontakt.simplex@pm.me>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/SimplexDE/goTasks/cmd/goTasks"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete Task-ID",
	Short: "Complete a Task",
	Long: `Complete a Task`,
	Aliases: []string{
		"done",
		"c",
	},
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := &gotasks.Task{}

		_id, err := strconv.Atoi(args[0])

		if err != nil {
			panic(err)
		}

		err2 := task.Complete(int32(_id))
		
		if err2 != nil {
			panic(err2)
		}

		fmt.Println("Task completed!")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
