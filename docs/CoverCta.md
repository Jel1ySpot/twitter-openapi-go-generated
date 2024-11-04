# CoverCta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** |  | [optional] 
**ButtonStyle** | Pointer to **string** |  | [optional] 
**Callbacks** | [**[]Callback**](Callback.md) |  | 
**ClientEventInfo** | [**CtaClientEventInfo**](CtaClientEventInfo.md) |  | 
**CtaBehavior** | [**TimelineCoverBehavior**](TimelineCoverBehavior.md) |  | 

## Methods

### NewCoverCta

`func NewCoverCta(callbacks []Callback, clientEventInfo CtaClientEventInfo, ctaBehavior TimelineCoverBehavior, ) *CoverCta`

NewCoverCta instantiates a new CoverCta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoverCtaWithDefaults

`func NewCoverCtaWithDefaults() *CoverCta`

NewCoverCtaWithDefaults instantiates a new CoverCta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *CoverCta) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CoverCta) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CoverCta) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CoverCta) HasText() bool`

HasText returns a boolean if a field has been set.

### GetButtonStyle

`func (o *CoverCta) GetButtonStyle() string`

GetButtonStyle returns the ButtonStyle field if non-nil, zero value otherwise.

### GetButtonStyleOk

`func (o *CoverCta) GetButtonStyleOk() (*string, bool)`

GetButtonStyleOk returns a tuple with the ButtonStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonStyle

`func (o *CoverCta) SetButtonStyle(v string)`

SetButtonStyle sets ButtonStyle field to given value.

### HasButtonStyle

`func (o *CoverCta) HasButtonStyle() bool`

HasButtonStyle returns a boolean if a field has been set.

### GetCallbacks

`func (o *CoverCta) GetCallbacks() []Callback`

GetCallbacks returns the Callbacks field if non-nil, zero value otherwise.

### GetCallbacksOk

`func (o *CoverCta) GetCallbacksOk() (*[]Callback, bool)`

GetCallbacksOk returns a tuple with the Callbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbacks

`func (o *CoverCta) SetCallbacks(v []Callback)`

SetCallbacks sets Callbacks field to given value.


### GetClientEventInfo

`func (o *CoverCta) GetClientEventInfo() CtaClientEventInfo`

GetClientEventInfo returns the ClientEventInfo field if non-nil, zero value otherwise.

### GetClientEventInfoOk

`func (o *CoverCta) GetClientEventInfoOk() (*CtaClientEventInfo, bool)`

GetClientEventInfoOk returns a tuple with the ClientEventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEventInfo

`func (o *CoverCta) SetClientEventInfo(v CtaClientEventInfo)`

SetClientEventInfo sets ClientEventInfo field to given value.


### GetCtaBehavior

`func (o *CoverCta) GetCtaBehavior() TimelineCoverBehavior`

GetCtaBehavior returns the CtaBehavior field if non-nil, zero value otherwise.

### GetCtaBehaviorOk

`func (o *CoverCta) GetCtaBehaviorOk() (*TimelineCoverBehavior, bool)`

GetCtaBehaviorOk returns a tuple with the CtaBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtaBehavior

`func (o *CoverCta) SetCtaBehavior(v TimelineCoverBehavior)`

SetCtaBehavior sets CtaBehavior field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


