# Todo List Tutorial

This is a simple server swagger tutorial that is based on [Todo List Tutorial](https://goswagger.io/tutorial/todo-list.html#todo-list-tutorial).

This also evaluate what go-swagger generated server can and cannot do.

CAN:
* API first design
    + design API with swagger.yml
    + then generate a template server code that provide all endpoint dummy handlers
    + then implement all the endpoint handlers
    + any change of endpoint (i.e. swagger.yml), just redo the generation of server code and implement the corresponding handlers
* middleware
    + support native net/http middleware
    + able to add global middleware to handle something like panic recovery, log, compression, etc.
    + able to add after-route middleware to handle something like timeout, session, etc.
* models
    + ??? how to use the generated models (e.g. User, ToDoItem) with DB without translation (i.e. no copying from one struct to another struct)?
    + ==> may use 'x-go-custom-tag: "string"' to add tag in the fields of the models
    + please refer to [Schema generation rules](https://goswagger.io/use/models/schemas.html) for detail
* authentication
    + support Basic Authentication, API key and Oauth 2 authenication
    + need to update swagger.yml to configure which endpoints need authentication
    + please refer to [Authentication sample](https://goswagger.io/tutorial/authentication/) for detail
* session
    + ??? how to handle session (i.e. cookies)?
* static file server
    + does not natively support file server
    + need to add a middleware to route between api (e.g. endpoint start with "/api/*" and file server (e.g. path with "/*")
* run-time response type
    + ??? support run-time determine response type (e.g. html or json)?

CANNOT:
* middleware
    + seems cannot insert middleware into specific path or endpoint

## Knowledge Base

* [Writing OpenAPI (Swagger) Specification Tutorial Series](https://apihandyman.io/writing-openapi-swagger-specification-tutorial-part-1-introduction/)
