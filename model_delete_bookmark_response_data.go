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

// checks if the DeleteBookmarkResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteBookmarkResponseData{}

// DeleteBookmarkResponseData struct for DeleteBookmarkResponseData
type DeleteBookmarkResponseData struct {
	TweetBookmarkDelete string `json:"tweet_bookmark_delete"`
}

type _DeleteBookmarkResponseData DeleteBookmarkResponseData

// NewDeleteBookmarkResponseData instantiates a new DeleteBookmarkResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteBookmarkResponseData(tweetBookmarkDelete string) *DeleteBookmarkResponseData {
	this := DeleteBookmarkResponseData{}
	this.TweetBookmarkDelete = tweetBookmarkDelete
	return &this
}

// NewDeleteBookmarkResponseDataWithDefaults instantiates a new DeleteBookmarkResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteBookmarkResponseDataWithDefaults() *DeleteBookmarkResponseData {
	this := DeleteBookmarkResponseData{}
	return &this
}

// GetTweetBookmarkDelete returns the TweetBookmarkDelete field value
func (o *DeleteBookmarkResponseData) GetTweetBookmarkDelete() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TweetBookmarkDelete
}

// GetTweetBookmarkDeleteOk returns a tuple with the TweetBookmarkDelete field value
// and a boolean to check if the value has been set.
func (o *DeleteBookmarkResponseData) GetTweetBookmarkDeleteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TweetBookmarkDelete, true
}

// SetTweetBookmarkDelete sets field value
func (o *DeleteBookmarkResponseData) SetTweetBookmarkDelete(v string) {
	o.TweetBookmarkDelete = v
}

func (o DeleteBookmarkResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteBookmarkResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tweet_bookmark_delete"] = o.TweetBookmarkDelete
	return toSerialize, nil
}

func (o *DeleteBookmarkResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tweet_bookmark_delete",
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

	varDeleteBookmarkResponseData := _DeleteBookmarkResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteBookmarkResponseData)

	if err != nil {
		return err
	}

	*o = DeleteBookmarkResponseData(varDeleteBookmarkResponseData)

	return err
}

type NullableDeleteBookmarkResponseData struct {
	value *DeleteBookmarkResponseData
	isSet bool
}

func (v NullableDeleteBookmarkResponseData) Get() *DeleteBookmarkResponseData {
	return v.value
}

func (v *NullableDeleteBookmarkResponseData) Set(val *DeleteBookmarkResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteBookmarkResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteBookmarkResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteBookmarkResponseData(val *DeleteBookmarkResponseData) *NullableDeleteBookmarkResponseData {
	return &NullableDeleteBookmarkResponseData{value: val, isSet: true}
}

func (v NullableDeleteBookmarkResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteBookmarkResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
