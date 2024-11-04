# TimelinePinEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entry** | [**TimelineAddEntry**](TimelineAddEntry.md) |  | 
**Type** | [**InstructionType**](InstructionType.md) |  | 

## Methods

### NewTimelinePinEntry

`func NewTimelinePinEntry(entry TimelineAddEntry, type_ InstructionType, ) *TimelinePinEntry`

NewTimelinePinEntry instantiates a new TimelinePinEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelinePinEntryWithDefaults

`func NewTimelinePinEntryWithDefaults() *TimelinePinEntry`

NewTimelinePinEntryWithDefaults instantiates a new TimelinePinEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntry

`func (o *TimelinePinEntry) GetEntry() TimelineAddEntry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *TimelinePinEntry) GetEntryOk() (*TimelineAddEntry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *TimelinePinEntry) SetEntry(v TimelineAddEntry)`

SetEntry sets Entry field to given value.


### GetType

`func (o *TimelinePinEntry) GetType() InstructionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimelinePinEntry) GetTypeOk() (*InstructionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimelinePinEntry) SetType(v InstructionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


