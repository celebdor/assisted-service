// Code generated by go-swagger; DO NOT EDIT.

package assisted_service_iso

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

// NewGetPresignedForAssistedServiceISOParams creates a new GetPresignedForAssistedServiceISOParams object
// with the default values initialized.
func NewGetPresignedForAssistedServiceISOParams() *GetPresignedForAssistedServiceISOParams {

	return &GetPresignedForAssistedServiceISOParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPresignedForAssistedServiceISOParamsWithTimeout creates a new GetPresignedForAssistedServiceISOParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPresignedForAssistedServiceISOParamsWithTimeout(timeout time.Duration) *GetPresignedForAssistedServiceISOParams {

	return &GetPresignedForAssistedServiceISOParams{

		timeout: timeout,
	}
}

// NewGetPresignedForAssistedServiceISOParamsWithContext creates a new GetPresignedForAssistedServiceISOParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPresignedForAssistedServiceISOParamsWithContext(ctx context.Context) *GetPresignedForAssistedServiceISOParams {

	return &GetPresignedForAssistedServiceISOParams{

		Context: ctx,
	}
}

// NewGetPresignedForAssistedServiceISOParamsWithHTTPClient creates a new GetPresignedForAssistedServiceISOParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPresignedForAssistedServiceISOParamsWithHTTPClient(client *http.Client) *GetPresignedForAssistedServiceISOParams {

	return &GetPresignedForAssistedServiceISOParams{
		HTTPClient: client,
	}
}

/*GetPresignedForAssistedServiceISOParams contains all the parameters to send to the API endpoint
for the get presigned for assisted service i s o operation typically these are written to a http.Request
*/
type GetPresignedForAssistedServiceISOParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get presigned for assisted service i s o params
func (o *GetPresignedForAssistedServiceISOParams) WithTimeout(timeout time.Duration) *GetPresignedForAssistedServiceISOParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get presigned for assisted service i s o params
func (o *GetPresignedForAssistedServiceISOParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get presigned for assisted service i s o params
func (o *GetPresignedForAssistedServiceISOParams) WithContext(ctx context.Context) *GetPresignedForAssistedServiceISOParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get presigned for assisted service i s o params
func (o *GetPresignedForAssistedServiceISOParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get presigned for assisted service i s o params
func (o *GetPresignedForAssistedServiceISOParams) WithHTTPClient(client *http.Client) *GetPresignedForAssistedServiceISOParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get presigned for assisted service i s o params
func (o *GetPresignedForAssistedServiceISOParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPresignedForAssistedServiceISOParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
