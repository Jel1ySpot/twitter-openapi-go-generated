# NoteTweetResultData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitySet** | [**Entities**](Entities.md) |  | 
**Id** | **string** |  | 
**Media** | Pointer to [**NoteTweetResultMedia**](NoteTweetResultMedia.md) |  | [optional] 
**Richtext** | Pointer to [**NoteTweetResultRichText**](NoteTweetResultRichText.md) |  | [optional] 
**Text** | **string** |  | 

## Methods

### NewNoteTweetResultData

`func NewNoteTweetResultData(entitySet Entities, id string, text string, ) *NoteTweetResultData`

NewNoteTweetResultData instantiates a new NoteTweetResultData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteTweetResultDataWithDefaults

`func NewNoteTweetResultDataWithDefaults() *NoteTweetResultData`

NewNoteTweetResultDataWithDefaults instantiates a new NoteTweetResultData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitySet

`func (o *NoteTweetResultData) GetEntitySet() Entities`

GetEntitySet returns the EntitySet field if non-nil, zero value otherwise.

### GetEntitySetOk

`func (o *NoteTweetResultData) GetEntitySetOk() (*Entities, bool)`

GetEntitySetOk returns a tuple with the EntitySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySet

`func (o *NoteTweetResultData) SetEntitySet(v Entities)`

SetEntitySet sets EntitySet field to given value.


### GetId

`func (o *NoteTweetResultData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NoteTweetResultData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NoteTweetResultData) SetId(v string)`

SetId sets Id field to given value.


### GetMedia

`func (o *NoteTweetResultData) GetMedia() NoteTweetResultMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *NoteTweetResultData) GetMediaOk() (*NoteTweetResultMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *NoteTweetResultData) SetMedia(v NoteTweetResultMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *NoteTweetResultData) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetRichtext

`func (o *NoteTweetResultData) GetRichtext() NoteTweetResultRichText`

GetRichtext returns the Richtext field if non-nil, zero value otherwise.

### GetRichtextOk

`func (o *NoteTweetResultData) GetRichtextOk() (*NoteTweetResultRichText, bool)`

GetRichtextOk returns a tuple with the Richtext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichtext

`func (o *NoteTweetResultData) SetRichtext(v NoteTweetResultRichText)`

SetRichtext sets Richtext field to given value.

### HasRichtext

`func (o *NoteTweetResultData) HasRichtext() bool`

HasRichtext returns a boolean if a field has been set.

### GetText

`func (o *NoteTweetResultData) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *NoteTweetResultData) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *NoteTweetResultData) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


