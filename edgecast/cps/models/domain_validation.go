// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainValidation domain validation
//
// swagger:model DomainValidation
type DomainValidation struct {

	// domain id
	DomainID int32 `json:"domain_id,omitempty"`

	// domain names
	DomainNames []string `json:"domain_names"`

	// status
	// Enum: [Unknown Pending Rejected Approved NA]
	Status string `json:"status,omitempty"`
}

// Validate validates this domain validation
func (m *DomainValidation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var domainValidationTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Pending","Rejected","Approved","NA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainValidationTypeStatusPropEnum = append(domainValidationTypeStatusPropEnum, v)
	}
}

const (

	// DomainValidationStatusUnknown captures enum value "Unknown"
	DomainValidationStatusUnknown string = "Unknown"

	// DomainValidationStatusPending captures enum value "Pending"
	DomainValidationStatusPending string = "Pending"

	// DomainValidationStatusRejected captures enum value "Rejected"
	DomainValidationStatusRejected string = "Rejected"

	// DomainValidationStatusApproved captures enum value "Approved"
	DomainValidationStatusApproved string = "Approved"

	// DomainValidationStatusNA captures enum value "NA"
	DomainValidationStatusNA string = "NA"
)

// prop value enum
func (m *DomainValidation) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, domainValidationTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DomainValidation) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain validation based on context it is used
func (m *DomainValidation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainValidation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainValidation) UnmarshalBinary(b []byte) error {
	var res DomainValidation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
