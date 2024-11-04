# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** |  | 
**Extensions** | [**ErrorExtensions**](ErrorExtensions.md) |  | 
**Kind** | **string** |  | 
**Locations** | [**[]Location**](Location.md) |  | 
**Message** | **string** |  | 
**Name** | **string** |  | 
**Path** | **[]string** |  | 
**RetryAfter** | Pointer to **int32** |  | [optional] 
**Source** | **string** |  | 
**Tracing** | [**Tracing**](Tracing.md) |  | 

## Methods

### NewError

`func NewError(code int32, extensions ErrorExtensions, kind string, locations []Location, message string, name string, path []string, source string, tracing Tracing, ) *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Error) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Error) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Error) SetCode(v int32)`

SetCode sets Code field to given value.


### GetExtensions

`func (o *Error) GetExtensions() ErrorExtensions`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Error) GetExtensionsOk() (*ErrorExtensions, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Error) SetExtensions(v ErrorExtensions)`

SetExtensions sets Extensions field to given value.


### GetKind

`func (o *Error) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Error) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Error) SetKind(v string)`

SetKind sets Kind field to given value.


### GetLocations

`func (o *Error) GetLocations() []Location`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *Error) GetLocationsOk() (*[]Location, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *Error) SetLocations(v []Location)`

SetLocations sets Locations field to given value.


### GetMessage

`func (o *Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetName

`func (o *Error) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Error) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Error) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *Error) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Error) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Error) SetPath(v []string)`

SetPath sets Path field to given value.


### GetRetryAfter

`func (o *Error) GetRetryAfter() int32`

GetRetryAfter returns the RetryAfter field if non-nil, zero value otherwise.

### GetRetryAfterOk

`func (o *Error) GetRetryAfterOk() (*int32, bool)`

GetRetryAfterOk returns a tuple with the RetryAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfter

`func (o *Error) SetRetryAfter(v int32)`

SetRetryAfter sets RetryAfter field to given value.

### HasRetryAfter

`func (o *Error) HasRetryAfter() bool`

HasRetryAfter returns a boolean if a field has been set.

### GetSource

`func (o *Error) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Error) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Error) SetSource(v string)`

SetSource sets Source field to given value.


### GetTracing

`func (o *Error) GetTracing() Tracing`

GetTracing returns the Tracing field if non-nil, zero value otherwise.

### GetTracingOk

`func (o *Error) GetTracingOk() (*Tracing, bool)`

GetTracingOk returns a tuple with the Tracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracing

`func (o *Error) SetTracing(v Tracing)`

SetTracing sets Tracing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


