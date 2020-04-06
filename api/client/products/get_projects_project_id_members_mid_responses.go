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

// GetProjectsProjectIDMembersMidReader is a Reader for the GetProjectsProjectIDMembersMid structure.
type GetProjectsProjectIDMembersMidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDMembersMidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsProjectIDMembersMidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectsProjectIDMembersMidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectsProjectIDMembersMidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectsProjectIDMembersMidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectsProjectIDMembersMidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectsProjectIDMembersMidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProjectsProjectIDMembersMidOK creates a GetProjectsProjectIDMembersMidOK with default headers values
func NewGetProjectsProjectIDMembersMidOK() *GetProjectsProjectIDMembersMidOK {
	return &GetProjectsProjectIDMembersMidOK{}
}

/*GetProjectsProjectIDMembersMidOK handles this case with default header values.

Project member retrieved successfully.
*/
type GetProjectsProjectIDMembersMidOK struct {
	Payload *models.ProjectMemberEntity
}

func (o *GetProjectsProjectIDMembersMidOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/members/{mid}][%d] getProjectsProjectIdMembersMidOK  %+v", 200, o.Payload)
}

func (o *GetProjectsProjectIDMembersMidOK) GetPayload() *models.ProjectMemberEntity {
	return o.Payload
}

func (o *GetProjectsProjectIDMembersMidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectMemberEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDMembersMidBadRequest creates a GetProjectsProjectIDMembersMidBadRequest with default headers values
func NewGetProjectsProjectIDMembersMidBadRequest() *GetProjectsProjectIDMembersMidBadRequest {
	return &GetProjectsProjectIDMembersMidBadRequest{}
}

/*GetProjectsProjectIDMembersMidBadRequest handles this case with default header values.

Illegal format of project member or invalid project id, member id.
*/
type GetProjectsProjectIDMembersMidBadRequest struct {
}

func (o *GetProjectsProjectIDMembersMidBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/members/{mid}][%d] getProjectsProjectIdMembersMidBadRequest ", 400)
}

func (o *GetProjectsProjectIDMembersMidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDMembersMidUnauthorized creates a GetProjectsProjectIDMembersMidUnauthorized with default headers values
func NewGetProjectsProjectIDMembersMidUnauthorized() *GetProjectsProjectIDMembersMidUnauthorized {
	return &GetProjectsProjectIDMembersMidUnauthorized{}
}

/*GetProjectsProjectIDMembersMidUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetProjectsProjectIDMembersMidUnauthorized struct {
}

func (o *GetProjectsProjectIDMembersMidUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/members/{mid}][%d] getProjectsProjectIdMembersMidUnauthorized ", 401)
}

func (o *GetProjectsProjectIDMembersMidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDMembersMidForbidden creates a GetProjectsProjectIDMembersMidForbidden with default headers values
func NewGetProjectsProjectIDMembersMidForbidden() *GetProjectsProjectIDMembersMidForbidden {
	return &GetProjectsProjectIDMembersMidForbidden{}
}

/*GetProjectsProjectIDMembersMidForbidden handles this case with default header values.

User in session does not have permission to the project.
*/
type GetProjectsProjectIDMembersMidForbidden struct {
}

func (o *GetProjectsProjectIDMembersMidForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/members/{mid}][%d] getProjectsProjectIdMembersMidForbidden ", 403)
}

func (o *GetProjectsProjectIDMembersMidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDMembersMidNotFound creates a GetProjectsProjectIDMembersMidNotFound with default headers values
func NewGetProjectsProjectIDMembersMidNotFound() *GetProjectsProjectIDMembersMidNotFound {
	return &GetProjectsProjectIDMembersMidNotFound{}
}

/*GetProjectsProjectIDMembersMidNotFound handles this case with default header values.

Project or projet member does not exist.
*/
type GetProjectsProjectIDMembersMidNotFound struct {
}

func (o *GetProjectsProjectIDMembersMidNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/members/{mid}][%d] getProjectsProjectIdMembersMidNotFound ", 404)
}

func (o *GetProjectsProjectIDMembersMidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDMembersMidInternalServerError creates a GetProjectsProjectIDMembersMidInternalServerError with default headers values
func NewGetProjectsProjectIDMembersMidInternalServerError() *GetProjectsProjectIDMembersMidInternalServerError {
	return &GetProjectsProjectIDMembersMidInternalServerError{}
}

/*GetProjectsProjectIDMembersMidInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetProjectsProjectIDMembersMidInternalServerError struct {
}

func (o *GetProjectsProjectIDMembersMidInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/members/{mid}][%d] getProjectsProjectIdMembersMidInternalServerError ", 500)
}

func (o *GetProjectsProjectIDMembersMidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
