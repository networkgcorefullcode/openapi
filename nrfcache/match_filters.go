// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 Infosys Limited
// SPDX-FileCopyrightText: 2024 Canonical Ltd.
/*
 *  Match the NF profiles based on the parameters
 */

// This file contains apis to match the nf profiles based on the parameters provided in the
// Nnrf_NFDiscovery.SearchNFInstancesParamOpts. There is a match function provided for each NF type
// which must be updated with logic to compare profiles based on the applicable params in
// Nnrf_NFDiscovery.SearchNFInstancesParamOpts

package nrfcache

import (
	"encoding/json"
	"regexp"

	"github.com/omec-project/openapi/Nnrf_NFDiscovery"
	"github.com/omec-project/openapi/logger"
	"github.com/omec-project/openapi/models"
)

// compareExtSnssai compares two Snssai structs for equality, ignoring slice fields that cannot be compared directly.
func compareExtSnssai(a, b models.Snssai) bool {
	// Compare Sst and Sd fields (add more fields if needed)
	if a.Sst != b.Sst {
		return false
	}
	if a.Sd != b.Sd {
		return false
	}
	// If there are other comparable fields, add them here.
	return true
}

type MatchFilter func(profile *models.NfProfile, opts *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) (bool, error)

type MatchFilters map[models.NfType]MatchFilter

var matchFilters = MatchFilters{
	models.NfType_SMF:  MatchSmfProfile,
	models.NfType_AUSF: MatchAusfProfile,
	models.NfType_PCF:  MatchPcfProfile,
	models.NfType_NSSF: MatchNssfProfile,
	models.NfType_UDM:  MatchUdmProfile,
	models.NfType_AMF:  MatchAmfProfile,
}

func MatchSmfProfile(profile *models.NfProfile, opts *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) (bool, error) {
	if opts.ServiceNames.IsSet() {
		reqServiceNames := opts.ServiceNames.Value().([]models.ServiceName)
		matchCount := 0
		for _, sn := range reqServiceNames {
			for i := 0; i < len(*profile.NfServices); i++ {
				if (*profile.NfServices)[i].ServiceName == sn {
					matchCount++
					break
				}
			}
		}

		if matchCount == 0 {
			return false, nil
		}
	}

	if opts.Snssais.IsSet() {
		reqSnssais := opts.Snssais.Value().([]string)
		matchCount := 0

		for _, reqSnssai := range reqSnssais {
			var snssai models.Snssai
			err := json.Unmarshal([]byte(reqSnssai), &snssai)
			if err != nil {
				logger.NrfcacheLog.Errorf("error Unmarshaling nssai: %+v", err)
				return false, err
			}

			// Snssai in the smfInfo has priority
			if profile.SmfInfo != nil && profile.SmfInfo.SNssaiSmfInfoList != nil {
				for _, s := range *profile.SmfInfo.SNssaiSmfInfoList {
					if s.SNssai != nil && compareExtSnssai(*s.SNssai, snssai) {
						matchCount++
					}
				}
			} else if profile.AllowedNssais != nil {
				for _, s := range *profile.AllowedNssais {
					if compareExtSnssai(s, snssai) {
						matchCount++
					}
				}
			}
		}

		// if at least one matching snssai has been found
		if matchCount == 0 {
			return false, nil
		}
	}

	// validate dnn
	if opts.Dnn.IsSet() {
		// if a dnn is provided by the upper layer, check for the exact match
		// or wild card match
		dnnMatched := false

		if profile.SmfInfo != nil && profile.SmfInfo.SNssaiSmfInfoList != nil {
		matchDnnLoop:
			for _, s := range *profile.SmfInfo.SNssaiSmfInfoList {
				if s.DnnSmfInfoList != nil {
					for _, d := range *s.DnnSmfInfoList {
						if d.Dnn == opts.Dnn.Value() || d.Dnn == "*" {
							dnnMatched = true
							break matchDnnLoop
						}
					}
				}
			}
		}

		if !dnnMatched {
			return false, nil
		}
	}
	logger.NrfcacheLog.Infof("smf match found, nfInstance Id %v", profile.NfInstanceId)
	return true, nil
}

func MatchSupiRange(supi string, supiRange []models.SupiRange) bool {
	matchFound := false
	for _, s := range supiRange {
		if len(s.Pattern) > 0 {
			r, err := regexp.Compile(s.Pattern)
			if err != nil {
				logger.NrfcacheLog.Errorf("parsing pattern error: %v", err)
				return false
			}
			if r.MatchString(supi) {
				matchFound = true
				break
			}
		} else if s.Start <= supi && supi <= s.End {
			matchFound = true
			break
		}
	}

	return matchFound
}

func MatchAusfProfile(profile *models.NfProfile, opts *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) (bool, error) {
	matchFound := true
	if opts.Supi.IsSet() {
		if profile.AusfInfo != nil && len(profile.AusfInfo.SupiRanges) > 0 {
			matchFound = MatchSupiRange(opts.Supi.Value(), profile.AusfInfo.SupiRanges)
		}
	}
	logger.NrfcacheLog.Infof("ausf match found = %v", matchFound)
	return matchFound, nil
}

func MatchNssfProfile(profile *models.NfProfile, opts *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) (bool, error) {
	logger.NrfcacheLog.Infoln("nssf match found")
	return true, nil
}

func MatchAmfProfile(profile *models.NfProfile, opts *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) (bool, error) {
	if opts.TargetPlmnList.IsSet() {
		if profile.PlmnList != nil {
			plmnMatchCount := 0

			targetPlmnList := opts.TargetPlmnList.Value().([]string)
			for _, targetPlmn := range targetPlmnList {
				var plmn models.PlmnId
				err := json.Unmarshal([]byte(targetPlmn), &plmn)
				if err != nil {
					logger.NrfcacheLog.Errorf("error Unmarshaling plmn: %+v", err)
					return false, err
				}

				for _, profilePlmn := range *profile.PlmnList {
					if profilePlmn == plmn {
						plmnMatchCount++
						break
					}
				}
			}
			if plmnMatchCount == 0 {
				return false, nil
			}
		}
	}

	if profile.AmfInfo != nil {
		if opts.Guami.IsSet() {
			if profile.AmfInfo.GuamiList != nil {
				guamiMatchCount := 0

				guamiList := opts.Guami.Value().([]string)
				for _, guami := range guamiList {
					var guamiOpt models.Guami
					err := json.Unmarshal([]byte(guami), &guamiOpt)
					if err != nil {
						logger.NrfcacheLog.Errorf("error Unmarshaling guami: %+v", err)
						return false, err
					}

					for _, guami := range *profile.AmfInfo.GuamiList {
						if guamiOpt == guami {
							guamiMatchCount++
							break
						}
					}
				}
				if guamiMatchCount == 0 {
					return false, nil
				}
			}
		}

		if opts.AmfRegionId.IsSet() {
			if len(profile.AmfInfo.AmfRegionId) > 0 {
				if profile.AmfInfo.AmfRegionId != opts.AmfRegionId.Value() {
					return false, nil
				}
			}
		}

		if opts.AmfSetId.IsSet() {
			if len(profile.AmfInfo.AmfSetId) > 0 {
				if profile.AmfInfo.AmfSetId != opts.AmfSetId.Value() {
					return false, nil
				}
			}
		}

		if opts.TargetNfInstanceId.IsSet() {
			if profile.NfInstanceId != "" {
				if profile.NfInstanceId != opts.TargetNfInstanceId.Value() {
					return false, nil
				}
			}
		}
	}
	logger.NrfcacheLog.Infof("amf match found = %v", profile.NfInstanceId)
	return true, nil
}

func MatchPcfProfile(profile *models.NfProfile, opts *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) (bool, error) {
	matchFound := true
	if opts.Supi.IsSet() {
		if profile.PcfInfo != nil && len(profile.PcfInfo.SupiRanges) > 0 {
			matchFound = MatchSupiRange(opts.Supi.Value(), profile.PcfInfo.SupiRanges)
		}
	}
	logger.NrfcacheLog.Infof("pcf match found = %v", matchFound)
	return matchFound, nil
}

func MatchUdmProfile(profile *models.NfProfile, opts *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) (bool, error) {
	matchFound := true
	if opts.Supi.IsSet() {
		if profile.UdmInfo != nil && len(profile.UdmInfo.SupiRanges) > 0 {
			matchFound = MatchSupiRange(opts.Supi.Value(), profile.UdmInfo.SupiRanges)
		}
	}
	logger.NrfcacheLog.Infof("udm match found = %v", matchFound)
	return matchFound, nil
}
