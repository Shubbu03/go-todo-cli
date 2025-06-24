package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

type TodoStatus string

const (
	TodoStatusPending    TodoStatus = "pending"
	TodoStatusInProgress TodoStatus = "in-progress"
	TodoStatusCompleted  TodoStatus = "completed"
)

type Todo struct {
	ID      int        `json:"id"`
	Title   string     `json:"title"`
	DueDate time.Time  `json:"due_date"`
	Status  TodoStatus `json:"status"`
}

func CreateTodo(title string, due string) {
	var todos []Todo
	const fileName = "todos.json"

	file, err := os.Open(fileName)
	if err == nil {
		defer file.Close()
		json.NewDecoder(file).Decode(&todos)
	}

	maxID := 0
	for _, t := range todos {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	var dueDate time.Time
	if due != "" {
		dueDate, err = time.Parse(time.RFC3339, due)
		if err != nil {
			dueDate, err = time.Parse("2006-01-02", due)
			if err != nil {
				dueDate = time.Now().Add(24 * time.Hour)
			}
		}
	} else {
		dueDate = time.Now().Add(24 * time.Hour)
	}

	todo := Todo{
		ID:      maxID + 1,
		Title:   title,
		DueDate: dueDate,
		Status:  TodoStatusPending,
	}
	todos = append(todos, todo)

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(todos)
}

func ViewTodos() {
	const fileName = "todos.json"
	var todos []Todo

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("No todos found.")
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&todos)

	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}

	sort.Slice(todos, func(i, j int) bool {
		return todos[i].DueDate.Before(todos[j].DueDate)
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Title", "Due Date", "Status"})

	for _, t := range todos {
		var statusColored string
		switch t.Status {
		case TodoStatusPending:
			statusColored = color.New(color.FgRed).Sprint(string(t.Status))
		case TodoStatusInProgress:
			statusColored = color.New(color.FgYellow).Sprint(string(t.Status))
		case TodoStatusCompleted:
			statusColored = color.New(color.FgGreen).Sprint(string(t.Status))
		default:
			statusColored = string(t.Status)
		}
		table.Append([]string{
			fmt.Sprintf("%d", t.ID),
			t.Title,
			t.DueDate.Format("2006-01-02 15:04"),
			statusColored,
		})
	}
	table.Render()
}

func ViewTodoByID(todoID string) {
	const fileName = "todos.json"
	var todos []Todo

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("No todos found.")
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&todos)

	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}

	id, err := strconv.Atoi(todoID)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	found := false
	for _, t := range todos {
		if t.ID == id {
			fmt.Printf("ID: %d\nTitle: %s\nDue Date: %s\nStatus: %s\n", t.ID, t.Title, t.DueDate.Format("2006-01-02 15:04"), t.Status)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Todo not found.")
	}
}

func DeleteTodo(todoID string) {
	const fileName = "todos.json"
	var todos []Todo

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("No todos found.")
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&todos)

	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}

	id, err := strconv.Atoi(todoID)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	newTodos := make([]Todo, 0, len(todos))
	found := false
	for _, t := range todos {
		if t.ID == id {
			found = true
			continue
		}
		newTodos = append(newTodos, t)
	}

	if !found {
		fmt.Println("Todo not found.")
		return
	}

	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Failed to update todos.")
		return
	}
	defer f.Close()
	json.NewEncoder(f).Encode(newTodos)
	fmt.Println("Todo deleted successfully.")
}

func UpdateTodoStatus(todoID string, status string) {
	const fileName = "todos.json"
	var todos []Todo

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("No todos found.")
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&todos)

	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}

	id, err := strconv.Atoi(todoID)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	var newStatus TodoStatus
	switch status {
	case string(TodoStatusPending):
		newStatus = TodoStatusPending
	case string(TodoStatusInProgress):
		newStatus = TodoStatusInProgress
	case string(TodoStatusCompleted):
		newStatus = TodoStatusCompleted
	default:
		fmt.Println("Invalid status. Use 'pending', 'in-progress', or 'completed'.")
		return
	}

	found := false
	for i, t := range todos {
		if t.ID == id {
			todos[i].Status = newStatus
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Todo not found.")
		return
	}

	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Failed to update todos.")
		return
	}
	defer f.Close()
	json.NewEncoder(f).Encode(todos)
	fmt.Println("Todo status updated successfully.")
}
