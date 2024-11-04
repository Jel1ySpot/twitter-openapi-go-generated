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

// checks if the TweetFavoritersResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetFavoritersResponseData{}

// TweetFavoritersResponseData struct for TweetFavoritersResponseData
type TweetFavoritersResponseData struct {
	FavoritersTimeline TimelineV2 `json:"favoriters_timeline"`
}

type _TweetFavoritersResponseData TweetFavoritersResponseData

// NewTweetFavoritersResponseData instantiates a new TweetFavoritersResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetFavoritersResponseData(favoritersTimeline TimelineV2) *TweetFavoritersResponseData {
	this := TweetFavoritersResponseData{}
	this.FavoritersTimeline = favoritersTimeline
	return &this
}

// NewTweetFavoritersResponseDataWithDefaults instantiates a new TweetFavoritersResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetFavoritersResponseDataWithDefaults() *TweetFavoritersResponseData {
	this := TweetFavoritersResponseData{}
	return &this
}

// GetFavoritersTimeline returns the FavoritersTimeline field value
func (o *TweetFavoritersResponseData) GetFavoritersTimeline() TimelineV2 {
	if o == nil {
		var ret TimelineV2
		return ret
	}

	return o.FavoritersTimeline
}

// GetFavoritersTimelineOk returns a tuple with the FavoritersTimeline field value
// and a boolean to check if the value has been set.
func (o *TweetFavoritersResponseData) GetFavoritersTimelineOk() (*TimelineV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FavoritersTimeline, true
}

// SetFavoritersTimeline sets field value
func (o *TweetFavoritersResponseData) SetFavoritersTimeline(v TimelineV2) {
	o.FavoritersTimeline = v
}

func (o TweetFavoritersResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetFavoritersResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["favoriters_timeline"] = o.FavoritersTimeline
	return toSerialize, nil
}

func (o *TweetFavoritersResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"favoriters_timeline",
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

	varTweetFavoritersResponseData := _TweetFavoritersResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTweetFavoritersResponseData)

	if err != nil {
		return err
	}

	*o = TweetFavoritersResponseData(varTweetFavoritersResponseData)

	return err
}

type NullableTweetFavoritersResponseData struct {
	value *TweetFavoritersResponseData
	isSet bool
}

func (v NullableTweetFavoritersResponseData) Get() *TweetFavoritersResponseData {
	return v.value
}

func (v *NullableTweetFavoritersResponseData) Set(val *TweetFavoritersResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetFavoritersResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetFavoritersResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetFavoritersResponseData(val *TweetFavoritersResponseData) *NullableTweetFavoritersResponseData {
	return &NullableTweetFavoritersResponseData{value: val, isSet: true}
}

func (v NullableTweetFavoritersResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetFavoritersResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


