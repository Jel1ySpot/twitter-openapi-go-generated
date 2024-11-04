# TimelineAddEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**ContentUnion**](ContentUnion.md) |  | 
**EntryId** | **string** |  | 
**SortIndex** | **string** |  | 

## Methods

### NewTimelineAddEntry

`func NewTimelineAddEntry(content ContentUnion, entryId string, sortIndex string, ) *TimelineAddEntry`

NewTimelineAddEntry instantiates a new TimelineAddEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineAddEntryWithDefaults

`func NewTimelineAddEntryWithDefaults() *TimelineAddEntry`

NewTimelineAddEntryWithDefaults instantiates a new TimelineAddEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *TimelineAddEntry) GetContent() ContentUnion`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *TimelineAddEntry) GetContentOk() (*ContentUnion, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *TimelineAddEntry) SetContent(v ContentUnion)`

SetContent sets Content field to given value.


### GetEntryId

`func (o *TimelineAddEntry) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *TimelineAddEntry) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *TimelineAddEntry) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetSortIndex

`func (o *TimelineAddEntry) GetSortIndex() string`

GetSortIndex returns the SortIndex field if non-nil, zero value otherwise.

### GetSortIndexOk

`func (o *TimelineAddEntry) GetSortIndexOk() (*string, bool)`

GetSortIndexOk returns a tuple with the SortIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortIndex

`func (o *TimelineAddEntry) SetSortIndex(v string)`

SetSortIndex sets SortIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


