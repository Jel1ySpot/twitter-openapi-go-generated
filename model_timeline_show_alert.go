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

// checks if the TimelineShowAlert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelineShowAlert{}

// TimelineShowAlert struct for TimelineShowAlert
type TimelineShowAlert struct {
	AlertType         *string                   `json:"alertType,omitempty"`
	ColorConfig       map[string]interface{}    `json:"colorConfig,omitempty"`
	DisplayDurationMs *int32                    `json:"displayDurationMs,omitempty"`
	DisplayLocation   *string                   `json:"displayLocation,omitempty"`
	IconDisplayInfo   map[string]interface{}    `json:"iconDisplayInfo,omitempty"`
	RichText          TimelineShowAlertRichText `json:"richText"`
	TriggerDelayMs    *int32                    `json:"triggerDelayMs,omitempty"`
	Type              InstructionType           `json:"type"`
	UsersResults      []UserResults             `json:"usersResults"`
}

type _TimelineShowAlert TimelineShowAlert

// NewTimelineShowAlert instantiates a new TimelineShowAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineShowAlert(richText TimelineShowAlertRichText, type_ InstructionType, usersResults []UserResults) *TimelineShowAlert {
	this := TimelineShowAlert{}
	this.RichText = richText
	this.Type = type_
	this.UsersResults = usersResults
	return &this
}

// NewTimelineShowAlertWithDefaults instantiates a new TimelineShowAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineShowAlertWithDefaults() *TimelineShowAlert {
	this := TimelineShowAlert{}
	return &this
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *TimelineShowAlert) GetAlertType() string {
	if o == nil || IsNil(o.AlertType) {
		var ret string
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineShowAlert) GetAlertTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AlertType) {
		return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *TimelineShowAlert) HasAlertType() bool {
	if o != nil && !IsNil(o.AlertType) {
		return true
	}

	return false
}

// SetAlertType gets a reference to the given string and assigns it to the AlertType field.
func (o *TimelineShowAlert) SetAlertType(v string) {
	o.AlertType = &v
}

// GetColorConfig returns the ColorConfig field value if set, zero value otherwise.
func (o *TimelineShowAlert) GetColorConfig() map[string]interface{} {
	if o == nil || IsNil(o.ColorConfig) {
		var ret map[string]interface{}
		return ret
	}
	return o.ColorConfig
}

// GetColorConfigOk returns a tuple with the ColorConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineShowAlert) GetColorConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ColorConfig) {
		return map[string]interface{}{}, false
	}
	return o.ColorConfig, true
}

// HasColorConfig returns a boolean if a field has been set.
func (o *TimelineShowAlert) HasColorConfig() bool {
	if o != nil && !IsNil(o.ColorConfig) {
		return true
	}

	return false
}

// SetColorConfig gets a reference to the given map[string]interface{} and assigns it to the ColorConfig field.
func (o *TimelineShowAlert) SetColorConfig(v map[string]interface{}) {
	o.ColorConfig = v
}

// GetDisplayDurationMs returns the DisplayDurationMs field value if set, zero value otherwise.
func (o *TimelineShowAlert) GetDisplayDurationMs() int32 {
	if o == nil || IsNil(o.DisplayDurationMs) {
		var ret int32
		return ret
	}
	return *o.DisplayDurationMs
}

// GetDisplayDurationMsOk returns a tuple with the DisplayDurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineShowAlert) GetDisplayDurationMsOk() (*int32, bool) {
	if o == nil || IsNil(o.DisplayDurationMs) {
		return nil, false
	}
	return o.DisplayDurationMs, true
}

// HasDisplayDurationMs returns a boolean if a field has been set.
func (o *TimelineShowAlert) HasDisplayDurationMs() bool {
	if o != nil && !IsNil(o.DisplayDurationMs) {
		return true
	}

	return false
}

// SetDisplayDurationMs gets a reference to the given int32 and assigns it to the DisplayDurationMs field.
func (o *TimelineShowAlert) SetDisplayDurationMs(v int32) {
	o.DisplayDurationMs = &v
}

// GetDisplayLocation returns the DisplayLocation field value if set, zero value otherwise.
func (o *TimelineShowAlert) GetDisplayLocation() string {
	if o == nil || IsNil(o.DisplayLocation) {
		var ret string
		return ret
	}
	return *o.DisplayLocation
}

// GetDisplayLocationOk returns a tuple with the DisplayLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineShowAlert) GetDisplayLocationOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayLocation) {
		return nil, false
	}
	return o.DisplayLocation, true
}

// HasDisplayLocation returns a boolean if a field has been set.
func (o *TimelineShowAlert) HasDisplayLocation() bool {
	if o != nil && !IsNil(o.DisplayLocation) {
		return true
	}

	return false
}

// SetDisplayLocation gets a reference to the given string and assigns it to the DisplayLocation field.
func (o *TimelineShowAlert) SetDisplayLocation(v string) {
	o.DisplayLocation = &v
}

// GetIconDisplayInfo returns the IconDisplayInfo field value if set, zero value otherwise.
func (o *TimelineShowAlert) GetIconDisplayInfo() map[string]interface{} {
	if o == nil || IsNil(o.IconDisplayInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.IconDisplayInfo
}

// GetIconDisplayInfoOk returns a tuple with the IconDisplayInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineShowAlert) GetIconDisplayInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.IconDisplayInfo) {
		return map[string]interface{}{}, false
	}
	return o.IconDisplayInfo, true
}

// HasIconDisplayInfo returns a boolean if a field has been set.
func (o *TimelineShowAlert) HasIconDisplayInfo() bool {
	if o != nil && !IsNil(o.IconDisplayInfo) {
		return true
	}

	return false
}

// SetIconDisplayInfo gets a reference to the given map[string]interface{} and assigns it to the IconDisplayInfo field.
func (o *TimelineShowAlert) SetIconDisplayInfo(v map[string]interface{}) {
	o.IconDisplayInfo = v
}

// GetRichText returns the RichText field value
func (o *TimelineShowAlert) GetRichText() TimelineShowAlertRichText {
	if o == nil {
		var ret TimelineShowAlertRichText
		return ret
	}

	return o.RichText
}

// GetRichTextOk returns a tuple with the RichText field value
// and a boolean to check if the value has been set.
func (o *TimelineShowAlert) GetRichTextOk() (*TimelineShowAlertRichText, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RichText, true
}

// SetRichText sets field value
func (o *TimelineShowAlert) SetRichText(v TimelineShowAlertRichText) {
	o.RichText = v
}

// GetTriggerDelayMs returns the TriggerDelayMs field value if set, zero value otherwise.
func (o *TimelineShowAlert) GetTriggerDelayMs() int32 {
	if o == nil || IsNil(o.TriggerDelayMs) {
		var ret int32
		return ret
	}
	return *o.TriggerDelayMs
}

// GetTriggerDelayMsOk returns a tuple with the TriggerDelayMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineShowAlert) GetTriggerDelayMsOk() (*int32, bool) {
	if o == nil || IsNil(o.TriggerDelayMs) {
		return nil, false
	}
	return o.TriggerDelayMs, true
}

// HasTriggerDelayMs returns a boolean if a field has been set.
func (o *TimelineShowAlert) HasTriggerDelayMs() bool {
	if o != nil && !IsNil(o.TriggerDelayMs) {
		return true
	}

	return false
}

// SetTriggerDelayMs gets a reference to the given int32 and assigns it to the TriggerDelayMs field.
func (o *TimelineShowAlert) SetTriggerDelayMs(v int32) {
	o.TriggerDelayMs = &v
}

// GetType returns the Type field value
func (o *TimelineShowAlert) GetType() InstructionType {
	if o == nil {
		var ret InstructionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TimelineShowAlert) GetTypeOk() (*InstructionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TimelineShowAlert) SetType(v InstructionType) {
	o.Type = v
}

// GetUsersResults returns the UsersResults field value
func (o *TimelineShowAlert) GetUsersResults() []UserResults {
	if o == nil {
		var ret []UserResults
		return ret
	}

	return o.UsersResults
}

// GetUsersResultsOk returns a tuple with the UsersResults field value
// and a boolean to check if the value has been set.
func (o *TimelineShowAlert) GetUsersResultsOk() ([]UserResults, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsersResults, true
}

// SetUsersResults sets field value
func (o *TimelineShowAlert) SetUsersResults(v []UserResults) {
	o.UsersResults = v
}

func (o TimelineShowAlert) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimelineShowAlert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlertType) {
		toSerialize["alertType"] = o.AlertType
	}
	if !IsNil(o.ColorConfig) {
		toSerialize["colorConfig"] = o.ColorConfig
	}
	if !IsNil(o.DisplayDurationMs) {
		toSerialize["displayDurationMs"] = o.DisplayDurationMs
	}
	if !IsNil(o.DisplayLocation) {
		toSerialize["displayLocation"] = o.DisplayLocation
	}
	if !IsNil(o.IconDisplayInfo) {
		toSerialize["iconDisplayInfo"] = o.IconDisplayInfo
	}
	toSerialize["richText"] = o.RichText
	if !IsNil(o.TriggerDelayMs) {
		toSerialize["triggerDelayMs"] = o.TriggerDelayMs
	}
	toSerialize["type"] = o.Type
	toSerialize["usersResults"] = o.UsersResults
	return toSerialize, nil
}

func (o *TimelineShowAlert) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"richText",
		"type",
		"usersResults",
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

	varTimelineShowAlert := _TimelineShowAlert{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTimelineShowAlert)

	if err != nil {
		return err
	}

	*o = TimelineShowAlert(varTimelineShowAlert)

	return err
}

type NullableTimelineShowAlert struct {
	value *TimelineShowAlert
	isSet bool
}

func (v NullableTimelineShowAlert) Get() *TimelineShowAlert {
	return v.value
}

func (v *NullableTimelineShowAlert) Set(val *TimelineShowAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineShowAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineShowAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineShowAlert(val *TimelineShowAlert) *NullableTimelineShowAlert {
	return &NullableTimelineShowAlert{value: val, isSet: true}
}

func (v NullableTimelineShowAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineShowAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
