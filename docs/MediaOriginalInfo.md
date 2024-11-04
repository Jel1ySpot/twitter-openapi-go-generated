# MediaOriginalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FocusRects** | Pointer to [**[]MediaOriginalInfoFocusRect**](MediaOriginalInfoFocusRect.md) |  | [optional] 
**Height** | **int32** |  | 
**Width** | **int32** |  | 

## Methods

### NewMediaOriginalInfo

`func NewMediaOriginalInfo(height int32, width int32, ) *MediaOriginalInfo`

NewMediaOriginalInfo instantiates a new MediaOriginalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaOriginalInfoWithDefaults

`func NewMediaOriginalInfoWithDefaults() *MediaOriginalInfo`

NewMediaOriginalInfoWithDefaults instantiates a new MediaOriginalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFocusRects

`func (o *MediaOriginalInfo) GetFocusRects() []MediaOriginalInfoFocusRect`

GetFocusRects returns the FocusRects field if non-nil, zero value otherwise.

### GetFocusRectsOk

`func (o *MediaOriginalInfo) GetFocusRectsOk() (*[]MediaOriginalInfoFocusRect, bool)`

GetFocusRectsOk returns a tuple with the FocusRects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFocusRects

`func (o *MediaOriginalInfo) SetFocusRects(v []MediaOriginalInfoFocusRect)`

SetFocusRects sets FocusRects field to given value.

### HasFocusRects

`func (o *MediaOriginalInfo) HasFocusRects() bool`

HasFocusRects returns a boolean if a field has been set.

### GetHeight

`func (o *MediaOriginalInfo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *MediaOriginalInfo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *MediaOriginalInfo) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetWidth

`func (o *MediaOriginalInfo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *MediaOriginalInfo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *MediaOriginalInfo) SetWidth(v int32)`

SetWidth sets Width field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


