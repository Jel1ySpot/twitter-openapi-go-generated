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

// checks if the TimelineHalfCover type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelineHalfCover{}

// TimelineHalfCover struct for TimelineHalfCover
type TimelineHalfCover struct {
	Dismissible bool `json:"dismissible"`
	HalfCoverDisplayType string `json:"halfCoverDisplayType"`
	ImpressionCallbacks []Callback `json:"impressionCallbacks"`
	PrimaryCoverCta CoverCta `json:"primaryCoverCta"`
	PrimaryText Text `json:"primaryText"`
	SecondaryText Text `json:"secondaryText"`
	Type string `json:"type"`
}

type _TimelineHalfCover TimelineHalfCover

// NewTimelineHalfCover instantiates a new TimelineHalfCover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineHalfCover(dismissible bool, halfCoverDisplayType string, impressionCallbacks []Callback, primaryCoverCta CoverCta, primaryText Text, secondaryText Text, type_ string) *TimelineHalfCover {
	this := TimelineHalfCover{}
	this.Dismissible = dismissible
	this.HalfCoverDisplayType = halfCoverDisplayType
	this.ImpressionCallbacks = impressionCallbacks
	this.PrimaryCoverCta = primaryCoverCta
	this.PrimaryText = primaryText
	this.SecondaryText = secondaryText
	this.Type = type_
	return &this
}

// NewTimelineHalfCoverWithDefaults instantiates a new TimelineHalfCover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineHalfCoverWithDefaults() *TimelineHalfCover {
	this := TimelineHalfCover{}
	return &this
}

// GetDismissible returns the Dismissible field value
func (o *TimelineHalfCover) GetDismissible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Dismissible
}

// GetDismissibleOk returns a tuple with the Dismissible field value
// and a boolean to check if the value has been set.
func (o *TimelineHalfCover) GetDismissibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dismissible, true
}

// SetDismissible sets field value
func (o *TimelineHalfCover) SetDismissible(v bool) {
	o.Dismissible = v
}

// GetHalfCoverDisplayType returns the HalfCoverDisplayType field value
func (o *TimelineHalfCover) GetHalfCoverDisplayType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HalfCoverDisplayType
}

// GetHalfCoverDisplayTypeOk returns a tuple with the HalfCoverDisplayType field value
// and a boolean to check if the value has been set.
func (o *TimelineHalfCover) GetHalfCoverDisplayTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HalfCoverDisplayType, true
}

// SetHalfCoverDisplayType sets field value
func (o *TimelineHalfCover) SetHalfCoverDisplayType(v string) {
	o.HalfCoverDisplayType = v
}

// GetImpressionCallbacks returns the ImpressionCallbacks field value
func (o *TimelineHalfCover) GetImpressionCallbacks() []Callback {
	if o == nil {
		var ret []Callback
		return ret
	}

	return o.ImpressionCallbacks
}

// GetImpressionCallbacksOk returns a tuple with the ImpressionCallbacks field value
// and a boolean to check if the value has been set.
func (o *TimelineHalfCover) GetImpressionCallbacksOk() ([]Callback, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImpressionCallbacks, true
}

// SetImpressionCallbacks sets field value
func (o *TimelineHalfCover) SetImpressionCallbacks(v []Callback) {
	o.ImpressionCallbacks = v
}

// GetPrimaryCoverCta returns the PrimaryCoverCta field value
func (o *TimelineHalfCover) GetPrimaryCoverCta() CoverCta {
	if o == nil {
		var ret CoverCta
		return ret
	}

	return o.PrimaryCoverCta
}

// GetPrimaryCoverCtaOk returns a tuple with the PrimaryCoverCta field value
// and a boolean to check if the value has been set.
func (o *TimelineHalfCover) GetPrimaryCoverCtaOk() (*CoverCta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryCoverCta, true
}

// SetPrimaryCoverCta sets field value
func (o *TimelineHalfCover) SetPrimaryCoverCta(v CoverCta) {
	o.PrimaryCoverCta = v
}

// GetPrimaryText returns the PrimaryText field value
func (o *TimelineHalfCover) GetPrimaryText() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.PrimaryText
}

// GetPrimaryTextOk returns a tuple with the PrimaryText field value
// and a boolean to check if the value has been set.
func (o *TimelineHalfCover) GetPrimaryTextOk() (*Text, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryText, true
}

// SetPrimaryText sets field value
func (o *TimelineHalfCover) SetPrimaryText(v Text) {
	o.PrimaryText = v
}

// GetSecondaryText returns the SecondaryText field value
func (o *TimelineHalfCover) GetSecondaryText() Text {
	if o == nil {
		var ret Text
		return ret
	}

	return o.SecondaryText
}

// GetSecondaryTextOk returns a tuple with the SecondaryText field value
// and a boolean to check if the value has been set.
func (o *TimelineHalfCover) GetSecondaryTextOk() (*Text, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecondaryText, true
}

// SetSecondaryText sets field value
func (o *TimelineHalfCover) SetSecondaryText(v Text) {
	o.SecondaryText = v
}

// GetType returns the Type field value
func (o *TimelineHalfCover) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TimelineHalfCover) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TimelineHalfCover) SetType(v string) {
	o.Type = v
}

func (o TimelineHalfCover) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimelineHalfCover) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dismissible"] = o.Dismissible
	toSerialize["halfCoverDisplayType"] = o.HalfCoverDisplayType
	toSerialize["impressionCallbacks"] = o.ImpressionCallbacks
	toSerialize["primaryCoverCta"] = o.PrimaryCoverCta
	toSerialize["primaryText"] = o.PrimaryText
	toSerialize["secondaryText"] = o.SecondaryText
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *TimelineHalfCover) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dismissible",
		"halfCoverDisplayType",
		"impressionCallbacks",
		"primaryCoverCta",
		"primaryText",
		"secondaryText",
		"type",
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

	varTimelineHalfCover := _TimelineHalfCover{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTimelineHalfCover)

	if err != nil {
		return err
	}

	*o = TimelineHalfCover(varTimelineHalfCover)

	return err
}

type NullableTimelineHalfCover struct {
	value *TimelineHalfCover
	isSet bool
}

func (v NullableTimelineHalfCover) Get() *TimelineHalfCover {
	return v.value
}

func (v *NullableTimelineHalfCover) Set(val *TimelineHalfCover) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineHalfCover) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineHalfCover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineHalfCover(val *TimelineHalfCover) *NullableTimelineHalfCover {
	return &NullableTimelineHalfCover{value: val, isSet: true}
}

func (v NullableTimelineHalfCover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineHalfCover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


