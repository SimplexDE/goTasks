package gotasks

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type Task struct {
	id      int32
	task    string
	created int64
	done    bool
}

func (t *Task) Add(name string) error {
	nextID := int32(1)

	file, err := os.OpenFile("tasks.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

		// Open the file in read mode to find the last ID
		readFile, err := os.Open("tasks.csv")
		if err != nil {
			return err
		}
		defer readFile.Close()

		reader := csv.NewReader(readFile)

		// Iterate through the rows to find the last ID
		for {
			record, err := reader.Read()
			if err != nil {
				break
			}
			id, err := strconv.Atoi(record[0])
			if err != nil {
				return err
			}
			if int32(id) >= nextID {
				nextID = int32(id) + 1
			}

	}

	// Create a new task with the incremented ID
	newTask := &Task{
		id:      nextID,
		task:    name,
		created: time.Now().Unix(),
		done:    false,
	}

	// Write the new task record to the CSV file
	record := []string{
		strconv.FormatInt(int64(newTask.id), 10),
		newTask.task,
		strconv.FormatInt(newTask.created, 10),
		strconv.FormatBool(newTask.done),
	}

	if err := writer.Write(record); err != nil {
		return err
	}
	writer.Flush() // Ensure all data is written

	return nil
}

// Complete marks a task as completed by setting the `done` field to `true` in "tasks.csv"
func (t *Task) Complete(taskID int32) error {
	// Read all tasks from the CSV file
	file, err := os.Open("tasks.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Update the `done` field for the task with the given ID
	for i, record := range records {
		if i == 0 {
			// Skip the header row
			continue
		}
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return err
		}
		if int32(id) == taskID {
			// Set `done` to true for the matched task ID
			record[3] = "true"
			break
		}
	}

	// Write the updated records back to the CSV file
	file, err = os.OpenFile("tasks.csv", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}