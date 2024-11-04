# Url

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayUrl** | **string** |  | 
**ExpandedUrl** | Pointer to **string** |  | [optional] 
**Indices** | **[]int32** |  | 
**Url** | **string** |  | 

## Methods

### NewUrl

`func NewUrl(displayUrl string, indices []int32, url string, ) *Url`

NewUrl instantiates a new Url object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlWithDefaults

`func NewUrlWithDefaults() *Url`

NewUrlWithDefaults instantiates a new Url object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayUrl

`func (o *Url) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *Url) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *Url) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetExpandedUrl

`func (o *Url) GetExpandedUrl() string`

GetExpandedUrl returns the ExpandedUrl field if non-nil, zero value otherwise.

### GetExpandedUrlOk

`func (o *Url) GetExpandedUrlOk() (*string, bool)`

GetExpandedUrlOk returns a tuple with the ExpandedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandedUrl

`func (o *Url) SetExpandedUrl(v string)`

SetExpandedUrl sets ExpandedUrl field to given value.

### HasExpandedUrl

`func (o *Url) HasExpandedUrl() bool`

HasExpandedUrl returns a boolean if a field has been set.

### GetIndices

`func (o *Url) GetIndices() []int32`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *Url) GetIndicesOk() (*[]int32, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *Url) SetIndices(v []int32)`

SetIndices sets Indices field to given value.


### GetUrl

`func (o *Url) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Url) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Url) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


