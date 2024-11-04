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
)

// checks if the TweetEditControl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetEditControl{}

// TweetEditControl struct for TweetEditControl
type TweetEditControl struct {
	EditControlInitial *TweetEditControlInitial `json:"edit_control_initial,omitempty"`
	EditTweetIds []string `json:"edit_tweet_ids,omitempty"`
	EditableUntilMsecs *string `json:"editable_until_msecs,omitempty" validate:"regexp=^[0-9]+$"`
	EditsRemaining *string `json:"edits_remaining,omitempty" validate:"regexp=^[0-9]+$"`
	InitialTweetId *string `json:"initial_tweet_id,omitempty" validate:"regexp=^[0-9]+$"`
	IsEditEligible *bool `json:"is_edit_eligible,omitempty"`
}

// NewTweetEditControl instantiates a new TweetEditControl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetEditControl() *TweetEditControl {
	this := TweetEditControl{}
	return &this
}

// NewTweetEditControlWithDefaults instantiates a new TweetEditControl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetEditControlWithDefaults() *TweetEditControl {
	this := TweetEditControl{}
	return &this
}

// GetEditControlInitial returns the EditControlInitial field value if set, zero value otherwise.
func (o *TweetEditControl) GetEditControlInitial() TweetEditControlInitial {
	if o == nil || IsNil(o.EditControlInitial) {
		var ret TweetEditControlInitial
		return ret
	}
	return *o.EditControlInitial
}

// GetEditControlInitialOk returns a tuple with the EditControlInitial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetEditControl) GetEditControlInitialOk() (*TweetEditControlInitial, bool) {
	if o == nil || IsNil(o.EditControlInitial) {
		return nil, false
	}
	return o.EditControlInitial, true
}

// HasEditControlInitial returns a boolean if a field has been set.
func (o *TweetEditControl) HasEditControlInitial() bool {
	if o != nil && !IsNil(o.EditControlInitial) {
		return true
	}

	return false
}

// SetEditControlInitial gets a reference to the given TweetEditControlInitial and assigns it to the EditControlInitial field.
func (o *TweetEditControl) SetEditControlInitial(v TweetEditControlInitial) {
	o.EditControlInitial = &v
}

// GetEditTweetIds returns the EditTweetIds field value if set, zero value otherwise.
func (o *TweetEditControl) GetEditTweetIds() []string {
	if o == nil || IsNil(o.EditTweetIds) {
		var ret []string
		return ret
	}
	return o.EditTweetIds
}

// GetEditTweetIdsOk returns a tuple with the EditTweetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetEditControl) GetEditTweetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.EditTweetIds) {
		return nil, false
	}
	return o.EditTweetIds, true
}

// HasEditTweetIds returns a boolean if a field has been set.
func (o *TweetEditControl) HasEditTweetIds() bool {
	if o != nil && !IsNil(o.EditTweetIds) {
		return true
	}

	return false
}

// SetEditTweetIds gets a reference to the given []string and assigns it to the EditTweetIds field.
func (o *TweetEditControl) SetEditTweetIds(v []string) {
	o.EditTweetIds = v
}

// GetEditableUntilMsecs returns the EditableUntilMsecs field value if set, zero value otherwise.
func (o *TweetEditControl) GetEditableUntilMsecs() string {
	if o == nil || IsNil(o.EditableUntilMsecs) {
		var ret string
		return ret
	}
	return *o.EditableUntilMsecs
}

// GetEditableUntilMsecsOk returns a tuple with the EditableUntilMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetEditControl) GetEditableUntilMsecsOk() (*string, bool) {
	if o == nil || IsNil(o.EditableUntilMsecs) {
		return nil, false
	}
	return o.EditableUntilMsecs, true
}

// HasEditableUntilMsecs returns a boolean if a field has been set.
func (o *TweetEditControl) HasEditableUntilMsecs() bool {
	if o != nil && !IsNil(o.EditableUntilMsecs) {
		return true
	}

	return false
}

// SetEditableUntilMsecs gets a reference to the given string and assigns it to the EditableUntilMsecs field.
func (o *TweetEditControl) SetEditableUntilMsecs(v string) {
	o.EditableUntilMsecs = &v
}

// GetEditsRemaining returns the EditsRemaining field value if set, zero value otherwise.
func (o *TweetEditControl) GetEditsRemaining() string {
	if o == nil || IsNil(o.EditsRemaining) {
		var ret string
		return ret
	}
	return *o.EditsRemaining
}

// GetEditsRemainingOk returns a tuple with the EditsRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetEditControl) GetEditsRemainingOk() (*string, bool) {
	if o == nil || IsNil(o.EditsRemaining) {
		return nil, false
	}
	return o.EditsRemaining, true
}

// HasEditsRemaining returns a boolean if a field has been set.
func (o *TweetEditControl) HasEditsRemaining() bool {
	if o != nil && !IsNil(o.EditsRemaining) {
		return true
	}

	return false
}

// SetEditsRemaining gets a reference to the given string and assigns it to the EditsRemaining field.
func (o *TweetEditControl) SetEditsRemaining(v string) {
	o.EditsRemaining = &v
}

// GetInitialTweetId returns the InitialTweetId field value if set, zero value otherwise.
func (o *TweetEditControl) GetInitialTweetId() string {
	if o == nil || IsNil(o.InitialTweetId) {
		var ret string
		return ret
	}
	return *o.InitialTweetId
}

// GetInitialTweetIdOk returns a tuple with the InitialTweetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetEditControl) GetInitialTweetIdOk() (*string, bool) {
	if o == nil || IsNil(o.InitialTweetId) {
		return nil, false
	}
	return o.InitialTweetId, true
}

// HasInitialTweetId returns a boolean if a field has been set.
func (o *TweetEditControl) HasInitialTweetId() bool {
	if o != nil && !IsNil(o.InitialTweetId) {
		return true
	}

	return false
}

// SetInitialTweetId gets a reference to the given string and assigns it to the InitialTweetId field.
func (o *TweetEditControl) SetInitialTweetId(v string) {
	o.InitialTweetId = &v
}

// GetIsEditEligible returns the IsEditEligible field value if set, zero value otherwise.
func (o *TweetEditControl) GetIsEditEligible() bool {
	if o == nil || IsNil(o.IsEditEligible) {
		var ret bool
		return ret
	}
	return *o.IsEditEligible
}

// GetIsEditEligibleOk returns a tuple with the IsEditEligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetEditControl) GetIsEditEligibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEditEligible) {
		return nil, false
	}
	return o.IsEditEligible, true
}

// HasIsEditEligible returns a boolean if a field has been set.
func (o *TweetEditControl) HasIsEditEligible() bool {
	if o != nil && !IsNil(o.IsEditEligible) {
		return true
	}

	return false
}

// SetIsEditEligible gets a reference to the given bool and assigns it to the IsEditEligible field.
func (o *TweetEditControl) SetIsEditEligible(v bool) {
	o.IsEditEligible = &v
}

func (o TweetEditControl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetEditControl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EditControlInitial) {
		toSerialize["edit_control_initial"] = o.EditControlInitial
	}
	if !IsNil(o.EditTweetIds) {
		toSerialize["edit_tweet_ids"] = o.EditTweetIds
	}
	if !IsNil(o.EditableUntilMsecs) {
		toSerialize["editable_until_msecs"] = o.EditableUntilMsecs
	}
	if !IsNil(o.EditsRemaining) {
		toSerialize["edits_remaining"] = o.EditsRemaining
	}
	if !IsNil(o.InitialTweetId) {
		toSerialize["initial_tweet_id"] = o.InitialTweetId
	}
	if !IsNil(o.IsEditEligible) {
		toSerialize["is_edit_eligible"] = o.IsEditEligible
	}
	return toSerialize, nil
}

type NullableTweetEditControl struct {
	value *TweetEditControl
	isSet bool
}

func (v NullableTweetEditControl) Get() *TweetEditControl {
	return v.value
}

func (v *NullableTweetEditControl) Set(val *TweetEditControl) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetEditControl) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetEditControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetEditControl(val *TweetEditControl) *NullableTweetEditControl {
	return &NullableTweetEditControl{value: val, isSet: true}
}

func (v NullableTweetEditControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetEditControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


