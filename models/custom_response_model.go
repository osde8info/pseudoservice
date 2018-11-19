// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustomResponseModel custom response model
// swagger:model CustomResponseModel
type CustomResponseModel struct {

	// results
	Results []string `json:"results"`

	// Number that was used to generate these results
	Seed int64 `json:"seed,omitempty"`
}

// Validate validates this custom response model
func (m *CustomResponseModel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomResponseModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomResponseModel) UnmarshalBinary(b []byte) error {
	var res CustomResponseModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}