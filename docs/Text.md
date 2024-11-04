# Text

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]TextEntity**](TextEntity.md) |  | 
**Text** | **string** |  | 

## Methods

### NewText

`func NewText(entities []TextEntity, text string, ) *Text`

NewText instantiates a new Text object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextWithDefaults

`func NewTextWithDefaults() *Text`

NewTextWithDefaults instantiates a new Text object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *Text) GetEntities() []TextEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *Text) GetEntitiesOk() (*[]TextEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *Text) SetEntities(v []TextEntity)`

SetEntities sets Entities field to given value.


### GetText

`func (o *Text) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Text) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Text) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


