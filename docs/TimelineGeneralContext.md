# TimelineGeneralContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextType** | Pointer to **string** |  | [optional] 
**LandingUrl** | Pointer to [**SocialContextLandingUrl**](SocialContextLandingUrl.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**SocialContextUnionType**](SocialContextUnionType.md) |  | [optional] 

## Methods

### NewTimelineGeneralContext

`func NewTimelineGeneralContext() *TimelineGeneralContext`

NewTimelineGeneralContext instantiates a new TimelineGeneralContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineGeneralContextWithDefaults

`func NewTimelineGeneralContextWithDefaults() *TimelineGeneralContext`

NewTimelineGeneralContextWithDefaults instantiates a new TimelineGeneralContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextType

`func (o *TimelineGeneralContext) GetContextType() string`

GetContextType returns the ContextType field if non-nil, zero value otherwise.

### GetContextTypeOk

`func (o *TimelineGeneralContext) GetContextTypeOk() (*string, bool)`

GetContextTypeOk returns a tuple with the ContextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextType

`func (o *TimelineGeneralContext) SetContextType(v string)`

SetContextType sets ContextType field to given value.

### HasContextType

`func (o *TimelineGeneralContext) HasContextType() bool`

HasContextType returns a boolean if a field has been set.

### GetLandingUrl

`func (o *TimelineGeneralContext) GetLandingUrl() SocialContextLandingUrl`

GetLandingUrl returns the LandingUrl field if non-nil, zero value otherwise.

### GetLandingUrlOk

`func (o *TimelineGeneralContext) GetLandingUrlOk() (*SocialContextLandingUrl, bool)`

GetLandingUrlOk returns a tuple with the LandingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingUrl

`func (o *TimelineGeneralContext) SetLandingUrl(v SocialContextLandingUrl)`

SetLandingUrl sets LandingUrl field to given value.

### HasLandingUrl

`func (o *TimelineGeneralContext) HasLandingUrl() bool`

HasLandingUrl returns a boolean if a field has been set.

### GetText

`func (o *TimelineGeneralContext) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TimelineGeneralContext) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TimelineGeneralContext) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TimelineGeneralContext) HasText() bool`

HasText returns a boolean if a field has been set.

### GetType

`func (o *TimelineGeneralContext) GetType() SocialContextUnionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimelineGeneralContext) GetTypeOk() (*SocialContextUnionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimelineGeneralContext) SetType(v SocialContextUnionType)`

SetType sets Type field to given value.

### HasType

`func (o *TimelineGeneralContext) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


