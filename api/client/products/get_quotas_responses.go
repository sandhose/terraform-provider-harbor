// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sandhose/terraform-provider-harbor/api/models"
)

// GetQuotasReader is a Reader for the GetQuotas structure.
type GetQuotasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQuotasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQuotasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetQuotasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQuotasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQuotasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetQuotasOK creates a GetQuotasOK with default headers values
func NewGetQuotasOK() *GetQuotasOK {
	return &GetQuotasOK{}
}

/*GetQuotasOK handles this case with default header values.

Successfully retrieved the quotas.
*/
type GetQuotasOK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of access logs
	 */
	XTotalCount int64

	Payload []*models.Quota
}

func (o *GetQuotasOK) Error() string {
	return fmt.Sprintf("[GET /quotas][%d] getQuotasOK  %+v", 200, o.Payload)
}

func (o *GetQuotasOK) GetPayload() []*models.Quota {
	return o.Payload
}

func (o *GetQuotasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt64(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int64", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQuotasUnauthorized creates a GetQuotasUnauthorized with default headers values
func NewGetQuotasUnauthorized() *GetQuotasUnauthorized {
	return &GetQuotasUnauthorized{}
}

/*GetQuotasUnauthorized handles this case with default header values.

User is not authenticated.
*/
type GetQuotasUnauthorized struct {
}

func (o *GetQuotasUnauthorized) Error() string {
	return fmt.Sprintf("[GET /quotas][%d] getQuotasUnauthorized ", 401)
}

func (o *GetQuotasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetQuotasForbidden creates a GetQuotasForbidden with default headers values
func NewGetQuotasForbidden() *GetQuotasForbidden {
	return &GetQuotasForbidden{}
}

/*GetQuotasForbidden handles this case with default header values.

User does not have permission to call this API.
*/
type GetQuotasForbidden struct {
}

func (o *GetQuotasForbidden) Error() string {
	return fmt.Sprintf("[GET /quotas][%d] getQuotasForbidden ", 403)
}

func (o *GetQuotasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetQuotasInternalServerError creates a GetQuotasInternalServerError with default headers values
func NewGetQuotasInternalServerError() *GetQuotasInternalServerError {
	return &GetQuotasInternalServerError{}
}

/*GetQuotasInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetQuotasInternalServerError struct {
}

func (o *GetQuotasInternalServerError) Error() string {
	return fmt.Sprintf("[GET /quotas][%d] getQuotasInternalServerError ", 500)
}

func (o *GetQuotasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
