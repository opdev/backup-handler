// Code generated by goa v3.7.3, DO NOT EDIT.
//
// Backup Service client HTTP transport
//
// Command:
// $ goa gen github.com/opdev/backup-handler/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the Backup Service service endpoint HTTP clients.
type Client struct {
	// Create Doer is the HTTP client used to make requests to the create endpoint.
	CreateDoer goahttp.Doer

	// Get Doer is the HTTP client used to make requests to the get endpoint.
	GetDoer goahttp.Doer

	// Update Doer is the HTTP client used to make requests to the update endpoint.
	UpdateDoer goahttp.Doer

	// Delete Doer is the HTTP client used to make requests to the delete endpoint.
	DeleteDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the Backup Service service
// servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		CreateDoer:          doer,
		GetDoer:             doer,
		UpdateDoer:          doer,
		DeleteDoer:          doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Create returns an endpoint that makes HTTP requests to the Backup Service
// service create server.
func (c *Client) Create() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateRequest(c.encoder)
		decodeResponse = DecodeCreateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("Backup Service", "create", err)
		}
		return decodeResponse(resp)
	}
}

// Get returns an endpoint that makes HTTP requests to the Backup Service
// service get server.
func (c *Client) Get() goa.Endpoint {
	var (
		decodeResponse = DecodeGetResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("Backup Service", "get", err)
		}
		return decodeResponse(resp)
	}
}

// Update returns an endpoint that makes HTTP requests to the Backup Service
// service update server.
func (c *Client) Update() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateRequest(c.encoder)
		decodeResponse = DecodeUpdateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("Backup Service", "update", err)
		}
		return decodeResponse(resp)
	}
}

// Delete returns an endpoint that makes HTTP requests to the Backup Service
// service delete server.
func (c *Client) Delete() goa.Endpoint {
	var (
		decodeResponse = DecodeDeleteResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("Backup Service", "delete", err)
		}
		return decodeResponse(resp)
	}
}
