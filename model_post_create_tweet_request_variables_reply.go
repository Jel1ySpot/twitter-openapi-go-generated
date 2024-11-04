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

// checks if the PostCreateTweetRequestVariablesReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCreateTweetRequestVariablesReply{}

// PostCreateTweetRequestVariablesReply struct for PostCreateTweetRequestVariablesReply
type PostCreateTweetRequestVariablesReply struct {
	ExcludeReplyUserIds []map[string]interface{} `json:"exclude_reply_user_ids"`
	InReplyToTweetId    string                   `json:"in_reply_to_tweet_id"`
}

type _PostCreateTweetRequestVariablesReply PostCreateTweetRequestVariablesReply

// NewPostCreateTweetRequestVariablesReply instantiates a new PostCreateTweetRequestVariablesReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCreateTweetRequestVariablesReply(excludeReplyUserIds []map[string]interface{}, inReplyToTweetId string) *PostCreateTweetRequestVariablesReply {
	this := PostCreateTweetRequestVariablesReply{}
	this.ExcludeReplyUserIds = excludeReplyUserIds
	this.InReplyToTweetId = inReplyToTweetId
	return &this
}

// NewPostCreateTweetRequestVariablesReplyWithDefaults instantiates a new PostCreateTweetRequestVariablesReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCreateTweetRequestVariablesReplyWithDefaults() *PostCreateTweetRequestVariablesReply {
	this := PostCreateTweetRequestVariablesReply{}
	var inReplyToTweetId string = "1111111111111111111"
	this.InReplyToTweetId = inReplyToTweetId
	return &this
}

// GetExcludeReplyUserIds returns the ExcludeReplyUserIds field value
func (o *PostCreateTweetRequestVariablesReply) GetExcludeReplyUserIds() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.ExcludeReplyUserIds
}

// GetExcludeReplyUserIdsOk returns a tuple with the ExcludeReplyUserIds field value
// and a boolean to check if the value has been set.
func (o *PostCreateTweetRequestVariablesReply) GetExcludeReplyUserIdsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExcludeReplyUserIds, true
}

// SetExcludeReplyUserIds sets field value
func (o *PostCreateTweetRequestVariablesReply) SetExcludeReplyUserIds(v []map[string]interface{}) {
	o.ExcludeReplyUserIds = v
}

// GetInReplyToTweetId returns the InReplyToTweetId field value
func (o *PostCreateTweetRequestVariablesReply) GetInReplyToTweetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InReplyToTweetId
}

// GetInReplyToTweetIdOk returns a tuple with the InReplyToTweetId field value
// and a boolean to check if the value has been set.
func (o *PostCreateTweetRequestVariablesReply) GetInReplyToTweetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InReplyToTweetId, true
}

// SetInReplyToTweetId sets field value
func (o *PostCreateTweetRequestVariablesReply) SetInReplyToTweetId(v string) {
	o.InReplyToTweetId = v
}

func (o PostCreateTweetRequestVariablesReply) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCreateTweetRequestVariablesReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exclude_reply_user_ids"] = o.ExcludeReplyUserIds
	toSerialize["in_reply_to_tweet_id"] = o.InReplyToTweetId
	return toSerialize, nil
}

func (o *PostCreateTweetRequestVariablesReply) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"exclude_reply_user_ids",
		"in_reply_to_tweet_id",
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

	varPostCreateTweetRequestVariablesReply := _PostCreateTweetRequestVariablesReply{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCreateTweetRequestVariablesReply)

	if err != nil {
		return err
	}

	*o = PostCreateTweetRequestVariablesReply(varPostCreateTweetRequestVariablesReply)

	return err
}

type NullablePostCreateTweetRequestVariablesReply struct {
	value *PostCreateTweetRequestVariablesReply
	isSet bool
}

func (v NullablePostCreateTweetRequestVariablesReply) Get() *PostCreateTweetRequestVariablesReply {
	return v.value
}

func (v *NullablePostCreateTweetRequestVariablesReply) Set(val *PostCreateTweetRequestVariablesReply) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCreateTweetRequestVariablesReply) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCreateTweetRequestVariablesReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCreateTweetRequestVariablesReply(val *PostCreateTweetRequestVariablesReply) *NullablePostCreateTweetRequestVariablesReply {
	return &NullablePostCreateTweetRequestVariablesReply{value: val, isSet: true}
}

func (v NullablePostCreateTweetRequestVariablesReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCreateTweetRequestVariablesReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
