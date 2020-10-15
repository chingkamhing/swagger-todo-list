// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new todos API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for todos API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddOne(params *AddOneParams) (*AddOneCreated, error)

	DeleteOne(params *DeleteOneParams) (*DeleteOneNoContent, error)

	FindTodos(params *FindTodosParams) (*FindTodosOK, error)

	UpdateOne(params *UpdateOneParams) (*UpdateOneOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddOne add one API
*/
func (a *Client) AddOne(params *AddOneParams) (*AddOneCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addOne",
		Method:             "POST",
		PathPattern:        "/todo/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddOneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddOneCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddOneDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteOne delete one API
*/
func (a *Client) DeleteOne(params *DeleteOneParams) (*DeleteOneNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOne",
		Method:             "DELETE",
		PathPattern:        "/todo/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteOneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteOneNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteOneDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FindTodos find todos API
*/
func (a *Client) FindTodos(params *FindTodosParams) (*FindTodosOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindTodosParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findTodos",
		Method:             "GET",
		PathPattern:        "/todo/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FindTodosReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindTodosOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindTodosDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateOne update one API
*/
func (a *Client) UpdateOne(params *UpdateOneParams) (*UpdateOneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOne",
		Method:             "PUT",
		PathPattern:        "/todo/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateOneOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateOneDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
