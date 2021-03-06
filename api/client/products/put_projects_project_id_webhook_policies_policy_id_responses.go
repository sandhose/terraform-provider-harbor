// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutProjectsProjectIDWebhookPoliciesPolicyIDReader is a Reader for the PutProjectsProjectIDWebhookPoliciesPolicyID structure.
type PutProjectsProjectIDWebhookPoliciesPolicyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutProjectsProjectIDWebhookPoliciesPolicyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutProjectsProjectIDWebhookPoliciesPolicyIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutProjectsProjectIDWebhookPoliciesPolicyIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutProjectsProjectIDWebhookPoliciesPolicyIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutProjectsProjectIDWebhookPoliciesPolicyIDOK creates a PutProjectsProjectIDWebhookPoliciesPolicyIDOK with default headers values
func NewPutProjectsProjectIDWebhookPoliciesPolicyIDOK() *PutProjectsProjectIDWebhookPoliciesPolicyIDOK {
	return &PutProjectsProjectIDWebhookPoliciesPolicyIDOK{}
}

/*PutProjectsProjectIDWebhookPoliciesPolicyIDOK handles this case with default header values.

Update webhook policy successfully.
*/
type PutProjectsProjectIDWebhookPoliciesPolicyIDOK struct {
}

func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/webhook/policies/{policy_id}][%d] putProjectsProjectIdWebhookPoliciesPolicyIdOK ", 200)
}

func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDWebhookPoliciesPolicyIDBadRequest creates a PutProjectsProjectIDWebhookPoliciesPolicyIDBadRequest with default headers values
func NewPutProjectsProjectIDWebhookPoliciesPolicyIDBadRequest() *PutProjectsProjectIDWebhookPoliciesPolicyIDBadRequest {
	return &PutProjectsProjectIDWebhookPoliciesPolicyIDBadRequest{}
}

/*PutProjectsProjectIDWebhookPoliciesPolicyIDBadRequest handles this case with default header values.

Illegal format of provided ID value.
*/
type PutProjectsProjectIDWebhookPoliciesPolicyIDBadRequest struct {
}

func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/webhook/policies/{policy_id}][%d] putProjectsProjectIdWebhookPoliciesPolicyIdBadRequest ", 400)
}

func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized creates a PutProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized with default headers values
func NewPutProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized() *PutProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized {
	return &PutProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized{}
}

/*PutProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized handles this case with default header values.

User need to log in first.
*/
type PutProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized struct {
}

func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/webhook/policies/{policy_id}][%d] putProjectsProjectIdWebhookPoliciesPolicyIdUnauthorized ", 401)
}

func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDWebhookPoliciesPolicyIDForbidden creates a PutProjectsProjectIDWebhookPoliciesPolicyIDForbidden with default headers values
func NewPutProjectsProjectIDWebhookPoliciesPolicyIDForbidden() *PutProjectsProjectIDWebhookPoliciesPolicyIDForbidden {
	return &PutProjectsProjectIDWebhookPoliciesPolicyIDForbidden{}
}

/*PutProjectsProjectIDWebhookPoliciesPolicyIDForbidden handles this case with default header values.

User have no permission to update webhook policy of the project.
*/
type PutProjectsProjectIDWebhookPoliciesPolicyIDForbidden struct {
}

func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/webhook/policies/{policy_id}][%d] putProjectsProjectIdWebhookPoliciesPolicyIdForbidden ", 403)
}

func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDWebhookPoliciesPolicyIDNotFound creates a PutProjectsProjectIDWebhookPoliciesPolicyIDNotFound with default headers values
func NewPutProjectsProjectIDWebhookPoliciesPolicyIDNotFound() *PutProjectsProjectIDWebhookPoliciesPolicyIDNotFound {
	return &PutProjectsProjectIDWebhookPoliciesPolicyIDNotFound{}
}

/*PutProjectsProjectIDWebhookPoliciesPolicyIDNotFound handles this case with default header values.

Webhook policy ID does not exist.
*/
type PutProjectsProjectIDWebhookPoliciesPolicyIDNotFound struct {
}

func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/webhook/policies/{policy_id}][%d] putProjectsProjectIdWebhookPoliciesPolicyIdNotFound ", 404)
}

func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError creates a PutProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError with default headers values
func NewPutProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError() *PutProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError {
	return &PutProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError{}
}

/*PutProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError handles this case with default header values.

Internal server errors.
*/
type PutProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError struct {
}

func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/webhook/policies/{policy_id}][%d] putProjectsProjectIdWebhookPoliciesPolicyIdInternalServerError ", 500)
}

func (o *PutProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
