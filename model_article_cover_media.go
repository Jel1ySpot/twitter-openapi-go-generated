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

// checks if the ArticleCoverMedia type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArticleCoverMedia{}

// ArticleCoverMedia struct for ArticleCoverMedia
type ArticleCoverMedia struct {
	Id string `json:"id"`
	MediaId string `json:"media_id" validate:"regexp=^[0-9]+$"`
	MediaInfo ArticleCoverMediaInfo `json:"media_info"`
	MediaKey string `json:"media_key"`
}

type _ArticleCoverMedia ArticleCoverMedia

// NewArticleCoverMedia instantiates a new ArticleCoverMedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArticleCoverMedia(id string, mediaId string, mediaInfo ArticleCoverMediaInfo, mediaKey string) *ArticleCoverMedia {
	this := ArticleCoverMedia{}
	this.Id = id
	this.MediaId = mediaId
	this.MediaInfo = mediaInfo
	this.MediaKey = mediaKey
	return &this
}

// NewArticleCoverMediaWithDefaults instantiates a new ArticleCoverMedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArticleCoverMediaWithDefaults() *ArticleCoverMedia {
	this := ArticleCoverMedia{}
	return &this
}

// GetId returns the Id field value
func (o *ArticleCoverMedia) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArticleCoverMedia) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArticleCoverMedia) SetId(v string) {
	o.Id = v
}

// GetMediaId returns the MediaId field value
func (o *ArticleCoverMedia) GetMediaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MediaId
}

// GetMediaIdOk returns a tuple with the MediaId field value
// and a boolean to check if the value has been set.
func (o *ArticleCoverMedia) GetMediaIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaId, true
}

// SetMediaId sets field value
func (o *ArticleCoverMedia) SetMediaId(v string) {
	o.MediaId = v
}

// GetMediaInfo returns the MediaInfo field value
func (o *ArticleCoverMedia) GetMediaInfo() ArticleCoverMediaInfo {
	if o == nil {
		var ret ArticleCoverMediaInfo
		return ret
	}

	return o.MediaInfo
}

// GetMediaInfoOk returns a tuple with the MediaInfo field value
// and a boolean to check if the value has been set.
func (o *ArticleCoverMedia) GetMediaInfoOk() (*ArticleCoverMediaInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaInfo, true
}

// SetMediaInfo sets field value
func (o *ArticleCoverMedia) SetMediaInfo(v ArticleCoverMediaInfo) {
	o.MediaInfo = v
}

// GetMediaKey returns the MediaKey field value
func (o *ArticleCoverMedia) GetMediaKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MediaKey
}

// GetMediaKeyOk returns a tuple with the MediaKey field value
// and a boolean to check if the value has been set.
func (o *ArticleCoverMedia) GetMediaKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaKey, true
}

// SetMediaKey sets field value
func (o *ArticleCoverMedia) SetMediaKey(v string) {
	o.MediaKey = v
}

func (o ArticleCoverMedia) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArticleCoverMedia) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["media_id"] = o.MediaId
	toSerialize["media_info"] = o.MediaInfo
	toSerialize["media_key"] = o.MediaKey
	return toSerialize, nil
}

func (o *ArticleCoverMedia) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"media_id",
		"media_info",
		"media_key",
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

	varArticleCoverMedia := _ArticleCoverMedia{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArticleCoverMedia)

	if err != nil {
		return err
	}

	*o = ArticleCoverMedia(varArticleCoverMedia)

	return err
}

type NullableArticleCoverMedia struct {
	value *ArticleCoverMedia
	isSet bool
}

func (v NullableArticleCoverMedia) Get() *ArticleCoverMedia {
	return v.value
}

func (v *NullableArticleCoverMedia) Set(val *ArticleCoverMedia) {
	v.value = val
	v.isSet = true
}

func (v NullableArticleCoverMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableArticleCoverMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArticleCoverMedia(val *ArticleCoverMedia) *NullableArticleCoverMedia {
	return &NullableArticleCoverMedia{value: val, isSet: true}
}

func (v NullableArticleCoverMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArticleCoverMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


