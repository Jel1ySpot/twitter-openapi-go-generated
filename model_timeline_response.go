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

// checks if the TimelineResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelineResponse{}

// TimelineResponse struct for TimelineResponse
type TimelineResponse struct {
	Data HomeTimelineResponseData `json:"data"`
}

type _TimelineResponse TimelineResponse

// NewTimelineResponse instantiates a new TimelineResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineResponse(data HomeTimelineResponseData) *TimelineResponse {
	this := TimelineResponse{}
	this.Data = data
	return &this
}

// NewTimelineResponseWithDefaults instantiates a new TimelineResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineResponseWithDefaults() *TimelineResponse {
	this := TimelineResponse{}
	return &this
}

// GetData returns the Data field value
func (o *TimelineResponse) GetData() HomeTimelineResponseData {
	if o == nil {
		var ret HomeTimelineResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TimelineResponse) GetDataOk() (*HomeTimelineResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TimelineResponse) SetData(v HomeTimelineResponseData) {
	o.Data = v
}

func (o TimelineResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimelineResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *TimelineResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varTimelineResponse := _TimelineResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTimelineResponse)

	if err != nil {
		return err
	}

	*o = TimelineResponse(varTimelineResponse)

	return err
}

type NullableTimelineResponse struct {
	value *TimelineResponse
	isSet bool
}

func (v NullableTimelineResponse) Get() *TimelineResponse {
	return v.value
}

func (v *NullableTimelineResponse) Set(val *TimelineResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineResponse(val *TimelineResponse) *NullableTimelineResponse {
	return &NullableTimelineResponse{value: val, isSet: true}
}

func (v NullableTimelineResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


