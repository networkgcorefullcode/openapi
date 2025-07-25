// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * NRF NFManagement Service
 *
 * LmfInfo type definition
 *
 * API version: 1.4.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// LmfInfo represents LMF information
type LmfInfo struct {
	ServingClientTypes     []ExternalClientType `json:"servingClientTypes,omitempty"`
	LmfId                  *string              `json:"lmfId,omitempty"`
	ServingAccessTypes     []AccessType         `json:"servingAccessTypes,omitempty"`
	ServingAnNodeTypes     []AnNodeType         `json:"servingAnNodeTypes,omitempty"`
	ServingRatTypes        []RatType            `json:"servingRatTypes,omitempty"`
	TaiList                []Tai                `json:"taiList,omitempty"`
	TaiRangeList           []TaiRange           `json:"taiRangeList,omitempty"`
	SupportedGADShapes     []SupportedGADShapes `json:"supportedGADShapes,omitempty"`
	PruExistenceInfo       *PruExistenceInfo    `json:"pruExistenceInfo,omitempty"`
	PruSupportInd          *bool                `json:"pruSupportInd,omitempty"`
	RangingslposSupportInd *bool                `json:"rangingslposSupportInd,omitempty"`
	// user plane positioning capability is supported by the LMF
	UpPositioningInd *bool `json:"upPositioningInd,omitempty"`
}
