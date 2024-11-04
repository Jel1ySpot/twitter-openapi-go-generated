# ModuleEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientEventInfo** | Pointer to [**ClientEventInfo**](ClientEventInfo.md) |  | [optional] 
**FeedbackInfo** | Pointer to [**FeedbackInfo**](FeedbackInfo.md) |  | [optional] 
**ItemContent** | [**ItemContentUnion**](ItemContentUnion.md) |  | 

## Methods

### NewModuleEntry

`func NewModuleEntry(itemContent ItemContentUnion, ) *ModuleEntry`

NewModuleEntry instantiates a new ModuleEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleEntryWithDefaults

`func NewModuleEntryWithDefaults() *ModuleEntry`

NewModuleEntryWithDefaults instantiates a new ModuleEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientEventInfo

`func (o *ModuleEntry) GetClientEventInfo() ClientEventInfo`

GetClientEventInfo returns the ClientEventInfo field if non-nil, zero value otherwise.

### GetClientEventInfoOk

`func (o *ModuleEntry) GetClientEventInfoOk() (*ClientEventInfo, bool)`

GetClientEventInfoOk returns a tuple with the ClientEventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEventInfo

`func (o *ModuleEntry) SetClientEventInfo(v ClientEventInfo)`

SetClientEventInfo sets ClientEventInfo field to given value.

### HasClientEventInfo

`func (o *ModuleEntry) HasClientEventInfo() bool`

HasClientEventInfo returns a boolean if a field has been set.

### GetFeedbackInfo

`func (o *ModuleEntry) GetFeedbackInfo() FeedbackInfo`

GetFeedbackInfo returns the FeedbackInfo field if non-nil, zero value otherwise.

### GetFeedbackInfoOk

`func (o *ModuleEntry) GetFeedbackInfoOk() (*FeedbackInfo, bool)`

GetFeedbackInfoOk returns a tuple with the FeedbackInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackInfo

`func (o *ModuleEntry) SetFeedbackInfo(v FeedbackInfo)`

SetFeedbackInfo sets FeedbackInfo field to given value.

### HasFeedbackInfo

`func (o *ModuleEntry) HasFeedbackInfo() bool`

HasFeedbackInfo returns a boolean if a field has been set.

### GetItemContent

`func (o *ModuleEntry) GetItemContent() ItemContentUnion`

GetItemContent returns the ItemContent field if non-nil, zero value otherwise.

### GetItemContentOk

`func (o *ModuleEntry) GetItemContentOk() (*ItemContentUnion, bool)`

GetItemContentOk returns a tuple with the ItemContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemContent

`func (o *ModuleEntry) SetItemContent(v ItemContentUnion)`

SetItemContent sets ItemContent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


