// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package dcv

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NewDcvGetCertificateDomainDetailsParams creates a new DcvGetCertificateDomainDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcvGetCertificateDomainDetailsParams() DcvGetCertificateDomainDetailsParams {
	return DcvGetCertificateDomainDetailsParams{}
}

// DcvGetCertificateDomainDetailsParams contains all the parameters to send to the API
// endpoint for the dcv get certificate domain details operation. Typically these are written
// to a http.Request.
type DcvGetCertificateDomainDetailsParams struct {

	// DomainIds.
	DomainIds *string

	// ID.
	//
	// Format: int64
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcv get certificate domain details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcvGetCertificateDomainDetailsParams) WithDefaults() *DcvGetCertificateDomainDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcv get certificate domain details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcvGetCertificateDomainDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WriteToRequest extracts parameters and sets for the request to be consumed
func WriteToRequestDcvGetCertificateDomainDetailsParams(o DcvGetCertificateDomainDetailsParams) (RequestParameters, error) {

	var res []error

	params := NewRequestParameters()

	if o.DomainIds != nil {

		// query param domain_ids
		var qrDomainIds string

		if o.DomainIds != nil {
			qrDomainIds = *o.DomainIds
		}
		qDomainIds := qrDomainIds
		if qDomainIds != "" {

			params.QueryParams["domain_ids"] = qDomainIds
		}
	}

	// path param id
	params.PathParams["id"] = swag.FormatInt64(o.ID)

	if len(res) > 0 {
		return *params, errors.CompositeValidationError(res...)
	}
	return *params, nil
}
