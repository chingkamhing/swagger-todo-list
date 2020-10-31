package main

import (
	"log"

	"github.com/go-openapi/loads"

	"example.com/todo-list/restapi"
	"example.com/todo-list/restapi/operations"
)

var server *restapi.Server

func init() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}
	checkError := func(err error) {
		if err != nil {
			defer server.Shutdown()
			log.Fatalln(err)
		}
	}

	api := operations.NewTodoListAPI(swaggerSpec)
	server = restapi.NewServer(api)
	cmd, err := parser.AddCommand("server",
		"Todo server",
		"The server command run todo-list in server mode.",
		server)
	checkError(err)
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := cmd.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		checkError(err)
	}
}
