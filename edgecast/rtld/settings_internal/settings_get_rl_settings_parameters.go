// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package settings_internal

// This file was generated by codegen-sdk-go.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NewSettingsGetRlSettingsParams creates a new SettingsGetRlSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSettingsGetRlSettingsParams() *SettingsGetRlSettingsParams {
	return &SettingsGetRlSettingsParams{}
}

/* SettingsGetRlSettingsParams contains all the parameters to send to the API endpoint
   for the settings get rl settings operation.

   Typically these are written to a http.Request.
*/
type SettingsGetRlSettingsParams struct {

	// Page.
	//
	// Format: int32
	Page *int32

	// PageSize.
	//
	// Format: int32
	PageSize *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the settings get rl settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingsGetRlSettingsParams) WithDefaults() *SettingsGetRlSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the settings get rl settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingsGetRlSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WriteToRequest extracts parameters and sets for the request to be consumed
func WriteToRequestSettingsGetRlSettingsParams(o *SettingsGetRlSettingsParams) (RequestParameters, error) {

	var res []error

	params := NewRequestParameters()

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			params.QueryParams["page"] = qPage
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			params.QueryParams["page_size"] = qPageSize
		}
	}

	if len(res) > 0 {
		return *params, errors.CompositeValidationError(res...)
	}
	return *params, nil
}
