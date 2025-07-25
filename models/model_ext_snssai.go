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

// Snssai The sdRanges and wildcardSd attributes shall be exclusive from each other. If one of these attributes is present,  the sd attribute shall also be present and it shall contain one Slice Differentiator value within the range of SD  (if the sdRanges attribute is present) or with any value (if the wildcardSd attribute is present).
type Snssai struct {
	// Unsigned integer, within the range 0 to 255, representing the Slice/Service Type.  It indicates the expected Network Slice behaviour in terms of features and services. Values 0 to 127 correspond to the standardized SST range. Values 128 to 255 correspond  to the Operator-specific range. See clause 28.4.2 of 3GPP TS 23.003. Standardized values are defined in clause 5.15.2.2 of 3GPP TS 23.501.
	Sst int32 `json:"sst"`
	// 3-octet string, representing the Slice Differentiator, in hexadecimal representation. Each character in the string shall take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character representing the 4 most significant bits of the SD shall appear first in the string, and the character representing the 4 least significant bit of the SD shall appear last in the string. This is an optional parameter that complements the Slice/Service type(s) to allow to  differentiate amongst multiple Network Slices of the same Slice/Service type. This IE shall be absent if no SD value is associated with the SST.
	Sd string `json:"sd,omitempty"`
	// When present, it shall contain the range(s) of Slice Differentiator values supported for the Slice/Service Type value indicated in the sst attribute of the Snssai data type
	SdRanges []SdRange `json:"sdRanges,omitempty"`
	// When present, it shall be set to true, to indicate that all SD values are supported for the Slice/Service Type value indicated in the sst attribute of the Snssai data type.
	WildcardSd *bool `json:"wildcardSd,omitempty"`
}

type _ExtSnssai Snssai

// NewExtSnssai instantiates a new Snssai object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtSnssai(sst int32) *Snssai {
	this := Snssai{}
	this.Sst = sst
	return &this
}

// NewExtSnssaiWithDefaults instantiates a new Snssai object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtSnssaiWithDefaults() *Snssai {
	this := Snssai{}
	return &this
}

// GetSst returns the Sst field value
func (o *Snssai) GetSst() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sst
}

// GetSstOk returns a tuple with the Sst field value
// and a boolean to check if the value has been set.
func (o *Snssai) GetSstOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sst, true
}

// SetSst sets field value
func (o *Snssai) SetSst(v int32) {
	o.Sst = v
}

// GetSd returns the Sd field value if set, zero value otherwise.
func (o *Snssai) GetSd() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Sd
}

// GetSdOk returns a tuple with the Sd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snssai) GetSdOk() (*string, bool) {
	if o == nil || o.Sd == "" {
		return nil, false
	}
	return &o.Sd, true
}

// HasSd returns a boolean if a field has been set.
func (o *Snssai) HasSd() bool {
	if o != nil && o.Sd != "" {
		return true
	}

	return false
}

// SetSd gets a reference to the given string and assigns it to the Sd field.
func (o *Snssai) SetSd(v string) {
	o.Sd = v
}

// GetSdRanges returns the SdRanges field value if set, zero value otherwise.
func (o *Snssai) GetSdRanges() []SdRange {
	if o == nil || IsNil(o.SdRanges) {
		var ret []SdRange
		return ret
	}
	return o.SdRanges
}

// GetSdRangesOk returns a tuple with the SdRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snssai) GetSdRangesOk() ([]SdRange, bool) {
	if o == nil || IsNil(o.SdRanges) {
		return nil, false
	}
	return o.SdRanges, true
}

// HasSdRanges returns a boolean if a field has been set.
func (o *Snssai) HasSdRanges() bool {
	if o != nil && !IsNil(o.SdRanges) {
		return true
	}

	return false
}

// SetSdRanges gets a reference to the given []SdRange and assigns it to the SdRanges field.
func (o *Snssai) SetSdRanges(v []SdRange) {
	o.SdRanges = v
}

// GetWildcardSd returns the WildcardSd field value if set, zero value otherwise.
func (o *Snssai) GetWildcardSd() bool {
	if o == nil || IsNil(o.WildcardSd) {
		var ret bool
		return ret
	}
	return *o.WildcardSd
}

// GetWildcardSdOk returns a tuple with the WildcardSd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snssai) GetWildcardSdOk() (*bool, bool) {
	if o == nil || IsNil(o.WildcardSd) {
		return nil, false
	}
	return o.WildcardSd, true
}

// HasWildcardSd returns a boolean if a field has been set.
func (o *Snssai) HasWildcardSd() bool {
	if o != nil && !IsNil(o.WildcardSd) {
		return true
	}

	return false
}

// SetWildcardSd gets a reference to the given bool and assigns it to the WildcardSd field.
func (o *Snssai) SetWildcardSd(v bool) {
	o.WildcardSd = &v
}

func (o Snssai) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Snssai) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sst"] = o.Sst
	if o.Sd != "" {
		toSerialize["sd"] = o.Sd
	}
	if !IsNil(o.SdRanges) {
		toSerialize["sdRanges"] = o.SdRanges
	}
	if !IsNil(o.WildcardSd) {
		toSerialize["wildcardSd"] = o.WildcardSd
	}
	return toSerialize, nil
}

func (o *Snssai) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sst",
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

	varExtSnssai := _ExtSnssai{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExtSnssai)

	if err != nil {
		return err
	}

	*o = Snssai(varExtSnssai)

	return err
}

type NullableExtSnssai struct {
	value *Snssai
	isSet bool
}

func (v NullableExtSnssai) Get() *Snssai {
	return v.value
}

func (v *NullableExtSnssai) Set(val *Snssai) {
	v.value = val
	v.isSet = true
}

func (v NullableExtSnssai) IsSet() bool {
	return v.isSet
}

func (v *NullableExtSnssai) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtSnssai(val *Snssai) *NullableExtSnssai {
	return &NullableExtSnssai{value: val, isSet: true}
}

func (v NullableExtSnssai) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtSnssai) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
