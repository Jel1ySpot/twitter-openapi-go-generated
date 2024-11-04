# TimelineTimelineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Typename** | [**TypeName**](TypeName.md) |  | 
**ClientEventInfo** | Pointer to [**ClientEventInfo**](ClientEventInfo.md) |  | [optional] 
**EntryType** | [**ContentEntryType**](ContentEntryType.md) |  | 
**FeedbackInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**ItemContent** | [**ItemContentUnion**](ItemContentUnion.md) |  | 

## Methods

### NewTimelineTimelineItem

`func NewTimelineTimelineItem(typename TypeName, entryType ContentEntryType, itemContent ItemContentUnion, ) *TimelineTimelineItem`

NewTimelineTimelineItem instantiates a new TimelineTimelineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineTimelineItemWithDefaults

`func NewTimelineTimelineItemWithDefaults() *TimelineTimelineItem`

NewTimelineTimelineItemWithDefaults instantiates a new TimelineTimelineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypename

`func (o *TimelineTimelineItem) GetTypename() TypeName`

GetTypename returns the Typename field if non-nil, zero value otherwise.

### GetTypenameOk

`func (o *TimelineTimelineItem) GetTypenameOk() (*TypeName, bool)`

GetTypenameOk returns a tuple with the Typename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypename

`func (o *TimelineTimelineItem) SetTypename(v TypeName)`

SetTypename sets Typename field to given value.


### GetClientEventInfo

`func (o *TimelineTimelineItem) GetClientEventInfo() ClientEventInfo`

GetClientEventInfo returns the ClientEventInfo field if non-nil, zero value otherwise.

### GetClientEventInfoOk

`func (o *TimelineTimelineItem) GetClientEventInfoOk() (*ClientEventInfo, bool)`

GetClientEventInfoOk returns a tuple with the ClientEventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEventInfo

`func (o *TimelineTimelineItem) SetClientEventInfo(v ClientEventInfo)`

SetClientEventInfo sets ClientEventInfo field to given value.

### HasClientEventInfo

`func (o *TimelineTimelineItem) HasClientEventInfo() bool`

HasClientEventInfo returns a boolean if a field has been set.

### GetEntryType

`func (o *TimelineTimelineItem) GetEntryType() ContentEntryType`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *TimelineTimelineItem) GetEntryTypeOk() (*ContentEntryType, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *TimelineTimelineItem) SetEntryType(v ContentEntryType)`

SetEntryType sets EntryType field to given value.


### GetFeedbackInfo

`func (o *TimelineTimelineItem) GetFeedbackInfo() map[string]interface{}`

GetFeedbackInfo returns the FeedbackInfo field if non-nil, zero value otherwise.

### GetFeedbackInfoOk

`func (o *TimelineTimelineItem) GetFeedbackInfoOk() (*map[string]interface{}, bool)`

GetFeedbackInfoOk returns a tuple with the FeedbackInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackInfo

`func (o *TimelineTimelineItem) SetFeedbackInfo(v map[string]interface{})`

SetFeedbackInfo sets FeedbackInfo field to given value.

### HasFeedbackInfo

`func (o *TimelineTimelineItem) HasFeedbackInfo() bool`

HasFeedbackInfo returns a boolean if a field has been set.

### GetItemContent

`func (o *TimelineTimelineItem) GetItemContent() ItemContentUnion`

GetItemContent returns the ItemContent field if non-nil, zero value otherwise.

### GetItemContentOk

`func (o *TimelineTimelineItem) GetItemContentOk() (*ItemContentUnion, bool)`

GetItemContentOk returns a tuple with the ItemContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemContent

`func (o *TimelineTimelineItem) SetItemContent(v ItemContentUnion)`

SetItemContent sets ItemContent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


