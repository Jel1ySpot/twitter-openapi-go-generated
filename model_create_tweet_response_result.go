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

// checks if the CreateTweetResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTweetResponseResult{}

// CreateTweetResponseResult struct for CreateTweetResponseResult
type CreateTweetResponseResult struct {
	TweetResults CreateTweet `json:"tweet_results"`
}

type _CreateTweetResponseResult CreateTweetResponseResult

// NewCreateTweetResponseResult instantiates a new CreateTweetResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTweetResponseResult(tweetResults CreateTweet) *CreateTweetResponseResult {
	this := CreateTweetResponseResult{}
	this.TweetResults = tweetResults
	return &this
}

// NewCreateTweetResponseResultWithDefaults instantiates a new CreateTweetResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTweetResponseResultWithDefaults() *CreateTweetResponseResult {
	this := CreateTweetResponseResult{}
	return &this
}

// GetTweetResults returns the TweetResults field value
func (o *CreateTweetResponseResult) GetTweetResults() CreateTweet {
	if o == nil {
		var ret CreateTweet
		return ret
	}

	return o.TweetResults
}

// GetTweetResultsOk returns a tuple with the TweetResults field value
// and a boolean to check if the value has been set.
func (o *CreateTweetResponseResult) GetTweetResultsOk() (*CreateTweet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TweetResults, true
}

// SetTweetResults sets field value
func (o *CreateTweetResponseResult) SetTweetResults(v CreateTweet) {
	o.TweetResults = v
}

func (o CreateTweetResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTweetResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tweet_results"] = o.TweetResults
	return toSerialize, nil
}

func (o *CreateTweetResponseResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tweet_results",
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

	varCreateTweetResponseResult := _CreateTweetResponseResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTweetResponseResult)

	if err != nil {
		return err
	}

	*o = CreateTweetResponseResult(varCreateTweetResponseResult)

	return err
}

type NullableCreateTweetResponseResult struct {
	value *CreateTweetResponseResult
	isSet bool
}

func (v NullableCreateTweetResponseResult) Get() *CreateTweetResponseResult {
	return v.value
}

func (v *NullableCreateTweetResponseResult) Set(val *CreateTweetResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTweetResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTweetResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTweetResponseResult(val *CreateTweetResponseResult) *NullableCreateTweetResponseResult {
	return &NullableCreateTweetResponseResult{value: val, isSet: true}
}

func (v NullableCreateTweetResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTweetResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

