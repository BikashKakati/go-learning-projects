package main

func main() {
	todoList := TodoList{}
	storage := NewStorage[TodoList]("todo-list.json")
	storage.Load(&todoList)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todoList)
	storage.Save(todoList)
}
