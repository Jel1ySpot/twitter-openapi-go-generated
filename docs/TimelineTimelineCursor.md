# TimelineTimelineCursor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Typename** | [**TypeName**](TypeName.md) |  | 
**CursorType** | [**CursorType**](CursorType.md) |  | 
**DisplayTreatment** | Pointer to [**DisplayTreatment**](DisplayTreatment.md) |  | [optional] 
**EntryType** | Pointer to [**ContentEntryType**](ContentEntryType.md) |  | [optional] 
**ItemType** | Pointer to [**ContentEntryType**](ContentEntryType.md) |  | [optional] 
**StopOnEmptyResponse** | Pointer to **bool** |  | [optional] 
**Value** | **string** |  | 

## Methods

### NewTimelineTimelineCursor

`func NewTimelineTimelineCursor(typename TypeName, cursorType CursorType, value string, ) *TimelineTimelineCursor`

NewTimelineTimelineCursor instantiates a new TimelineTimelineCursor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineTimelineCursorWithDefaults

`func NewTimelineTimelineCursorWithDefaults() *TimelineTimelineCursor`

NewTimelineTimelineCursorWithDefaults instantiates a new TimelineTimelineCursor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypename

`func (o *TimelineTimelineCursor) GetTypename() TypeName`

GetTypename returns the Typename field if non-nil, zero value otherwise.

### GetTypenameOk

`func (o *TimelineTimelineCursor) GetTypenameOk() (*TypeName, bool)`

GetTypenameOk returns a tuple with the Typename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypename

`func (o *TimelineTimelineCursor) SetTypename(v TypeName)`

SetTypename sets Typename field to given value.


### GetCursorType

`func (o *TimelineTimelineCursor) GetCursorType() CursorType`

GetCursorType returns the CursorType field if non-nil, zero value otherwise.

### GetCursorTypeOk

`func (o *TimelineTimelineCursor) GetCursorTypeOk() (*CursorType, bool)`

GetCursorTypeOk returns a tuple with the CursorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorType

`func (o *TimelineTimelineCursor) SetCursorType(v CursorType)`

SetCursorType sets CursorType field to given value.


### GetDisplayTreatment

`func (o *TimelineTimelineCursor) GetDisplayTreatment() DisplayTreatment`

GetDisplayTreatment returns the DisplayTreatment field if non-nil, zero value otherwise.

### GetDisplayTreatmentOk

`func (o *TimelineTimelineCursor) GetDisplayTreatmentOk() (*DisplayTreatment, bool)`

GetDisplayTreatmentOk returns a tuple with the DisplayTreatment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTreatment

`func (o *TimelineTimelineCursor) SetDisplayTreatment(v DisplayTreatment)`

SetDisplayTreatment sets DisplayTreatment field to given value.

### HasDisplayTreatment

`func (o *TimelineTimelineCursor) HasDisplayTreatment() bool`

HasDisplayTreatment returns a boolean if a field has been set.

### GetEntryType

`func (o *TimelineTimelineCursor) GetEntryType() ContentEntryType`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *TimelineTimelineCursor) GetEntryTypeOk() (*ContentEntryType, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *TimelineTimelineCursor) SetEntryType(v ContentEntryType)`

SetEntryType sets EntryType field to given value.

### HasEntryType

`func (o *TimelineTimelineCursor) HasEntryType() bool`

HasEntryType returns a boolean if a field has been set.

### GetItemType

`func (o *TimelineTimelineCursor) GetItemType() ContentEntryType`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *TimelineTimelineCursor) GetItemTypeOk() (*ContentEntryType, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *TimelineTimelineCursor) SetItemType(v ContentEntryType)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *TimelineTimelineCursor) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetStopOnEmptyResponse

`func (o *TimelineTimelineCursor) GetStopOnEmptyResponse() bool`

GetStopOnEmptyResponse returns the StopOnEmptyResponse field if non-nil, zero value otherwise.

### GetStopOnEmptyResponseOk

`func (o *TimelineTimelineCursor) GetStopOnEmptyResponseOk() (*bool, bool)`

GetStopOnEmptyResponseOk returns a tuple with the StopOnEmptyResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOnEmptyResponse

`func (o *TimelineTimelineCursor) SetStopOnEmptyResponse(v bool)`

SetStopOnEmptyResponse sets StopOnEmptyResponse field to given value.

### HasStopOnEmptyResponse

`func (o *TimelineTimelineCursor) HasStopOnEmptyResponse() bool`

HasStopOnEmptyResponse returns a boolean if a field has been set.

### GetValue

`func (o *TimelineTimelineCursor) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TimelineTimelineCursor) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TimelineTimelineCursor) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


