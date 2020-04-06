// Code generated by go-swagger; DO NOT EDIT.

package products

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

	"github.com/sandhose/terraform-provider-harbor/api/models"
)

// NewPutInternalSwitchquotaParams creates a new PutInternalSwitchquotaParams object
// with the default values initialized.
func NewPutInternalSwitchquotaParams() *PutInternalSwitchquotaParams {
	var ()
	return &PutInternalSwitchquotaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutInternalSwitchquotaParamsWithTimeout creates a new PutInternalSwitchquotaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutInternalSwitchquotaParamsWithTimeout(timeout time.Duration) *PutInternalSwitchquotaParams {
	var ()
	return &PutInternalSwitchquotaParams{

		timeout: timeout,
	}
}

// NewPutInternalSwitchquotaParamsWithContext creates a new PutInternalSwitchquotaParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutInternalSwitchquotaParamsWithContext(ctx context.Context) *PutInternalSwitchquotaParams {
	var ()
	return &PutInternalSwitchquotaParams{

		Context: ctx,
	}
}

// NewPutInternalSwitchquotaParamsWithHTTPClient creates a new PutInternalSwitchquotaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutInternalSwitchquotaParamsWithHTTPClient(client *http.Client) *PutInternalSwitchquotaParams {
	var ()
	return &PutInternalSwitchquotaParams{
		HTTPClient: client,
	}
}

/*PutInternalSwitchquotaParams contains all the parameters to send to the API endpoint
for the put internal switchquota operation typically these are written to a http.Request
*/
type PutInternalSwitchquotaParams struct {

	/*Switcher*/
	Switcher *models.QuotaSwitcher

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put internal switchquota params
func (o *PutInternalSwitchquotaParams) WithTimeout(timeout time.Duration) *PutInternalSwitchquotaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put internal switchquota params
func (o *PutInternalSwitchquotaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put internal switchquota params
func (o *PutInternalSwitchquotaParams) WithContext(ctx context.Context) *PutInternalSwitchquotaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put internal switchquota params
func (o *PutInternalSwitchquotaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put internal switchquota params
func (o *PutInternalSwitchquotaParams) WithHTTPClient(client *http.Client) *PutInternalSwitchquotaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put internal switchquota params
func (o *PutInternalSwitchquotaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSwitcher adds the switcher to the put internal switchquota params
func (o *PutInternalSwitchquotaParams) WithSwitcher(switcher *models.QuotaSwitcher) *PutInternalSwitchquotaParams {
	o.SetSwitcher(switcher)
	return o
}

// SetSwitcher adds the switcher to the put internal switchquota params
func (o *PutInternalSwitchquotaParams) SetSwitcher(switcher *models.QuotaSwitcher) {
	o.Switcher = switcher
}

// WriteToRequest writes these params to a swagger request
func (o *PutInternalSwitchquotaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Switcher != nil {
		if err := r.SetBodyParam(o.Switcher); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
