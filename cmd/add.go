package cmd

import (
	"fmt"
	"time"
	"taskmanager/storage"
	
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task title]",
	Short: "Add a new task",
	Long: "Create a new task with the specified title.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd*cobra.Command, args []string){
		title := args[0]
		tasks, err := storage.LoadTasks()
		if err != nil{
			fmt.Println("Error loading tasks:", err)
			return
		}
		newTask := storage.Task{
			ID:        storage.NextID(tasks),
			Title:     title,
			Done:      false,
			CreatedAt: time.Now().Format(time.RFC3339),
		}
		tasks = append(tasks, newTask)
		if err := storage.SaveTasks(tasks); err != nil {
			fmt.Println("Error saving task:", err)
			return
		}
		fmt.Printf("Task added: %s\n", title)
	},
}

func init(){
	rootCmd.AddCommand(addCmd)
}
