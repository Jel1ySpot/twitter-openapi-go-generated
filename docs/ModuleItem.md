# ModuleItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**Item** | [**ModuleEntry**](ModuleEntry.md) |  | 

## Methods

### NewModuleItem

`func NewModuleItem(entryId string, item ModuleEntry, ) *ModuleItem`

NewModuleItem instantiates a new ModuleItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleItemWithDefaults

`func NewModuleItemWithDefaults() *ModuleItem`

NewModuleItemWithDefaults instantiates a new ModuleItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *ModuleItem) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *ModuleItem) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *ModuleItem) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetItem

`func (o *ModuleItem) GetItem() ModuleEntry`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ModuleItem) GetItemOk() (*ModuleEntry, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ModuleItem) SetItem(v ModuleEntry)`

SetItem sets Item field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


