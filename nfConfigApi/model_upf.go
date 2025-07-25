// SPDX-FileCopyrightText: 2025 Canonical Ltd
//
// SPDX-License-Identifier: Apache-2.0
//

/*
WebConsole NFConfig API

API for managing access, mobility, policy, session and PLMN information.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nfConfigApi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Upf struct for Upf
type Upf struct {
	Hostname string `json:"hostname"`
	Port     *int32 `json:"port,omitempty"`
}

type _Upf Upf

// NewUpf instantiates a new Upf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpf(hostname string) *Upf {
	this := Upf{}
	this.Hostname = hostname
	return &this
}

// NewUpfWithDefaults instantiates a new Upf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpfWithDefaults() *Upf {
	this := Upf{}
	return &this
}

// GetHostname returns the Hostname field value
func (o *Upf) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *Upf) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *Upf) SetHostname(v string) {
	o.Hostname = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *Upf) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Upf) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *Upf) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *Upf) SetPort(v int32) {
	o.Port = &v
}

func (o Upf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Upf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hostname"] = o.Hostname
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

func (o *Upf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hostname",
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

	varUpf := _Upf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpf)

	if err != nil {
		return err
	}

	*o = Upf(varUpf)

	return err
}

type NullableUpf struct {
	value *Upf
	isSet bool
}

func (v NullableUpf) Get() *Upf {
	return v.value
}

func (v *NullableUpf) Set(val *Upf) {
	v.value = val
	v.isSet = true
}

func (v NullableUpf) IsSet() bool {
	return v.isSet
}

func (v *NullableUpf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpf(val *Upf) *NullableUpf {
	return &NullableUpf{value: val, isSet: true}
}

func (v NullableUpf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
