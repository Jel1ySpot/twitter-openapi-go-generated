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

// checks if the PostCreateTweetRequestVariables type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCreateTweetRequestVariables{}

// PostCreateTweetRequestVariables struct for PostCreateTweetRequestVariables
type PostCreateTweetRequestVariables struct {
	AttachmentUrl          *string                                             `json:"attachment_url,omitempty"`
	ConversationControl    *PostCreateTweetRequestVariablesConversationControl `json:"conversation_control,omitempty"`
	DarkRequest            bool                                                `json:"dark_request"`
	DisallowedReplyOptions map[string]interface{}                              `json:"disallowed_reply_options,omitempty"`
	Media                  PostCreateTweetRequestVariablesMedia                `json:"media"`
	Reply                  *PostCreateTweetRequestVariablesReply               `json:"reply,omitempty"`
	SemanticAnnotationIds  []map[string]interface{}                            `json:"semantic_annotation_ids"`
	TweetText              string                                              `json:"tweet_text"`
}

type _PostCreateTweetRequestVariables PostCreateTweetRequestVariables

// NewPostCreateTweetRequestVariables instantiates a new PostCreateTweetRequestVariables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCreateTweetRequestVariables(darkRequest bool, media PostCreateTweetRequestVariablesMedia, semanticAnnotationIds []map[string]interface{}, tweetText string) *PostCreateTweetRequestVariables {
	this := PostCreateTweetRequestVariables{}
	var attachmentUrl string = "https://x.com/elonmusk/status/1349129669258448897"
	this.AttachmentUrl = &attachmentUrl
	this.DarkRequest = darkRequest
	this.Media = media
	this.SemanticAnnotationIds = semanticAnnotationIds
	this.TweetText = tweetText
	return &this
}

// NewPostCreateTweetRequestVariablesWithDefaults instantiates a new PostCreateTweetRequestVariables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCreateTweetRequestVariablesWithDefaults() *PostCreateTweetRequestVariables {
	this := PostCreateTweetRequestVariables{}
	var attachmentUrl string = "https://x.com/elonmusk/status/1349129669258448897"
	this.AttachmentUrl = &attachmentUrl
	var darkRequest bool = false
	this.DarkRequest = darkRequest
	var tweetText string = "test"
	this.TweetText = tweetText
	return &this
}

// GetAttachmentUrl returns the AttachmentUrl field value if set, zero value otherwise.
func (o *PostCreateTweetRequestVariables) GetAttachmentUrl() string {
	if o == nil || IsNil(o.AttachmentUrl) {
		var ret string
		return ret
	}
	return *o.AttachmentUrl
}

// GetAttachmentUrlOk returns a tuple with the AttachmentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateTweetRequestVariables) GetAttachmentUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AttachmentUrl) {
		return nil, false
	}
	return o.AttachmentUrl, true
}

// HasAttachmentUrl returns a boolean if a field has been set.
func (o *PostCreateTweetRequestVariables) HasAttachmentUrl() bool {
	if o != nil && !IsNil(o.AttachmentUrl) {
		return true
	}

	return false
}

// SetAttachmentUrl gets a reference to the given string and assigns it to the AttachmentUrl field.
func (o *PostCreateTweetRequestVariables) SetAttachmentUrl(v string) {
	o.AttachmentUrl = &v
}

// GetConversationControl returns the ConversationControl field value if set, zero value otherwise.
func (o *PostCreateTweetRequestVariables) GetConversationControl() PostCreateTweetRequestVariablesConversationControl {
	if o == nil || IsNil(o.ConversationControl) {
		var ret PostCreateTweetRequestVariablesConversationControl
		return ret
	}
	return *o.ConversationControl
}

// GetConversationControlOk returns a tuple with the ConversationControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateTweetRequestVariables) GetConversationControlOk() (*PostCreateTweetRequestVariablesConversationControl, bool) {
	if o == nil || IsNil(o.ConversationControl) {
		return nil, false
	}
	return o.ConversationControl, true
}

// HasConversationControl returns a boolean if a field has been set.
func (o *PostCreateTweetRequestVariables) HasConversationControl() bool {
	if o != nil && !IsNil(o.ConversationControl) {
		return true
	}

	return false
}

// SetConversationControl gets a reference to the given PostCreateTweetRequestVariablesConversationControl and assigns it to the ConversationControl field.
func (o *PostCreateTweetRequestVariables) SetConversationControl(v PostCreateTweetRequestVariablesConversationControl) {
	o.ConversationControl = &v
}

// GetDarkRequest returns the DarkRequest field value
func (o *PostCreateTweetRequestVariables) GetDarkRequest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DarkRequest
}

// GetDarkRequestOk returns a tuple with the DarkRequest field value
// and a boolean to check if the value has been set.
func (o *PostCreateTweetRequestVariables) GetDarkRequestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DarkRequest, true
}

// SetDarkRequest sets field value
func (o *PostCreateTweetRequestVariables) SetDarkRequest(v bool) {
	o.DarkRequest = v
}

// GetDisallowedReplyOptions returns the DisallowedReplyOptions field value if set, zero value otherwise.
func (o *PostCreateTweetRequestVariables) GetDisallowedReplyOptions() map[string]interface{} {
	if o == nil || IsNil(o.DisallowedReplyOptions) {
		var ret map[string]interface{}
		return ret
	}
	return o.DisallowedReplyOptions
}

// GetDisallowedReplyOptionsOk returns a tuple with the DisallowedReplyOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateTweetRequestVariables) GetDisallowedReplyOptionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DisallowedReplyOptions) {
		return map[string]interface{}{}, false
	}
	return o.DisallowedReplyOptions, true
}

// HasDisallowedReplyOptions returns a boolean if a field has been set.
func (o *PostCreateTweetRequestVariables) HasDisallowedReplyOptions() bool {
	if o != nil && !IsNil(o.DisallowedReplyOptions) {
		return true
	}

	return false
}

// SetDisallowedReplyOptions gets a reference to the given map[string]interface{} and assigns it to the DisallowedReplyOptions field.
func (o *PostCreateTweetRequestVariables) SetDisallowedReplyOptions(v map[string]interface{}) {
	o.DisallowedReplyOptions = v
}

// GetMedia returns the Media field value
func (o *PostCreateTweetRequestVariables) GetMedia() PostCreateTweetRequestVariablesMedia {
	if o == nil {
		var ret PostCreateTweetRequestVariablesMedia
		return ret
	}

	return o.Media
}

// GetMediaOk returns a tuple with the Media field value
// and a boolean to check if the value has been set.
func (o *PostCreateTweetRequestVariables) GetMediaOk() (*PostCreateTweetRequestVariablesMedia, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Media, true
}

// SetMedia sets field value
func (o *PostCreateTweetRequestVariables) SetMedia(v PostCreateTweetRequestVariablesMedia) {
	o.Media = v
}

// GetReply returns the Reply field value if set, zero value otherwise.
func (o *PostCreateTweetRequestVariables) GetReply() PostCreateTweetRequestVariablesReply {
	if o == nil || IsNil(o.Reply) {
		var ret PostCreateTweetRequestVariablesReply
		return ret
	}
	return *o.Reply
}

// GetReplyOk returns a tuple with the Reply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateTweetRequestVariables) GetReplyOk() (*PostCreateTweetRequestVariablesReply, bool) {
	if o == nil || IsNil(o.Reply) {
		return nil, false
	}
	return o.Reply, true
}

// HasReply returns a boolean if a field has been set.
func (o *PostCreateTweetRequestVariables) HasReply() bool {
	if o != nil && !IsNil(o.Reply) {
		return true
	}

	return false
}

// SetReply gets a reference to the given PostCreateTweetRequestVariablesReply and assigns it to the Reply field.
func (o *PostCreateTweetRequestVariables) SetReply(v PostCreateTweetRequestVariablesReply) {
	o.Reply = &v
}

// GetSemanticAnnotationIds returns the SemanticAnnotationIds field value
func (o *PostCreateTweetRequestVariables) GetSemanticAnnotationIds() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.SemanticAnnotationIds
}

// GetSemanticAnnotationIdsOk returns a tuple with the SemanticAnnotationIds field value
// and a boolean to check if the value has been set.
func (o *PostCreateTweetRequestVariables) GetSemanticAnnotationIdsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.SemanticAnnotationIds, true
}

// SetSemanticAnnotationIds sets field value
func (o *PostCreateTweetRequestVariables) SetSemanticAnnotationIds(v []map[string]interface{}) {
	o.SemanticAnnotationIds = v
}

// GetTweetText returns the TweetText field value
func (o *PostCreateTweetRequestVariables) GetTweetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TweetText
}

// GetTweetTextOk returns a tuple with the TweetText field value
// and a boolean to check if the value has been set.
func (o *PostCreateTweetRequestVariables) GetTweetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TweetText, true
}

// SetTweetText sets field value
func (o *PostCreateTweetRequestVariables) SetTweetText(v string) {
	o.TweetText = v
}

func (o PostCreateTweetRequestVariables) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCreateTweetRequestVariables) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttachmentUrl) {
		toSerialize["attachment_url"] = o.AttachmentUrl
	}
	if !IsNil(o.ConversationControl) {
		toSerialize["conversation_control"] = o.ConversationControl
	}
	toSerialize["dark_request"] = o.DarkRequest
	if !IsNil(o.DisallowedReplyOptions) {
		toSerialize["disallowed_reply_options"] = o.DisallowedReplyOptions
	}
	toSerialize["media"] = o.Media
	if !IsNil(o.Reply) {
		toSerialize["reply"] = o.Reply
	}
	toSerialize["semantic_annotation_ids"] = o.SemanticAnnotationIds
	toSerialize["tweet_text"] = o.TweetText
	return toSerialize, nil
}

func (o *PostCreateTweetRequestVariables) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dark_request",
		"media",
		"semantic_annotation_ids",
		"tweet_text",
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

	varPostCreateTweetRequestVariables := _PostCreateTweetRequestVariables{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCreateTweetRequestVariables)

	if err != nil {
		return err
	}

	*o = PostCreateTweetRequestVariables(varPostCreateTweetRequestVariables)

	return err
}

type NullablePostCreateTweetRequestVariables struct {
	value *PostCreateTweetRequestVariables
	isSet bool
}

func (v NullablePostCreateTweetRequestVariables) Get() *PostCreateTweetRequestVariables {
	return v.value
}

func (v *NullablePostCreateTweetRequestVariables) Set(val *PostCreateTweetRequestVariables) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCreateTweetRequestVariables) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCreateTweetRequestVariables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCreateTweetRequestVariables(val *PostCreateTweetRequestVariables) *NullablePostCreateTweetRequestVariables {
	return &NullablePostCreateTweetRequestVariables{value: val, isSet: true}
}

func (v NullablePostCreateTweetRequestVariables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCreateTweetRequestVariables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
