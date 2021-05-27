// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewServiceBrokerAuthInfoTokenParams creates a new ServiceBrokerAuthInfoTokenParams object
// with the default values initialized.
func NewServiceBrokerAuthInfoTokenParams() *ServiceBrokerAuthInfoTokenParams {

	return &ServiceBrokerAuthInfoTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceBrokerAuthInfoTokenParamsWithTimeout creates a new ServiceBrokerAuthInfoTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceBrokerAuthInfoTokenParamsWithTimeout(timeout time.Duration) *ServiceBrokerAuthInfoTokenParams {

	return &ServiceBrokerAuthInfoTokenParams{

		timeout: timeout,
	}
}

// NewServiceBrokerAuthInfoTokenParamsWithContext creates a new ServiceBrokerAuthInfoTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceBrokerAuthInfoTokenParamsWithContext(ctx context.Context) *ServiceBrokerAuthInfoTokenParams {

	return &ServiceBrokerAuthInfoTokenParams{

		Context: ctx,
	}
}

// NewServiceBrokerAuthInfoTokenParamsWithHTTPClient creates a new ServiceBrokerAuthInfoTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceBrokerAuthInfoTokenParamsWithHTTPClient(client *http.Client) *ServiceBrokerAuthInfoTokenParams {

	return &ServiceBrokerAuthInfoTokenParams{
		HTTPClient: client,
	}
}

/*ServiceBrokerAuthInfoTokenParams contains all the parameters to send to the API endpoint
for the service broker auth info token operation typically these are written to a http.Request
*/
type ServiceBrokerAuthInfoTokenParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service broker auth info token params
func (o *ServiceBrokerAuthInfoTokenParams) WithTimeout(timeout time.Duration) *ServiceBrokerAuthInfoTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service broker auth info token params
func (o *ServiceBrokerAuthInfoTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service broker auth info token params
func (o *ServiceBrokerAuthInfoTokenParams) WithContext(ctx context.Context) *ServiceBrokerAuthInfoTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service broker auth info token params
func (o *ServiceBrokerAuthInfoTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service broker auth info token params
func (o *ServiceBrokerAuthInfoTokenParams) WithHTTPClient(client *http.Client) *ServiceBrokerAuthInfoTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service broker auth info token params
func (o *ServiceBrokerAuthInfoTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceBrokerAuthInfoTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
