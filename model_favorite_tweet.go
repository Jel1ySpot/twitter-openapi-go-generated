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

// checks if the FavoriteTweet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FavoriteTweet{}

// FavoriteTweet struct for FavoriteTweet
type FavoriteTweet struct {
	FavoriteTweet string `json:"favorite_tweet"`
}

type _FavoriteTweet FavoriteTweet

// NewFavoriteTweet instantiates a new FavoriteTweet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFavoriteTweet(favoriteTweet string) *FavoriteTweet {
	this := FavoriteTweet{}
	this.FavoriteTweet = favoriteTweet
	return &this
}

// NewFavoriteTweetWithDefaults instantiates a new FavoriteTweet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFavoriteTweetWithDefaults() *FavoriteTweet {
	this := FavoriteTweet{}
	return &this
}

// GetFavoriteTweet returns the FavoriteTweet field value
func (o *FavoriteTweet) GetFavoriteTweet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FavoriteTweet
}

// GetFavoriteTweetOk returns a tuple with the FavoriteTweet field value
// and a boolean to check if the value has been set.
func (o *FavoriteTweet) GetFavoriteTweetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FavoriteTweet, true
}

// SetFavoriteTweet sets field value
func (o *FavoriteTweet) SetFavoriteTweet(v string) {
	o.FavoriteTweet = v
}

func (o FavoriteTweet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FavoriteTweet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["favorite_tweet"] = o.FavoriteTweet
	return toSerialize, nil
}

func (o *FavoriteTweet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"favorite_tweet",
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

	varFavoriteTweet := _FavoriteTweet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFavoriteTweet)

	if err != nil {
		return err
	}

	*o = FavoriteTweet(varFavoriteTweet)

	return err
}

type NullableFavoriteTweet struct {
	value *FavoriteTweet
	isSet bool
}

func (v NullableFavoriteTweet) Get() *FavoriteTweet {
	return v.value
}

func (v *NullableFavoriteTweet) Set(val *FavoriteTweet) {
	v.value = val
	v.isSet = true
}

func (v NullableFavoriteTweet) IsSet() bool {
	return v.isSet
}

func (v *NullableFavoriteTweet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFavoriteTweet(val *FavoriteTweet) *NullableFavoriteTweet {
	return &NullableFavoriteTweet{value: val, isSet: true}
}

func (v NullableFavoriteTweet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFavoriteTweet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
