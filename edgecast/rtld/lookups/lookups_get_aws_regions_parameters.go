// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package lookups

// This file was generated by codegen-sdk-go.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
)

// NewLookupsGetAwsRegionsParams creates a new LookupsGetAwsRegionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLookupsGetAwsRegionsParams() *LookupsGetAwsRegionsParams {
	return &LookupsGetAwsRegionsParams{}
}

/* LookupsGetAwsRegionsParams contains all the parameters to send to the API endpoint
   for the lookups get aws regions operation.

   Typically these are written to a http.Request.
*/
type LookupsGetAwsRegionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lookups get aws regions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LookupsGetAwsRegionsParams) WithDefaults() *LookupsGetAwsRegionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lookups get aws regions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LookupsGetAwsRegionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WriteToRequest extracts parameters and sets for the request to be consumed
func WriteToRequestLookupsGetAwsRegionsParams(o *LookupsGetAwsRegionsParams) (RequestParameters, error) {

	var res []error

	params := NewRequestParameters()

	if len(res) > 0 {
		return *params, errors.CompositeValidationError(res...)
	}
	return *params, nil
}
