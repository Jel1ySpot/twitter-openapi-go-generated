# TextEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromIndex** | **int32** |  | 
**Ref** | [**TextEntityRef**](TextEntityRef.md) |  | 
**ToIndex** | **int32** |  | 

## Methods

### NewTextEntity

`func NewTextEntity(fromIndex int32, ref TextEntityRef, toIndex int32, ) *TextEntity`

NewTextEntity instantiates a new TextEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextEntityWithDefaults

`func NewTextEntityWithDefaults() *TextEntity`

NewTextEntityWithDefaults instantiates a new TextEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromIndex

`func (o *TextEntity) GetFromIndex() int32`

GetFromIndex returns the FromIndex field if non-nil, zero value otherwise.

### GetFromIndexOk

`func (o *TextEntity) GetFromIndexOk() (*int32, bool)`

GetFromIndexOk returns a tuple with the FromIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromIndex

`func (o *TextEntity) SetFromIndex(v int32)`

SetFromIndex sets FromIndex field to given value.


### GetRef

`func (o *TextEntity) GetRef() TextEntityRef`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *TextEntity) GetRefOk() (*TextEntityRef, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *TextEntity) SetRef(v TextEntityRef)`

SetRef sets Ref field to given value.


### GetToIndex

`func (o *TextEntity) GetToIndex() int32`

GetToIndex returns the ToIndex field if non-nil, zero value otherwise.

### GetToIndexOk

`func (o *TextEntity) GetToIndexOk() (*int32, bool)`

GetToIndexOk returns a tuple with the ToIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToIndex

`func (o *TextEntity) SetToIndex(v int32)`

SetToIndex sets ToIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


