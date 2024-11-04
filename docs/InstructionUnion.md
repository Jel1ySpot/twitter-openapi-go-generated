# InstructionUnion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | [**[]TimelineAddEntry**](TimelineAddEntry.md) |  | 
**Type** | [**InstructionType**](InstructionType.md) |  | 
**ModuleEntryId** | **string** |  | 
**ModuleItems** | [**[]ModuleItem**](ModuleItem.md) |  | 
**Prepend** | Pointer to **bool** |  | [optional] 
**Entry** | [**TimelineAddEntry**](TimelineAddEntry.md) |  | 
**EntryIdToReplace** | **string** |  | 
**AlertType** | Pointer to **string** |  | [optional] 
**ColorConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**DisplayDurationMs** | Pointer to **int32** |  | [optional] 
**DisplayLocation** | Pointer to **string** |  | [optional] 
**IconDisplayInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**RichText** | [**TimelineShowAlertRichText**](TimelineShowAlertRichText.md) |  | 
**TriggerDelayMs** | Pointer to **int32** |  | [optional] 
**UsersResults** | [**[]UserResults**](UserResults.md) |  | 
**Direction** | **string** |  | 
**ClientEventInfo** | [**ClientEventInfo**](ClientEventInfo.md) |  | 
**Cover** | [**TimelineHalfCover**](TimelineHalfCover.md) |  | 

## Methods

### NewInstructionUnion

`func NewInstructionUnion(entries []TimelineAddEntry, type_ InstructionType, moduleEntryId string, moduleItems []ModuleItem, entry TimelineAddEntry, entryIdToReplace string, richText TimelineShowAlertRichText, usersResults []UserResults, direction string, clientEventInfo ClientEventInfo, cover TimelineHalfCover, ) *InstructionUnion`

NewInstructionUnion instantiates a new InstructionUnion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstructionUnionWithDefaults

`func NewInstructionUnionWithDefaults() *InstructionUnion`

NewInstructionUnionWithDefaults instantiates a new InstructionUnion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *InstructionUnion) GetEntries() []TimelineAddEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *InstructionUnion) GetEntriesOk() (*[]TimelineAddEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *InstructionUnion) SetEntries(v []TimelineAddEntry)`

SetEntries sets Entries field to given value.


### GetType

`func (o *InstructionUnion) GetType() InstructionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstructionUnion) GetTypeOk() (*InstructionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstructionUnion) SetType(v InstructionType)`

SetType sets Type field to given value.


### GetModuleEntryId

`func (o *InstructionUnion) GetModuleEntryId() string`

GetModuleEntryId returns the ModuleEntryId field if non-nil, zero value otherwise.

### GetModuleEntryIdOk

`func (o *InstructionUnion) GetModuleEntryIdOk() (*string, bool)`

GetModuleEntryIdOk returns a tuple with the ModuleEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleEntryId

`func (o *InstructionUnion) SetModuleEntryId(v string)`

SetModuleEntryId sets ModuleEntryId field to given value.


### GetModuleItems

`func (o *InstructionUnion) GetModuleItems() []ModuleItem`

GetModuleItems returns the ModuleItems field if non-nil, zero value otherwise.

### GetModuleItemsOk

`func (o *InstructionUnion) GetModuleItemsOk() (*[]ModuleItem, bool)`

GetModuleItemsOk returns a tuple with the ModuleItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleItems

`func (o *InstructionUnion) SetModuleItems(v []ModuleItem)`

SetModuleItems sets ModuleItems field to given value.


### GetPrepend

`func (o *InstructionUnion) GetPrepend() bool`

GetPrepend returns the Prepend field if non-nil, zero value otherwise.

### GetPrependOk

`func (o *InstructionUnion) GetPrependOk() (*bool, bool)`

GetPrependOk returns a tuple with the Prepend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepend

`func (o *InstructionUnion) SetPrepend(v bool)`

SetPrepend sets Prepend field to given value.

### HasPrepend

`func (o *InstructionUnion) HasPrepend() bool`

HasPrepend returns a boolean if a field has been set.

### GetEntry

`func (o *InstructionUnion) GetEntry() TimelineAddEntry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *InstructionUnion) GetEntryOk() (*TimelineAddEntry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *InstructionUnion) SetEntry(v TimelineAddEntry)`

SetEntry sets Entry field to given value.


### GetEntryIdToReplace

`func (o *InstructionUnion) GetEntryIdToReplace() string`

GetEntryIdToReplace returns the EntryIdToReplace field if non-nil, zero value otherwise.

### GetEntryIdToReplaceOk

`func (o *InstructionUnion) GetEntryIdToReplaceOk() (*string, bool)`

GetEntryIdToReplaceOk returns a tuple with the EntryIdToReplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryIdToReplace

`func (o *InstructionUnion) SetEntryIdToReplace(v string)`

SetEntryIdToReplace sets EntryIdToReplace field to given value.


### GetAlertType

`func (o *InstructionUnion) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *InstructionUnion) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *InstructionUnion) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *InstructionUnion) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetColorConfig

`func (o *InstructionUnion) GetColorConfig() map[string]interface{}`

GetColorConfig returns the ColorConfig field if non-nil, zero value otherwise.

### GetColorConfigOk

`func (o *InstructionUnion) GetColorConfigOk() (*map[string]interface{}, bool)`

GetColorConfigOk returns a tuple with the ColorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorConfig

`func (o *InstructionUnion) SetColorConfig(v map[string]interface{})`

SetColorConfig sets ColorConfig field to given value.

### HasColorConfig

`func (o *InstructionUnion) HasColorConfig() bool`

HasColorConfig returns a boolean if a field has been set.

### GetDisplayDurationMs

`func (o *InstructionUnion) GetDisplayDurationMs() int32`

GetDisplayDurationMs returns the DisplayDurationMs field if non-nil, zero value otherwise.

### GetDisplayDurationMsOk

`func (o *InstructionUnion) GetDisplayDurationMsOk() (*int32, bool)`

GetDisplayDurationMsOk returns a tuple with the DisplayDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDurationMs

`func (o *InstructionUnion) SetDisplayDurationMs(v int32)`

SetDisplayDurationMs sets DisplayDurationMs field to given value.

### HasDisplayDurationMs

`func (o *InstructionUnion) HasDisplayDurationMs() bool`

HasDisplayDurationMs returns a boolean if a field has been set.

### GetDisplayLocation

`func (o *InstructionUnion) GetDisplayLocation() string`

GetDisplayLocation returns the DisplayLocation field if non-nil, zero value otherwise.

### GetDisplayLocationOk

`func (o *InstructionUnion) GetDisplayLocationOk() (*string, bool)`

GetDisplayLocationOk returns a tuple with the DisplayLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLocation

`func (o *InstructionUnion) SetDisplayLocation(v string)`

SetDisplayLocation sets DisplayLocation field to given value.

### HasDisplayLocation

`func (o *InstructionUnion) HasDisplayLocation() bool`

HasDisplayLocation returns a boolean if a field has been set.

### GetIconDisplayInfo

`func (o *InstructionUnion) GetIconDisplayInfo() map[string]interface{}`

GetIconDisplayInfo returns the IconDisplayInfo field if non-nil, zero value otherwise.

### GetIconDisplayInfoOk

`func (o *InstructionUnion) GetIconDisplayInfoOk() (*map[string]interface{}, bool)`

GetIconDisplayInfoOk returns a tuple with the IconDisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconDisplayInfo

`func (o *InstructionUnion) SetIconDisplayInfo(v map[string]interface{})`

SetIconDisplayInfo sets IconDisplayInfo field to given value.

### HasIconDisplayInfo

`func (o *InstructionUnion) HasIconDisplayInfo() bool`

HasIconDisplayInfo returns a boolean if a field has been set.

### GetRichText

`func (o *InstructionUnion) GetRichText() TimelineShowAlertRichText`

GetRichText returns the RichText field if non-nil, zero value otherwise.

### GetRichTextOk

`func (o *InstructionUnion) GetRichTextOk() (*TimelineShowAlertRichText, bool)`

GetRichTextOk returns a tuple with the RichText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichText

`func (o *InstructionUnion) SetRichText(v TimelineShowAlertRichText)`

SetRichText sets RichText field to given value.


### GetTriggerDelayMs

`func (o *InstructionUnion) GetTriggerDelayMs() int32`

GetTriggerDelayMs returns the TriggerDelayMs field if non-nil, zero value otherwise.

### GetTriggerDelayMsOk

`func (o *InstructionUnion) GetTriggerDelayMsOk() (*int32, bool)`

GetTriggerDelayMsOk returns a tuple with the TriggerDelayMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerDelayMs

`func (o *InstructionUnion) SetTriggerDelayMs(v int32)`

SetTriggerDelayMs sets TriggerDelayMs field to given value.

### HasTriggerDelayMs

`func (o *InstructionUnion) HasTriggerDelayMs() bool`

HasTriggerDelayMs returns a boolean if a field has been set.

### GetUsersResults

`func (o *InstructionUnion) GetUsersResults() []UserResults`

GetUsersResults returns the UsersResults field if non-nil, zero value otherwise.

### GetUsersResultsOk

`func (o *InstructionUnion) GetUsersResultsOk() (*[]UserResults, bool)`

GetUsersResultsOk returns a tuple with the UsersResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersResults

`func (o *InstructionUnion) SetUsersResults(v []UserResults)`

SetUsersResults sets UsersResults field to given value.


### GetDirection

`func (o *InstructionUnion) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *InstructionUnion) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *InstructionUnion) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetClientEventInfo

`func (o *InstructionUnion) GetClientEventInfo() ClientEventInfo`

GetClientEventInfo returns the ClientEventInfo field if non-nil, zero value otherwise.

### GetClientEventInfoOk

`func (o *InstructionUnion) GetClientEventInfoOk() (*ClientEventInfo, bool)`

GetClientEventInfoOk returns a tuple with the ClientEventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEventInfo

`func (o *InstructionUnion) SetClientEventInfo(v ClientEventInfo)`

SetClientEventInfo sets ClientEventInfo field to given value.


### GetCover

`func (o *InstructionUnion) GetCover() TimelineHalfCover`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *InstructionUnion) GetCoverOk() (*TimelineHalfCover, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *InstructionUnion) SetCover(v TimelineHalfCover)`

SetCover sets Cover field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


