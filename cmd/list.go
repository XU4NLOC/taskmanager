package cmd

import (
	"fmt"
	"taskmanager/storage"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "Display a list of all tasks with their details.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := storage.LoadTasks()
		if err != nil {
			fmt.Println("Error fetching tasks:", err)
			return
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		fmt.Println("ID\tTitle\tDone\tCreated At")
		for _, task := range tasks {
			fmt.Printf("%d\t%s\t%t\t%s\n", task.ID, task.Title, task.Done, task.CreatedAt)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
