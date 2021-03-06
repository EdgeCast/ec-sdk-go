package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/edgecname"
	"github.com/EdgeCast/ec-sdk-go/edgecast/shared/enums"
)

func main() {
	// Setup - fill in the below variables before running this code
	accountNumber := "MY_ACCOUNT_NUMBER"
	apiToken := "MY_API_TOKEN"

	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = apiToken
	edgecnameService, err := edgecname.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating Edgecname Service: %v\n", err)
		return
	}

	// Create Edge CNAME
	cname := edgecname.EdgeCname{
		Name:                "test001.sharedectest.com",
		DirPath:             "/my/origin/path",
		EnableCustomReports: 1,
		OriginID:            -1,
		MediaTypeID:         3,
	}

	addParams := edgecname.NewAddEdgeCnameParams()
	addParams.AccountNumber = accountNumber
	addParams.EdgeCname = cname

	edgeCnameID, err := edgecnameService.AddEdgeCname(*addParams)

	if err != nil {
		fmt.Printf("error creating Edge CNAME: %v\n", err)
		return
	}

	// Get all Edge CNAMEs by Platform
	getAllParms := edgecname.NewGetAllEdgeCnameParams()
	getAllParms.AccountNumber = accountNumber
	getAllParms.Platform = enums.HttpLarge

	edgeCnames, err := edgecnameService.GetAllEdgeCnames(*getAllParms)

	if err != nil {
		fmt.Printf("error retrieving all Edge CNAMEs: %v\n", err)
		return
	}

	fmt.Printf("Edge CNAMEs retrieved: %v", edgeCnames)

	// Get a single Edge CNAME
	getParams := edgecname.NewGetEdgeCnameParams()
	getParams.AccountNumber = accountNumber
	getParams.EdgeCnameID = *edgeCnameID

	edgeCnameObj, err := edgecnameService.GetEdgeCname(*getParams)

	if err != nil {
		fmt.Printf("error retrieving Edge CNAME: %v\n", err)
		return
	}

	// Get Edge CNAME propgation status
	getStatusParams := edgecname.NewGetEdgeCnamePropagationStatusParams()
	getStatusParams.AccountNumber = accountNumber
	getStatusParams.EdgeCnameID = *edgeCnameID

	propagationStatus, err := edgecnameService.GetEdgeCnamePropagationStatus(
		*getStatusParams,
	)

	if err != nil {
		fmt.Printf("error retrieving Edge CNAME propagation status: %v\n", err)
		return
	}

	fmt.Printf("Edge CNAME propagation status: %v\n", propagationStatus)

	// Update Edge CNAME
	edgeCnameObj.EnableCustomReports = 0
	updateParams := edgecname.NewUpdateEdgeCnameParams()
	updateParams.AccountNumber = accountNumber
	updateParams.EdgeCname = *edgeCnameObj

	_, err = edgecnameService.UpdateEdgeCname(*updateParams)

	if err != nil {
		fmt.Printf("error updating Edge CNAME: %v\n", err)
		return
	}

	// Delete Edge CNAME
	deleteParams := edgecname.NewDeleteEdgeCnameParams()
	deleteParams.AccountNumber = accountNumber
	deleteParams.EdgeCname = *edgeCnameObj

	err = edgecnameService.DeleteEdgeCname(*deleteParams)

	if err != nil {
		fmt.Printf("error deleting Edge CNAME: %v\n", err)
		return
	}
}
