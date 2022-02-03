// Code generated by go-swagger; DO NOT EDIT.

package rtldmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RtldPlatformSettingDto rtld platform setting dto
//
// swagger:model RtldPlatformSettingDto
type RtldPlatformSettingDto struct {

	// aws s3
	AwsS3 *RtldAwsSettingDto `json:"aws_s3,omitempty"`

	// azure blob storage
	AzureBlobStorage *RtldAzureSettingDto `json:"azure_blob_storage,omitempty"`

	// custom cookies
	CustomCookies []string `json:"custom_cookies"`

	// custom request headers
	CustomRequestHeaders []string `json:"custom_request_headers"`

	// custom response headers
	CustomResponseHeaders []string `json:"custom_response_headers"`

	// customer id
	CustomerID int32 `json:"customer_id,omitempty"`

	// datadog
	Datadog *RtldDatadogSettingDto `json:"datadog,omitempty"`

	// delivery method
	DeliveryMethod string `json:"delivery_method,omitempty"`

	// downsampling rate
	DownsamplingRate float64 `json:"downsampling_rate,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// excluded host filter
	ExcludedHostFilter []string `json:"excluded_host_filter"`

	// fields
	Fields []string `json:"fields"`

	// host filter
	HostFilter []string `json:"host_filter"`

	// http post
	HTTPPost *RtldHTTPPostSettingDto `json:"http_post,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// log format
	LogFormat string `json:"log_format,omitempty"`

	// splunk enterprise
	SplunkEnterprise *RtldSplunkSettingDto `json:"splunk_enterprise,omitempty"`

	// status code filter
	StatusCodeFilter []string `json:"status_code_filter"`

	// sumo logic
	SumoLogic *RtldSumoLogicSettingDto `json:"sumo_logic,omitempty"`
}

// Validate validates this rtld platform setting dto
func (m *RtldPlatformSettingDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsS3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureBlobStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatadog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPPost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSplunkEnterprise(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSumoLogic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RtldPlatformSettingDto) validateAwsS3(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsS3) { // not required
		return nil
	}

	if m.AwsS3 != nil {
		if err := m.AwsS3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_s3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_s3")
			}
			return err
		}
	}

	return nil
}

func (m *RtldPlatformSettingDto) validateAzureBlobStorage(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureBlobStorage) { // not required
		return nil
	}

	if m.AzureBlobStorage != nil {
		if err := m.AzureBlobStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure_blob_storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure_blob_storage")
			}
			return err
		}
	}

	return nil
}

func (m *RtldPlatformSettingDto) validateDatadog(formats strfmt.Registry) error {
	if swag.IsZero(m.Datadog) { // not required
		return nil
	}

	if m.Datadog != nil {
		if err := m.Datadog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datadog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datadog")
			}
			return err
		}
	}

	return nil
}

func (m *RtldPlatformSettingDto) validateHTTPPost(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPPost) { // not required
		return nil
	}

	if m.HTTPPost != nil {
		if err := m.HTTPPost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_post")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_post")
			}
			return err
		}
	}

	return nil
}

func (m *RtldPlatformSettingDto) validateSplunkEnterprise(formats strfmt.Registry) error {
	if swag.IsZero(m.SplunkEnterprise) { // not required
		return nil
	}

	if m.SplunkEnterprise != nil {
		if err := m.SplunkEnterprise.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("splunk_enterprise")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("splunk_enterprise")
			}
			return err
		}
	}

	return nil
}

func (m *RtldPlatformSettingDto) validateSumoLogic(formats strfmt.Registry) error {
	if swag.IsZero(m.SumoLogic) { // not required
		return nil
	}

	if m.SumoLogic != nil {
		if err := m.SumoLogic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sumo_logic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sumo_logic")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rtld platform setting dto based on the context it is used
func (m *RtldPlatformSettingDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsS3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureBlobStorage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatadog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTPPost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSplunkEnterprise(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSumoLogic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RtldPlatformSettingDto) contextValidateAwsS3(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsS3 != nil {
		if err := m.AwsS3.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_s3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_s3")
			}
			return err
		}
	}

	return nil
}

func (m *RtldPlatformSettingDto) contextValidateAzureBlobStorage(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureBlobStorage != nil {
		if err := m.AzureBlobStorage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure_blob_storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure_blob_storage")
			}
			return err
		}
	}

	return nil
}

func (m *RtldPlatformSettingDto) contextValidateDatadog(ctx context.Context, formats strfmt.Registry) error {

	if m.Datadog != nil {
		if err := m.Datadog.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datadog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datadog")
			}
			return err
		}
	}

	return nil
}

func (m *RtldPlatformSettingDto) contextValidateHTTPPost(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTPPost != nil {
		if err := m.HTTPPost.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_post")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_post")
			}
			return err
		}
	}

	return nil
}

func (m *RtldPlatformSettingDto) contextValidateSplunkEnterprise(ctx context.Context, formats strfmt.Registry) error {

	if m.SplunkEnterprise != nil {
		if err := m.SplunkEnterprise.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("splunk_enterprise")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("splunk_enterprise")
			}
			return err
		}
	}

	return nil
}

func (m *RtldPlatformSettingDto) contextValidateSumoLogic(ctx context.Context, formats strfmt.Registry) error {

	if m.SumoLogic != nil {
		if err := m.SumoLogic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sumo_logic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sumo_logic")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RtldPlatformSettingDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RtldPlatformSettingDto) UnmarshalBinary(b []byte) error {
	var res RtldPlatformSettingDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}