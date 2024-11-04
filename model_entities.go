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

// checks if the Entities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Entities{}

// Entities struct for Entities
type Entities struct {
	Hashtags []map[string]interface{} `json:"hashtags"`
	Media []Media `json:"media,omitempty"`
	Symbols []map[string]interface{} `json:"symbols"`
	Timestamps []Timestamp `json:"timestamps,omitempty"`
	Urls []Url `json:"urls"`
	UserMentions []map[string]interface{} `json:"user_mentions"`
}

type _Entities Entities

// NewEntities instantiates a new Entities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntities(hashtags []map[string]interface{}, symbols []map[string]interface{}, urls []Url, userMentions []map[string]interface{}) *Entities {
	this := Entities{}
	this.Hashtags = hashtags
	this.Symbols = symbols
	this.Urls = urls
	this.UserMentions = userMentions
	return &this
}

// NewEntitiesWithDefaults instantiates a new Entities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitiesWithDefaults() *Entities {
	this := Entities{}
	return &this
}

// GetHashtags returns the Hashtags field value
func (o *Entities) GetHashtags() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Hashtags
}

// GetHashtagsOk returns a tuple with the Hashtags field value
// and a boolean to check if the value has been set.
func (o *Entities) GetHashtagsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hashtags, true
}

// SetHashtags sets field value
func (o *Entities) SetHashtags(v []map[string]interface{}) {
	o.Hashtags = v
}

// GetMedia returns the Media field value if set, zero value otherwise.
func (o *Entities) GetMedia() []Media {
	if o == nil || IsNil(o.Media) {
		var ret []Media
		return ret
	}
	return o.Media
}

// GetMediaOk returns a tuple with the Media field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entities) GetMediaOk() ([]Media, bool) {
	if o == nil || IsNil(o.Media) {
		return nil, false
	}
	return o.Media, true
}

// HasMedia returns a boolean if a field has been set.
func (o *Entities) HasMedia() bool {
	if o != nil && !IsNil(o.Media) {
		return true
	}

	return false
}

// SetMedia gets a reference to the given []Media and assigns it to the Media field.
func (o *Entities) SetMedia(v []Media) {
	o.Media = v
}

// GetSymbols returns the Symbols field value
func (o *Entities) GetSymbols() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value
// and a boolean to check if the value has been set.
func (o *Entities) GetSymbolsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Symbols, true
}

// SetSymbols sets field value
func (o *Entities) SetSymbols(v []map[string]interface{}) {
	o.Symbols = v
}

// GetTimestamps returns the Timestamps field value if set, zero value otherwise.
func (o *Entities) GetTimestamps() []Timestamp {
	if o == nil || IsNil(o.Timestamps) {
		var ret []Timestamp
		return ret
	}
	return o.Timestamps
}

// GetTimestampsOk returns a tuple with the Timestamps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entities) GetTimestampsOk() ([]Timestamp, bool) {
	if o == nil || IsNil(o.Timestamps) {
		return nil, false
	}
	return o.Timestamps, true
}

// HasTimestamps returns a boolean if a field has been set.
func (o *Entities) HasTimestamps() bool {
	if o != nil && !IsNil(o.Timestamps) {
		return true
	}

	return false
}

// SetTimestamps gets a reference to the given []Timestamp and assigns it to the Timestamps field.
func (o *Entities) SetTimestamps(v []Timestamp) {
	o.Timestamps = v
}

// GetUrls returns the Urls field value
func (o *Entities) GetUrls() []Url {
	if o == nil {
		var ret []Url
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *Entities) GetUrlsOk() ([]Url, bool) {
	if o == nil {
		return nil, false
	}
	return o.Urls, true
}

// SetUrls sets field value
func (o *Entities) SetUrls(v []Url) {
	o.Urls = v
}

// GetUserMentions returns the UserMentions field value
func (o *Entities) GetUserMentions() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.UserMentions
}

// GetUserMentionsOk returns a tuple with the UserMentions field value
// and a boolean to check if the value has been set.
func (o *Entities) GetUserMentionsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserMentions, true
}

// SetUserMentions sets field value
func (o *Entities) SetUserMentions(v []map[string]interface{}) {
	o.UserMentions = v
}

func (o Entities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Entities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hashtags"] = o.Hashtags
	if !IsNil(o.Media) {
		toSerialize["media"] = o.Media
	}
	toSerialize["symbols"] = o.Symbols
	if !IsNil(o.Timestamps) {
		toSerialize["timestamps"] = o.Timestamps
	}
	toSerialize["urls"] = o.Urls
	toSerialize["user_mentions"] = o.UserMentions
	return toSerialize, nil
}

func (o *Entities) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hashtags",
		"symbols",
		"urls",
		"user_mentions",
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

	varEntities := _Entities{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEntities)

	if err != nil {
		return err
	}

	*o = Entities(varEntities)

	return err
}

type NullableEntities struct {
	value *Entities
	isSet bool
}

func (v NullableEntities) Get() *Entities {
	return v.value
}

func (v *NullableEntities) Set(val *Entities) {
	v.value = val
	v.isSet = true
}

func (v NullableEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntities(val *Entities) *NullableEntities {
	return &NullableEntities{value: val, isSet: true}
}

func (v NullableEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

