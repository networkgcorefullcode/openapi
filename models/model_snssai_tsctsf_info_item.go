/*
NRF NFManagement Service

NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.4.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// SnssaiTsctsfInfoItem Set of parameters supported by TSCTSF for a given S-NSSAI
type SnssaiTsctsfInfoItem struct {
	SNssai      Snssai              `json:"sNssai"`
	DnnInfoList []DnnTsctsfInfoItem `json:"dnnInfoList"`
}

type _SnssaiTsctsfInfoItem SnssaiTsctsfInfoItem

// NewSnssaiTsctsfInfoItem instantiates a new SnssaiTsctsfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssaiTsctsfInfoItem(sNssai Snssai, dnnInfoList []DnnTsctsfInfoItem) *SnssaiTsctsfInfoItem {
	this := SnssaiTsctsfInfoItem{}
	this.SNssai = sNssai
	this.DnnInfoList = dnnInfoList
	return &this
}

// NewSnssaiTsctsfInfoItemWithDefaults instantiates a new SnssaiTsctsfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiTsctsfInfoItemWithDefaults() *SnssaiTsctsfInfoItem {
	this := SnssaiTsctsfInfoItem{}
	return &this
}

// GetSNssai returns the SNssai field value
func (o *SnssaiTsctsfInfoItem) GetSNssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.SNssai
}

// GetSNssaiOk returns a tuple with the SNssai field value
// and a boolean to check if the value has been set.
func (o *SnssaiTsctsfInfoItem) GetSNssaiOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SNssai, true
}

// SetSNssai sets field value
func (o *SnssaiTsctsfInfoItem) SetSNssai(v Snssai) {
	o.SNssai = v
}

// GetDnnInfoList returns the DnnInfoList field value
func (o *SnssaiTsctsfInfoItem) GetDnnInfoList() []DnnTsctsfInfoItem {
	if o == nil {
		var ret []DnnTsctsfInfoItem
		return ret
	}

	return o.DnnInfoList
}

// GetDnnInfoListOk returns a tuple with the DnnInfoList field value
// and a boolean to check if the value has been set.
func (o *SnssaiTsctsfInfoItem) GetDnnInfoListOk() ([]DnnTsctsfInfoItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnnInfoList, true
}

// SetDnnInfoList sets field value
func (o *SnssaiTsctsfInfoItem) SetDnnInfoList(v []DnnTsctsfInfoItem) {
	o.DnnInfoList = v
}

func (o SnssaiTsctsfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnssaiTsctsfInfoItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sNssai"] = o.SNssai
	toSerialize["dnnInfoList"] = o.DnnInfoList
	return toSerialize, nil
}

func (o *SnssaiTsctsfInfoItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sNssai",
		"dnnInfoList",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSnssaiTsctsfInfoItem := _SnssaiTsctsfInfoItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSnssaiTsctsfInfoItem)

	if err != nil {
		return err
	}

	*o = SnssaiTsctsfInfoItem(varSnssaiTsctsfInfoItem)

	return err
}

type NullableSnssaiTsctsfInfoItem struct {
	value *SnssaiTsctsfInfoItem
	isSet bool
}

func (v NullableSnssaiTsctsfInfoItem) Get() *SnssaiTsctsfInfoItem {
	return v.value
}

func (v *NullableSnssaiTsctsfInfoItem) Set(val *SnssaiTsctsfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssaiTsctsfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssaiTsctsfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssaiTsctsfInfoItem(val *SnssaiTsctsfInfoItem) *NullableSnssaiTsctsfInfoItem {
	return &NullableSnssaiTsctsfInfoItem{value: val, isSet: true}
}

func (v NullableSnssaiTsctsfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssaiTsctsfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
