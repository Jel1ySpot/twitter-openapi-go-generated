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

// checks if the BookmarksResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookmarksResponseData{}

// BookmarksResponseData struct for BookmarksResponseData
type BookmarksResponseData struct {
	BookmarkTimelineV2 BookmarksTimeline `json:"bookmark_timeline_v2"`
}

type _BookmarksResponseData BookmarksResponseData

// NewBookmarksResponseData instantiates a new BookmarksResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarksResponseData(bookmarkTimelineV2 BookmarksTimeline) *BookmarksResponseData {
	this := BookmarksResponseData{}
	this.BookmarkTimelineV2 = bookmarkTimelineV2
	return &this
}

// NewBookmarksResponseDataWithDefaults instantiates a new BookmarksResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarksResponseDataWithDefaults() *BookmarksResponseData {
	this := BookmarksResponseData{}
	return &this
}

// GetBookmarkTimelineV2 returns the BookmarkTimelineV2 field value
func (o *BookmarksResponseData) GetBookmarkTimelineV2() BookmarksTimeline {
	if o == nil {
		var ret BookmarksTimeline
		return ret
	}

	return o.BookmarkTimelineV2
}

// GetBookmarkTimelineV2Ok returns a tuple with the BookmarkTimelineV2 field value
// and a boolean to check if the value has been set.
func (o *BookmarksResponseData) GetBookmarkTimelineV2Ok() (*BookmarksTimeline, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BookmarkTimelineV2, true
}

// SetBookmarkTimelineV2 sets field value
func (o *BookmarksResponseData) SetBookmarkTimelineV2(v BookmarksTimeline) {
	o.BookmarkTimelineV2 = v
}

func (o BookmarksResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookmarksResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bookmark_timeline_v2"] = o.BookmarkTimelineV2
	return toSerialize, nil
}

func (o *BookmarksResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bookmark_timeline_v2",
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

	varBookmarksResponseData := _BookmarksResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBookmarksResponseData)

	if err != nil {
		return err
	}

	*o = BookmarksResponseData(varBookmarksResponseData)

	return err
}

type NullableBookmarksResponseData struct {
	value *BookmarksResponseData
	isSet bool
}

func (v NullableBookmarksResponseData) Get() *BookmarksResponseData {
	return v.value
}

func (v *NullableBookmarksResponseData) Set(val *BookmarksResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmarksResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmarksResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmarksResponseData(val *BookmarksResponseData) *NullableBookmarksResponseData {
	return &NullableBookmarksResponseData{value: val, isSet: true}
}

func (v NullableBookmarksResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmarksResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
