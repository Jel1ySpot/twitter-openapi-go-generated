# TimelineAddEntries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | [**[]TimelineAddEntry**](TimelineAddEntry.md) |  | 
**Type** | [**InstructionType**](InstructionType.md) |  | 

## Methods

### NewTimelineAddEntries

`func NewTimelineAddEntries(entries []TimelineAddEntry, type_ InstructionType, ) *TimelineAddEntries`

NewTimelineAddEntries instantiates a new TimelineAddEntries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineAddEntriesWithDefaults

`func NewTimelineAddEntriesWithDefaults() *TimelineAddEntries`

NewTimelineAddEntriesWithDefaults instantiates a new TimelineAddEntries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *TimelineAddEntries) GetEntries() []TimelineAddEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *TimelineAddEntries) GetEntriesOk() (*[]TimelineAddEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *TimelineAddEntries) SetEntries(v []TimelineAddEntry)`

SetEntries sets Entries field to given value.


### GetType

`func (o *TimelineAddEntries) GetType() InstructionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimelineAddEntries) GetTypeOk() (*InstructionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimelineAddEntries) SetType(v InstructionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


