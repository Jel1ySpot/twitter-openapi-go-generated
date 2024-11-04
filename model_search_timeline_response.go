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

// checks if the SearchTimelineResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchTimelineResponse{}

// SearchTimelineResponse struct for SearchTimelineResponse
type SearchTimelineResponse struct {
	Data SearchTimelineData `json:"data"`
}

type _SearchTimelineResponse SearchTimelineResponse

// NewSearchTimelineResponse instantiates a new SearchTimelineResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchTimelineResponse(data SearchTimelineData) *SearchTimelineResponse {
	this := SearchTimelineResponse{}
	this.Data = data
	return &this
}

// NewSearchTimelineResponseWithDefaults instantiates a new SearchTimelineResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchTimelineResponseWithDefaults() *SearchTimelineResponse {
	this := SearchTimelineResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SearchTimelineResponse) GetData() SearchTimelineData {
	if o == nil {
		var ret SearchTimelineData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SearchTimelineResponse) GetDataOk() (*SearchTimelineData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SearchTimelineResponse) SetData(v SearchTimelineData) {
	o.Data = v
}

func (o SearchTimelineResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchTimelineResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SearchTimelineResponse) UnmarshalJSON(data []byte) (err error) {
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

	varSearchTimelineResponse := _SearchTimelineResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchTimelineResponse)

	if err != nil {
		return err
	}

	*o = SearchTimelineResponse(varSearchTimelineResponse)

	return err
}

type NullableSearchTimelineResponse struct {
	value *SearchTimelineResponse
	isSet bool
}

func (v NullableSearchTimelineResponse) Get() *SearchTimelineResponse {
	return v.value
}

func (v *NullableSearchTimelineResponse) Set(val *SearchTimelineResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchTimelineResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchTimelineResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchTimelineResponse(val *SearchTimelineResponse) *NullableSearchTimelineResponse {
	return &NullableSearchTimelineResponse{value: val, isSet: true}
}

func (v NullableSearchTimelineResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchTimelineResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


