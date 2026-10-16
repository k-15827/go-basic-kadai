package main

import (
	"fmt"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

func (todo *Todo) Complete() {
	fmt.Printf("ID: %dのTodoを完了に更新します\n", todo.ID)
	todo.Completed = true
}

func printTodos(todos []Todo) {
	for _, todo := range todos {
		if todo.Completed {
			fmt.Printf("[完　了] (ID: %d) %s\n", todo.ID, todo.Title)
		} else {
			fmt.Printf("[未完了] (ID: %d) %s\n", todo.ID, todo.Title)
		}
	}
	fmt.Printf("")
}

func main() {
	todos := []Todo{
		{
			ID:        1,
			Title:     "学習計画",
			Completed: false,
		},
		{
			ID:        2,
			Title:     "環境構築",
			Completed: false,
		},
		{
			ID:        3,
			Title:     "基礎文法",
			Completed: false,
		},
	}

	printTodos(todos)

	for i, todo := range todos {
		if todo.ID == 1 || todo.ID == 2 {
			todos[i].Complete()
		}
	}

	printTodos(todos)
}
