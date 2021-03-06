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

// DomainDcv domain dcv
//
// swagger:model DomainDcv
type DomainDcv struct {

	// dcv method
	// Enum: [Email DnsCnameToken DnsTxtToken]
	DcvMethod string `json:"dcv_method,omitempty"`

	// domain id
	DomainID int64 `json:"domain_id,omitempty"`
}

// Validate validates this domain dcv
func (m *DomainDcv) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDcvMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var domainDcvTypeDcvMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Email","DnsCnameToken","DnsTxtToken"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainDcvTypeDcvMethodPropEnum = append(domainDcvTypeDcvMethodPropEnum, v)
	}
}

const (

	// DomainDcvDcvMethodEmail captures enum value "Email"
	DomainDcvDcvMethodEmail string = "Email"

	// DomainDcvDcvMethodDNSCnameToken captures enum value "DnsCnameToken"
	DomainDcvDcvMethodDNSCnameToken string = "DnsCnameToken"

	// DomainDcvDcvMethodDNSTxtToken captures enum value "DnsTxtToken"
	DomainDcvDcvMethodDNSTxtToken string = "DnsTxtToken"
)

// prop value enum
func (m *DomainDcv) validateDcvMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, domainDcvTypeDcvMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DomainDcv) validateDcvMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.DcvMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateDcvMethodEnum("dcv_method", "body", m.DcvMethod); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain dcv based on context it is used
func (m *DomainDcv) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainDcv) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainDcv) UnmarshalBinary(b []byte) error {
	var res DomainDcv
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
