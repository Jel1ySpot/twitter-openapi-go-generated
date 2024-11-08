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

// checks if the CommunityRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunityRelationship{}

// CommunityRelationship struct for CommunityRelationship
type CommunityRelationship struct {
	Actions         CommunityActions       `json:"actions"`
	Id              string                 `json:"id"`
	ModerationState map[string]interface{} `json:"moderation_state"`
	RestId          string                 `json:"rest_id" validate:"regexp=^[0-9]+$"`
}

type _CommunityRelationship CommunityRelationship

// NewCommunityRelationship instantiates a new CommunityRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunityRelationship(actions CommunityActions, id string, moderationState map[string]interface{}, restId string) *CommunityRelationship {
	this := CommunityRelationship{}
	this.Actions = actions
	this.Id = id
	this.ModerationState = moderationState
	this.RestId = restId
	return &this
}

// NewCommunityRelationshipWithDefaults instantiates a new CommunityRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunityRelationshipWithDefaults() *CommunityRelationship {
	this := CommunityRelationship{}
	return &this
}

// GetActions returns the Actions field value
func (o *CommunityRelationship) GetActions() CommunityActions {
	if o == nil {
		var ret CommunityActions
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *CommunityRelationship) GetActionsOk() (*CommunityActions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actions, true
}

// SetActions sets field value
func (o *CommunityRelationship) SetActions(v CommunityActions) {
	o.Actions = v
}

// GetId returns the Id field value
func (o *CommunityRelationship) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CommunityRelationship) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CommunityRelationship) SetId(v string) {
	o.Id = v
}

// GetModerationState returns the ModerationState field value
func (o *CommunityRelationship) GetModerationState() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ModerationState
}

// GetModerationStateOk returns a tuple with the ModerationState field value
// and a boolean to check if the value has been set.
func (o *CommunityRelationship) GetModerationStateOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ModerationState, true
}

// SetModerationState sets field value
func (o *CommunityRelationship) SetModerationState(v map[string]interface{}) {
	o.ModerationState = v
}

// GetRestId returns the RestId field value
func (o *CommunityRelationship) GetRestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RestId
}

// GetRestIdOk returns a tuple with the RestId field value
// and a boolean to check if the value has been set.
func (o *CommunityRelationship) GetRestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestId, true
}

// SetRestId sets field value
func (o *CommunityRelationship) SetRestId(v string) {
	o.RestId = v
}

func (o CommunityRelationship) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunityRelationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actions"] = o.Actions
	toSerialize["id"] = o.Id
	toSerialize["moderation_state"] = o.ModerationState
	toSerialize["rest_id"] = o.RestId
	return toSerialize, nil
}

func (o *CommunityRelationship) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"actions",
		"id",
		"moderation_state",
		"rest_id",
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

	varCommunityRelationship := _CommunityRelationship{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommunityRelationship)

	if err != nil {
		return err
	}

	*o = CommunityRelationship(varCommunityRelationship)

	return err
}

type NullableCommunityRelationship struct {
	value *CommunityRelationship
	isSet bool
}

func (v NullableCommunityRelationship) Get() *CommunityRelationship {
	return v.value
}

func (v *NullableCommunityRelationship) Set(val *CommunityRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunityRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunityRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunityRelationship(val *CommunityRelationship) *NullableCommunityRelationship {
	return &NullableCommunityRelationship{value: val, isSet: true}
}

func (v NullableCommunityRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunityRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
