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

// checks if the TimelineTimelineCursor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelineTimelineCursor{}

// TimelineTimelineCursor struct for TimelineTimelineCursor
type TimelineTimelineCursor struct {
	Typename            TypeName          `json:"__typename"`
	CursorType          CursorType        `json:"cursorType"`
	DisplayTreatment    *DisplayTreatment `json:"displayTreatment,omitempty"`
	EntryType           *ContentEntryType `json:"entryType,omitempty"`
	ItemType            *ContentEntryType `json:"itemType,omitempty"`
	StopOnEmptyResponse *bool             `json:"stopOnEmptyResponse,omitempty"`
	Value               string            `json:"value"`
}

type _TimelineTimelineCursor TimelineTimelineCursor

// NewTimelineTimelineCursor instantiates a new TimelineTimelineCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineTimelineCursor(typename TypeName, cursorType CursorType, value string) *TimelineTimelineCursor {
	this := TimelineTimelineCursor{}
	this.Typename = typename
	this.CursorType = cursorType
	this.Value = value
	return &this
}

// NewTimelineTimelineCursorWithDefaults instantiates a new TimelineTimelineCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineTimelineCursorWithDefaults() *TimelineTimelineCursor {
	this := TimelineTimelineCursor{}
	return &this
}

// GetTypename returns the Typename field value
func (o *TimelineTimelineCursor) GetTypename() TypeName {
	if o == nil {
		var ret TypeName
		return ret
	}

	return o.Typename
}

// GetTypenameOk returns a tuple with the Typename field value
// and a boolean to check if the value has been set.
func (o *TimelineTimelineCursor) GetTypenameOk() (*TypeName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Typename, true
}

// SetTypename sets field value
func (o *TimelineTimelineCursor) SetTypename(v TypeName) {
	o.Typename = v
}

// GetCursorType returns the CursorType field value
func (o *TimelineTimelineCursor) GetCursorType() CursorType {
	if o == nil {
		var ret CursorType
		return ret
	}

	return o.CursorType
}

// GetCursorTypeOk returns a tuple with the CursorType field value
// and a boolean to check if the value has been set.
func (o *TimelineTimelineCursor) GetCursorTypeOk() (*CursorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CursorType, true
}

// SetCursorType sets field value
func (o *TimelineTimelineCursor) SetCursorType(v CursorType) {
	o.CursorType = v
}

// GetDisplayTreatment returns the DisplayTreatment field value if set, zero value otherwise.
func (o *TimelineTimelineCursor) GetDisplayTreatment() DisplayTreatment {
	if o == nil || IsNil(o.DisplayTreatment) {
		var ret DisplayTreatment
		return ret
	}
	return *o.DisplayTreatment
}

// GetDisplayTreatmentOk returns a tuple with the DisplayTreatment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineTimelineCursor) GetDisplayTreatmentOk() (*DisplayTreatment, bool) {
	if o == nil || IsNil(o.DisplayTreatment) {
		return nil, false
	}
	return o.DisplayTreatment, true
}

// HasDisplayTreatment returns a boolean if a field has been set.
func (o *TimelineTimelineCursor) HasDisplayTreatment() bool {
	if o != nil && !IsNil(o.DisplayTreatment) {
		return true
	}

	return false
}

// SetDisplayTreatment gets a reference to the given DisplayTreatment and assigns it to the DisplayTreatment field.
func (o *TimelineTimelineCursor) SetDisplayTreatment(v DisplayTreatment) {
	o.DisplayTreatment = &v
}

// GetEntryType returns the EntryType field value if set, zero value otherwise.
func (o *TimelineTimelineCursor) GetEntryType() ContentEntryType {
	if o == nil || IsNil(o.EntryType) {
		var ret ContentEntryType
		return ret
	}
	return *o.EntryType
}

// GetEntryTypeOk returns a tuple with the EntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineTimelineCursor) GetEntryTypeOk() (*ContentEntryType, bool) {
	if o == nil || IsNil(o.EntryType) {
		return nil, false
	}
	return o.EntryType, true
}

// HasEntryType returns a boolean if a field has been set.
func (o *TimelineTimelineCursor) HasEntryType() bool {
	if o != nil && !IsNil(o.EntryType) {
		return true
	}

	return false
}

// SetEntryType gets a reference to the given ContentEntryType and assigns it to the EntryType field.
func (o *TimelineTimelineCursor) SetEntryType(v ContentEntryType) {
	o.EntryType = &v
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *TimelineTimelineCursor) GetItemType() ContentEntryType {
	if o == nil || IsNil(o.ItemType) {
		var ret ContentEntryType
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineTimelineCursor) GetItemTypeOk() (*ContentEntryType, bool) {
	if o == nil || IsNil(o.ItemType) {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *TimelineTimelineCursor) HasItemType() bool {
	if o != nil && !IsNil(o.ItemType) {
		return true
	}

	return false
}

// SetItemType gets a reference to the given ContentEntryType and assigns it to the ItemType field.
func (o *TimelineTimelineCursor) SetItemType(v ContentEntryType) {
	o.ItemType = &v
}

// GetStopOnEmptyResponse returns the StopOnEmptyResponse field value if set, zero value otherwise.
func (o *TimelineTimelineCursor) GetStopOnEmptyResponse() bool {
	if o == nil || IsNil(o.StopOnEmptyResponse) {
		var ret bool
		return ret
	}
	return *o.StopOnEmptyResponse
}

// GetStopOnEmptyResponseOk returns a tuple with the StopOnEmptyResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineTimelineCursor) GetStopOnEmptyResponseOk() (*bool, bool) {
	if o == nil || IsNil(o.StopOnEmptyResponse) {
		return nil, false
	}
	return o.StopOnEmptyResponse, true
}

// HasStopOnEmptyResponse returns a boolean if a field has been set.
func (o *TimelineTimelineCursor) HasStopOnEmptyResponse() bool {
	if o != nil && !IsNil(o.StopOnEmptyResponse) {
		return true
	}

	return false
}

// SetStopOnEmptyResponse gets a reference to the given bool and assigns it to the StopOnEmptyResponse field.
func (o *TimelineTimelineCursor) SetStopOnEmptyResponse(v bool) {
	o.StopOnEmptyResponse = &v
}

// GetValue returns the Value field value
func (o *TimelineTimelineCursor) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TimelineTimelineCursor) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TimelineTimelineCursor) SetValue(v string) {
	o.Value = v
}

func (o TimelineTimelineCursor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimelineTimelineCursor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["__typename"] = o.Typename
	toSerialize["cursorType"] = o.CursorType
	if !IsNil(o.DisplayTreatment) {
		toSerialize["displayTreatment"] = o.DisplayTreatment
	}
	if !IsNil(o.EntryType) {
		toSerialize["entryType"] = o.EntryType
	}
	if !IsNil(o.ItemType) {
		toSerialize["itemType"] = o.ItemType
	}
	if !IsNil(o.StopOnEmptyResponse) {
		toSerialize["stopOnEmptyResponse"] = o.StopOnEmptyResponse
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *TimelineTimelineCursor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"__typename",
		"cursorType",
		"value",
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

	varTimelineTimelineCursor := _TimelineTimelineCursor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTimelineTimelineCursor)

	if err != nil {
		return err
	}

	*o = TimelineTimelineCursor(varTimelineTimelineCursor)

	return err
}

type NullableTimelineTimelineCursor struct {
	value *TimelineTimelineCursor
	isSet bool
}

func (v NullableTimelineTimelineCursor) Get() *TimelineTimelineCursor {
	return v.value
}

func (v *NullableTimelineTimelineCursor) Set(val *TimelineTimelineCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineTimelineCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineTimelineCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineTimelineCursor(val *TimelineTimelineCursor) *NullableTimelineTimelineCursor {
	return &NullableTimelineTimelineCursor{value: val, isSet: true}
}

func (v NullableTimelineTimelineCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineTimelineCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
