# Todo List Tutorial

This is a simple server swagger tutorial that is based on [Todo List Tutorial](https://goswagger.io/tutorial/todo-list.html#todo-list-tutorial).

This also evaluate what goswagger generated server can and cannot do.

CAN:
* API first design
    + design API with swagger.yml
    + then generate a template server code that provide all endpoint dummy handlers
    + then implement all the endpoint handlers
    + any change of endpoint (i.e. swagger.yml), just redo the generation of server code and implement the corresponding handlers
* add middleware
    + support net/http style middleware
    + 

CANNOT:
