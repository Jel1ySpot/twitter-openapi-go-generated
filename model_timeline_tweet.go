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

// checks if the TimelineTweet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelineTweet{}

// TimelineTweet struct for TimelineTweet
type TimelineTweet struct {
	Typename TypeName `json:"__typename"`
	Highlights *Highlight `json:"highlights,omitempty"`
	ItemType ContentItemType `json:"itemType"`
	PromotedMetadata map[string]interface{} `json:"promotedMetadata,omitempty"`
	SocialContext *SocialContextUnion `json:"socialContext,omitempty"`
	TweetDisplayType string `json:"tweetDisplayType"`
	TweetResults ItemResult `json:"tweet_results"`
}

type _TimelineTweet TimelineTweet

// NewTimelineTweet instantiates a new TimelineTweet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineTweet(typename TypeName, itemType ContentItemType, tweetDisplayType string, tweetResults ItemResult) *TimelineTweet {
	this := TimelineTweet{}
	this.Typename = typename
	this.ItemType = itemType
	this.TweetDisplayType = tweetDisplayType
	this.TweetResults = tweetResults
	return &this
}

// NewTimelineTweetWithDefaults instantiates a new TimelineTweet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineTweetWithDefaults() *TimelineTweet {
	this := TimelineTweet{}
	return &this
}

// GetTypename returns the Typename field value
func (o *TimelineTweet) GetTypename() TypeName {
	if o == nil {
		var ret TypeName
		return ret
	}

	return o.Typename
}

// GetTypenameOk returns a tuple with the Typename field value
// and a boolean to check if the value has been set.
func (o *TimelineTweet) GetTypenameOk() (*TypeName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Typename, true
}

// SetTypename sets field value
func (o *TimelineTweet) SetTypename(v TypeName) {
	o.Typename = v
}

// GetHighlights returns the Highlights field value if set, zero value otherwise.
func (o *TimelineTweet) GetHighlights() Highlight {
	if o == nil || IsNil(o.Highlights) {
		var ret Highlight
		return ret
	}
	return *o.Highlights
}

// GetHighlightsOk returns a tuple with the Highlights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineTweet) GetHighlightsOk() (*Highlight, bool) {
	if o == nil || IsNil(o.Highlights) {
		return nil, false
	}
	return o.Highlights, true
}

// HasHighlights returns a boolean if a field has been set.
func (o *TimelineTweet) HasHighlights() bool {
	if o != nil && !IsNil(o.Highlights) {
		return true
	}

	return false
}

// SetHighlights gets a reference to the given Highlight and assigns it to the Highlights field.
func (o *TimelineTweet) SetHighlights(v Highlight) {
	o.Highlights = &v
}

// GetItemType returns the ItemType field value
func (o *TimelineTweet) GetItemType() ContentItemType {
	if o == nil {
		var ret ContentItemType
		return ret
	}

	return o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value
// and a boolean to check if the value has been set.
func (o *TimelineTweet) GetItemTypeOk() (*ContentItemType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemType, true
}

// SetItemType sets field value
func (o *TimelineTweet) SetItemType(v ContentItemType) {
	o.ItemType = v
}

// GetPromotedMetadata returns the PromotedMetadata field value if set, zero value otherwise.
func (o *TimelineTweet) GetPromotedMetadata() map[string]interface{} {
	if o == nil || IsNil(o.PromotedMetadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.PromotedMetadata
}

// GetPromotedMetadataOk returns a tuple with the PromotedMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineTweet) GetPromotedMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PromotedMetadata) {
		return map[string]interface{}{}, false
	}
	return o.PromotedMetadata, true
}

// HasPromotedMetadata returns a boolean if a field has been set.
func (o *TimelineTweet) HasPromotedMetadata() bool {
	if o != nil && !IsNil(o.PromotedMetadata) {
		return true
	}

	return false
}

// SetPromotedMetadata gets a reference to the given map[string]interface{} and assigns it to the PromotedMetadata field.
func (o *TimelineTweet) SetPromotedMetadata(v map[string]interface{}) {
	o.PromotedMetadata = v
}

// GetSocialContext returns the SocialContext field value if set, zero value otherwise.
func (o *TimelineTweet) GetSocialContext() SocialContextUnion {
	if o == nil || IsNil(o.SocialContext) {
		var ret SocialContextUnion
		return ret
	}
	return *o.SocialContext
}

// GetSocialContextOk returns a tuple with the SocialContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineTweet) GetSocialContextOk() (*SocialContextUnion, bool) {
	if o == nil || IsNil(o.SocialContext) {
		return nil, false
	}
	return o.SocialContext, true
}

// HasSocialContext returns a boolean if a field has been set.
func (o *TimelineTweet) HasSocialContext() bool {
	if o != nil && !IsNil(o.SocialContext) {
		return true
	}

	return false
}

// SetSocialContext gets a reference to the given SocialContextUnion and assigns it to the SocialContext field.
func (o *TimelineTweet) SetSocialContext(v SocialContextUnion) {
	o.SocialContext = &v
}

// GetTweetDisplayType returns the TweetDisplayType field value
func (o *TimelineTweet) GetTweetDisplayType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TweetDisplayType
}

// GetTweetDisplayTypeOk returns a tuple with the TweetDisplayType field value
// and a boolean to check if the value has been set.
func (o *TimelineTweet) GetTweetDisplayTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TweetDisplayType, true
}

// SetTweetDisplayType sets field value
func (o *TimelineTweet) SetTweetDisplayType(v string) {
	o.TweetDisplayType = v
}

// GetTweetResults returns the TweetResults field value
func (o *TimelineTweet) GetTweetResults() ItemResult {
	if o == nil {
		var ret ItemResult
		return ret
	}

	return o.TweetResults
}

// GetTweetResultsOk returns a tuple with the TweetResults field value
// and a boolean to check if the value has been set.
func (o *TimelineTweet) GetTweetResultsOk() (*ItemResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TweetResults, true
}

// SetTweetResults sets field value
func (o *TimelineTweet) SetTweetResults(v ItemResult) {
	o.TweetResults = v
}

func (o TimelineTweet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimelineTweet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["__typename"] = o.Typename
	if !IsNil(o.Highlights) {
		toSerialize["highlights"] = o.Highlights
	}
	toSerialize["itemType"] = o.ItemType
	if !IsNil(o.PromotedMetadata) {
		toSerialize["promotedMetadata"] = o.PromotedMetadata
	}
	if !IsNil(o.SocialContext) {
		toSerialize["socialContext"] = o.SocialContext
	}
	toSerialize["tweetDisplayType"] = o.TweetDisplayType
	toSerialize["tweet_results"] = o.TweetResults
	return toSerialize, nil
}

func (o *TimelineTweet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"__typename",
		"itemType",
		"tweetDisplayType",
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

	varTimelineTweet := _TimelineTweet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTimelineTweet)

	if err != nil {
		return err
	}

	*o = TimelineTweet(varTimelineTweet)

	return err
}

type NullableTimelineTweet struct {
	value *TimelineTweet
	isSet bool
}

func (v NullableTimelineTweet) Get() *TimelineTweet {
	return v.value
}

func (v *NullableTimelineTweet) Set(val *TimelineTweet) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineTweet) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineTweet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineTweet(val *TimelineTweet) *NullableTimelineTweet {
	return &NullableTimelineTweet{value: val, isSet: true}
}

func (v NullableTimelineTweet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineTweet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

