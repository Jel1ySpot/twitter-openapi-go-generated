# TimelineShowAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertType** | Pointer to **string** |  | [optional] 
**ColorConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**DisplayDurationMs** | Pointer to **int32** |  | [optional] 
**DisplayLocation** | Pointer to **string** |  | [optional] 
**IconDisplayInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**RichText** | [**TimelineShowAlertRichText**](TimelineShowAlertRichText.md) |  | 
**TriggerDelayMs** | Pointer to **int32** |  | [optional] 
**Type** | [**InstructionType**](InstructionType.md) |  | 
**UsersResults** | [**[]UserResults**](UserResults.md) |  | 

## Methods

### NewTimelineShowAlert

`func NewTimelineShowAlert(richText TimelineShowAlertRichText, type_ InstructionType, usersResults []UserResults, ) *TimelineShowAlert`

NewTimelineShowAlert instantiates a new TimelineShowAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineShowAlertWithDefaults

`func NewTimelineShowAlertWithDefaults() *TimelineShowAlert`

NewTimelineShowAlertWithDefaults instantiates a new TimelineShowAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertType

`func (o *TimelineShowAlert) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *TimelineShowAlert) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *TimelineShowAlert) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *TimelineShowAlert) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetColorConfig

`func (o *TimelineShowAlert) GetColorConfig() map[string]interface{}`

GetColorConfig returns the ColorConfig field if non-nil, zero value otherwise.

### GetColorConfigOk

`func (o *TimelineShowAlert) GetColorConfigOk() (*map[string]interface{}, bool)`

GetColorConfigOk returns a tuple with the ColorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorConfig

`func (o *TimelineShowAlert) SetColorConfig(v map[string]interface{})`

SetColorConfig sets ColorConfig field to given value.

### HasColorConfig

`func (o *TimelineShowAlert) HasColorConfig() bool`

HasColorConfig returns a boolean if a field has been set.

### GetDisplayDurationMs

`func (o *TimelineShowAlert) GetDisplayDurationMs() int32`

GetDisplayDurationMs returns the DisplayDurationMs field if non-nil, zero value otherwise.

### GetDisplayDurationMsOk

`func (o *TimelineShowAlert) GetDisplayDurationMsOk() (*int32, bool)`

GetDisplayDurationMsOk returns a tuple with the DisplayDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDurationMs

`func (o *TimelineShowAlert) SetDisplayDurationMs(v int32)`

SetDisplayDurationMs sets DisplayDurationMs field to given value.

### HasDisplayDurationMs

`func (o *TimelineShowAlert) HasDisplayDurationMs() bool`

HasDisplayDurationMs returns a boolean if a field has been set.

### GetDisplayLocation

`func (o *TimelineShowAlert) GetDisplayLocation() string`

GetDisplayLocation returns the DisplayLocation field if non-nil, zero value otherwise.

### GetDisplayLocationOk

`func (o *TimelineShowAlert) GetDisplayLocationOk() (*string, bool)`

GetDisplayLocationOk returns a tuple with the DisplayLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLocation

`func (o *TimelineShowAlert) SetDisplayLocation(v string)`

SetDisplayLocation sets DisplayLocation field to given value.

### HasDisplayLocation

`func (o *TimelineShowAlert) HasDisplayLocation() bool`

HasDisplayLocation returns a boolean if a field has been set.

### GetIconDisplayInfo

`func (o *TimelineShowAlert) GetIconDisplayInfo() map[string]interface{}`

GetIconDisplayInfo returns the IconDisplayInfo field if non-nil, zero value otherwise.

### GetIconDisplayInfoOk

`func (o *TimelineShowAlert) GetIconDisplayInfoOk() (*map[string]interface{}, bool)`

GetIconDisplayInfoOk returns a tuple with the IconDisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconDisplayInfo

`func (o *TimelineShowAlert) SetIconDisplayInfo(v map[string]interface{})`

SetIconDisplayInfo sets IconDisplayInfo field to given value.

### HasIconDisplayInfo

`func (o *TimelineShowAlert) HasIconDisplayInfo() bool`

HasIconDisplayInfo returns a boolean if a field has been set.

### GetRichText

`func (o *TimelineShowAlert) GetRichText() TimelineShowAlertRichText`

GetRichText returns the RichText field if non-nil, zero value otherwise.

### GetRichTextOk

`func (o *TimelineShowAlert) GetRichTextOk() (*TimelineShowAlertRichText, bool)`

GetRichTextOk returns a tuple with the RichText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichText

`func (o *TimelineShowAlert) SetRichText(v TimelineShowAlertRichText)`

SetRichText sets RichText field to given value.


### GetTriggerDelayMs

`func (o *TimelineShowAlert) GetTriggerDelayMs() int32`

GetTriggerDelayMs returns the TriggerDelayMs field if non-nil, zero value otherwise.

### GetTriggerDelayMsOk

`func (o *TimelineShowAlert) GetTriggerDelayMsOk() (*int32, bool)`

GetTriggerDelayMsOk returns a tuple with the TriggerDelayMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerDelayMs

`func (o *TimelineShowAlert) SetTriggerDelayMs(v int32)`

SetTriggerDelayMs sets TriggerDelayMs field to given value.

### HasTriggerDelayMs

`func (o *TimelineShowAlert) HasTriggerDelayMs() bool`

HasTriggerDelayMs returns a boolean if a field has been set.

### GetType

`func (o *TimelineShowAlert) GetType() InstructionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimelineShowAlert) GetTypeOk() (*InstructionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimelineShowAlert) SetType(v InstructionType)`

SetType sets Type field to given value.


### GetUsersResults

`func (o *TimelineShowAlert) GetUsersResults() []UserResults`

GetUsersResults returns the UsersResults field if non-nil, zero value otherwise.

### GetUsersResultsOk

`func (o *TimelineShowAlert) GetUsersResultsOk() (*[]UserResults, bool)`

GetUsersResultsOk returns a tuple with the UsersResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersResults

`func (o *TimelineShowAlert) SetUsersResults(v []UserResults)`

SetUsersResults sets UsersResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


