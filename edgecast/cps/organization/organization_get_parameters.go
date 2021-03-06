// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package organization

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NewOrganizationGetParams creates a new OrganizationGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationGetParams() OrganizationGetParams {
	return OrganizationGetParams{}
}

// OrganizationGetParams contains all the parameters to send to the API
// endpoint for the organization get operation. Typically these are written
// to a http.Request.
type OrganizationGetParams struct {

	// ID.
	//
	// Format: int64
	ID int64

	// Name.
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the organization get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationGetParams) WithDefaults() *OrganizationGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organization get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WriteToRequest extracts parameters and sets for the request to be consumed
func WriteToRequestOrganizationGetParams(o OrganizationGetParams) (RequestParameters, error) {

	var res []error

	params := NewRequestParameters()

	// path param id
	params.PathParams["id"] = swag.FormatInt64(o.ID)

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			params.QueryParams["name"] = qName
		}
	}

	if len(res) > 0 {
		return *params, errors.CompositeValidationError(res...)
	}
	return *params, nil
}
