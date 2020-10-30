package main

import (
	"context"
	"errors"
	"fmt"

	"example.com/todo-list/client"
	"example.com/todo-list/client/todos"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ClientOption define flags option struct
type ClientOption struct {
	Schemes  []string `long:"scheme" description:"the listeners to enable, this can be repeated and defaults to the schemes in the swagger spec"`
	Host     string   `long:"host" default:"localhost" env:"HOST" description:"the IP to listen on"`
	Port     int      `long:"port" default:"8888" env:"PORT" description:"the port to listen on for insecure connections, defaults to a random value"`
	BasePath string   `long:"base" default:"/" description:"base endpoint path"`
}

var clientOption = &ClientOption{}

var errAuthenticateRequest = errors.New("authenticate request error")

func init() {
	parser.AddCommand("client",
		"Todo client",
		"The client command to get todo list.",
		clientOption)
}

// AuthenticateRequest implement ClientAuthInfoWriter interface to make client request
func (option *ClientOption) AuthenticateRequest(request runtime.ClientRequest, registry strfmt.Registry) error {
	// FIXME, make the code compile first, need to add basic authentication logic
	request.SetHeaderParam("Authorization", "Bearer MySecureAPIKey")
	request.SetQueryParam("since", "1604056678")
	request.SetQueryParam("limit", "10")
	return nil
}

// Execute implement jessevdk/go-flags command interface
func (option *ClientOption) Execute(args []string) error {
	cfg := &client.TransportConfig{
		Host:     fmt.Sprintf("%v:%v", clientOption.Host, clientOption.Port),
		BasePath: clientOption.BasePath,
		Schemes:  clientOption.Schemes,
	}
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := &todos.FindTodosParams{Context: context.Background()}
	response, err := c.Todos.FindTodos(params, clientOption)
	if err != nil {
		return err
	}
	for _, todo := range response.Payload {
		fmt.Printf("%4d: Description=%q Completed=%v\n", todo.ID, *todo.Description, todo.Completed)
	}
	return nil
}
