/*
Copyright © 2024 Simplex <kontakt.simplex@pm.me>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/SimplexDE/goTasks/cmd/goTasks"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears the task list",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := gotasks.ClearTasks()
		if err != nil {
			panic(err)
		}
		fmt.Println("Cleared all Tasks")
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
