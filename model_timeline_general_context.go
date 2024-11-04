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
)

// checks if the TimelineGeneralContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelineGeneralContext{}

// TimelineGeneralContext struct for TimelineGeneralContext
type TimelineGeneralContext struct {
	ContextType *string                  `json:"contextType,omitempty"`
	LandingUrl  *SocialContextLandingUrl `json:"landingUrl,omitempty"`
	Text        *string                  `json:"text,omitempty"`
	Type        *SocialContextUnionType  `json:"type,omitempty"`
}

// NewTimelineGeneralContext instantiates a new TimelineGeneralContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineGeneralContext() *TimelineGeneralContext {
	this := TimelineGeneralContext{}
	return &this
}

// NewTimelineGeneralContextWithDefaults instantiates a new TimelineGeneralContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineGeneralContextWithDefaults() *TimelineGeneralContext {
	this := TimelineGeneralContext{}
	return &this
}

// GetContextType returns the ContextType field value if set, zero value otherwise.
func (o *TimelineGeneralContext) GetContextType() string {
	if o == nil || IsNil(o.ContextType) {
		var ret string
		return ret
	}
	return *o.ContextType
}

// GetContextTypeOk returns a tuple with the ContextType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineGeneralContext) GetContextTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContextType) {
		return nil, false
	}
	return o.ContextType, true
}

// HasContextType returns a boolean if a field has been set.
func (o *TimelineGeneralContext) HasContextType() bool {
	if o != nil && !IsNil(o.ContextType) {
		return true
	}

	return false
}

// SetContextType gets a reference to the given string and assigns it to the ContextType field.
func (o *TimelineGeneralContext) SetContextType(v string) {
	o.ContextType = &v
}

// GetLandingUrl returns the LandingUrl field value if set, zero value otherwise.
func (o *TimelineGeneralContext) GetLandingUrl() SocialContextLandingUrl {
	if o == nil || IsNil(o.LandingUrl) {
		var ret SocialContextLandingUrl
		return ret
	}
	return *o.LandingUrl
}

// GetLandingUrlOk returns a tuple with the LandingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineGeneralContext) GetLandingUrlOk() (*SocialContextLandingUrl, bool) {
	if o == nil || IsNil(o.LandingUrl) {
		return nil, false
	}
	return o.LandingUrl, true
}

// HasLandingUrl returns a boolean if a field has been set.
func (o *TimelineGeneralContext) HasLandingUrl() bool {
	if o != nil && !IsNil(o.LandingUrl) {
		return true
	}

	return false
}

// SetLandingUrl gets a reference to the given SocialContextLandingUrl and assigns it to the LandingUrl field.
func (o *TimelineGeneralContext) SetLandingUrl(v SocialContextLandingUrl) {
	o.LandingUrl = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *TimelineGeneralContext) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineGeneralContext) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *TimelineGeneralContext) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *TimelineGeneralContext) SetText(v string) {
	o.Text = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TimelineGeneralContext) GetType() SocialContextUnionType {
	if o == nil || IsNil(o.Type) {
		var ret SocialContextUnionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineGeneralContext) GetTypeOk() (*SocialContextUnionType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TimelineGeneralContext) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SocialContextUnionType and assigns it to the Type field.
func (o *TimelineGeneralContext) SetType(v SocialContextUnionType) {
	o.Type = &v
}

func (o TimelineGeneralContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimelineGeneralContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContextType) {
		toSerialize["contextType"] = o.ContextType
	}
	if !IsNil(o.LandingUrl) {
		toSerialize["landingUrl"] = o.LandingUrl
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTimelineGeneralContext struct {
	value *TimelineGeneralContext
	isSet bool
}

func (v NullableTimelineGeneralContext) Get() *TimelineGeneralContext {
	return v.value
}

func (v *NullableTimelineGeneralContext) Set(val *TimelineGeneralContext) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineGeneralContext) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineGeneralContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineGeneralContext(val *TimelineGeneralContext) *NullableTimelineGeneralContext {
	return &NullableTimelineGeneralContext{value: val, isSet: true}
}

func (v NullableTimelineGeneralContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineGeneralContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
