// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutSystemCVEWhitelistReader is a Reader for the PutSystemCVEWhitelist structure.
type PutSystemCVEWhitelistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSystemCVEWhitelistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutSystemCVEWhitelistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutSystemCVEWhitelistUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutSystemCVEWhitelistForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutSystemCVEWhitelistInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutSystemCVEWhitelistOK creates a PutSystemCVEWhitelistOK with default headers values
func NewPutSystemCVEWhitelistOK() *PutSystemCVEWhitelistOK {
	return &PutSystemCVEWhitelistOK{}
}

/*PutSystemCVEWhitelistOK handles this case with default header values.

Successfully updated the CVE whitelist.
*/
type PutSystemCVEWhitelistOK struct {
}

func (o *PutSystemCVEWhitelistOK) Error() string {
	return fmt.Sprintf("[PUT /system/CVEWhitelist][%d] putSystemCVEWhitelistOK ", 200)
}

func (o *PutSystemCVEWhitelistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSystemCVEWhitelistUnauthorized creates a PutSystemCVEWhitelistUnauthorized with default headers values
func NewPutSystemCVEWhitelistUnauthorized() *PutSystemCVEWhitelistUnauthorized {
	return &PutSystemCVEWhitelistUnauthorized{}
}

/*PutSystemCVEWhitelistUnauthorized handles this case with default header values.

User is not authenticated.
*/
type PutSystemCVEWhitelistUnauthorized struct {
}

func (o *PutSystemCVEWhitelistUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /system/CVEWhitelist][%d] putSystemCVEWhitelistUnauthorized ", 401)
}

func (o *PutSystemCVEWhitelistUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSystemCVEWhitelistForbidden creates a PutSystemCVEWhitelistForbidden with default headers values
func NewPutSystemCVEWhitelistForbidden() *PutSystemCVEWhitelistForbidden {
	return &PutSystemCVEWhitelistForbidden{}
}

/*PutSystemCVEWhitelistForbidden handles this case with default header values.

User does not have permission to call this API.
*/
type PutSystemCVEWhitelistForbidden struct {
}

func (o *PutSystemCVEWhitelistForbidden) Error() string {
	return fmt.Sprintf("[PUT /system/CVEWhitelist][%d] putSystemCVEWhitelistForbidden ", 403)
}

func (o *PutSystemCVEWhitelistForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSystemCVEWhitelistInternalServerError creates a PutSystemCVEWhitelistInternalServerError with default headers values
func NewPutSystemCVEWhitelistInternalServerError() *PutSystemCVEWhitelistInternalServerError {
	return &PutSystemCVEWhitelistInternalServerError{}
}

/*PutSystemCVEWhitelistInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type PutSystemCVEWhitelistInternalServerError struct {
}

func (o *PutSystemCVEWhitelistInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /system/CVEWhitelist][%d] putSystemCVEWhitelistInternalServerError ", 500)
}

func (o *PutSystemCVEWhitelistInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
