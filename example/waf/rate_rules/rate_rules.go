// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/rate"
)

// Demonstrates the usage of WAF Rate Rules
//
// Usage:
// go run rate.go
//
// For detailed information about Rate Rules in WAF, please refer to:
// https://docs.edgecast.com/cdn/#Web-Security/Rate-Rules.htm
func main() {

	// Setup - fill in the below variables before running this code
	accountNumber := "MY_ACCOUNT_NUMBER"
	apiToken := "MY_API_TOKEN"

	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = apiToken
	wafService, err := waf.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	// First, we'll set up a new rule to use in this example
	rule := setupRateRule(accountNumber)

	fmt.Println("")
	fmt.Println("**** CREATE ****")
	fmt.Println("")
	fmt.Printf("Creating Rate Rule: %+v\n", rule)
	ruleID, err := wafService.Rate.AddRateRule(
		rate.AddRateRuleParams{
			AccountNumber: accountNumber,
			RateRule:      rule,
		})

	if err != nil {
		fmt.Printf("failed to create Rate Rule: %+v\n", err)
		return
	} else {
		fmt.Printf("successfully created Rate Rule: %+v\n", ruleID)
	}

	fmt.Println("")
	fmt.Println("**** GET ****")
	fmt.Println("")
	getResponse, err := wafService.Rate.GetRateRule(
		rate.GetRateRuleParams{
			AccountNumber: accountNumber,
			RateRuleID:    ruleID,
		})

	if err != nil {
		fmt.Printf("Failed to retrieve Rate Rule: %+v\n", err)
		return
	} else {
		fmt.Printf("Successfully retrieved Rate Rule: %+v\n", getResponse)
	}

	fmt.Println("")
	fmt.Println("**** GET ALL ****")
	fmt.Println("")
	getAllResponse, err := wafService.Rate.GetAllRateRules(
		rate.GetAllRateRulesParams{
			AccountNumber: accountNumber,
		})

	if err != nil {
		fmt.Printf("Failed to retrieve all Rate Rules: %+v\n", err)
		return
	} else {
		fmt.Printf(
			"Successfully retrieved all Rate Rules: %+v\n",
			getAllResponse)
	}

	fmt.Println("")
	fmt.Println("**** UPDATE ****")
	fmt.Println("")
	rule.Name = "Updated rule from example"

	err = wafService.Rate.UpdateRateRule(
		rate.UpdateRateRuleParams{
			AccountNumber: accountNumber,
			RateRule:      rule,
			RateRuleID:    ruleID,
		})

	if err != nil {
		fmt.Printf("Failed to update Rate Rule: %+v\n", err)
		return
	} else {
		fmt.Println("Successfully updated Rate Rule")
	}

	fmt.Println("")
	fmt.Println("**** DELETE ****")
	fmt.Println("")
	err = wafService.Rate.DeleteRateRule(
		rate.DeleteRateRuleParams{
			AccountNumber: accountNumber,
			RateRuleID:    ruleID,
		})
	if err != nil {
		fmt.Printf("Failed to delete Rate Rule: %+v\n", err)
	} else {
		fmt.Println("Successfully deleted Rate Rule")
	}
}

func setupRateRule(accountNumber string) rate.RateRule {
	return rate.RateRule{
		Name:        "Rate Rule 1",
		Keys:        []string{"IP", "USER_AGENT"},
		DurationSec: 5,
		Num:         10,
		CustomerID:  accountNumber,
		ConditionGroups: []rate.ConditionGroup{
			{
				Name: "Group 1",
				Conditions: []rate.Condition{
					{
						Target: rate.Target{
							Type: "REQUEST_METHOD",
						},
						OP: rate.OP{
							Type:   "EM",
							Values: []string{"POST", "GET"},
						},
					},
					{
						Target: rate.Target{
							Type: "REMOTE_ADDR",
						},
						OP: rate.OP{
							Type:   "IPMATCH",
							Values: []string{"10.10.2.3", "10.10.2.4"},
						},
					},
				},
			},
			{
				Name: "Group 2",
				Conditions: []rate.Condition{
					{
						Target: rate.Target{
							Type:  "REQUEST_HEADERS",
							Value: "User-Agent",
						},
						OP: rate.OP{
							Type: "EM",
							Values: []string{
								"Mozilla/5.0", "Chrome/91.0.4472.114"},
						},
					},
					{
						Target: rate.Target{
							Type: "FILE_EXT",
						},
						OP: rate.OP{
							Type:  "RX",
							Value: "(.*?)\\.(jpg|gif|doc|pdf)$",
						},
					},
				},
			},
		},
	}
}
