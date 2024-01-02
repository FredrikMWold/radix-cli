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

	"github.com/equinor/radix-cli/generated-client/models"
)

// NewTriggerPipelinePromoteParams creates a new TriggerPipelinePromoteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTriggerPipelinePromoteParams() *TriggerPipelinePromoteParams {
	return &TriggerPipelinePromoteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTriggerPipelinePromoteParamsWithTimeout creates a new TriggerPipelinePromoteParams object
// with the ability to set a timeout on a request.
func NewTriggerPipelinePromoteParamsWithTimeout(timeout time.Duration) *TriggerPipelinePromoteParams {
	return &TriggerPipelinePromoteParams{
		timeout: timeout,
	}
}

// NewTriggerPipelinePromoteParamsWithContext creates a new TriggerPipelinePromoteParams object
// with the ability to set a context for a request.
func NewTriggerPipelinePromoteParamsWithContext(ctx context.Context) *TriggerPipelinePromoteParams {
	return &TriggerPipelinePromoteParams{
		Context: ctx,
	}
}

// NewTriggerPipelinePromoteParamsWithHTTPClient creates a new TriggerPipelinePromoteParams object
// with the ability to set a custom HTTPClient for a request.
func NewTriggerPipelinePromoteParamsWithHTTPClient(client *http.Client) *TriggerPipelinePromoteParams {
	return &TriggerPipelinePromoteParams{
		HTTPClient: client,
	}
}

/*
TriggerPipelinePromoteParams contains all the parameters to send to the API endpoint

	for the trigger pipeline promote operation.

	Typically these are written to a http.Request.
*/
type TriggerPipelinePromoteParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of a comma-seperated list of test groups (Required if Impersonate-User is set)
	*/
	ImpersonateGroup *string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* PipelineParametersPromote.

	   Pipeline parameters
	*/
	PipelineParametersPromote *models.PipelineParametersPromote

	/* AppName.

	   Name of application
	*/
	AppName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the trigger pipeline promote params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriggerPipelinePromoteParams) WithDefaults() *TriggerPipelinePromoteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the trigger pipeline promote params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriggerPipelinePromoteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the trigger pipeline promote params
func (o *TriggerPipelinePromoteParams) WithTimeout(timeout time.Duration) *TriggerPipelinePromoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trigger pipeline promote params
func (o *TriggerPipelinePromoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trigger pipeline promote params
func (o *TriggerPipelinePromoteParams) WithContext(ctx context.Context) *TriggerPipelinePromoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trigger pipeline promote params
func (o *TriggerPipelinePromoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trigger pipeline promote params
func (o *TriggerPipelinePromoteParams) WithHTTPClient(client *http.Client) *TriggerPipelinePromoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trigger pipeline promote params
func (o *TriggerPipelinePromoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the trigger pipeline promote params
func (o *TriggerPipelinePromoteParams) WithImpersonateGroup(impersonateGroup *string) *TriggerPipelinePromoteParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the trigger pipeline promote params
func (o *TriggerPipelinePromoteParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the trigger pipeline promote params
func (o *TriggerPipelinePromoteParams) WithImpersonateUser(impersonateUser *string) *TriggerPipelinePromoteParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the trigger pipeline promote params
func (o *TriggerPipelinePromoteParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithPipelineParametersPromote adds the pipelineParametersPromote to the trigger pipeline promote params
func (o *TriggerPipelinePromoteParams) WithPipelineParametersPromote(pipelineParametersPromote *models.PipelineParametersPromote) *TriggerPipelinePromoteParams {
	o.SetPipelineParametersPromote(pipelineParametersPromote)
	return o
}

// SetPipelineParametersPromote adds the pipelineParametersPromote to the trigger pipeline promote params
func (o *TriggerPipelinePromoteParams) SetPipelineParametersPromote(pipelineParametersPromote *models.PipelineParametersPromote) {
	o.PipelineParametersPromote = pipelineParametersPromote
}

// WithAppName adds the appName to the trigger pipeline promote params
func (o *TriggerPipelinePromoteParams) WithAppName(appName string) *TriggerPipelinePromoteParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the trigger pipeline promote params
func (o *TriggerPipelinePromoteParams) SetAppName(appName string) {
	o.AppName = appName
}

// WriteToRequest writes these params to a swagger request
func (o *TriggerPipelinePromoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ImpersonateGroup != nil {

		// header param Impersonate-Group
		if err := r.SetHeaderParam("Impersonate-Group", *o.ImpersonateGroup); err != nil {
			return err
		}
	}

	if o.ImpersonateUser != nil {

		// header param Impersonate-User
		if err := r.SetHeaderParam("Impersonate-User", *o.ImpersonateUser); err != nil {
			return err
		}
	}
	if o.PipelineParametersPromote != nil {
		if err := r.SetBodyParam(o.PipelineParametersPromote); err != nil {
			return err
		}
	}

	// path param appName
	if err := r.SetPathParam("appName", o.AppName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
