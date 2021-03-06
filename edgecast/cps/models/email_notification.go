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

// EmailNotification email notification
//
// swagger:model EmailNotification
type EmailNotification struct {

	// emails
	Emails []string `json:"emails"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// notification type
	// Enum: [CertificateRenewal CertificateExpiring PendingValidations]
	NotificationType string `json:"notification_type,omitempty"`
}

// Validate validates this email notification
func (m *EmailNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotificationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var emailNotificationTypeNotificationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CertificateRenewal","CertificateExpiring","PendingValidations"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emailNotificationTypeNotificationTypePropEnum = append(emailNotificationTypeNotificationTypePropEnum, v)
	}
}

const (

	// EmailNotificationNotificationTypeCertificateRenewal captures enum value "CertificateRenewal"
	EmailNotificationNotificationTypeCertificateRenewal string = "CertificateRenewal"

	// EmailNotificationNotificationTypeCertificateExpiring captures enum value "CertificateExpiring"
	EmailNotificationNotificationTypeCertificateExpiring string = "CertificateExpiring"

	// EmailNotificationNotificationTypePendingValidations captures enum value "PendingValidations"
	EmailNotificationNotificationTypePendingValidations string = "PendingValidations"
)

// prop value enum
func (m *EmailNotification) validateNotificationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emailNotificationTypeNotificationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmailNotification) validateNotificationType(formats strfmt.Registry) error {
	if swag.IsZero(m.NotificationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateNotificationTypeEnum("notification_type", "body", m.NotificationType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this email notification based on context it is used
func (m *EmailNotification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmailNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailNotification) UnmarshalBinary(b []byte) error {
	var res EmailNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
