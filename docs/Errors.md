# Errors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ErrorsData**](ErrorsData.md) |  | [optional] 
**Errors** | [**[]Error**](Error.md) |  | 

## Methods

### NewErrors

`func NewErrors(errors []Error, ) *Errors`

NewErrors instantiates a new Errors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsWithDefaults

`func NewErrorsWithDefaults() *Errors`

NewErrorsWithDefaults instantiates a new Errors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Errors) GetData() ErrorsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Errors) GetDataOk() (*ErrorsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Errors) SetData(v ErrorsData)`

SetData sets Data field to given value.

### HasData

`func (o *Errors) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *Errors) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Errors) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Errors) SetErrors(v []Error)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


