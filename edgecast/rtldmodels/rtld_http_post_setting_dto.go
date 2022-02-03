// Code generated by go-swagger; DO NOT EDIT.

package rtldmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RtldHTTPPostSettingDto rtld Http post setting dto
//
// swagger:model RtldHttpPostSettingDto
type RtldHTTPPostSettingDto struct {

	// authentication type
	AuthenticationType string `json:"authentication_type,omitempty"`

	// destination endpoint
	DestinationEndpoint string `json:"destination_endpoint,omitempty"`

	// masked password
	MaskedPassword string `json:"masked_password,omitempty"`

	// masked token
	MaskedToken string `json:"masked_token,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this rtld Http post setting dto
func (m *RtldHTTPPostSettingDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rtld Http post setting dto based on context it is used
func (m *RtldHTTPPostSettingDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RtldHTTPPostSettingDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RtldHTTPPostSettingDto) UnmarshalBinary(b []byte) error {
	var res RtldHTTPPostSettingDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}