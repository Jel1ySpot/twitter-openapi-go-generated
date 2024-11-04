# TimelineAddToModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModuleEntryId** | **string** |  | 
**ModuleItems** | [**[]ModuleItem**](ModuleItem.md) |  | 
**Prepend** | Pointer to **bool** |  | [optional] 
**Type** | [**InstructionType**](InstructionType.md) |  | 

## Methods

### NewTimelineAddToModule

`func NewTimelineAddToModule(moduleEntryId string, moduleItems []ModuleItem, type_ InstructionType, ) *TimelineAddToModule`

NewTimelineAddToModule instantiates a new TimelineAddToModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineAddToModuleWithDefaults

`func NewTimelineAddToModuleWithDefaults() *TimelineAddToModule`

NewTimelineAddToModuleWithDefaults instantiates a new TimelineAddToModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModuleEntryId

`func (o *TimelineAddToModule) GetModuleEntryId() string`

GetModuleEntryId returns the ModuleEntryId field if non-nil, zero value otherwise.

### GetModuleEntryIdOk

`func (o *TimelineAddToModule) GetModuleEntryIdOk() (*string, bool)`

GetModuleEntryIdOk returns a tuple with the ModuleEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleEntryId

`func (o *TimelineAddToModule) SetModuleEntryId(v string)`

SetModuleEntryId sets ModuleEntryId field to given value.


### GetModuleItems

`func (o *TimelineAddToModule) GetModuleItems() []ModuleItem`

GetModuleItems returns the ModuleItems field if non-nil, zero value otherwise.

### GetModuleItemsOk

`func (o *TimelineAddToModule) GetModuleItemsOk() (*[]ModuleItem, bool)`

GetModuleItemsOk returns a tuple with the ModuleItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleItems

`func (o *TimelineAddToModule) SetModuleItems(v []ModuleItem)`

SetModuleItems sets ModuleItems field to given value.


### GetPrepend

`func (o *TimelineAddToModule) GetPrepend() bool`

GetPrepend returns the Prepend field if non-nil, zero value otherwise.

### GetPrependOk

`func (o *TimelineAddToModule) GetPrependOk() (*bool, bool)`

GetPrependOk returns a tuple with the Prepend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepend

`func (o *TimelineAddToModule) SetPrepend(v bool)`

SetPrepend sets Prepend field to given value.

### HasPrepend

`func (o *TimelineAddToModule) HasPrepend() bool`

HasPrepend returns a boolean if a field has been set.

### GetType

`func (o *TimelineAddToModule) GetType() InstructionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimelineAddToModule) GetTypeOk() (*InstructionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimelineAddToModule) SetType(v InstructionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


