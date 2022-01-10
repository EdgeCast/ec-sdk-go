package origin

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/shared/ecmodels"
)

// GetAllOrigins retrieves a list of customer origin configurations associated
// with the provided platform.
func (svc *OriginService) GetAllOrigins(
	params GetAllOriginsParams,
) (*[]OriginGetOK, error) {
	request, err := svc.BuildRequest(
		"GET",
		fmt.Sprintf("v2/mcc/customers/%s/origins/%s",
			params.AccountNumber,
			params.MediaTypeID.String(),
		),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("GetAllOrigins: %v", err)
	}

	parsedResponse := &[]OriginGetOK{}
	_, err = svc.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("GetAllOrigins: %v", err)
	}

	return parsedResponse, nil
}

// AddOrigin adds a customer origin to the specified platform.
func (svc *OriginService) AddOrigin(params AddOriginParams) (*int, error) {
	request, err := svc.BuildRequest(
		"POST",
		fmt.Sprintf(
			"v2/mcc/customers/%s/origins/%s", params.AccountNumber,
			params.MediaTypeID.String(),
		),
		params.Origin,
	)
	if err != nil {
		return nil, fmt.Errorf("AddOrigin: %v", err)
	}

	parsedResponse := &AddUpdateOriginOK{}
	_, err = svc.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("AddOrigin: %v", err)
	}

	return &parsedResponse.CustomerOriginID, nil
}

// GetOrigin retrieves the properties of a customer origin configuration.
func (svc *OriginService) GetOrigin(params GetOriginParams) (*OriginGetOK, error) {
	request, err := svc.BuildRequest(
		"GET",
		fmt.Sprintf("v2/mcc/customers/%s/origins/%s/%d",
			params.AccountNumber,
			params.MediaTypeID.String(),
			params.CustomerOriginID,
		),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("GetOrigin: %v", err)
	}

	parsedResponse := &OriginGetOK{}
	_, err = svc.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("GetOrigin: %v", err)
	}

	return parsedResponse, nil
}

// UpdateOrigin sets the properties for a customer origin.
func (svc *OriginService) UpdateOrigin(params UpdateOriginParams) (*int, error) {
	request, err := svc.BuildRequest(
		"PUT",
		fmt.Sprintf(
			"v2/mcc/customers/%s/origins/%s/%d",
			params.AccountNumber,
			params.Origin.MediaTypeID,
			params.Origin.ID,
		),
		params.Origin,
	)
	if err != nil {
		return nil, fmt.Errorf("UpdateOrigin: %v", err)
	}

	parsedResponse := &AddUpdateOriginOK{}
	_, err = svc.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("UpdateOrigin: %v", err)
	}

	return &parsedResponse.CustomerOriginID, nil
}

// DeleteOrigin deletes a customer origin.
func (svc *OriginService) DeleteOrigin(params DeleteOriginParams) error {
	request, err := svc.BuildRequest(
		"DELETE",
		fmt.Sprintf("v2/mcc/customers/%s/origins/%d",
			params.AccountNumber,
			params.Origin.ID,
		),
		nil,
	)
	if err != nil {
		return fmt.Errorf("DeleteOrigin: %v", err)
	}

	_, err = svc.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("DeleteOrigin: %v", err)
	}

	return nil
}

// GetCDNIPBlocks retrieves a list of IPv4 and IPv6 blocks used by our CDN
// service. Ensure that our CDN may communicate with your web servers by
// allowlisting these IP blocks on your firewall.
func (svc *OriginService) GetCDNIPBlocks() (*CDNIPBlocksOK, error) {
	request, err := svc.BuildRequest("GET", "v2/mcc/customers/superblocks", nil)
	if err != nil {
		return nil, fmt.Errorf("GetCDNIPBlocks: %v", err)
	}

	parsedResponse := &CDNIPBlocksOK{}
	_, err = svc.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("GetCDNIPBlocks: %v", err)
	}

	return parsedResponse, nil
}

// GetOriginPropagationStatus retrieves a list of IPv4 and IPv6 blocks used by
// our CDN service. Ensure that our CDN may communicate with your web servers by
// allowlisting these IP blocks on your firewall.
func (svc *OriginService) GetOriginPropagationStatus(
	params GetOriginPropagationStatusParams,
) (*ecmodels.PropagationStatus, error) {
	request, err := svc.BuildRequest(
		"GET",
		fmt.Sprintf(
			"v2/mcc/customers/%s/origins/%d/status",
			params.AccountNumber,
			params.CustomerOriginID,
		),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("GetOriginPropagationStatus: %v", err)
	}

	parsedResponse := &ecmodels.PropagationStatus{}
	_, err = svc.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("GetOriginPropagationStatus: %v", err)
	}

	return parsedResponse, nil
}

// GetOriginShieldPOPs lists the available Origin Shield locations for the
// specified platform. This list consists of the name, POP code, and region for
// each POP that can provide Origin Shield protection to a customer origin
// server. These abbreviations can then be used to set or to interpret Origin
// Shield settings for a customer origin.
// This applies to HTTPLarge and HTTPSmall platform origins
func (svc *OriginService) GetOriginShieldPOPs(
	params GetOriginShieldPOPsParams,
) (*[]ShieldPOP, error) {
	request, err := svc.BuildRequest(
		"GET",
		fmt.Sprintf(
			"v2/mcc/customers/%s/origins/%s/shieldpops",
			params.AccountNumber,
			params.MediaTypeID.String(),
		),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("GetOriginShieldPOPs: %v", err)
	}

	parsedResponse := &[]ShieldPOP{}
	_, err = svc.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("GetOriginShieldPOPs: %v", err)
	}

	return parsedResponse, nil
}

// ReselectADNGateways reevaluates and defines a failover list of ADN gateways
// for the specified customer origin configuration.
// This applies only to ADN platform origins
func (svc *OriginService) ReselectADNGateways(
	params ReselectADNGatewaysParams,
) error {
	request, err := svc.BuildRequest(
		"PUT",
		fmt.Sprintf(
			"v2/mcc/customers/%s/origins/%s/%v/reselect",
			params.AccountNumber,
			params.MediaTypeID.String(),
			params.CustomerOriginID,
		),
		nil,
	)
	if err != nil {
		return fmt.Errorf("ReselectADNGateways: %v", err)
	}
	_, err = svc.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("ReselectADNGateways: %v", err)
	}

	return nil
}
