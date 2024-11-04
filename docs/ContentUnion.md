# ContentUnion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Typename** | [**TypeName**](TypeName.md) |  | 
**ClientEventInfo** | **map[string]interface{}** |  | 
**EntryType** | [**ContentEntryType**](ContentEntryType.md) |  | 
**FeedbackInfo** | Pointer to [**FeedbackInfo**](FeedbackInfo.md) |  | [optional] 
**ItemContent** | [**ItemContentUnion**](ItemContentUnion.md) |  | 
**DisplayType** | [**DisplayType**](DisplayType.md) |  | 
**Footer** | Pointer to **map[string]interface{}** |  | [optional] 
**Header** | Pointer to **map[string]interface{}** |  | [optional] 
**Items** | Pointer to [**[]ModuleItem**](ModuleItem.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CursorType** | [**CursorType**](CursorType.md) |  | 
**DisplayTreatment** | Pointer to [**DisplayTreatment**](DisplayTreatment.md) |  | [optional] 
**ItemType** | Pointer to [**ContentEntryType**](ContentEntryType.md) |  | [optional] 
**StopOnEmptyResponse** | Pointer to **bool** |  | [optional] 
**Value** | **string** |  | 

## Methods

### NewContentUnion

`func NewContentUnion(typename TypeName, clientEventInfo map[string]interface{}, entryType ContentEntryType, itemContent ItemContentUnion, displayType DisplayType, cursorType CursorType, value string, ) *ContentUnion`

NewContentUnion instantiates a new ContentUnion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentUnionWithDefaults

`func NewContentUnionWithDefaults() *ContentUnion`

NewContentUnionWithDefaults instantiates a new ContentUnion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypename

`func (o *ContentUnion) GetTypename() TypeName`

GetTypename returns the Typename field if non-nil, zero value otherwise.

### GetTypenameOk

`func (o *ContentUnion) GetTypenameOk() (*TypeName, bool)`

GetTypenameOk returns a tuple with the Typename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypename

`func (o *ContentUnion) SetTypename(v TypeName)`

SetTypename sets Typename field to given value.


### GetClientEventInfo

`func (o *ContentUnion) GetClientEventInfo() map[string]interface{}`

GetClientEventInfo returns the ClientEventInfo field if non-nil, zero value otherwise.

### GetClientEventInfoOk

`func (o *ContentUnion) GetClientEventInfoOk() (*map[string]interface{}, bool)`

GetClientEventInfoOk returns a tuple with the ClientEventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEventInfo

`func (o *ContentUnion) SetClientEventInfo(v map[string]interface{})`

SetClientEventInfo sets ClientEventInfo field to given value.


### GetEntryType

`func (o *ContentUnion) GetEntryType() ContentEntryType`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *ContentUnion) GetEntryTypeOk() (*ContentEntryType, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *ContentUnion) SetEntryType(v ContentEntryType)`

SetEntryType sets EntryType field to given value.


### GetFeedbackInfo

`func (o *ContentUnion) GetFeedbackInfo() FeedbackInfo`

GetFeedbackInfo returns the FeedbackInfo field if non-nil, zero value otherwise.

### GetFeedbackInfoOk

`func (o *ContentUnion) GetFeedbackInfoOk() (*FeedbackInfo, bool)`

GetFeedbackInfoOk returns a tuple with the FeedbackInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackInfo

`func (o *ContentUnion) SetFeedbackInfo(v FeedbackInfo)`

SetFeedbackInfo sets FeedbackInfo field to given value.

### HasFeedbackInfo

`func (o *ContentUnion) HasFeedbackInfo() bool`

HasFeedbackInfo returns a boolean if a field has been set.

### GetItemContent

`func (o *ContentUnion) GetItemContent() ItemContentUnion`

GetItemContent returns the ItemContent field if non-nil, zero value otherwise.

### GetItemContentOk

`func (o *ContentUnion) GetItemContentOk() (*ItemContentUnion, bool)`

GetItemContentOk returns a tuple with the ItemContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemContent

`func (o *ContentUnion) SetItemContent(v ItemContentUnion)`

SetItemContent sets ItemContent field to given value.


### GetDisplayType

`func (o *ContentUnion) GetDisplayType() DisplayType`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *ContentUnion) GetDisplayTypeOk() (*DisplayType, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *ContentUnion) SetDisplayType(v DisplayType)`

SetDisplayType sets DisplayType field to given value.


### GetFooter

`func (o *ContentUnion) GetFooter() map[string]interface{}`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *ContentUnion) GetFooterOk() (*map[string]interface{}, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *ContentUnion) SetFooter(v map[string]interface{})`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *ContentUnion) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetHeader

`func (o *ContentUnion) GetHeader() map[string]interface{}`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ContentUnion) GetHeaderOk() (*map[string]interface{}, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ContentUnion) SetHeader(v map[string]interface{})`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ContentUnion) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetItems

`func (o *ContentUnion) GetItems() []ModuleItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContentUnion) GetItemsOk() (*[]ModuleItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContentUnion) SetItems(v []ModuleItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *ContentUnion) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMetadata

`func (o *ContentUnion) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ContentUnion) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ContentUnion) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ContentUnion) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCursorType

`func (o *ContentUnion) GetCursorType() CursorType`

GetCursorType returns the CursorType field if non-nil, zero value otherwise.

### GetCursorTypeOk

`func (o *ContentUnion) GetCursorTypeOk() (*CursorType, bool)`

GetCursorTypeOk returns a tuple with the CursorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorType

`func (o *ContentUnion) SetCursorType(v CursorType)`

SetCursorType sets CursorType field to given value.


### GetDisplayTreatment

`func (o *ContentUnion) GetDisplayTreatment() DisplayTreatment`

GetDisplayTreatment returns the DisplayTreatment field if non-nil, zero value otherwise.

### GetDisplayTreatmentOk

`func (o *ContentUnion) GetDisplayTreatmentOk() (*DisplayTreatment, bool)`

GetDisplayTreatmentOk returns a tuple with the DisplayTreatment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTreatment

`func (o *ContentUnion) SetDisplayTreatment(v DisplayTreatment)`

SetDisplayTreatment sets DisplayTreatment field to given value.

### HasDisplayTreatment

`func (o *ContentUnion) HasDisplayTreatment() bool`

HasDisplayTreatment returns a boolean if a field has been set.

### GetItemType

`func (o *ContentUnion) GetItemType() ContentEntryType`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *ContentUnion) GetItemTypeOk() (*ContentEntryType, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *ContentUnion) SetItemType(v ContentEntryType)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *ContentUnion) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetStopOnEmptyResponse

`func (o *ContentUnion) GetStopOnEmptyResponse() bool`

GetStopOnEmptyResponse returns the StopOnEmptyResponse field if non-nil, zero value otherwise.

### GetStopOnEmptyResponseOk

`func (o *ContentUnion) GetStopOnEmptyResponseOk() (*bool, bool)`

GetStopOnEmptyResponseOk returns a tuple with the StopOnEmptyResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOnEmptyResponse

`func (o *ContentUnion) SetStopOnEmptyResponse(v bool)`

SetStopOnEmptyResponse sets StopOnEmptyResponse field to given value.

### HasStopOnEmptyResponse

`func (o *ContentUnion) HasStopOnEmptyResponse() bool`

HasStopOnEmptyResponse returns a boolean if a field has been set.

### GetValue

`func (o *ContentUnion) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContentUnion) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContentUnion) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


