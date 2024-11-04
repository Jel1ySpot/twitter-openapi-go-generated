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

// checks if the ArticleResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArticleResults{}

// ArticleResults struct for ArticleResults
type ArticleResults struct {
	Result ArticleResult `json:"result"`
}

type _ArticleResults ArticleResults

// NewArticleResults instantiates a new ArticleResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArticleResults(result ArticleResult) *ArticleResults {
	this := ArticleResults{}
	this.Result = result
	return &this
}

// NewArticleResultsWithDefaults instantiates a new ArticleResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArticleResultsWithDefaults() *ArticleResults {
	this := ArticleResults{}
	return &this
}

// GetResult returns the Result field value
func (o *ArticleResults) GetResult() ArticleResult {
	if o == nil {
		var ret ArticleResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ArticleResults) GetResultOk() (*ArticleResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *ArticleResults) SetResult(v ArticleResult) {
	o.Result = v
}

func (o ArticleResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArticleResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *ArticleResults) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
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

	varArticleResults := _ArticleResults{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArticleResults)

	if err != nil {
		return err
	}

	*o = ArticleResults(varArticleResults)

	return err
}

type NullableArticleResults struct {
	value *ArticleResults
	isSet bool
}

func (v NullableArticleResults) Get() *ArticleResults {
	return v.value
}

func (v *NullableArticleResults) Set(val *ArticleResults) {
	v.value = val
	v.isSet = true
}

func (v NullableArticleResults) IsSet() bool {
	return v.isSet
}

func (v *NullableArticleResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArticleResults(val *ArticleResults) *NullableArticleResults {
	return &NullableArticleResults{value: val, isSet: true}
}

func (v NullableArticleResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArticleResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


