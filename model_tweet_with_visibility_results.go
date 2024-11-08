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

// checks if the TweetWithVisibilityResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetWithVisibilityResults{}

// TweetWithVisibilityResults struct for TweetWithVisibilityResults
type TweetWithVisibilityResults struct {
	Typename               TypeName                `json:"__typename"`
	LimitedActionResults   map[string]interface{}  `json:"limitedActionResults,omitempty"`
	MediaVisibilityResults *MediaVisibilityResults `json:"mediaVisibilityResults,omitempty"`
	Tweet                  Tweet                   `json:"tweet"`
	TweetInterstitial      *TweetInterstitial      `json:"tweetInterstitial,omitempty"`
}

type _TweetWithVisibilityResults TweetWithVisibilityResults

// NewTweetWithVisibilityResults instantiates a new TweetWithVisibilityResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetWithVisibilityResults(typename TypeName, tweet Tweet) *TweetWithVisibilityResults {
	this := TweetWithVisibilityResults{}
	this.Typename = typename
	this.Tweet = tweet
	return &this
}

// NewTweetWithVisibilityResultsWithDefaults instantiates a new TweetWithVisibilityResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetWithVisibilityResultsWithDefaults() *TweetWithVisibilityResults {
	this := TweetWithVisibilityResults{}
	return &this
}

// GetTypename returns the Typename field value
func (o *TweetWithVisibilityResults) GetTypename() TypeName {
	if o == nil {
		var ret TypeName
		return ret
	}

	return o.Typename
}

// GetTypenameOk returns a tuple with the Typename field value
// and a boolean to check if the value has been set.
func (o *TweetWithVisibilityResults) GetTypenameOk() (*TypeName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Typename, true
}

// SetTypename sets field value
func (o *TweetWithVisibilityResults) SetTypename(v TypeName) {
	o.Typename = v
}

// GetLimitedActionResults returns the LimitedActionResults field value if set, zero value otherwise.
func (o *TweetWithVisibilityResults) GetLimitedActionResults() map[string]interface{} {
	if o == nil || IsNil(o.LimitedActionResults) {
		var ret map[string]interface{}
		return ret
	}
	return o.LimitedActionResults
}

// GetLimitedActionResultsOk returns a tuple with the LimitedActionResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetWithVisibilityResults) GetLimitedActionResultsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.LimitedActionResults) {
		return map[string]interface{}{}, false
	}
	return o.LimitedActionResults, true
}

// HasLimitedActionResults returns a boolean if a field has been set.
func (o *TweetWithVisibilityResults) HasLimitedActionResults() bool {
	if o != nil && !IsNil(o.LimitedActionResults) {
		return true
	}

	return false
}

// SetLimitedActionResults gets a reference to the given map[string]interface{} and assigns it to the LimitedActionResults field.
func (o *TweetWithVisibilityResults) SetLimitedActionResults(v map[string]interface{}) {
	o.LimitedActionResults = v
}

// GetMediaVisibilityResults returns the MediaVisibilityResults field value if set, zero value otherwise.
func (o *TweetWithVisibilityResults) GetMediaVisibilityResults() MediaVisibilityResults {
	if o == nil || IsNil(o.MediaVisibilityResults) {
		var ret MediaVisibilityResults
		return ret
	}
	return *o.MediaVisibilityResults
}

// GetMediaVisibilityResultsOk returns a tuple with the MediaVisibilityResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetWithVisibilityResults) GetMediaVisibilityResultsOk() (*MediaVisibilityResults, bool) {
	if o == nil || IsNil(o.MediaVisibilityResults) {
		return nil, false
	}
	return o.MediaVisibilityResults, true
}

// HasMediaVisibilityResults returns a boolean if a field has been set.
func (o *TweetWithVisibilityResults) HasMediaVisibilityResults() bool {
	if o != nil && !IsNil(o.MediaVisibilityResults) {
		return true
	}

	return false
}

// SetMediaVisibilityResults gets a reference to the given MediaVisibilityResults and assigns it to the MediaVisibilityResults field.
func (o *TweetWithVisibilityResults) SetMediaVisibilityResults(v MediaVisibilityResults) {
	o.MediaVisibilityResults = &v
}

// GetTweet returns the Tweet field value
func (o *TweetWithVisibilityResults) GetTweet() Tweet {
	if o == nil {
		var ret Tweet
		return ret
	}

	return o.Tweet
}

// GetTweetOk returns a tuple with the Tweet field value
// and a boolean to check if the value has been set.
func (o *TweetWithVisibilityResults) GetTweetOk() (*Tweet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tweet, true
}

// SetTweet sets field value
func (o *TweetWithVisibilityResults) SetTweet(v Tweet) {
	o.Tweet = v
}

// GetTweetInterstitial returns the TweetInterstitial field value if set, zero value otherwise.
func (o *TweetWithVisibilityResults) GetTweetInterstitial() TweetInterstitial {
	if o == nil || IsNil(o.TweetInterstitial) {
		var ret TweetInterstitial
		return ret
	}
	return *o.TweetInterstitial
}

// GetTweetInterstitialOk returns a tuple with the TweetInterstitial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetWithVisibilityResults) GetTweetInterstitialOk() (*TweetInterstitial, bool) {
	if o == nil || IsNil(o.TweetInterstitial) {
		return nil, false
	}
	return o.TweetInterstitial, true
}

// HasTweetInterstitial returns a boolean if a field has been set.
func (o *TweetWithVisibilityResults) HasTweetInterstitial() bool {
	if o != nil && !IsNil(o.TweetInterstitial) {
		return true
	}

	return false
}

// SetTweetInterstitial gets a reference to the given TweetInterstitial and assigns it to the TweetInterstitial field.
func (o *TweetWithVisibilityResults) SetTweetInterstitial(v TweetInterstitial) {
	o.TweetInterstitial = &v
}

func (o TweetWithVisibilityResults) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetWithVisibilityResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["__typename"] = o.Typename
	if !IsNil(o.LimitedActionResults) {
		toSerialize["limitedActionResults"] = o.LimitedActionResults
	}
	if !IsNil(o.MediaVisibilityResults) {
		toSerialize["mediaVisibilityResults"] = o.MediaVisibilityResults
	}
	toSerialize["tweet"] = o.Tweet
	if !IsNil(o.TweetInterstitial) {
		toSerialize["tweetInterstitial"] = o.TweetInterstitial
	}
	return toSerialize, nil
}

func (o *TweetWithVisibilityResults) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"__typename",
		"tweet",
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

	varTweetWithVisibilityResults := _TweetWithVisibilityResults{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTweetWithVisibilityResults)

	if err != nil {
		return err
	}

	*o = TweetWithVisibilityResults(varTweetWithVisibilityResults)

	return err
}

type NullableTweetWithVisibilityResults struct {
	value *TweetWithVisibilityResults
	isSet bool
}

func (v NullableTweetWithVisibilityResults) Get() *TweetWithVisibilityResults {
	return v.value
}

func (v *NullableTweetWithVisibilityResults) Set(val *TweetWithVisibilityResults) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetWithVisibilityResults) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetWithVisibilityResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetWithVisibilityResults(val *TweetWithVisibilityResults) *NullableTweetWithVisibilityResults {
	return &NullableTweetWithVisibilityResults{value: val, isSet: true}
}

func (v NullableTweetWithVisibilityResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetWithVisibilityResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
