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

// checks if the UserVerificationInfoReasonDescriptionEntitiesRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserVerificationInfoReasonDescriptionEntitiesRef{}

// UserVerificationInfoReasonDescriptionEntitiesRef struct for UserVerificationInfoReasonDescriptionEntitiesRef
type UserVerificationInfoReasonDescriptionEntitiesRef struct {
	Url     string `json:"url"`
	UrlType string `json:"url_type"`
}

type _UserVerificationInfoReasonDescriptionEntitiesRef UserVerificationInfoReasonDescriptionEntitiesRef

// NewUserVerificationInfoReasonDescriptionEntitiesRef instantiates a new UserVerificationInfoReasonDescriptionEntitiesRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserVerificationInfoReasonDescriptionEntitiesRef(url string, urlType string) *UserVerificationInfoReasonDescriptionEntitiesRef {
	this := UserVerificationInfoReasonDescriptionEntitiesRef{}
	this.Url = url
	this.UrlType = urlType
	return &this
}

// NewUserVerificationInfoReasonDescriptionEntitiesRefWithDefaults instantiates a new UserVerificationInfoReasonDescriptionEntitiesRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserVerificationInfoReasonDescriptionEntitiesRefWithDefaults() *UserVerificationInfoReasonDescriptionEntitiesRef {
	this := UserVerificationInfoReasonDescriptionEntitiesRef{}
	return &this
}

// GetUrl returns the Url field value
func (o *UserVerificationInfoReasonDescriptionEntitiesRef) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *UserVerificationInfoReasonDescriptionEntitiesRef) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *UserVerificationInfoReasonDescriptionEntitiesRef) SetUrl(v string) {
	o.Url = v
}

// GetUrlType returns the UrlType field value
func (o *UserVerificationInfoReasonDescriptionEntitiesRef) GetUrlType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UrlType
}

// GetUrlTypeOk returns a tuple with the UrlType field value
// and a boolean to check if the value has been set.
func (o *UserVerificationInfoReasonDescriptionEntitiesRef) GetUrlTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UrlType, true
}

// SetUrlType sets field value
func (o *UserVerificationInfoReasonDescriptionEntitiesRef) SetUrlType(v string) {
	o.UrlType = v
}

func (o UserVerificationInfoReasonDescriptionEntitiesRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserVerificationInfoReasonDescriptionEntitiesRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["url_type"] = o.UrlType
	return toSerialize, nil
}

func (o *UserVerificationInfoReasonDescriptionEntitiesRef) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"url_type",
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

	varUserVerificationInfoReasonDescriptionEntitiesRef := _UserVerificationInfoReasonDescriptionEntitiesRef{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserVerificationInfoReasonDescriptionEntitiesRef)

	if err != nil {
		return err
	}

	*o = UserVerificationInfoReasonDescriptionEntitiesRef(varUserVerificationInfoReasonDescriptionEntitiesRef)

	return err
}

type NullableUserVerificationInfoReasonDescriptionEntitiesRef struct {
	value *UserVerificationInfoReasonDescriptionEntitiesRef
	isSet bool
}

func (v NullableUserVerificationInfoReasonDescriptionEntitiesRef) Get() *UserVerificationInfoReasonDescriptionEntitiesRef {
	return v.value
}

func (v *NullableUserVerificationInfoReasonDescriptionEntitiesRef) Set(val *UserVerificationInfoReasonDescriptionEntitiesRef) {
	v.value = val
	v.isSet = true
}

func (v NullableUserVerificationInfoReasonDescriptionEntitiesRef) IsSet() bool {
	return v.isSet
}

func (v *NullableUserVerificationInfoReasonDescriptionEntitiesRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserVerificationInfoReasonDescriptionEntitiesRef(val *UserVerificationInfoReasonDescriptionEntitiesRef) *NullableUserVerificationInfoReasonDescriptionEntitiesRef {
	return &NullableUserVerificationInfoReasonDescriptionEntitiesRef{value: val, isSet: true}
}

func (v NullableUserVerificationInfoReasonDescriptionEntitiesRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserVerificationInfoReasonDescriptionEntitiesRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
