/*
Copyright © 2024 Simplex <kontakt.simplex@pm.me>
*/
package cmd

import (
	"fmt"

	"github.com/SimplexDE/goTasks/cmd/goTasks"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete Task-ID",
	Short: "Delete a task out of the list",
	Long: ``,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := &gotasks.Task{}

		err2 := task.Delete(args[0])

		if err2 != nil {
			panic(err2)
		}

		fmt.Println("Deleted Task")
	},
}


func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
