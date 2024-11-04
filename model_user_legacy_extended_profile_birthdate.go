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

// checks if the UserLegacyExtendedProfileBirthdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserLegacyExtendedProfileBirthdate{}

// UserLegacyExtendedProfileBirthdate struct for UserLegacyExtendedProfileBirthdate
type UserLegacyExtendedProfileBirthdate struct {
	Day int32 `json:"day"`
	Month int32 `json:"month"`
	Visibility string `json:"visibility"`
	Year int32 `json:"year"`
	YearVisibility string `json:"year_visibility"`
}

type _UserLegacyExtendedProfileBirthdate UserLegacyExtendedProfileBirthdate

// NewUserLegacyExtendedProfileBirthdate instantiates a new UserLegacyExtendedProfileBirthdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLegacyExtendedProfileBirthdate(day int32, month int32, visibility string, year int32, yearVisibility string) *UserLegacyExtendedProfileBirthdate {
	this := UserLegacyExtendedProfileBirthdate{}
	this.Day = day
	this.Month = month
	this.Visibility = visibility
	this.Year = year
	this.YearVisibility = yearVisibility
	return &this
}

// NewUserLegacyExtendedProfileBirthdateWithDefaults instantiates a new UserLegacyExtendedProfileBirthdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLegacyExtendedProfileBirthdateWithDefaults() *UserLegacyExtendedProfileBirthdate {
	this := UserLegacyExtendedProfileBirthdate{}
	return &this
}

// GetDay returns the Day field value
func (o *UserLegacyExtendedProfileBirthdate) GetDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *UserLegacyExtendedProfileBirthdate) GetDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value
func (o *UserLegacyExtendedProfileBirthdate) SetDay(v int32) {
	o.Day = v
}

// GetMonth returns the Month field value
func (o *UserLegacyExtendedProfileBirthdate) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *UserLegacyExtendedProfileBirthdate) GetMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *UserLegacyExtendedProfileBirthdate) SetMonth(v int32) {
	o.Month = v
}

// GetVisibility returns the Visibility field value
func (o *UserLegacyExtendedProfileBirthdate) GetVisibility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *UserLegacyExtendedProfileBirthdate) GetVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *UserLegacyExtendedProfileBirthdate) SetVisibility(v string) {
	o.Visibility = v
}

// GetYear returns the Year field value
func (o *UserLegacyExtendedProfileBirthdate) GetYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *UserLegacyExtendedProfileBirthdate) GetYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *UserLegacyExtendedProfileBirthdate) SetYear(v int32) {
	o.Year = v
}

// GetYearVisibility returns the YearVisibility field value
func (o *UserLegacyExtendedProfileBirthdate) GetYearVisibility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YearVisibility
}

// GetYearVisibilityOk returns a tuple with the YearVisibility field value
// and a boolean to check if the value has been set.
func (o *UserLegacyExtendedProfileBirthdate) GetYearVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.YearVisibility, true
}

// SetYearVisibility sets field value
func (o *UserLegacyExtendedProfileBirthdate) SetYearVisibility(v string) {
	o.YearVisibility = v
}

func (o UserLegacyExtendedProfileBirthdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserLegacyExtendedProfileBirthdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["day"] = o.Day
	toSerialize["month"] = o.Month
	toSerialize["visibility"] = o.Visibility
	toSerialize["year"] = o.Year
	toSerialize["year_visibility"] = o.YearVisibility
	return toSerialize, nil
}

func (o *UserLegacyExtendedProfileBirthdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"day",
		"month",
		"visibility",
		"year",
		"year_visibility",
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

	varUserLegacyExtendedProfileBirthdate := _UserLegacyExtendedProfileBirthdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserLegacyExtendedProfileBirthdate)

	if err != nil {
		return err
	}

	*o = UserLegacyExtendedProfileBirthdate(varUserLegacyExtendedProfileBirthdate)

	return err
}

type NullableUserLegacyExtendedProfileBirthdate struct {
	value *UserLegacyExtendedProfileBirthdate
	isSet bool
}

func (v NullableUserLegacyExtendedProfileBirthdate) Get() *UserLegacyExtendedProfileBirthdate {
	return v.value
}

func (v *NullableUserLegacyExtendedProfileBirthdate) Set(val *UserLegacyExtendedProfileBirthdate) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLegacyExtendedProfileBirthdate) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLegacyExtendedProfileBirthdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLegacyExtendedProfileBirthdate(val *UserLegacyExtendedProfileBirthdate) *NullableUserLegacyExtendedProfileBirthdate {
	return &NullableUserLegacyExtendedProfileBirthdate{value: val, isSet: true}
}

func (v NullableUserLegacyExtendedProfileBirthdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLegacyExtendedProfileBirthdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


