// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StatisticMap statistic map
// swagger:model StatisticMap
type StatisticMap struct {

	// The count of the private projects which the user is a member of.
	PrivateProjectCount int32 `json:"private_project_count,omitempty"`

	// The count of the private repositories belonging to the projects which the user is a member of.
	PrivateRepoCount int32 `json:"private_repo_count,omitempty"`

	// The count of the public projects.
	PublicProjectCount int32 `json:"public_project_count,omitempty"`

	// The count of the public repositories belonging to the public projects which the user is a member of.
	PublicRepoCount int32 `json:"public_repo_count,omitempty"`

	// The count of the total projects, only be seen when the is admin.
	TotalProjectCount int32 `json:"total_project_count,omitempty"`

	// The count of the total repositories, only be seen when the user is admin.
	TotalRepoCount int32 `json:"total_repo_count,omitempty"`
}

// Validate validates this statistic map
func (m *StatisticMap) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StatisticMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatisticMap) UnmarshalBinary(b []byte) error {
	var res StatisticMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
