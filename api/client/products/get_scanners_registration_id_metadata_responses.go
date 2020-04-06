// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sandhose/terraform-provider-harbor/api/models"
)

// GetScannersRegistrationIDMetadataReader is a Reader for the GetScannersRegistrationIDMetadata structure.
type GetScannersRegistrationIDMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScannersRegistrationIDMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScannersRegistrationIDMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetScannersRegistrationIDMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScannersRegistrationIDMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScannersRegistrationIDMetadataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetScannersRegistrationIDMetadataOK creates a GetScannersRegistrationIDMetadataOK with default headers values
func NewGetScannersRegistrationIDMetadataOK() *GetScannersRegistrationIDMetadataOK {
	return &GetScannersRegistrationIDMetadataOK{}
}

/*GetScannersRegistrationIDMetadataOK handles this case with default header values.

The metadata of the specified scanner adapter
*/
type GetScannersRegistrationIDMetadataOK struct {
	Payload *models.ScannerAdapterMetadata
}

func (o *GetScannersRegistrationIDMetadataOK) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}/metadata][%d] getScannersRegistrationIdMetadataOK  %+v", 200, o.Payload)
}

func (o *GetScannersRegistrationIDMetadataOK) GetPayload() *models.ScannerAdapterMetadata {
	return o.Payload
}

func (o *GetScannersRegistrationIDMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScannerAdapterMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScannersRegistrationIDMetadataUnauthorized creates a GetScannersRegistrationIDMetadataUnauthorized with default headers values
func NewGetScannersRegistrationIDMetadataUnauthorized() *GetScannersRegistrationIDMetadataUnauthorized {
	return &GetScannersRegistrationIDMetadataUnauthorized{}
}

/*GetScannersRegistrationIDMetadataUnauthorized handles this case with default header values.

Unauthorized request
*/
type GetScannersRegistrationIDMetadataUnauthorized struct {
}

func (o *GetScannersRegistrationIDMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}/metadata][%d] getScannersRegistrationIdMetadataUnauthorized ", 401)
}

func (o *GetScannersRegistrationIDMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScannersRegistrationIDMetadataForbidden creates a GetScannersRegistrationIDMetadataForbidden with default headers values
func NewGetScannersRegistrationIDMetadataForbidden() *GetScannersRegistrationIDMetadataForbidden {
	return &GetScannersRegistrationIDMetadataForbidden{}
}

/*GetScannersRegistrationIDMetadataForbidden handles this case with default header values.

Request is not allowed
*/
type GetScannersRegistrationIDMetadataForbidden struct {
}

func (o *GetScannersRegistrationIDMetadataForbidden) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}/metadata][%d] getScannersRegistrationIdMetadataForbidden ", 403)
}

func (o *GetScannersRegistrationIDMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScannersRegistrationIDMetadataInternalServerError creates a GetScannersRegistrationIDMetadataInternalServerError with default headers values
func NewGetScannersRegistrationIDMetadataInternalServerError() *GetScannersRegistrationIDMetadataInternalServerError {
	return &GetScannersRegistrationIDMetadataInternalServerError{}
}

/*GetScannersRegistrationIDMetadataInternalServerError handles this case with default header values.

Internal server error happened
*/
type GetScannersRegistrationIDMetadataInternalServerError struct {
}

func (o *GetScannersRegistrationIDMetadataInternalServerError) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}/metadata][%d] getScannersRegistrationIdMetadataInternalServerError ", 500)
}

func (o *GetScannersRegistrationIDMetadataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
