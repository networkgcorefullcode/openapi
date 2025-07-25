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

// DnnInfoItem Set of parameters supported by NF for a given DNN
type DnnInfoItem struct {
	Dnn DnnSmfInfoItemDnn `json:"dnn"`
}

type _DnnInfoItem DnnInfoItem

// NewDnnInfoItem instantiates a new DnnInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnnInfoItem(dnn DnnSmfInfoItemDnn) *DnnInfoItem {
	this := DnnInfoItem{}
	this.Dnn = dnn
	return &this
}

// NewDnnInfoItemWithDefaults instantiates a new DnnInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnnInfoItemWithDefaults() *DnnInfoItem {
	this := DnnInfoItem{}
	return &this
}

// GetDnn returns the Dnn field value
func (o *DnnInfoItem) GetDnn() DnnSmfInfoItemDnn {
	if o == nil {
		var ret DnnSmfInfoItemDnn
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *DnnInfoItem) GetDnnOk() (*DnnSmfInfoItemDnn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *DnnInfoItem) SetDnn(v DnnSmfInfoItemDnn) {
	o.Dnn = v
}

func (o DnnInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnnInfoItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dnn"] = o.Dnn
	return toSerialize, nil
}

func (o *DnnInfoItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dnn",
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

	varDnnInfoItem := _DnnInfoItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDnnInfoItem)

	if err != nil {
		return err
	}

	*o = DnnInfoItem(varDnnInfoItem)

	return err
}

type NullableDnnInfoItem struct {
	value *DnnInfoItem
	isSet bool
}

func (v NullableDnnInfoItem) Get() *DnnInfoItem {
	return v.value
}

func (v *NullableDnnInfoItem) Set(val *DnnInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDnnInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDnnInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnnInfoItem(val *DnnInfoItem) *NullableDnnInfoItem {
	return &NullableDnnInfoItem{value: val, isSet: true}
}

func (v NullableDnnInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnnInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
