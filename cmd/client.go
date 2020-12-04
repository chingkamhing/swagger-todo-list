package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"example.com/todo-list/client"
	"example.com/todo-list/client/todos"
	"example.com/todo-list/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// clientOptionStruct define client command flags option struct
type clientOptionStruct struct {
	Schemes  []string `long:"scheme" description:"the listeners to enable, this can be repeated and defaults to the schemes in the swagger spec"`
	Host     string   `long:"host" default:"localhost" env:"HOST" description:"the IP to listen on"`
	Port     int      `long:"port" default:"8888" env:"PORT" description:"the port to listen on for insecure connections, defaults to a random value"`
	BasePath string   `long:"base" default:"/" description:"base endpoint path"`
}

var clientOption = &clientOptionStruct{}

// getOptionStruct define get sub command flags option struct
type getOptionStruct struct{}

var getOption = &getOptionStruct{}

// updateOptionStruct define get sub command flags option struct
type updateOptionStruct struct{}

var updateOption = &updateOptionStruct{}

// deleteOptionStruct define get sub command flags option struct
type deleteOptionStruct struct{}

var deleteOption = &deleteOptionStruct{}

var errAuthenticateRequest = errors.New("authenticate request error")

func init() {
	checkError := func(err error) {
		if err != nil {
			log.Fatalf("checkError: %v\n", err)
		}
	}
	// client command
	cmdClient, err := parser.AddCommand("client",
		"Todo client commands",
		"The client commands to send different requests to todo server.",
		clientOption)
	checkError(err)

	// sub-commands
	{
		_, err := cmdClient.AddCommand("get",
			"Get todo list",
			"The client command to get todo list.",
			getOption)
		checkError(err)
	}
	{
		_, err := cmdClient.AddCommand("update",
			"Update a todo task",
			"The client command to update a todo task.",
			updateOption)
		checkError(err)
	}
	{
		_, err := cmdClient.AddCommand("delete",
			"Delete a todo task",
			"The client command to delete a todo task.",
			deleteOption)
		checkError(err)
	}
}

// client command Execute implementation function
func (option *clientOptionStruct) Execute(args []string) error {
	parser.WriteHelp(os.Stdout)
	return nil
}

// AuthenticateRequest implement ClientAuthInfoWriter interface to make client request
func (option *getOptionStruct) AuthenticateRequest(request runtime.ClientRequest, registry strfmt.Registry) error {
	request.SetHeaderParam("Authorization", "Bearer MySecureAPIKey")
	return nil
}

// get sub command Execute implementation function
func (option *getOptionStruct) Execute(args []string) error {
	cfg := &client.TransportConfig{
		Host:     fmt.Sprintf("%v:%v", clientOption.Host, clientOption.Port),
		BasePath: clientOption.BasePath,
		Schemes:  clientOption.Schemes,
	}
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := &todos.FindTodosParams{
		Context: context.Background(),
		Since:   swag.Int64(1604056678),
		Limit:   swag.Int32(10),
	}
	response, err := c.Todos.FindTodos(params, getOption)
	if err != nil {
		return err
	}
	for _, todo := range response.Payload {
		fmt.Printf("%4d: Description=%q Completed=%v\n", todo.ID, todo.Description, todo.Completed)
	}
	return nil
}

// AuthenticateRequest implement ClientAuthInfoWriter interface to make client request
func (option *updateOptionStruct) AuthenticateRequest(request runtime.ClientRequest, registry strfmt.Registry) error {
	request.SetHeaderParam("Authorization", "Bearer MySecureAPIKey")
	return nil
}

// update sub command Execute implementation function
func (option *updateOptionStruct) Execute(args []string) error {
	cfg := &client.TransportConfig{
		Host:     fmt.Sprintf("%v:%v", clientOption.Host, clientOption.Port),
		BasePath: clientOption.BasePath,
		Schemes:  clientOption.Schemes,
	}
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := &todos.UpdateOneParams{
		Context: context.Background(),
		Body: &models.Item{
			ID:          1,
			Description: "Todo task 1 (updated)",
			Completed:   true,
		},
	}
	response, err := c.Todos.UpdateOne(params, updateOption)
	if err != nil {
		log.Printf("err: %v\n", err)
		return err
	}
	fmt.Printf("Updated todo: %v\n", response.Payload.ID)
	return nil
}

// AuthenticateRequest implement ClientAuthInfoWriter interface to make client request
func (option *deleteOptionStruct) AuthenticateRequest(request runtime.ClientRequest, registry strfmt.Registry) error {
	request.SetHeaderParam("Authorization", "Bearer MySecureAPIKey")
	return nil
}

// delete sub command Execute implementation function
func (option *deleteOptionStruct) Execute(args []string) error {
	var deleteID int64 = 1
	cfg := &client.TransportConfig{
		Host:     fmt.Sprintf("%v:%v", clientOption.Host, clientOption.Port),
		BasePath: clientOption.BasePath,
		Schemes:  clientOption.Schemes,
	}
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := &todos.DeleteOneParams{
		Context: context.Background(),
		ID:      deleteID,
	}
	_, err := c.Todos.DeleteOne(params, deleteOption)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted todo: %v\n", deleteID)
	return nil
}
