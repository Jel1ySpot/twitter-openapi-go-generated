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

// checks if the ArticleLifecycleState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArticleLifecycleState{}

// ArticleLifecycleState struct for ArticleLifecycleState
type ArticleLifecycleState struct {
	ModifiedAtSecs int32 `json:"modified_at_secs"`
}

type _ArticleLifecycleState ArticleLifecycleState

// NewArticleLifecycleState instantiates a new ArticleLifecycleState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArticleLifecycleState(modifiedAtSecs int32) *ArticleLifecycleState {
	this := ArticleLifecycleState{}
	this.ModifiedAtSecs = modifiedAtSecs
	return &this
}

// NewArticleLifecycleStateWithDefaults instantiates a new ArticleLifecycleState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArticleLifecycleStateWithDefaults() *ArticleLifecycleState {
	this := ArticleLifecycleState{}
	return &this
}

// GetModifiedAtSecs returns the ModifiedAtSecs field value
func (o *ArticleLifecycleState) GetModifiedAtSecs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ModifiedAtSecs
}

// GetModifiedAtSecsOk returns a tuple with the ModifiedAtSecs field value
// and a boolean to check if the value has been set.
func (o *ArticleLifecycleState) GetModifiedAtSecsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAtSecs, true
}

// SetModifiedAtSecs sets field value
func (o *ArticleLifecycleState) SetModifiedAtSecs(v int32) {
	o.ModifiedAtSecs = v
}

func (o ArticleLifecycleState) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArticleLifecycleState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["modified_at_secs"] = o.ModifiedAtSecs
	return toSerialize, nil
}

func (o *ArticleLifecycleState) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"modified_at_secs",
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

	varArticleLifecycleState := _ArticleLifecycleState{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArticleLifecycleState)

	if err != nil {
		return err
	}

	*o = ArticleLifecycleState(varArticleLifecycleState)

	return err
}

type NullableArticleLifecycleState struct {
	value *ArticleLifecycleState
	isSet bool
}

func (v NullableArticleLifecycleState) Get() *ArticleLifecycleState {
	return v.value
}

func (v *NullableArticleLifecycleState) Set(val *ArticleLifecycleState) {
	v.value = val
	v.isSet = true
}

func (v NullableArticleLifecycleState) IsSet() bool {
	return v.isSet
}

func (v *NullableArticleLifecycleState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArticleLifecycleState(val *ArticleLifecycleState) *NullableArticleLifecycleState {
	return &NullableArticleLifecycleState{value: val, isSet: true}
}

func (v NullableArticleLifecycleState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArticleLifecycleState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
