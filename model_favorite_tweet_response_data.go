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

// checks if the FavoriteTweetResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FavoriteTweetResponseData{}

// FavoriteTweetResponseData struct for FavoriteTweetResponseData
type FavoriteTweetResponseData struct {
	Data FavoriteTweet `json:"data"`
}

type _FavoriteTweetResponseData FavoriteTweetResponseData

// NewFavoriteTweetResponseData instantiates a new FavoriteTweetResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFavoriteTweetResponseData(data FavoriteTweet) *FavoriteTweetResponseData {
	this := FavoriteTweetResponseData{}
	this.Data = data
	return &this
}

// NewFavoriteTweetResponseDataWithDefaults instantiates a new FavoriteTweetResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFavoriteTweetResponseDataWithDefaults() *FavoriteTweetResponseData {
	this := FavoriteTweetResponseData{}
	return &this
}

// GetData returns the Data field value
func (o *FavoriteTweetResponseData) GetData() FavoriteTweet {
	if o == nil {
		var ret FavoriteTweet
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FavoriteTweetResponseData) GetDataOk() (*FavoriteTweet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FavoriteTweetResponseData) SetData(v FavoriteTweet) {
	o.Data = v
}

func (o FavoriteTweetResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FavoriteTweetResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *FavoriteTweetResponseData) UnmarshalJSON(data []byte) (err error) {
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

	varFavoriteTweetResponseData := _FavoriteTweetResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFavoriteTweetResponseData)

	if err != nil {
		return err
	}

	*o = FavoriteTweetResponseData(varFavoriteTweetResponseData)

	return err
}

type NullableFavoriteTweetResponseData struct {
	value *FavoriteTweetResponseData
	isSet bool
}

func (v NullableFavoriteTweetResponseData) Get() *FavoriteTweetResponseData {
	return v.value
}

func (v *NullableFavoriteTweetResponseData) Set(val *FavoriteTweetResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableFavoriteTweetResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableFavoriteTweetResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFavoriteTweetResponseData(val *FavoriteTweetResponseData) *NullableFavoriteTweetResponseData {
	return &NullableFavoriteTweetResponseData{value: val, isSet: true}
}

func (v NullableFavoriteTweetResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFavoriteTweetResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


