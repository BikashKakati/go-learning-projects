package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Edit   string
	Delete int
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cmdFlags := CmdFlags{}

	flag.StringVar(&cmdFlags.Add, "add", "", "Add a new todo")
	flag.StringVar(&cmdFlags.Edit, "edit", "", "Edit a todo by giving index and specify the new todo")
	flag.IntVar(&cmdFlags.Delete, "delete", -1, "Specify a todo index to delete the todo")
	flag.IntVar(&cmdFlags.Toggle, "toggle", -1, "Specify a todo index to change the completion status")
	flag.BoolVar(&cmdFlags.List, "list", false, "List all todos")

	flag.Parse()
	return &cmdFlags
}

func (cmdFalgs *CmdFlags) Execute(todoList *TodoList) {
	switch {
	case cmdFalgs.List:
		todoList.printTodoList()
	case cmdFalgs.Add != "":
		todoList.addTodo(cmdFalgs.Add)
	case cmdFalgs.Edit != "":
		editTodoCmdParts := strings.SplitN(cmdFalgs.Edit, ":", 2)
		if len(editTodoCmdParts) != 2 {
			fmt.Println("Error: invalid format to edit todo, please use index:new_todo")
			os.Exit(1)
		}

		index, err := strconv.Atoi(editTodoCmdParts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		todoList.editTodo(index, editTodoCmdParts[1])
	case cmdFalgs.Toggle != -1:
		todoList.toggleTodo(cmdFalgs.Toggle)
	case cmdFalgs.Delete != -1:
		todoList.deleteTodo(cmdFalgs.Delete)
	default:
		fmt.Println("Invalid command!")
	}
}
