# MediaVideoInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AspectRatio** | **[]int32** |  | 
**DurationMillis** | Pointer to **int32** |  | [optional] 
**Variants** | [**[]MediaVideoInfoVariant**](MediaVideoInfoVariant.md) |  | 

## Methods

### NewMediaVideoInfo

`func NewMediaVideoInfo(aspectRatio []int32, variants []MediaVideoInfoVariant, ) *MediaVideoInfo`

NewMediaVideoInfo instantiates a new MediaVideoInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaVideoInfoWithDefaults

`func NewMediaVideoInfoWithDefaults() *MediaVideoInfo`

NewMediaVideoInfoWithDefaults instantiates a new MediaVideoInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAspectRatio

`func (o *MediaVideoInfo) GetAspectRatio() []int32`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *MediaVideoInfo) GetAspectRatioOk() (*[]int32, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *MediaVideoInfo) SetAspectRatio(v []int32)`

SetAspectRatio sets AspectRatio field to given value.


### GetDurationMillis

`func (o *MediaVideoInfo) GetDurationMillis() int32`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *MediaVideoInfo) GetDurationMillisOk() (*int32, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *MediaVideoInfo) SetDurationMillis(v int32)`

SetDurationMillis sets DurationMillis field to given value.

### HasDurationMillis

`func (o *MediaVideoInfo) HasDurationMillis() bool`

HasDurationMillis returns a boolean if a field has been set.

### GetVariants

`func (o *MediaVideoInfo) GetVariants() []MediaVideoInfoVariant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *MediaVideoInfo) GetVariantsOk() (*[]MediaVideoInfoVariant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *MediaVideoInfo) SetVariants(v []MediaVideoInfoVariant)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


