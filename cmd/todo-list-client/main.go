package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"example.com/todo-list/client"
	"example.com/todo-list/client/todos"
)

type myClient struct{}

var errAuthenticateRequest = errors.New("authenticate request error")

func (client *myClient) AuthenticateRequest(request runtime.ClientRequest, registry strfmt.Registry) error {
	// FIXME, make the code compile first, need to add basic authentication logic
	return errAuthenticateRequest
}

func main() {
	flag.Parse()
	cfg := &client.TransportConfig{
		Host:     "localhost:8888",
		BasePath: "/",
		Schemes:  []string{"http"},
	}
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := &todos.FindTodosParams{Context: context.Background()}
	infoWriter := &myClient{}
	response, err := c.Todos.FindTodos(params, infoWriter)
	if err != nil {
		log.Fatal(err)
	}
	for _, todo := range response.Payload {
		fmt.Printf("%4d: Description=%q Completed=%v\n", todo.ID, *todo.Description, todo.Completed)
	}
}
