# TimelineTimelineModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Typename** | [**TypeName**](TypeName.md) |  | 
**ClientEventInfo** | **map[string]interface{}** |  | 
**DisplayType** | [**DisplayType**](DisplayType.md) |  | 
**EntryType** | [**ContentEntryType**](ContentEntryType.md) |  | 
**FeedbackInfo** | Pointer to [**FeedbackInfo**](FeedbackInfo.md) |  | [optional] 
**Footer** | Pointer to **map[string]interface{}** |  | [optional] 
**Header** | Pointer to **map[string]interface{}** |  | [optional] 
**Items** | Pointer to [**[]ModuleItem**](ModuleItem.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTimelineTimelineModule

`func NewTimelineTimelineModule(typename TypeName, clientEventInfo map[string]interface{}, displayType DisplayType, entryType ContentEntryType, ) *TimelineTimelineModule`

NewTimelineTimelineModule instantiates a new TimelineTimelineModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineTimelineModuleWithDefaults

`func NewTimelineTimelineModuleWithDefaults() *TimelineTimelineModule`

NewTimelineTimelineModuleWithDefaults instantiates a new TimelineTimelineModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypename

`func (o *TimelineTimelineModule) GetTypename() TypeName`

GetTypename returns the Typename field if non-nil, zero value otherwise.

### GetTypenameOk

`func (o *TimelineTimelineModule) GetTypenameOk() (*TypeName, bool)`

GetTypenameOk returns a tuple with the Typename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypename

`func (o *TimelineTimelineModule) SetTypename(v TypeName)`

SetTypename sets Typename field to given value.


### GetClientEventInfo

`func (o *TimelineTimelineModule) GetClientEventInfo() map[string]interface{}`

GetClientEventInfo returns the ClientEventInfo field if non-nil, zero value otherwise.

### GetClientEventInfoOk

`func (o *TimelineTimelineModule) GetClientEventInfoOk() (*map[string]interface{}, bool)`

GetClientEventInfoOk returns a tuple with the ClientEventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEventInfo

`func (o *TimelineTimelineModule) SetClientEventInfo(v map[string]interface{})`

SetClientEventInfo sets ClientEventInfo field to given value.


### GetDisplayType

`func (o *TimelineTimelineModule) GetDisplayType() DisplayType`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *TimelineTimelineModule) GetDisplayTypeOk() (*DisplayType, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *TimelineTimelineModule) SetDisplayType(v DisplayType)`

SetDisplayType sets DisplayType field to given value.


### GetEntryType

`func (o *TimelineTimelineModule) GetEntryType() ContentEntryType`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *TimelineTimelineModule) GetEntryTypeOk() (*ContentEntryType, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *TimelineTimelineModule) SetEntryType(v ContentEntryType)`

SetEntryType sets EntryType field to given value.


### GetFeedbackInfo

`func (o *TimelineTimelineModule) GetFeedbackInfo() FeedbackInfo`

GetFeedbackInfo returns the FeedbackInfo field if non-nil, zero value otherwise.

### GetFeedbackInfoOk

`func (o *TimelineTimelineModule) GetFeedbackInfoOk() (*FeedbackInfo, bool)`

GetFeedbackInfoOk returns a tuple with the FeedbackInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackInfo

`func (o *TimelineTimelineModule) SetFeedbackInfo(v FeedbackInfo)`

SetFeedbackInfo sets FeedbackInfo field to given value.

### HasFeedbackInfo

`func (o *TimelineTimelineModule) HasFeedbackInfo() bool`

HasFeedbackInfo returns a boolean if a field has been set.

### GetFooter

`func (o *TimelineTimelineModule) GetFooter() map[string]interface{}`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *TimelineTimelineModule) GetFooterOk() (*map[string]interface{}, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *TimelineTimelineModule) SetFooter(v map[string]interface{})`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *TimelineTimelineModule) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetHeader

`func (o *TimelineTimelineModule) GetHeader() map[string]interface{}`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *TimelineTimelineModule) GetHeaderOk() (*map[string]interface{}, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *TimelineTimelineModule) SetHeader(v map[string]interface{})`

SetHeader sets Header field to given value.

### HasHeader

`func (o *TimelineTimelineModule) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetItems

`func (o *TimelineTimelineModule) GetItems() []ModuleItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TimelineTimelineModule) GetItemsOk() (*[]ModuleItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TimelineTimelineModule) SetItems(v []ModuleItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *TimelineTimelineModule) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMetadata

`func (o *TimelineTimelineModule) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TimelineTimelineModule) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TimelineTimelineModule) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TimelineTimelineModule) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


