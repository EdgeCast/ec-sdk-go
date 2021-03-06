// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package cps

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

import (
	"fmt"
	"net/url"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecauth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"

	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/appendix"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/certificate"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/customer"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/dcv"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/organization"
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/task"
)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/sec/cps"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

//New creates a new certificate provisioning API v2 client
func New(config edgecast.SDKConfig) (*CpsService, error) {

	apiURL, err := url.Parse(config.BaseAPIURL.String() + DefaultBasePath)
	if err != nil {
		return nil, fmt.Errorf("CpsService.New(): %v", err)
	}

	// OAuth2 authentication
	authProvider, err := ecauth.NewIDSAuthorizationProvider(config.BaseIDSURL, ecauth.OAuth2Credentials(config.IDSCredentials))
	if err != nil {

		//Token authentication
		authTokenProvider, err := ecauth.NewTokenAuthorizationProvider(config.APIToken)
		if err != nil {
			return nil, fmt.Errorf("CpsService.New(): %v", err)
		}
		c := ecclient.New(ecclient.ClientConfig{
			BaseAPIURL:   *apiURL,
			UserAgent:    config.UserAgent,
			Logger:       config.Logger,
			AuthProvider: authTokenProvider,
		})

		return &CpsService{
			client:       c,
			Logger:       config.Logger,
			Appendix:     appendix.New(c, c.Config.BaseAPIURL.String()),
			Certificate:  certificate.New(c, c.Config.BaseAPIURL.String()),
			Customer:     customer.New(c, c.Config.BaseAPIURL.String()),
			Dcv:          dcv.New(c, c.Config.BaseAPIURL.String()),
			Organization: organization.New(c, c.Config.BaseAPIURL.String()),
			Task:         task.New(c, c.Config.BaseAPIURL.String()),
		}, nil

	} else {

		c := ecclient.New(ecclient.ClientConfig{
			BaseAPIURL:   *apiURL,
			UserAgent:    config.UserAgent,
			Logger:       config.Logger,
			AuthProvider: authProvider,
		})

		return &CpsService{
			client:       c,
			Logger:       config.Logger,
			Appendix:     appendix.New(c, c.Config.BaseAPIURL.String()),
			Certificate:  certificate.New(c, c.Config.BaseAPIURL.String()),
			Customer:     customer.New(c, c.Config.BaseAPIURL.String()),
			Dcv:          dcv.New(c, c.Config.BaseAPIURL.String()),
			Organization: organization.New(c, c.Config.BaseAPIURL.String()),
			Task:         task.New(c, c.Config.BaseAPIURL.String()),
		}, nil
	}
}

// CpsService is a client for certificate provisioning API v2
type CpsService struct {
	Appendix appendix.ClientService

	Certificate certificate.ClientService

	Customer customer.ClientService

	Dcv dcv.ClientService

	Organization organization.ClientService

	Task task.ClientService

	client ecclient.APIClient

	clientConfig ecclient.ClientConfig

	Logger eclog.Logger
}
