// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutLabelsIDReader is a Reader for the PutLabelsID structure.
type PutLabelsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLabelsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutLabelsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutLabelsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutLabelsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutLabelsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutLabelsIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutLabelsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutLabelsIDOK creates a PutLabelsIDOK with default headers values
func NewPutLabelsIDOK() *PutLabelsIDOK {
	return &PutLabelsIDOK{}
}

/*PutLabelsIDOK handles this case with default header values.

Update successfully.
*/
type PutLabelsIDOK struct {
}

func (o *PutLabelsIDOK) Error() string {
	return fmt.Sprintf("[PUT /labels/{id}][%d] putLabelsIdOK ", 200)
}

func (o *PutLabelsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLabelsIDBadRequest creates a PutLabelsIDBadRequest with default headers values
func NewPutLabelsIDBadRequest() *PutLabelsIDBadRequest {
	return &PutLabelsIDBadRequest{}
}

/*PutLabelsIDBadRequest handles this case with default header values.

Invalid parameters.
*/
type PutLabelsIDBadRequest struct {
}

func (o *PutLabelsIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /labels/{id}][%d] putLabelsIdBadRequest ", 400)
}

func (o *PutLabelsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLabelsIDUnauthorized creates a PutLabelsIDUnauthorized with default headers values
func NewPutLabelsIDUnauthorized() *PutLabelsIDUnauthorized {
	return &PutLabelsIDUnauthorized{}
}

/*PutLabelsIDUnauthorized handles this case with default header values.

User need to log in first.
*/
type PutLabelsIDUnauthorized struct {
}

func (o *PutLabelsIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /labels/{id}][%d] putLabelsIdUnauthorized ", 401)
}

func (o *PutLabelsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLabelsIDNotFound creates a PutLabelsIDNotFound with default headers values
func NewPutLabelsIDNotFound() *PutLabelsIDNotFound {
	return &PutLabelsIDNotFound{}
}

/*PutLabelsIDNotFound handles this case with default header values.

The resource does not exist.
*/
type PutLabelsIDNotFound struct {
}

func (o *PutLabelsIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /labels/{id}][%d] putLabelsIdNotFound ", 404)
}

func (o *PutLabelsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLabelsIDConflict creates a PutLabelsIDConflict with default headers values
func NewPutLabelsIDConflict() *PutLabelsIDConflict {
	return &PutLabelsIDConflict{}
}

/*PutLabelsIDConflict handles this case with default header values.

The label with the same name already exists.
*/
type PutLabelsIDConflict struct {
}

func (o *PutLabelsIDConflict) Error() string {
	return fmt.Sprintf("[PUT /labels/{id}][%d] putLabelsIdConflict ", 409)
}

func (o *PutLabelsIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLabelsIDInternalServerError creates a PutLabelsIDInternalServerError with default headers values
func NewPutLabelsIDInternalServerError() *PutLabelsIDInternalServerError {
	return &PutLabelsIDInternalServerError{}
}

/*PutLabelsIDInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type PutLabelsIDInternalServerError struct {
}

func (o *PutLabelsIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /labels/{id}][%d] putLabelsIdInternalServerError ", 500)
}

func (o *PutLabelsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
