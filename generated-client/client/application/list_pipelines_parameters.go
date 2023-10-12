// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListPipelinesParams creates a new ListPipelinesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPipelinesParams() *ListPipelinesParams {
	return &ListPipelinesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPipelinesParamsWithTimeout creates a new ListPipelinesParams object
// with the ability to set a timeout on a request.
func NewListPipelinesParamsWithTimeout(timeout time.Duration) *ListPipelinesParams {
	return &ListPipelinesParams{
		timeout: timeout,
	}
}

// NewListPipelinesParamsWithContext creates a new ListPipelinesParams object
// with the ability to set a context for a request.
func NewListPipelinesParamsWithContext(ctx context.Context) *ListPipelinesParams {
	return &ListPipelinesParams{
		Context: ctx,
	}
}

// NewListPipelinesParamsWithHTTPClient creates a new ListPipelinesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPipelinesParamsWithHTTPClient(client *http.Client) *ListPipelinesParams {
	return &ListPipelinesParams{
		HTTPClient: client,
	}
}

/*
ListPipelinesParams contains all the parameters to send to the API endpoint

	for the list pipelines operation.

	Typically these are written to a http.Request.
*/
type ListPipelinesParams struct {

	/* AppName.

	   Name of application
	*/
	AppName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list pipelines params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPipelinesParams) WithDefaults() *ListPipelinesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list pipelines params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPipelinesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list pipelines params
func (o *ListPipelinesParams) WithTimeout(timeout time.Duration) *ListPipelinesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list pipelines params
func (o *ListPipelinesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list pipelines params
func (o *ListPipelinesParams) WithContext(ctx context.Context) *ListPipelinesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list pipelines params
func (o *ListPipelinesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list pipelines params
func (o *ListPipelinesParams) WithHTTPClient(client *http.Client) *ListPipelinesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list pipelines params
func (o *ListPipelinesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the list pipelines params
func (o *ListPipelinesParams) WithAppName(appName string) *ListPipelinesParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the list pipelines params
func (o *ListPipelinesParams) SetAppName(appName string) {
	o.AppName = appName
}

// WriteToRequest writes these params to a swagger request
func (o *ListPipelinesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appName
	if err := r.SetPathParam("appName", o.AppName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}