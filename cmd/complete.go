package cmd

import (
	"fmt"
	"taskmanager/storage"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [task ID]",
	Short: "Mark a task as completed",
	Long:  "Mark the task with the specified ID as completed.",
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		tasks, err := storage.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}
		for i, task := range tasks {
			if fmt.Sprintf("%d", task.ID) == title {
				if task.Done {
					fmt.Printf("Task %d is already completed.\n", task.ID)
					return
				}
				tasks[i].Done = true
				if err := storage.SaveTasks(tasks); err != nil {
					fmt.Println("Error saving tasks:", err)
					return
				}
				fmt.Printf("Task %d marked as completed.\n", task.ID)
				return
			}
		}
		fmt.Printf("Task with ID %s not found.\n", title)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
