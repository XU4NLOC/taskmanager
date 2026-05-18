package cmd

import (
	"fmt"
	"taskmanager/storage"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a task",
	Long:  "Delete the task with the specified ID.",
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		tasks, err := storage.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}
		for i := range tasks {
			if fmt.Sprintf("%d", tasks[i].ID) == title {
				tasks = append(tasks[:i], tasks[i+1:]...)
				break
			}
		}
		if err := storage.SaveTasks(tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
		fmt.Printf("Task with ID %s deleted.\n", title)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
