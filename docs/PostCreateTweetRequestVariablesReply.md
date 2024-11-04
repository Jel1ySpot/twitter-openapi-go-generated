# PostCreateTweetRequestVariablesReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeReplyUserIds** | **[]map[string]interface{}** |  | 
**InReplyToTweetId** | **string** |  | [default to "1111111111111111111"]

## Methods

### NewPostCreateTweetRequestVariablesReply

`func NewPostCreateTweetRequestVariablesReply(excludeReplyUserIds []map[string]interface{}, inReplyToTweetId string, ) *PostCreateTweetRequestVariablesReply`

NewPostCreateTweetRequestVariablesReply instantiates a new PostCreateTweetRequestVariablesReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCreateTweetRequestVariablesReplyWithDefaults

`func NewPostCreateTweetRequestVariablesReplyWithDefaults() *PostCreateTweetRequestVariablesReply`

NewPostCreateTweetRequestVariablesReplyWithDefaults instantiates a new PostCreateTweetRequestVariablesReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeReplyUserIds

`func (o *PostCreateTweetRequestVariablesReply) GetExcludeReplyUserIds() []map[string]interface{}`

GetExcludeReplyUserIds returns the ExcludeReplyUserIds field if non-nil, zero value otherwise.

### GetExcludeReplyUserIdsOk

`func (o *PostCreateTweetRequestVariablesReply) GetExcludeReplyUserIdsOk() (*[]map[string]interface{}, bool)`

GetExcludeReplyUserIdsOk returns a tuple with the ExcludeReplyUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeReplyUserIds

`func (o *PostCreateTweetRequestVariablesReply) SetExcludeReplyUserIds(v []map[string]interface{})`

SetExcludeReplyUserIds sets ExcludeReplyUserIds field to given value.


### GetInReplyToTweetId

`func (o *PostCreateTweetRequestVariablesReply) GetInReplyToTweetId() string`

GetInReplyToTweetId returns the InReplyToTweetId field if non-nil, zero value otherwise.

### GetInReplyToTweetIdOk

`func (o *PostCreateTweetRequestVariablesReply) GetInReplyToTweetIdOk() (*string, bool)`

GetInReplyToTweetIdOk returns a tuple with the InReplyToTweetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyToTweetId

`func (o *PostCreateTweetRequestVariablesReply) SetInReplyToTweetId(v string)`

SetInReplyToTweetId sets InReplyToTweetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


