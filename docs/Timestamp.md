# Timestamp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indices** | **[]int32** |  | 
**Seconds** | **int32** |  | 
**Text** | **string** |  | 

## Methods

### NewTimestamp

`func NewTimestamp(indices []int32, seconds int32, text string, ) *Timestamp`

NewTimestamp instantiates a new Timestamp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimestampWithDefaults

`func NewTimestampWithDefaults() *Timestamp`

NewTimestampWithDefaults instantiates a new Timestamp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndices

`func (o *Timestamp) GetIndices() []int32`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *Timestamp) GetIndicesOk() (*[]int32, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *Timestamp) SetIndices(v []int32)`

SetIndices sets Indices field to given value.


### GetSeconds

`func (o *Timestamp) GetSeconds() int32`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *Timestamp) GetSecondsOk() (*int32, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *Timestamp) SetSeconds(v int32)`

SetSeconds sets Seconds field to given value.


### GetText

`func (o *Timestamp) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Timestamp) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Timestamp) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


