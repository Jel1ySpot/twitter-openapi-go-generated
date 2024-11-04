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

// checks if the TweetCardLegacy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetCardLegacy{}

// TweetCardLegacy struct for TweetCardLegacy
type TweetCardLegacy struct {
	BindingValues []TweetCardLegacyBindingValue `json:"binding_values"`
	CardPlatform *TweetCardPlatformData `json:"card_platform,omitempty"`
	Name string `json:"name"`
	Url string `json:"url"`
	UserRefsResults []UserResults `json:"user_refs_results,omitempty"`
}

type _TweetCardLegacy TweetCardLegacy

// NewTweetCardLegacy instantiates a new TweetCardLegacy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetCardLegacy(bindingValues []TweetCardLegacyBindingValue, name string, url string) *TweetCardLegacy {
	this := TweetCardLegacy{}
	this.BindingValues = bindingValues
	this.Name = name
	this.Url = url
	return &this
}

// NewTweetCardLegacyWithDefaults instantiates a new TweetCardLegacy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetCardLegacyWithDefaults() *TweetCardLegacy {
	this := TweetCardLegacy{}
	return &this
}

// GetBindingValues returns the BindingValues field value
func (o *TweetCardLegacy) GetBindingValues() []TweetCardLegacyBindingValue {
	if o == nil {
		var ret []TweetCardLegacyBindingValue
		return ret
	}

	return o.BindingValues
}

// GetBindingValuesOk returns a tuple with the BindingValues field value
// and a boolean to check if the value has been set.
func (o *TweetCardLegacy) GetBindingValuesOk() ([]TweetCardLegacyBindingValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.BindingValues, true
}

// SetBindingValues sets field value
func (o *TweetCardLegacy) SetBindingValues(v []TweetCardLegacyBindingValue) {
	o.BindingValues = v
}

// GetCardPlatform returns the CardPlatform field value if set, zero value otherwise.
func (o *TweetCardLegacy) GetCardPlatform() TweetCardPlatformData {
	if o == nil || IsNil(o.CardPlatform) {
		var ret TweetCardPlatformData
		return ret
	}
	return *o.CardPlatform
}

// GetCardPlatformOk returns a tuple with the CardPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetCardLegacy) GetCardPlatformOk() (*TweetCardPlatformData, bool) {
	if o == nil || IsNil(o.CardPlatform) {
		return nil, false
	}
	return o.CardPlatform, true
}

// HasCardPlatform returns a boolean if a field has been set.
func (o *TweetCardLegacy) HasCardPlatform() bool {
	if o != nil && !IsNil(o.CardPlatform) {
		return true
	}

	return false
}

// SetCardPlatform gets a reference to the given TweetCardPlatformData and assigns it to the CardPlatform field.
func (o *TweetCardLegacy) SetCardPlatform(v TweetCardPlatformData) {
	o.CardPlatform = &v
}

// GetName returns the Name field value
func (o *TweetCardLegacy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TweetCardLegacy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TweetCardLegacy) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *TweetCardLegacy) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TweetCardLegacy) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TweetCardLegacy) SetUrl(v string) {
	o.Url = v
}

// GetUserRefsResults returns the UserRefsResults field value if set, zero value otherwise.
func (o *TweetCardLegacy) GetUserRefsResults() []UserResults {
	if o == nil || IsNil(o.UserRefsResults) {
		var ret []UserResults
		return ret
	}
	return o.UserRefsResults
}

// GetUserRefsResultsOk returns a tuple with the UserRefsResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetCardLegacy) GetUserRefsResultsOk() ([]UserResults, bool) {
	if o == nil || IsNil(o.UserRefsResults) {
		return nil, false
	}
	return o.UserRefsResults, true
}

// HasUserRefsResults returns a boolean if a field has been set.
func (o *TweetCardLegacy) HasUserRefsResults() bool {
	if o != nil && !IsNil(o.UserRefsResults) {
		return true
	}

	return false
}

// SetUserRefsResults gets a reference to the given []UserResults and assigns it to the UserRefsResults field.
func (o *TweetCardLegacy) SetUserRefsResults(v []UserResults) {
	o.UserRefsResults = v
}

func (o TweetCardLegacy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetCardLegacy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["binding_values"] = o.BindingValues
	if !IsNil(o.CardPlatform) {
		toSerialize["card_platform"] = o.CardPlatform
	}
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	if !IsNil(o.UserRefsResults) {
		toSerialize["user_refs_results"] = o.UserRefsResults
	}
	return toSerialize, nil
}

func (o *TweetCardLegacy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"binding_values",
		"name",
		"url",
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

	varTweetCardLegacy := _TweetCardLegacy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTweetCardLegacy)

	if err != nil {
		return err
	}

	*o = TweetCardLegacy(varTweetCardLegacy)

	return err
}

type NullableTweetCardLegacy struct {
	value *TweetCardLegacy
	isSet bool
}

func (v NullableTweetCardLegacy) Get() *TweetCardLegacy {
	return v.value
}

func (v *NullableTweetCardLegacy) Set(val *TweetCardLegacy) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetCardLegacy) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetCardLegacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetCardLegacy(val *TweetCardLegacy) *NullableTweetCardLegacy {
	return &NullableTweetCardLegacy{value: val, isSet: true}
}

func (v NullableTweetCardLegacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetCardLegacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

