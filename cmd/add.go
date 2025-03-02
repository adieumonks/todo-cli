/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

type Todo struct {
	Task string `json:"task"`
	Done bool   `json:"done"`
}

var task string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Long: `Add a new task to your TODO list. For example:

todo-cli add "Learn Go"`,
	Run: func(cmd *cobra.Command, args []string) {
		addTask(task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&task, "task", "t", "", "Task to add to your TODO list")
	addCmd.MarkFlagRequired("task")
}

func addTask(task string) {
	todos := loadTodos()
	todos = append(todos, Todo{Task: task, Done: false})
	saveTodos(todos)
	fmt.Println("Task added: ", task)
}

func loadTodos() []Todo {
	return []Todo{}
}

func saveTodos(todos []Todo) {
	data, _ := json.MarshalIndent(todos, "", "  ")
	fmt.Println(string(data))
}
