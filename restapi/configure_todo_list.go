// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	todoMiddleware "example.com/todo-list/middleware"
	"example.com/todo-list/models"
	"example.com/todo-list/restapi/operations"
	"example.com/todo-list/restapi/operations/todos"
)

// myPrincipal define the authenticated user's info
type myPrincipal struct {
	user string
}

// custom option flags
// note: please refer https://godoc.org/github.com/jessevdk/go-flags#hdr-Available_field_tags for option's tag usage
var option struct {
	// static file server's root path
	StaticFilePath string `long:"static" default:"./build" description:"Static file server's root path"`
}

var errUserSecurityAuth = errors.New(http.StatusUnauthorized, "user security authentication error")

// apiKey hard code token
const authenAPIKey = "Bearer MySecureAPIKey"

//go:generate swagger generate server --target ../../swagger-todo-list --name TodoList --spec ../swagger/swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.TodoListAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "configure flags",
			LongDescription:  "Todo configuration flags",
			Options:          &option,
		},
	}
}

func configureAPI(api *operations.TodoListAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the Authorization header is set with the Basic scheme
	api.UserSecurityAuth = func(token string) (interface{}, error) {
		log.Printf("UserSecurityAuth: %v\n", token)
		if token != authenAPIKey {
			return nil, errUserSecurityAuth
		}
		return myPrincipal{user: "admin"}, nil
	}

	api.TodosAddOneHandler = todos.AddOneHandlerFunc(func(params todos.AddOneParams, principal interface{}) middleware.Responder {
		api.Logger("add todo %+v\n", params.Body)
		return middleware.NotImplemented("operation todos.AddOne has not yet been implemented")
	})
	api.TodosDeleteOneHandler = todos.DeleteOneHandlerFunc(func(params todos.DeleteOneParams, principal interface{}) middleware.Responder {
		api.Logger("delete todo %+v\n", params.ID)
		return middleware.NotImplemented("operation todos.DestroyOne has not yet been implemented")
	})
	api.TodosFindTodosHandler = todos.FindTodosHandlerFunc(func(params todos.FindTodosParams, principal interface{}) middleware.Responder {
		payload := []*models.Item{
			{
				ID:          1,
				Description: swag.String("Todo task 1"),
				Completed:   false,
			},
			{
				ID:          2,
				Description: swag.String("Todo task 2"),
				Completed:   false,
			},
		}
		api.Logger("get todo %+v %+v %+v\n", principal, *params.Since, *params.Limit)
		return todos.NewFindTodosOK().WithPayload(payload)
	})
	api.TodosUpdateOneHandler = todos.UpdateOneHandlerFunc(func(params todos.UpdateOneParams, principal interface{}) middleware.Responder {
		api.Logger("update todo %+v\n", params.Body)
		return middleware.NotImplemented("operation todos.UpdateOne has not yet been implemented")
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
	log.Printf("configureServer()\n")
	switch scheme {
	case schemeHTTP:
		log.Printf("server, scheme: %v\n", scheme)
	case schemeHTTPS:
		log.Printf("server, scheme: %v\n", scheme)
	case schemeUnix:
		log.Printf("server, scheme: %v\n", scheme)
	}
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	log.Printf("setupMiddlewares()\n")
	return todoMiddleware.HandlerLogger(handler)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	log.Printf("setupGlobalMiddleware()\n")
	return todoMiddleware.Recover(
		todoMiddleware.Logger(
			todoMiddleware.StaticFileServer(option.StaticFilePath, handler),
		),
	)
}
