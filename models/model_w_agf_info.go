/*
NRF NFManagement Service

NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.4.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// WAgfInfo Information of the W-AGF end-points
type WAgfInfo struct {
	Ipv4EndpointAddresses []string   `json:"ipv4EndpointAddresses,omitempty"`
	Ipv6EndpointAddresses []Ipv6Addr `json:"ipv6EndpointAddresses,omitempty"`
	// Fully Qualified Domain Name
	EndpointFqdn *string `json:"endpointFqdn,omitempty"`
}

// NewWAgfInfo instantiates a new WAgfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWAgfInfo() *WAgfInfo {
	this := WAgfInfo{}
	return &this
}

// NewWAgfInfoWithDefaults instantiates a new WAgfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWAgfInfoWithDefaults() *WAgfInfo {
	this := WAgfInfo{}
	return &this
}

// GetIpv4EndpointAddresses returns the Ipv4EndpointAddresses field value if set, zero value otherwise.
func (o *WAgfInfo) GetIpv4EndpointAddresses() []string {
	if o == nil || IsNil(o.Ipv4EndpointAddresses) {
		var ret []string
		return ret
	}
	return o.Ipv4EndpointAddresses
}

// GetIpv4EndpointAddressesOk returns a tuple with the Ipv4EndpointAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAgfInfo) GetIpv4EndpointAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.Ipv4EndpointAddresses) {
		return nil, false
	}
	return o.Ipv4EndpointAddresses, true
}

// HasIpv4EndpointAddresses returns a boolean if a field has been set.
func (o *WAgfInfo) HasIpv4EndpointAddresses() bool {
	if o != nil && !IsNil(o.Ipv4EndpointAddresses) {
		return true
	}

	return false
}

// SetIpv4EndpointAddresses gets a reference to the given []string and assigns it to the Ipv4EndpointAddresses field.
func (o *WAgfInfo) SetIpv4EndpointAddresses(v []string) {
	o.Ipv4EndpointAddresses = v
}

// GetIpv6EndpointAddresses returns the Ipv6EndpointAddresses field value if set, zero value otherwise.
func (o *WAgfInfo) GetIpv6EndpointAddresses() []Ipv6Addr {
	if o == nil || IsNil(o.Ipv6EndpointAddresses) {
		var ret []Ipv6Addr
		return ret
	}
	return o.Ipv6EndpointAddresses
}

// GetIpv6EndpointAddressesOk returns a tuple with the Ipv6EndpointAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAgfInfo) GetIpv6EndpointAddressesOk() ([]Ipv6Addr, bool) {
	if o == nil || IsNil(o.Ipv6EndpointAddresses) {
		return nil, false
	}
	return o.Ipv6EndpointAddresses, true
}

// HasIpv6EndpointAddresses returns a boolean if a field has been set.
func (o *WAgfInfo) HasIpv6EndpointAddresses() bool {
	if o != nil && !IsNil(o.Ipv6EndpointAddresses) {
		return true
	}

	return false
}

// SetIpv6EndpointAddresses gets a reference to the given []Ipv6Addr and assigns it to the Ipv6EndpointAddresses field.
func (o *WAgfInfo) SetIpv6EndpointAddresses(v []Ipv6Addr) {
	o.Ipv6EndpointAddresses = v
}

// GetEndpointFqdn returns the EndpointFqdn field value if set, zero value otherwise.
func (o *WAgfInfo) GetEndpointFqdn() string {
	if o == nil || IsNil(o.EndpointFqdn) {
		var ret string
		return ret
	}
	return *o.EndpointFqdn
}

// GetEndpointFqdnOk returns a tuple with the EndpointFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAgfInfo) GetEndpointFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointFqdn) {
		return nil, false
	}
	return o.EndpointFqdn, true
}

// HasEndpointFqdn returns a boolean if a field has been set.
func (o *WAgfInfo) HasEndpointFqdn() bool {
	if o != nil && !IsNil(o.EndpointFqdn) {
		return true
	}

	return false
}

// SetEndpointFqdn gets a reference to the given string and assigns it to the EndpointFqdn field.
func (o *WAgfInfo) SetEndpointFqdn(v string) {
	o.EndpointFqdn = &v
}

func (o WAgfInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WAgfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ipv4EndpointAddresses) {
		toSerialize["ipv4EndpointAddresses"] = o.Ipv4EndpointAddresses
	}
	if !IsNil(o.Ipv6EndpointAddresses) {
		toSerialize["ipv6EndpointAddresses"] = o.Ipv6EndpointAddresses
	}
	if !IsNil(o.EndpointFqdn) {
		toSerialize["endpointFqdn"] = o.EndpointFqdn
	}
	return toSerialize, nil
}

type NullableWAgfInfo struct {
	value *WAgfInfo
	isSet bool
}

func (v NullableWAgfInfo) Get() *WAgfInfo {
	return v.value
}

func (v *NullableWAgfInfo) Set(val *WAgfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWAgfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWAgfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWAgfInfo(val *WAgfInfo) *NullableWAgfInfo {
	return &NullableWAgfInfo{value: val, isSet: true}
}

func (v NullableWAgfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWAgfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
