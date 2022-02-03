// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package settings_internal

// This file was generated by codegen-sdk-go.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// New creates a new settings internal API client.
func New(c ecclient.APIClient, baseAPIURL string) ClientService {
	return &Client{c, baseAPIURL}
}

/*
Client for settings internal API
*/
type Client struct {
	client     ecclient.APIClient
	baseAPIURL string
}

// ClientService is the interface for Client methods
type ClientService interface {
	SettingsGetRlSettings(params *SettingsGetRlSettingsParams) (*SettingsGetRlSettingsOK, error)

	SettingsGetSettingsByPlatform(params *SettingsGetSettingsByPlatformParams) (*SettingsGetSettingsByPlatformOK, error)

	SettingsGetWafSettings(params *SettingsGetWafSettingsParams) (*SettingsGetWafSettingsOK, error)
}

/*
  SettingsGetRlSettings settings get rl settings API
*/
func (a *Client) SettingsGetRlSettings(params *SettingsGetRlSettingsParams) (*SettingsGetRlSettingsOK, error) {
	// Validate the params before sending
	if params == nil {
		params = NewSettingsGetRlSettingsParams()
	}

	//Set parameters
	results, err := WriteToRequestSettingsGetRlSettingsParams(params)
	if err != nil {
		return nil, err
	}

	method, err := ecclient.ToHTTPMethod("GET")
	if err != nil {
		return nil, fmt.Errorf("SettingsGetRlSettings: %v", err)
	}

	parsedResponse := &SettingsGetRlSettingsOK{}

	_, err = a.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         method,
		Path:           a.baseAPIURL + "/v1.0/rl/settings",
		RawBody:        results.Body,
		PathParams:     results.PathParams,
		QueryParams:    results.QueryParams,
		ParsedResponse: parsedResponse,
	})

	if err != nil {
		return nil, fmt.Errorf("SettingsGetRlSettings: %v", err)
	}

	return parsedResponse, nil
}

/*
  SettingsGetSettingsByPlatform settings get settings by platform API
*/
func (a *Client) SettingsGetSettingsByPlatform(params *SettingsGetSettingsByPlatformParams) (*SettingsGetSettingsByPlatformOK, error) {
	// Validate the params before sending
	if params == nil {
		params = NewSettingsGetSettingsByPlatformParams()
	}

	//Set parameters
	results, err := WriteToRequestSettingsGetSettingsByPlatformParams(params)
	if err != nil {
		return nil, err
	}

	method, err := ecclient.ToHTTPMethod("GET")
	if err != nil {
		return nil, fmt.Errorf("SettingsGetSettingsByPlatform: %v", err)
	}

	parsedResponse := &SettingsGetSettingsByPlatformOK{}

	_, err = a.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         method,
		Path:           a.baseAPIURL + "/v1.0/platforms/{platformId}/settings",
		RawBody:        results.Body,
		PathParams:     results.PathParams,
		QueryParams:    results.QueryParams,
		ParsedResponse: parsedResponse,
	})

	if err != nil {
		return nil, fmt.Errorf("SettingsGetSettingsByPlatform: %v", err)
	}

	return parsedResponse, nil
}

/*
  SettingsGetWafSettings settings get waf settings API
*/
func (a *Client) SettingsGetWafSettings(params *SettingsGetWafSettingsParams) (*SettingsGetWafSettingsOK, error) {
	// Validate the params before sending
	if params == nil {
		params = NewSettingsGetWafSettingsParams()
	}

	//Set parameters
	results, err := WriteToRequestSettingsGetWafSettingsParams(params)
	if err != nil {
		return nil, err
	}

	method, err := ecclient.ToHTTPMethod("GET")
	if err != nil {
		return nil, fmt.Errorf("SettingsGetWafSettings: %v", err)
	}

	parsedResponse := &SettingsGetWafSettingsOK{}

	_, err = a.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         method,
		Path:           a.baseAPIURL + "/v1.0/waf/settings",
		RawBody:        results.Body,
		PathParams:     results.PathParams,
		QueryParams:    results.QueryParams,
		ParsedResponse: parsedResponse,
	})

	if err != nil {
		return nil, fmt.Errorf("SettingsGetWafSettings: %v", err)
	}

	return parsedResponse, nil
}

type RequestParameters struct {
	QueryParams map[string]string
	PathParams  map[string]string
	Body        interface{}
}

func NewRequestParameters() *RequestParameters {
	return &RequestParameters{
		QueryParams: make(map[string]string),
		PathParams:  make(map[string]string),
	}
}
