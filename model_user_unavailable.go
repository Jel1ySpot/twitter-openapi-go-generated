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

// checks if the UserUnavailable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserUnavailable{}

// UserUnavailable struct for UserUnavailable
type UserUnavailable struct {
	Typename TypeName `json:"__typename"`
	Message *string `json:"message,omitempty"`
	Reason string `json:"reason"`
}

type _UserUnavailable UserUnavailable

// NewUserUnavailable instantiates a new UserUnavailable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUnavailable(typename TypeName, reason string) *UserUnavailable {
	this := UserUnavailable{}
	this.Typename = typename
	this.Reason = reason
	return &this
}

// NewUserUnavailableWithDefaults instantiates a new UserUnavailable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUnavailableWithDefaults() *UserUnavailable {
	this := UserUnavailable{}
	return &this
}

// GetTypename returns the Typename field value
func (o *UserUnavailable) GetTypename() TypeName {
	if o == nil {
		var ret TypeName
		return ret
	}

	return o.Typename
}

// GetTypenameOk returns a tuple with the Typename field value
// and a boolean to check if the value has been set.
func (o *UserUnavailable) GetTypenameOk() (*TypeName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Typename, true
}

// SetTypename sets field value
func (o *UserUnavailable) SetTypename(v TypeName) {
	o.Typename = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UserUnavailable) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUnavailable) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UserUnavailable) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UserUnavailable) SetMessage(v string) {
	o.Message = &v
}

// GetReason returns the Reason field value
func (o *UserUnavailable) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *UserUnavailable) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *UserUnavailable) SetReason(v string) {
	o.Reason = v
}

func (o UserUnavailable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUnavailable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["__typename"] = o.Typename
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	toSerialize["reason"] = o.Reason
	return toSerialize, nil
}

func (o *UserUnavailable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"__typename",
		"reason",
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

	varUserUnavailable := _UserUnavailable{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserUnavailable)

	if err != nil {
		return err
	}

	*o = UserUnavailable(varUserUnavailable)

	return err
}

type NullableUserUnavailable struct {
	value *UserUnavailable
	isSet bool
}

func (v NullableUserUnavailable) Get() *UserUnavailable {
	return v.value
}

func (v *NullableUserUnavailable) Set(val *UserUnavailable) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUnavailable) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUnavailable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUnavailable(val *UserUnavailable) *NullableUserUnavailable {
	return &NullableUserUnavailable{value: val, isSet: true}
}

func (v NullableUserUnavailable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUnavailable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

