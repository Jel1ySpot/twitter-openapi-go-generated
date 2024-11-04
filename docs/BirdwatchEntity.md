# BirdwatchEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromIndex** | **int32** |  | 
**Ref** | [**BirdwatchEntityRef**](BirdwatchEntityRef.md) |  | 
**ToIndex** | **int32** |  | 

## Methods

### NewBirdwatchEntity

`func NewBirdwatchEntity(fromIndex int32, ref BirdwatchEntityRef, toIndex int32, ) *BirdwatchEntity`

NewBirdwatchEntity instantiates a new BirdwatchEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBirdwatchEntityWithDefaults

`func NewBirdwatchEntityWithDefaults() *BirdwatchEntity`

NewBirdwatchEntityWithDefaults instantiates a new BirdwatchEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromIndex

`func (o *BirdwatchEntity) GetFromIndex() int32`

GetFromIndex returns the FromIndex field if non-nil, zero value otherwise.

### GetFromIndexOk

`func (o *BirdwatchEntity) GetFromIndexOk() (*int32, bool)`

GetFromIndexOk returns a tuple with the FromIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromIndex

`func (o *BirdwatchEntity) SetFromIndex(v int32)`

SetFromIndex sets FromIndex field to given value.


### GetRef

`func (o *BirdwatchEntity) GetRef() BirdwatchEntityRef`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *BirdwatchEntity) GetRefOk() (*BirdwatchEntityRef, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *BirdwatchEntity) SetRef(v BirdwatchEntityRef)`

SetRef sets Ref field to given value.


### GetToIndex

`func (o *BirdwatchEntity) GetToIndex() int32`

GetToIndex returns the ToIndex field if non-nil, zero value otherwise.

### GetToIndexOk

`func (o *BirdwatchEntity) GetToIndexOk() (*int32, bool)`

GetToIndexOk returns a tuple with the ToIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToIndex

`func (o *BirdwatchEntity) SetToIndex(v int32)`

SetToIndex sets ToIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


