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

// checks if the TweetPreviousCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetPreviousCounts{}

// TweetPreviousCounts struct for TweetPreviousCounts
type TweetPreviousCounts struct {
	BookmarkCount int32 `json:"bookmark_count"`
	FavoriteCount int32 `json:"favorite_count"`
	QuoteCount    int32 `json:"quote_count"`
	ReplyCount    int32 `json:"reply_count"`
	RetweetCount  int32 `json:"retweet_count"`
}

type _TweetPreviousCounts TweetPreviousCounts

// NewTweetPreviousCounts instantiates a new TweetPreviousCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetPreviousCounts(bookmarkCount int32, favoriteCount int32, quoteCount int32, replyCount int32, retweetCount int32) *TweetPreviousCounts {
	this := TweetPreviousCounts{}
	this.BookmarkCount = bookmarkCount
	this.FavoriteCount = favoriteCount
	this.QuoteCount = quoteCount
	this.ReplyCount = replyCount
	this.RetweetCount = retweetCount
	return &this
}

// NewTweetPreviousCountsWithDefaults instantiates a new TweetPreviousCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetPreviousCountsWithDefaults() *TweetPreviousCounts {
	this := TweetPreviousCounts{}
	return &this
}

// GetBookmarkCount returns the BookmarkCount field value
func (o *TweetPreviousCounts) GetBookmarkCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BookmarkCount
}

// GetBookmarkCountOk returns a tuple with the BookmarkCount field value
// and a boolean to check if the value has been set.
func (o *TweetPreviousCounts) GetBookmarkCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BookmarkCount, true
}

// SetBookmarkCount sets field value
func (o *TweetPreviousCounts) SetBookmarkCount(v int32) {
	o.BookmarkCount = v
}

// GetFavoriteCount returns the FavoriteCount field value
func (o *TweetPreviousCounts) GetFavoriteCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FavoriteCount
}

// GetFavoriteCountOk returns a tuple with the FavoriteCount field value
// and a boolean to check if the value has been set.
func (o *TweetPreviousCounts) GetFavoriteCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FavoriteCount, true
}

// SetFavoriteCount sets field value
func (o *TweetPreviousCounts) SetFavoriteCount(v int32) {
	o.FavoriteCount = v
}

// GetQuoteCount returns the QuoteCount field value
func (o *TweetPreviousCounts) GetQuoteCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QuoteCount
}

// GetQuoteCountOk returns a tuple with the QuoteCount field value
// and a boolean to check if the value has been set.
func (o *TweetPreviousCounts) GetQuoteCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteCount, true
}

// SetQuoteCount sets field value
func (o *TweetPreviousCounts) SetQuoteCount(v int32) {
	o.QuoteCount = v
}

// GetReplyCount returns the ReplyCount field value
func (o *TweetPreviousCounts) GetReplyCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReplyCount
}

// GetReplyCountOk returns a tuple with the ReplyCount field value
// and a boolean to check if the value has been set.
func (o *TweetPreviousCounts) GetReplyCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplyCount, true
}

// SetReplyCount sets field value
func (o *TweetPreviousCounts) SetReplyCount(v int32) {
	o.ReplyCount = v
}

// GetRetweetCount returns the RetweetCount field value
func (o *TweetPreviousCounts) GetRetweetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RetweetCount
}

// GetRetweetCountOk returns a tuple with the RetweetCount field value
// and a boolean to check if the value has been set.
func (o *TweetPreviousCounts) GetRetweetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetweetCount, true
}

// SetRetweetCount sets field value
func (o *TweetPreviousCounts) SetRetweetCount(v int32) {
	o.RetweetCount = v
}

func (o TweetPreviousCounts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetPreviousCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bookmark_count"] = o.BookmarkCount
	toSerialize["favorite_count"] = o.FavoriteCount
	toSerialize["quote_count"] = o.QuoteCount
	toSerialize["reply_count"] = o.ReplyCount
	toSerialize["retweet_count"] = o.RetweetCount
	return toSerialize, nil
}

func (o *TweetPreviousCounts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bookmark_count",
		"favorite_count",
		"quote_count",
		"reply_count",
		"retweet_count",
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

	varTweetPreviousCounts := _TweetPreviousCounts{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTweetPreviousCounts)

	if err != nil {
		return err
	}

	*o = TweetPreviousCounts(varTweetPreviousCounts)

	return err
}

type NullableTweetPreviousCounts struct {
	value *TweetPreviousCounts
	isSet bool
}

func (v NullableTweetPreviousCounts) Get() *TweetPreviousCounts {
	return v.value
}

func (v *NullableTweetPreviousCounts) Set(val *TweetPreviousCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetPreviousCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetPreviousCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetPreviousCounts(val *TweetPreviousCounts) *NullableTweetPreviousCounts {
	return &NullableTweetPreviousCounts{value: val, isSet: true}
}

func (v NullableTweetPreviousCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetPreviousCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
