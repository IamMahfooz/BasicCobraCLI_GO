package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
	"task/db"
)

func init() {
	RootCmd.AddCommand(addCmd)
	RootCmd.AddCommand(doCmd)
	RootCmd.AddCommand(listCmd)

}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds task to the tasklist",
	//Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("cannot create ,got an error:", err)
		}
		fmt.Printf("Task sucessfully created with id number :%d\n", id)
	},
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task to be complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to pass argument: ", arg)
				return
			} else {
				ids = append(ids, id)
			}
			fmt.Println(ids)
		}
		//tasks, err := db.AllTask()
		//if err != nil {
		//	fmt.Println("something went wrong", err)
		//	return
		//}
		for _, id := range ids {
			//if id <= 0 || id >= tasks[len(tasks)-1].Key {
			//	fmt.Println("Invalid task number :", id)
			//	return
			//}
			//task := tasks[id]
			err := db.DeleteTask(id)
			if err != nil {
				fmt.Printf("Failed to mark the task %d as completed. Error %s", id, err)
			} else {
				fmt.Printf("Marked Task: %d as completed\n", id)
			}
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "displays all the task in the bucker",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTask()
		if err != nil {
			fmt.Println("something went wrong:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no task pending")
		}
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s ;KEY=%d\n", i+1, task.Value, task.Key)
		}
	},
}
