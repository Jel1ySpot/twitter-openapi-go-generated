# TimelineTopicContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionalityType** | Pointer to **string** |  | [optional] 
**Topic** | Pointer to [**TopicContext**](TopicContext.md) |  | [optional] 
**Type** | Pointer to [**SocialContextUnionType**](SocialContextUnionType.md) |  | [optional] 

## Methods

### NewTimelineTopicContext

`func NewTimelineTopicContext() *TimelineTopicContext`

NewTimelineTopicContext instantiates a new TimelineTopicContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineTopicContextWithDefaults

`func NewTimelineTopicContextWithDefaults() *TimelineTopicContext`

NewTimelineTopicContextWithDefaults instantiates a new TimelineTopicContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionalityType

`func (o *TimelineTopicContext) GetFunctionalityType() string`

GetFunctionalityType returns the FunctionalityType field if non-nil, zero value otherwise.

### GetFunctionalityTypeOk

`func (o *TimelineTopicContext) GetFunctionalityTypeOk() (*string, bool)`

GetFunctionalityTypeOk returns a tuple with the FunctionalityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionalityType

`func (o *TimelineTopicContext) SetFunctionalityType(v string)`

SetFunctionalityType sets FunctionalityType field to given value.

### HasFunctionalityType

`func (o *TimelineTopicContext) HasFunctionalityType() bool`

HasFunctionalityType returns a boolean if a field has been set.

### GetTopic

`func (o *TimelineTopicContext) GetTopic() TopicContext`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *TimelineTopicContext) GetTopicOk() (*TopicContext, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *TimelineTopicContext) SetTopic(v TopicContext)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *TimelineTopicContext) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetType

`func (o *TimelineTopicContext) GetType() SocialContextUnionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimelineTopicContext) GetTypeOk() (*SocialContextUnionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimelineTopicContext) SetType(v SocialContextUnionType)`

SetType sets Type field to given value.

### HasType

`func (o *TimelineTopicContext) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


