package main

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type TodoList []Todo

func (todoList *TodoList) addTodo(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todoList = append(*todoList, todo)
}

func (todoList *TodoList) validateIndex(id int) error {
	if id < 0 || id >= len(*todoList) {
		err := errors.New("Invalid Id, please check the Id")
		return err
	}
	return nil
}

func (todoList *TodoList) deleteTodo(id int) error {
	t := *todoList
	err := t.validateIndex(id)

	if err != nil {
		return err
	}

	*todoList = append(t[:id], t[id+1:]...)
	return nil
}

func (todoList *TodoList) toggleTodo(id int) error {
	t := *todoList
	err := t.validateIndex(id)

	if err != nil {
		return err
	}

	isCompleted := t[id].Completed
	if !isCompleted {
		completionTime := time.Now()
		t[id].CompletedAt = &completionTime
	}

	t[id].Completed = !isCompleted
	return nil
}

func (todoList *TodoList) editTodo(id int, newTitle string) error {
	t := *todoList
	err := t.validateIndex(id)

	if err != nil {
		return err
	}

	t[id].Title = newTitle
	return nil
}

func (todoList *TodoList) printTodoList() {
	table := table.New(os.Stdout)

	table.SetRowLines(false)
	table.SetHeaders("#Id", "Title", "Completed", "CreatedAt", "CompletedAt")

	for id, todo := range *todoList {
		completed := "❌"
		completedAt := ""

		if todo.Completed {
			completed = "✅"

			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(
			strconv.Itoa(id),
			todo.Title,
			completed,
			todo.CreatedAt.Format(time.RFC1123),
			completedAt,
		)
	}

	table.Render()
}
