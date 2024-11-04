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

// checks if the PostCreateTweetRequestVariablesMedia type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCreateTweetRequestVariablesMedia{}

// PostCreateTweetRequestVariablesMedia struct for PostCreateTweetRequestVariablesMedia
type PostCreateTweetRequestVariablesMedia struct {
	MediaEntities []PostCreateTweetRequestVariablesMediaMediaEntitiesInner `json:"media_entities,omitempty"`
	PossiblySensitive bool `json:"possibly_sensitive"`
}

type _PostCreateTweetRequestVariablesMedia PostCreateTweetRequestVariablesMedia

// NewPostCreateTweetRequestVariablesMedia instantiates a new PostCreateTweetRequestVariablesMedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCreateTweetRequestVariablesMedia(possiblySensitive bool) *PostCreateTweetRequestVariablesMedia {
	this := PostCreateTweetRequestVariablesMedia{}
	this.PossiblySensitive = possiblySensitive
	return &this
}

// NewPostCreateTweetRequestVariablesMediaWithDefaults instantiates a new PostCreateTweetRequestVariablesMedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCreateTweetRequestVariablesMediaWithDefaults() *PostCreateTweetRequestVariablesMedia {
	this := PostCreateTweetRequestVariablesMedia{}
	var possiblySensitive bool = false
	this.PossiblySensitive = possiblySensitive
	return &this
}

// GetMediaEntities returns the MediaEntities field value if set, zero value otherwise.
func (o *PostCreateTweetRequestVariablesMedia) GetMediaEntities() []PostCreateTweetRequestVariablesMediaMediaEntitiesInner {
	if o == nil || IsNil(o.MediaEntities) {
		var ret []PostCreateTweetRequestVariablesMediaMediaEntitiesInner
		return ret
	}
	return o.MediaEntities
}

// GetMediaEntitiesOk returns a tuple with the MediaEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateTweetRequestVariablesMedia) GetMediaEntitiesOk() ([]PostCreateTweetRequestVariablesMediaMediaEntitiesInner, bool) {
	if o == nil || IsNil(o.MediaEntities) {
		return nil, false
	}
	return o.MediaEntities, true
}

// HasMediaEntities returns a boolean if a field has been set.
func (o *PostCreateTweetRequestVariablesMedia) HasMediaEntities() bool {
	if o != nil && !IsNil(o.MediaEntities) {
		return true
	}

	return false
}

// SetMediaEntities gets a reference to the given []PostCreateTweetRequestVariablesMediaMediaEntitiesInner and assigns it to the MediaEntities field.
func (o *PostCreateTweetRequestVariablesMedia) SetMediaEntities(v []PostCreateTweetRequestVariablesMediaMediaEntitiesInner) {
	o.MediaEntities = v
}

// GetPossiblySensitive returns the PossiblySensitive field value
func (o *PostCreateTweetRequestVariablesMedia) GetPossiblySensitive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PossiblySensitive
}

// GetPossiblySensitiveOk returns a tuple with the PossiblySensitive field value
// and a boolean to check if the value has been set.
func (o *PostCreateTweetRequestVariablesMedia) GetPossiblySensitiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PossiblySensitive, true
}

// SetPossiblySensitive sets field value
func (o *PostCreateTweetRequestVariablesMedia) SetPossiblySensitive(v bool) {
	o.PossiblySensitive = v
}

func (o PostCreateTweetRequestVariablesMedia) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCreateTweetRequestVariablesMedia) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MediaEntities) {
		toSerialize["media_entities"] = o.MediaEntities
	}
	toSerialize["possibly_sensitive"] = o.PossiblySensitive
	return toSerialize, nil
}

func (o *PostCreateTweetRequestVariablesMedia) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"possibly_sensitive",
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

	varPostCreateTweetRequestVariablesMedia := _PostCreateTweetRequestVariablesMedia{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCreateTweetRequestVariablesMedia)

	if err != nil {
		return err
	}

	*o = PostCreateTweetRequestVariablesMedia(varPostCreateTweetRequestVariablesMedia)

	return err
}

type NullablePostCreateTweetRequestVariablesMedia struct {
	value *PostCreateTweetRequestVariablesMedia
	isSet bool
}

func (v NullablePostCreateTweetRequestVariablesMedia) Get() *PostCreateTweetRequestVariablesMedia {
	return v.value
}

func (v *NullablePostCreateTweetRequestVariablesMedia) Set(val *PostCreateTweetRequestVariablesMedia) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCreateTweetRequestVariablesMedia) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCreateTweetRequestVariablesMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCreateTweetRequestVariablesMedia(val *PostCreateTweetRequestVariablesMedia) *NullablePostCreateTweetRequestVariablesMedia {
	return &NullablePostCreateTweetRequestVariablesMedia{value: val, isSet: true}
}

func (v NullablePostCreateTweetRequestVariablesMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCreateTweetRequestVariablesMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


