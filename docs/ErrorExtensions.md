# ErrorExtensions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** |  | 
**Kind** | **string** |  | 
**Name** | **string** |  | 
**RetryAfter** | Pointer to **int32** |  | [optional] 
**Source** | **string** |  | 
**Tracing** | [**Tracing**](Tracing.md) |  | 

## Methods

### NewErrorExtensions

`func NewErrorExtensions(code int32, kind string, name string, source string, tracing Tracing, ) *ErrorExtensions`

NewErrorExtensions instantiates a new ErrorExtensions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorExtensionsWithDefaults

`func NewErrorExtensionsWithDefaults() *ErrorExtensions`

NewErrorExtensionsWithDefaults instantiates a new ErrorExtensions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorExtensions) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorExtensions) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorExtensions) SetCode(v int32)`

SetCode sets Code field to given value.


### GetKind

`func (o *ErrorExtensions) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ErrorExtensions) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ErrorExtensions) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *ErrorExtensions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ErrorExtensions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ErrorExtensions) SetName(v string)`

SetName sets Name field to given value.


### GetRetryAfter

`func (o *ErrorExtensions) GetRetryAfter() int32`

GetRetryAfter returns the RetryAfter field if non-nil, zero value otherwise.

### GetRetryAfterOk

`func (o *ErrorExtensions) GetRetryAfterOk() (*int32, bool)`

GetRetryAfterOk returns a tuple with the RetryAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfter

`func (o *ErrorExtensions) SetRetryAfter(v int32)`

SetRetryAfter sets RetryAfter field to given value.

### HasRetryAfter

`func (o *ErrorExtensions) HasRetryAfter() bool`

HasRetryAfter returns a boolean if a field has been set.

### GetSource

`func (o *ErrorExtensions) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ErrorExtensions) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ErrorExtensions) SetSource(v string)`

SetSource sets Source field to given value.


### GetTracing

`func (o *ErrorExtensions) GetTracing() Tracing`

GetTracing returns the Tracing field if non-nil, zero value otherwise.

### GetTracingOk

`func (o *ErrorExtensions) GetTracingOk() (*Tracing, bool)`

GetTracingOk returns a tuple with the Tracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracing

`func (o *ErrorExtensions) SetTracing(v Tracing)`

SetTracing sets Tracing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


