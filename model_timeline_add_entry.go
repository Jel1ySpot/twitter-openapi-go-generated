/*
Twitter OpenAPI

Twitter OpenAPI(Swagger) specification

API version: 0.0.1
Contact: yuki@yuki0311.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TimelineAddEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelineAddEntry{}

// TimelineAddEntry struct for TimelineAddEntry
type TimelineAddEntry struct {
	Content ContentUnion `json:"content"`
	EntryId string `json:"entryId" validate:"regexp=^(([a-z]+|[0-9]+|[0-9a-f]+)(-|$))+"`
	SortIndex string `json:"sortIndex" validate:"regexp=[0-9]+$"`
}

type _TimelineAddEntry TimelineAddEntry

// NewTimelineAddEntry instantiates a new TimelineAddEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineAddEntry(content ContentUnion, entryId string, sortIndex string) *TimelineAddEntry {
	this := TimelineAddEntry{}
	this.Content = content
	this.EntryId = entryId
	this.SortIndex = sortIndex
	return &this
}

// NewTimelineAddEntryWithDefaults instantiates a new TimelineAddEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineAddEntryWithDefaults() *TimelineAddEntry {
	this := TimelineAddEntry{}
	return &this
}

// GetContent returns the Content field value
func (o *TimelineAddEntry) GetContent() ContentUnion {
	if o == nil {
		var ret ContentUnion
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *TimelineAddEntry) GetContentOk() (*ContentUnion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *TimelineAddEntry) SetContent(v ContentUnion) {
	o.Content = v
}

// GetEntryId returns the EntryId field value
func (o *TimelineAddEntry) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *TimelineAddEntry) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *TimelineAddEntry) SetEntryId(v string) {
	o.EntryId = v
}

// GetSortIndex returns the SortIndex field value
func (o *TimelineAddEntry) GetSortIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SortIndex
}

// GetSortIndexOk returns a tuple with the SortIndex field value
// and a boolean to check if the value has been set.
func (o *TimelineAddEntry) GetSortIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortIndex, true
}

// SetSortIndex sets field value
func (o *TimelineAddEntry) SetSortIndex(v string) {
	o.SortIndex = v
}

func (o TimelineAddEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimelineAddEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content
	toSerialize["entryId"] = o.EntryId
	toSerialize["sortIndex"] = o.SortIndex
	return toSerialize, nil
}

func (o *TimelineAddEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
		"entryId",
		"sortIndex",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTimelineAddEntry := _TimelineAddEntry{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTimelineAddEntry)

	if err != nil {
		return err
	}

	*o = TimelineAddEntry(varTimelineAddEntry)

	return err
}

type NullableTimelineAddEntry struct {
	value *TimelineAddEntry
	isSet bool
}

func (v NullableTimelineAddEntry) Get() *TimelineAddEntry {
	return v.value
}

func (v *NullableTimelineAddEntry) Set(val *TimelineAddEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineAddEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineAddEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineAddEntry(val *TimelineAddEntry) *NullableTimelineAddEntry {
	return &NullableTimelineAddEntry{value: val, isSet: true}
}

func (v NullableTimelineAddEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineAddEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

