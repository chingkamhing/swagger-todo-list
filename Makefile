# build the source to binary
.PHONY: build
build:
	go build -o todo-list-server cmd/todo-list-server/*.go
	go build -o todo-list-client cmd/todo-list-client/*.go

# - generate swagger source codes
# - build the sources to binar
.PHONY: all
all: generate build

# go update libraries
.PHONY: update
update:
	go get -u ./...

# this command is used to generate an initial swagger.yml to begin with
# after the initial swagger.yml is generated, add definitions and paths as needed
# once all definitions and paths are defined, invoke "make generate" to generte server code
.PHONY: init
init:
	swagger init spec --title "A Todo list application" --description "From the todo list tutorial on goswagger.io" --version 1.0.0 --scheme http --consumes application/io.goswagger.examples.todo-list.v1+json --produces application/io.goswagger.examples.todo-list.v1+json

# validate if swagger.yml is valid
.PHONY: validate
validate:
	swagger validate ./swagger.yml

# generate server source code base on input of swagger.yml
.PHONY: generate
generate: validate
	swagger generate server --name TodoList --spec ./swagger.yml --principal interface{}
	swagger generate client --name TodoList --spec ./swagger.yml --principal interface{}

# clean all the binary and the generated code
# note: leave restapi/configure_todo_list.go untouch
.PHONY: clean
clean: validate
	rm -rf client cmd/todo-list-server models restapi/operations restapi/doc.go restapi/embedded_spec.go restapi/server.go
