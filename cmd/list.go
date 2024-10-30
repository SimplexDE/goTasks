/*
Copyright © 2024 Simplex <kontakt.simplex@pm.me>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Tasks",
	Long: `List all Tasks`,
	Aliases: []string{
		"l",
		"tasks",
	},
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile("tasks.csv")
		r := csv.NewReader(strings.NewReader(string(data)))

		if err != nil {
			panic(err)
		}

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 8, 0, '\t', 0)

		fmt.Fprintln(w, "ID\tTask\tCreated\tDone")

		for {
			record, err := r.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}

			timestampInt, err := strconv.ParseInt(record[2], 10, 64)
			if err != nil {
				fmt.Fprintf(w, "Error parsing timestamp: %v\n", err)
				return
			}

			// Convert the Unix timestamp to a `time.Time` object
			timestamp := time.Unix(timestampInt, 0)

			time := timediff.TimeDiff(timestamp)
			
			fmt.Fprintf(w, "%v\t%v\t%v\t%v", record[0], record[1], time, record[3])
			fmt.Fprintln(w)
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
