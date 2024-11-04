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

// checks if the PostCreateRetweetRequestVariables type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCreateRetweetRequestVariables{}

// PostCreateRetweetRequestVariables struct for PostCreateRetweetRequestVariables
type PostCreateRetweetRequestVariables struct {
	DarkRequest bool `json:"dark_request"`
	TweetId string `json:"tweet_id"`
}

type _PostCreateRetweetRequestVariables PostCreateRetweetRequestVariables

// NewPostCreateRetweetRequestVariables instantiates a new PostCreateRetweetRequestVariables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCreateRetweetRequestVariables(darkRequest bool, tweetId string) *PostCreateRetweetRequestVariables {
	this := PostCreateRetweetRequestVariables{}
	this.DarkRequest = darkRequest
	this.TweetId = tweetId
	return &this
}

// NewPostCreateRetweetRequestVariablesWithDefaults instantiates a new PostCreateRetweetRequestVariables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCreateRetweetRequestVariablesWithDefaults() *PostCreateRetweetRequestVariables {
	this := PostCreateRetweetRequestVariables{}
	var darkRequest bool = false
	this.DarkRequest = darkRequest
	var tweetId string = "1349129669258448897"
	this.TweetId = tweetId
	return &this
}

// GetDarkRequest returns the DarkRequest field value
func (o *PostCreateRetweetRequestVariables) GetDarkRequest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DarkRequest
}

// GetDarkRequestOk returns a tuple with the DarkRequest field value
// and a boolean to check if the value has been set.
func (o *PostCreateRetweetRequestVariables) GetDarkRequestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DarkRequest, true
}

// SetDarkRequest sets field value
func (o *PostCreateRetweetRequestVariables) SetDarkRequest(v bool) {
	o.DarkRequest = v
}

// GetTweetId returns the TweetId field value
func (o *PostCreateRetweetRequestVariables) GetTweetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TweetId
}

// GetTweetIdOk returns a tuple with the TweetId field value
// and a boolean to check if the value has been set.
func (o *PostCreateRetweetRequestVariables) GetTweetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TweetId, true
}

// SetTweetId sets field value
func (o *PostCreateRetweetRequestVariables) SetTweetId(v string) {
	o.TweetId = v
}

func (o PostCreateRetweetRequestVariables) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCreateRetweetRequestVariables) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dark_request"] = o.DarkRequest
	toSerialize["tweet_id"] = o.TweetId
	return toSerialize, nil
}

func (o *PostCreateRetweetRequestVariables) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dark_request",
		"tweet_id",
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

	varPostCreateRetweetRequestVariables := _PostCreateRetweetRequestVariables{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCreateRetweetRequestVariables)

	if err != nil {
		return err
	}

	*o = PostCreateRetweetRequestVariables(varPostCreateRetweetRequestVariables)

	return err
}

type NullablePostCreateRetweetRequestVariables struct {
	value *PostCreateRetweetRequestVariables
	isSet bool
}

func (v NullablePostCreateRetweetRequestVariables) Get() *PostCreateRetweetRequestVariables {
	return v.value
}

func (v *NullablePostCreateRetweetRequestVariables) Set(val *PostCreateRetweetRequestVariables) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCreateRetweetRequestVariables) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCreateRetweetRequestVariables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCreateRetweetRequestVariables(val *PostCreateRetweetRequestVariables) *NullablePostCreateRetweetRequestVariables {
	return &NullablePostCreateRetweetRequestVariables{value: val, isSet: true}
}

func (v NullablePostCreateRetweetRequestVariables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCreateRetweetRequestVariables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


