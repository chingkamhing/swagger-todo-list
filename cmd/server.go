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

	api := operations.NewTodoListAPI(swaggerSpec)
	server = restapi.NewServer(api)
	parser.AddCommand("server",
		"Todo server",
		"The server command run todo-list in server mode.",
		server)
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		// FIXME, should put the options under command "server"
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			defer server.Shutdown()
			log.Fatalln(err)
		}
	}
}
