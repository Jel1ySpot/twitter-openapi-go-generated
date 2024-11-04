# TimelineReplaceEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entry** | [**TimelineAddEntry**](TimelineAddEntry.md) |  | 
**EntryIdToReplace** | **string** |  | 
**Type** | [**InstructionType**](InstructionType.md) |  | 

## Methods

### NewTimelineReplaceEntry

`func NewTimelineReplaceEntry(entry TimelineAddEntry, entryIdToReplace string, type_ InstructionType, ) *TimelineReplaceEntry`

NewTimelineReplaceEntry instantiates a new TimelineReplaceEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineReplaceEntryWithDefaults

`func NewTimelineReplaceEntryWithDefaults() *TimelineReplaceEntry`

NewTimelineReplaceEntryWithDefaults instantiates a new TimelineReplaceEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntry

`func (o *TimelineReplaceEntry) GetEntry() TimelineAddEntry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *TimelineReplaceEntry) GetEntryOk() (*TimelineAddEntry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *TimelineReplaceEntry) SetEntry(v TimelineAddEntry)`

SetEntry sets Entry field to given value.


### GetEntryIdToReplace

`func (o *TimelineReplaceEntry) GetEntryIdToReplace() string`

GetEntryIdToReplace returns the EntryIdToReplace field if non-nil, zero value otherwise.

### GetEntryIdToReplaceOk

`func (o *TimelineReplaceEntry) GetEntryIdToReplaceOk() (*string, bool)`

GetEntryIdToReplaceOk returns a tuple with the EntryIdToReplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryIdToReplace

`func (o *TimelineReplaceEntry) SetEntryIdToReplace(v string)`

SetEntryIdToReplace sets EntryIdToReplace field to given value.


### GetType

`func (o *TimelineReplaceEntry) GetType() InstructionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimelineReplaceEntry) GetTypeOk() (*InstructionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimelineReplaceEntry) SetType(v InstructionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


