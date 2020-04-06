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

// GetRetentionsIDExecutionsReader is a Reader for the GetRetentionsIDExecutions structure.
type GetRetentionsIDExecutionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRetentionsIDExecutionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRetentionsIDExecutionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRetentionsIDExecutionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRetentionsIDExecutionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRetentionsIDExecutionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRetentionsIDExecutionsOK creates a GetRetentionsIDExecutionsOK with default headers values
func NewGetRetentionsIDExecutionsOK() *GetRetentionsIDExecutionsOK {
	return &GetRetentionsIDExecutionsOK{}
}

/*GetRetentionsIDExecutionsOK handles this case with default header values.

Get a Retention job successfully.
*/
type GetRetentionsIDExecutionsOK struct {
	Payload []*models.RetentionExecution
}

func (o *GetRetentionsIDExecutionsOK) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}/executions][%d] getRetentionsIdExecutionsOK  %+v", 200, o.Payload)
}

func (o *GetRetentionsIDExecutionsOK) GetPayload() []*models.RetentionExecution {
	return o.Payload
}

func (o *GetRetentionsIDExecutionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRetentionsIDExecutionsUnauthorized creates a GetRetentionsIDExecutionsUnauthorized with default headers values
func NewGetRetentionsIDExecutionsUnauthorized() *GetRetentionsIDExecutionsUnauthorized {
	return &GetRetentionsIDExecutionsUnauthorized{}
}

/*GetRetentionsIDExecutionsUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetRetentionsIDExecutionsUnauthorized struct {
}

func (o *GetRetentionsIDExecutionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}/executions][%d] getRetentionsIdExecutionsUnauthorized ", 401)
}

func (o *GetRetentionsIDExecutionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRetentionsIDExecutionsForbidden creates a GetRetentionsIDExecutionsForbidden with default headers values
func NewGetRetentionsIDExecutionsForbidden() *GetRetentionsIDExecutionsForbidden {
	return &GetRetentionsIDExecutionsForbidden{}
}

/*GetRetentionsIDExecutionsForbidden handles this case with default header values.

User have no permission.
*/
type GetRetentionsIDExecutionsForbidden struct {
}

func (o *GetRetentionsIDExecutionsForbidden) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}/executions][%d] getRetentionsIdExecutionsForbidden ", 403)
}

func (o *GetRetentionsIDExecutionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRetentionsIDExecutionsInternalServerError creates a GetRetentionsIDExecutionsInternalServerError with default headers values
func NewGetRetentionsIDExecutionsInternalServerError() *GetRetentionsIDExecutionsInternalServerError {
	return &GetRetentionsIDExecutionsInternalServerError{}
}

/*GetRetentionsIDExecutionsInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetRetentionsIDExecutionsInternalServerError struct {
}

func (o *GetRetentionsIDExecutionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}/executions][%d] getRetentionsIdExecutionsInternalServerError ", 500)
}

func (o *GetRetentionsIDExecutionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
