/*
Copyright © 2025 Shubham <thatcoderguyshubham@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-todo",
	Short: "A basic todo app for your cli",
	Long: `Welcome to the go-todo cli app. It is a basic todo app,
	You can create, view , delete and update the status of your todos with simple flags
	For the flag details type go-todo -help`,
}

var dueDateStr string

var createTodoCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create"},
	Short:   "Create new todo",
	Long:    "Create new todos and add due date. Default pending data for todos is 24h",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CreateTodo(args[0], dueDateStr)
	},
}

var viewTodosCmd = &cobra.Command{
	Use:   "view",
	Short: "View all todos",
	Long:  "Display all todos in a table, sorted by due date (ascending)",
	Run: func(cmd *cobra.Command, args []string) {
		ViewTodos()
	},
}

var viewTodoByIDCmd = &cobra.Command{
	Use:   "get",
	Short: "View todo by id",
	Long:  "Display todo by the id provided",
	Run: func(cmd *cobra.Command, args []string) {
		ViewTodoByID(args[0])
	},
}

var deleteTodoCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"remove"},
	Short:   "Delete any todo",
	Long:    "Delete any todo.Pass the id of todo to be deleted",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		DeleteTodo(args[0])
	},
}

var updateTodoStatusCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the status of a todo",
	Long:  "Update the status of a todo by providing its ID and the new status (pending, in-progress, completed)",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		UpdateTodoStatus(args[0], args[1])
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	createTodoCmd.Flags().StringVarP(&dueDateStr, "due", "d", "", "Due date for the todo (YYYY-MM-DD). Default is 24h from now.")
	rootCmd.AddCommand(createTodoCmd)
	rootCmd.AddCommand(viewTodosCmd)
	rootCmd.AddCommand(viewTodoByIDCmd)
	rootCmd.AddCommand(deleteTodoCmd)
	rootCmd.AddCommand(updateTodoStatusCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
