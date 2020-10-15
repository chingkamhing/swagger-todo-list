package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"example.com/todo-list/client"
	"example.com/todo-list/client/todos"
)

func main() {
	flag.Parse()
	cfg := &client.TransportConfig{
		Host:     "localhost:8888",
		BasePath: "/",
		Schemes:  []string{"http"},
	}
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := &todos.FindTodosParams{Context: context.Background()}
	response, err := c.Todos.FindTodos(params)
	if err != nil {
		log.Fatal(err)
	}
	for _, todo := range response.Payload {
		fmt.Printf("%4d: Description=%q Completed=%v\n", todo.ID, *todo.Description, todo.Completed)
	}
}
