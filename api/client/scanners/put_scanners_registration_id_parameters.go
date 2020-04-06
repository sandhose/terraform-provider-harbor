// Code generated by go-swagger; DO NOT EDIT.

package scanners

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

// NewPutScannersRegistrationIDParams creates a new PutScannersRegistrationIDParams object
// with the default values initialized.
func NewPutScannersRegistrationIDParams() *PutScannersRegistrationIDParams {
	var ()
	return &PutScannersRegistrationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutScannersRegistrationIDParamsWithTimeout creates a new PutScannersRegistrationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutScannersRegistrationIDParamsWithTimeout(timeout time.Duration) *PutScannersRegistrationIDParams {
	var ()
	return &PutScannersRegistrationIDParams{

		timeout: timeout,
	}
}

// NewPutScannersRegistrationIDParamsWithContext creates a new PutScannersRegistrationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutScannersRegistrationIDParamsWithContext(ctx context.Context) *PutScannersRegistrationIDParams {
	var ()
	return &PutScannersRegistrationIDParams{

		Context: ctx,
	}
}

// NewPutScannersRegistrationIDParamsWithHTTPClient creates a new PutScannersRegistrationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutScannersRegistrationIDParamsWithHTTPClient(client *http.Client) *PutScannersRegistrationIDParams {
	var ()
	return &PutScannersRegistrationIDParams{
		HTTPClient: client,
	}
}

/*PutScannersRegistrationIDParams contains all the parameters to send to the API endpoint
for the put scanners registration ID operation typically these are written to a http.Request
*/
type PutScannersRegistrationIDParams struct {

	/*Registration
	  A scanner registraiton to be updated.

	*/
	Registration *models.ScannerRegistrationReq
	/*RegistrationID
	  The scanner registration identifier.

	*/
	RegistrationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put scanners registration ID params
func (o *PutScannersRegistrationIDParams) WithTimeout(timeout time.Duration) *PutScannersRegistrationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put scanners registration ID params
func (o *PutScannersRegistrationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put scanners registration ID params
func (o *PutScannersRegistrationIDParams) WithContext(ctx context.Context) *PutScannersRegistrationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put scanners registration ID params
func (o *PutScannersRegistrationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put scanners registration ID params
func (o *PutScannersRegistrationIDParams) WithHTTPClient(client *http.Client) *PutScannersRegistrationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put scanners registration ID params
func (o *PutScannersRegistrationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegistration adds the registration to the put scanners registration ID params
func (o *PutScannersRegistrationIDParams) WithRegistration(registration *models.ScannerRegistrationReq) *PutScannersRegistrationIDParams {
	o.SetRegistration(registration)
	return o
}

// SetRegistration adds the registration to the put scanners registration ID params
func (o *PutScannersRegistrationIDParams) SetRegistration(registration *models.ScannerRegistrationReq) {
	o.Registration = registration
}

// WithRegistrationID adds the registrationID to the put scanners registration ID params
func (o *PutScannersRegistrationIDParams) WithRegistrationID(registrationID string) *PutScannersRegistrationIDParams {
	o.SetRegistrationID(registrationID)
	return o
}

// SetRegistrationID adds the registrationId to the put scanners registration ID params
func (o *PutScannersRegistrationIDParams) SetRegistrationID(registrationID string) {
	o.RegistrationID = registrationID
}

// WriteToRequest writes these params to a swagger request
func (o *PutScannersRegistrationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Registration != nil {
		if err := r.SetBodyParam(o.Registration); err != nil {
			return err
		}
	}

	// path param registration_id
	if err := r.SetPathParam("registration_id", o.RegistrationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
