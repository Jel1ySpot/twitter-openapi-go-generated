# PostCreateTweetRequestVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentUrl** | Pointer to **string** |  | [optional] [default to "https://x.com/elonmusk/status/1349129669258448897"]
**ConversationControl** | Pointer to [**PostCreateTweetRequestVariablesConversationControl**](PostCreateTweetRequestVariablesConversationControl.md) |  | [optional] 
**DarkRequest** | **bool** |  | [default to false]
**DisallowedReplyOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**Media** | [**PostCreateTweetRequestVariablesMedia**](PostCreateTweetRequestVariablesMedia.md) |  | 
**Reply** | Pointer to [**PostCreateTweetRequestVariablesReply**](PostCreateTweetRequestVariablesReply.md) |  | [optional] 
**SemanticAnnotationIds** | **[]map[string]interface{}** |  | 
**TweetText** | **string** |  | [default to "test"]

## Methods

### NewPostCreateTweetRequestVariables

`func NewPostCreateTweetRequestVariables(darkRequest bool, media PostCreateTweetRequestVariablesMedia, semanticAnnotationIds []map[string]interface{}, tweetText string, ) *PostCreateTweetRequestVariables`

NewPostCreateTweetRequestVariables instantiates a new PostCreateTweetRequestVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCreateTweetRequestVariablesWithDefaults

`func NewPostCreateTweetRequestVariablesWithDefaults() *PostCreateTweetRequestVariables`

NewPostCreateTweetRequestVariablesWithDefaults instantiates a new PostCreateTweetRequestVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentUrl

`func (o *PostCreateTweetRequestVariables) GetAttachmentUrl() string`

GetAttachmentUrl returns the AttachmentUrl field if non-nil, zero value otherwise.

### GetAttachmentUrlOk

`func (o *PostCreateTweetRequestVariables) GetAttachmentUrlOk() (*string, bool)`

GetAttachmentUrlOk returns a tuple with the AttachmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentUrl

`func (o *PostCreateTweetRequestVariables) SetAttachmentUrl(v string)`

SetAttachmentUrl sets AttachmentUrl field to given value.

### HasAttachmentUrl

`func (o *PostCreateTweetRequestVariables) HasAttachmentUrl() bool`

HasAttachmentUrl returns a boolean if a field has been set.

### GetConversationControl

`func (o *PostCreateTweetRequestVariables) GetConversationControl() PostCreateTweetRequestVariablesConversationControl`

GetConversationControl returns the ConversationControl field if non-nil, zero value otherwise.

### GetConversationControlOk

`func (o *PostCreateTweetRequestVariables) GetConversationControlOk() (*PostCreateTweetRequestVariablesConversationControl, bool)`

GetConversationControlOk returns a tuple with the ConversationControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationControl

`func (o *PostCreateTweetRequestVariables) SetConversationControl(v PostCreateTweetRequestVariablesConversationControl)`

SetConversationControl sets ConversationControl field to given value.

### HasConversationControl

`func (o *PostCreateTweetRequestVariables) HasConversationControl() bool`

HasConversationControl returns a boolean if a field has been set.

### GetDarkRequest

`func (o *PostCreateTweetRequestVariables) GetDarkRequest() bool`

GetDarkRequest returns the DarkRequest field if non-nil, zero value otherwise.

### GetDarkRequestOk

`func (o *PostCreateTweetRequestVariables) GetDarkRequestOk() (*bool, bool)`

GetDarkRequestOk returns a tuple with the DarkRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkRequest

`func (o *PostCreateTweetRequestVariables) SetDarkRequest(v bool)`

SetDarkRequest sets DarkRequest field to given value.


### GetDisallowedReplyOptions

`func (o *PostCreateTweetRequestVariables) GetDisallowedReplyOptions() map[string]interface{}`

GetDisallowedReplyOptions returns the DisallowedReplyOptions field if non-nil, zero value otherwise.

### GetDisallowedReplyOptionsOk

`func (o *PostCreateTweetRequestVariables) GetDisallowedReplyOptionsOk() (*map[string]interface{}, bool)`

GetDisallowedReplyOptionsOk returns a tuple with the DisallowedReplyOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowedReplyOptions

`func (o *PostCreateTweetRequestVariables) SetDisallowedReplyOptions(v map[string]interface{})`

SetDisallowedReplyOptions sets DisallowedReplyOptions field to given value.

### HasDisallowedReplyOptions

`func (o *PostCreateTweetRequestVariables) HasDisallowedReplyOptions() bool`

HasDisallowedReplyOptions returns a boolean if a field has been set.

### GetMedia

`func (o *PostCreateTweetRequestVariables) GetMedia() PostCreateTweetRequestVariablesMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *PostCreateTweetRequestVariables) GetMediaOk() (*PostCreateTweetRequestVariablesMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *PostCreateTweetRequestVariables) SetMedia(v PostCreateTweetRequestVariablesMedia)`

SetMedia sets Media field to given value.


### GetReply

`func (o *PostCreateTweetRequestVariables) GetReply() PostCreateTweetRequestVariablesReply`

GetReply returns the Reply field if non-nil, zero value otherwise.

### GetReplyOk

`func (o *PostCreateTweetRequestVariables) GetReplyOk() (*PostCreateTweetRequestVariablesReply, bool)`

GetReplyOk returns a tuple with the Reply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReply

`func (o *PostCreateTweetRequestVariables) SetReply(v PostCreateTweetRequestVariablesReply)`

SetReply sets Reply field to given value.

### HasReply

`func (o *PostCreateTweetRequestVariables) HasReply() bool`

HasReply returns a boolean if a field has been set.

### GetSemanticAnnotationIds

`func (o *PostCreateTweetRequestVariables) GetSemanticAnnotationIds() []map[string]interface{}`

GetSemanticAnnotationIds returns the SemanticAnnotationIds field if non-nil, zero value otherwise.

### GetSemanticAnnotationIdsOk

`func (o *PostCreateTweetRequestVariables) GetSemanticAnnotationIdsOk() (*[]map[string]interface{}, bool)`

GetSemanticAnnotationIdsOk returns a tuple with the SemanticAnnotationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemanticAnnotationIds

`func (o *PostCreateTweetRequestVariables) SetSemanticAnnotationIds(v []map[string]interface{})`

SetSemanticAnnotationIds sets SemanticAnnotationIds field to given value.


### GetTweetText

`func (o *PostCreateTweetRequestVariables) GetTweetText() string`

GetTweetText returns the TweetText field if non-nil, zero value otherwise.

### GetTweetTextOk

`func (o *PostCreateTweetRequestVariables) GetTweetTextOk() (*string, bool)`

GetTweetTextOk returns a tuple with the TweetText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweetText

`func (o *PostCreateTweetRequestVariables) SetTweetText(v string)`

SetTweetText sets TweetText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


