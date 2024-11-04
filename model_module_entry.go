/*
Twitter OpenAPI

Twitter OpenAPI(Swagger) specification

API version: 0.0.1
Contact: yuki@yuki0311.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ModuleEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleEntry{}

// ModuleEntry struct for ModuleEntry
type ModuleEntry struct {
	ClientEventInfo *ClientEventInfo `json:"clientEventInfo,omitempty"`
	FeedbackInfo    *FeedbackInfo    `json:"feedbackInfo,omitempty"`
	ItemContent     ItemContentUnion `json:"itemContent"`
}

type _ModuleEntry ModuleEntry

// NewModuleEntry instantiates a new ModuleEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleEntry(itemContent ItemContentUnion) *ModuleEntry {
	this := ModuleEntry{}
	this.ItemContent = itemContent
	return &this
}

// NewModuleEntryWithDefaults instantiates a new ModuleEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleEntryWithDefaults() *ModuleEntry {
	this := ModuleEntry{}
	return &this
}

// GetClientEventInfo returns the ClientEventInfo field value if set, zero value otherwise.
func (o *ModuleEntry) GetClientEventInfo() ClientEventInfo {
	if o == nil || IsNil(o.ClientEventInfo) {
		var ret ClientEventInfo
		return ret
	}
	return *o.ClientEventInfo
}

// GetClientEventInfoOk returns a tuple with the ClientEventInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleEntry) GetClientEventInfoOk() (*ClientEventInfo, bool) {
	if o == nil || IsNil(o.ClientEventInfo) {
		return nil, false
	}
	return o.ClientEventInfo, true
}

// HasClientEventInfo returns a boolean if a field has been set.
func (o *ModuleEntry) HasClientEventInfo() bool {
	if o != nil && !IsNil(o.ClientEventInfo) {
		return true
	}

	return false
}

// SetClientEventInfo gets a reference to the given ClientEventInfo and assigns it to the ClientEventInfo field.
func (o *ModuleEntry) SetClientEventInfo(v ClientEventInfo) {
	o.ClientEventInfo = &v
}

// GetFeedbackInfo returns the FeedbackInfo field value if set, zero value otherwise.
func (o *ModuleEntry) GetFeedbackInfo() FeedbackInfo {
	if o == nil || IsNil(o.FeedbackInfo) {
		var ret FeedbackInfo
		return ret
	}
	return *o.FeedbackInfo
}

// GetFeedbackInfoOk returns a tuple with the FeedbackInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleEntry) GetFeedbackInfoOk() (*FeedbackInfo, bool) {
	if o == nil || IsNil(o.FeedbackInfo) {
		return nil, false
	}
	return o.FeedbackInfo, true
}

// HasFeedbackInfo returns a boolean if a field has been set.
func (o *ModuleEntry) HasFeedbackInfo() bool {
	if o != nil && !IsNil(o.FeedbackInfo) {
		return true
	}

	return false
}

// SetFeedbackInfo gets a reference to the given FeedbackInfo and assigns it to the FeedbackInfo field.
func (o *ModuleEntry) SetFeedbackInfo(v FeedbackInfo) {
	o.FeedbackInfo = &v
}

// GetItemContent returns the ItemContent field value
func (o *ModuleEntry) GetItemContent() ItemContentUnion {
	if o == nil {
		var ret ItemContentUnion
		return ret
	}

	return o.ItemContent
}

// GetItemContentOk returns a tuple with the ItemContent field value
// and a boolean to check if the value has been set.
func (o *ModuleEntry) GetItemContentOk() (*ItemContentUnion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemContent, true
}

// SetItemContent sets field value
func (o *ModuleEntry) SetItemContent(v ItemContentUnion) {
	o.ItemContent = v
}

func (o ModuleEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientEventInfo) {
		toSerialize["clientEventInfo"] = o.ClientEventInfo
	}
	if !IsNil(o.FeedbackInfo) {
		toSerialize["feedbackInfo"] = o.FeedbackInfo
	}
	toSerialize["itemContent"] = o.ItemContent
	return toSerialize, nil
}

func (o *ModuleEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"itemContent",
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

	varModuleEntry := _ModuleEntry{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModuleEntry)

	if err != nil {
		return err
	}

	*o = ModuleEntry(varModuleEntry)

	return err
}

type NullableModuleEntry struct {
	value *ModuleEntry
	isSet bool
}

func (v NullableModuleEntry) Get() *ModuleEntry {
	return v.value
}

func (v *NullableModuleEntry) Set(val *ModuleEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleEntry(val *ModuleEntry) *NullableModuleEntry {
	return &NullableModuleEntry{value: val, isSet: true}
}

func (v NullableModuleEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
