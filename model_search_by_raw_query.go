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

// checks if the SearchByRawQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchByRawQuery{}

// SearchByRawQuery struct for SearchByRawQuery
type SearchByRawQuery struct {
	SearchTimeline SearchTimeline `json:"search_timeline"`
}

type _SearchByRawQuery SearchByRawQuery

// NewSearchByRawQuery instantiates a new SearchByRawQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchByRawQuery(searchTimeline SearchTimeline) *SearchByRawQuery {
	this := SearchByRawQuery{}
	this.SearchTimeline = searchTimeline
	return &this
}

// NewSearchByRawQueryWithDefaults instantiates a new SearchByRawQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchByRawQueryWithDefaults() *SearchByRawQuery {
	this := SearchByRawQuery{}
	return &this
}

// GetSearchTimeline returns the SearchTimeline field value
func (o *SearchByRawQuery) GetSearchTimeline() SearchTimeline {
	if o == nil {
		var ret SearchTimeline
		return ret
	}

	return o.SearchTimeline
}

// GetSearchTimelineOk returns a tuple with the SearchTimeline field value
// and a boolean to check if the value has been set.
func (o *SearchByRawQuery) GetSearchTimelineOk() (*SearchTimeline, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchTimeline, true
}

// SetSearchTimeline sets field value
func (o *SearchByRawQuery) SetSearchTimeline(v SearchTimeline) {
	o.SearchTimeline = v
}

func (o SearchByRawQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchByRawQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["search_timeline"] = o.SearchTimeline
	return toSerialize, nil
}

func (o *SearchByRawQuery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"search_timeline",
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

	varSearchByRawQuery := _SearchByRawQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchByRawQuery)

	if err != nil {
		return err
	}

	*o = SearchByRawQuery(varSearchByRawQuery)

	return err
}

type NullableSearchByRawQuery struct {
	value *SearchByRawQuery
	isSet bool
}

func (v NullableSearchByRawQuery) Get() *SearchByRawQuery {
	return v.value
}

func (v *NullableSearchByRawQuery) Set(val *SearchByRawQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchByRawQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchByRawQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchByRawQuery(val *SearchByRawQuery) *NullableSearchByRawQuery {
	return &NullableSearchByRawQuery{value: val, isSet: true}
}

func (v NullableSearchByRawQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchByRawQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
