// Code generated by go-swagger; DO NOT EDIT.

package job

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
	"github.com/go-openapi/swag"
)

// NewGetBatchesParams creates a new GetBatchesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBatchesParams() *GetBatchesParams {
	return &GetBatchesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBatchesParamsWithTimeout creates a new GetBatchesParams object
// with the ability to set a timeout on a request.
func NewGetBatchesParamsWithTimeout(timeout time.Duration) *GetBatchesParams {
	return &GetBatchesParams{
		timeout: timeout,
	}
}

// NewGetBatchesParamsWithContext creates a new GetBatchesParams object
// with the ability to set a context for a request.
func NewGetBatchesParamsWithContext(ctx context.Context) *GetBatchesParams {
	return &GetBatchesParams{
		Context: ctx,
	}
}

// NewGetBatchesParamsWithHTTPClient creates a new GetBatchesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBatchesParamsWithHTTPClient(client *http.Client) *GetBatchesParams {
	return &GetBatchesParams{
		HTTPClient: client,
	}
}

/*
GetBatchesParams contains all the parameters to send to the API endpoint

	for the get batches operation.

	Typically these are written to a http.Request.
*/
type GetBatchesParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)
	*/
	ImpersonateGroup []string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   Name of application
	*/
	AppName string

	/* EnvName.

	   Name of environment
	*/
	EnvName string

	/* JobComponentName.

	   Name of job-component
	*/
	JobComponentName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get batches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBatchesParams) WithDefaults() *GetBatchesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get batches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBatchesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get batches params
func (o *GetBatchesParams) WithTimeout(timeout time.Duration) *GetBatchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get batches params
func (o *GetBatchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get batches params
func (o *GetBatchesParams) WithContext(ctx context.Context) *GetBatchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get batches params
func (o *GetBatchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get batches params
func (o *GetBatchesParams) WithHTTPClient(client *http.Client) *GetBatchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get batches params
func (o *GetBatchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the get batches params
func (o *GetBatchesParams) WithImpersonateGroup(impersonateGroup []string) *GetBatchesParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the get batches params
func (o *GetBatchesParams) SetImpersonateGroup(impersonateGroup []string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the get batches params
func (o *GetBatchesParams) WithImpersonateUser(impersonateUser *string) *GetBatchesParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the get batches params
func (o *GetBatchesParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the get batches params
func (o *GetBatchesParams) WithAppName(appName string) *GetBatchesParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get batches params
func (o *GetBatchesParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithEnvName adds the envName to the get batches params
func (o *GetBatchesParams) WithEnvName(envName string) *GetBatchesParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the get batches params
func (o *GetBatchesParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WithJobComponentName adds the jobComponentName to the get batches params
func (o *GetBatchesParams) WithJobComponentName(jobComponentName string) *GetBatchesParams {
	o.SetJobComponentName(jobComponentName)
	return o
}

// SetJobComponentName adds the jobComponentName to the get batches params
func (o *GetBatchesParams) SetJobComponentName(jobComponentName string) {
	o.JobComponentName = jobComponentName
}

// WriteToRequest writes these params to a swagger request
func (o *GetBatchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ImpersonateGroup != nil {

		// binding items for Impersonate-Group
		joinedImpersonateGroup := o.bindParamImpersonateGroup(reg)

		// header array param Impersonate-Group
		if len(joinedImpersonateGroup) > 0 {
			if err := r.SetHeaderParam("Impersonate-Group", joinedImpersonateGroup[0]); err != nil {
				return err
			}
		}
	}

	if o.ImpersonateUser != nil {

		// header param Impersonate-User
		if err := r.SetHeaderParam("Impersonate-User", *o.ImpersonateUser); err != nil {
			return err
		}
	}

	// path param appName
	if err := r.SetPathParam("appName", o.AppName); err != nil {
		return err
	}

	// path param envName
	if err := r.SetPathParam("envName", o.EnvName); err != nil {
		return err
	}

	// path param jobComponentName
	if err := r.SetPathParam("jobComponentName", o.JobComponentName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetBatches binds the parameter Impersonate-Group
func (o *GetBatchesParams) bindParamImpersonateGroup(formats strfmt.Registry) []string {
	impersonateGroupIR := o.ImpersonateGroup

	var impersonateGroupIC []string
	for _, impersonateGroupIIR := range impersonateGroupIR { // explode []string

		impersonateGroupIIV := impersonateGroupIIR // string as string
		impersonateGroupIC = append(impersonateGroupIC, impersonateGroupIIV)
	}

	// items.CollectionFormat: ""
	impersonateGroupIS := swag.JoinByFormat(impersonateGroupIC, "")

	return impersonateGroupIS
}