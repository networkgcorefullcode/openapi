// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import "encoding/json"

type IdentityRange struct {
	Start   string `json:"start,omitempty" yaml:"start" bson:"start" mapstructure:"Start"`
	End     string `json:"end,omitempty" yaml:"end" bson:"end" mapstructure:"End"`
	Pattern string `json:"pattern,omitempty" yaml:"pattern" bson:"pattern" mapstructure:"Pattern"`
}

// NewIdentityRange instantiates a new IdentityRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityRange() *IdentityRange {
	this := IdentityRange{}
	return &this
}

// NewIdentityRangeWithDefaults instantiates a new IdentityRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityRangeWithDefaults() *IdentityRange {
	this := IdentityRange{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *IdentityRange) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRange) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *IdentityRange) HasStart() bool {
	return o != nil && o.Start != ""
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *IdentityRange) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *IdentityRange) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRange) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *IdentityRange) HasEnd() bool {
	return o != nil && o.End != ""
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *IdentityRange) SetEnd(v string) {
	o.End = v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *IdentityRange) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRange) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *IdentityRange) HasPattern() bool {
	return o != nil && o.Pattern != ""
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *IdentityRange) SetPattern(v string) {
	o.Pattern = v
}

func (o IdentityRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Start != "" {
		toSerialize["start"] = o.Start
	}
	if o.End != "" {
		toSerialize["end"] = o.End
	}
	if o.Pattern != "" {
		toSerialize["pattern"] = o.Pattern
	}
	return toSerialize, nil
}

type NullableIdentityRange struct {
	value *IdentityRange
	isSet bool
}

func (v NullableIdentityRange) Get() *IdentityRange {
	return v.value
}

func (v *NullableIdentityRange) Set(val *IdentityRange) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityRange) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityRange(val *IdentityRange) *NullableIdentityRange {
	return &NullableIdentityRange{value: val, isSet: true}
}

func (v NullableIdentityRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
