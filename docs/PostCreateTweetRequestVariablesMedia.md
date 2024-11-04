# PostCreateTweetRequestVariablesMedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaEntities** | Pointer to [**[]PostCreateTweetRequestVariablesMediaMediaEntitiesInner**](PostCreateTweetRequestVariablesMediaMediaEntitiesInner.md) |  | [optional] 
**PossiblySensitive** | **bool** |  | [default to false]

## Methods

### NewPostCreateTweetRequestVariablesMedia

`func NewPostCreateTweetRequestVariablesMedia(possiblySensitive bool, ) *PostCreateTweetRequestVariablesMedia`

NewPostCreateTweetRequestVariablesMedia instantiates a new PostCreateTweetRequestVariablesMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCreateTweetRequestVariablesMediaWithDefaults

`func NewPostCreateTweetRequestVariablesMediaWithDefaults() *PostCreateTweetRequestVariablesMedia`

NewPostCreateTweetRequestVariablesMediaWithDefaults instantiates a new PostCreateTweetRequestVariablesMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaEntities

`func (o *PostCreateTweetRequestVariablesMedia) GetMediaEntities() []PostCreateTweetRequestVariablesMediaMediaEntitiesInner`

GetMediaEntities returns the MediaEntities field if non-nil, zero value otherwise.

### GetMediaEntitiesOk

`func (o *PostCreateTweetRequestVariablesMedia) GetMediaEntitiesOk() (*[]PostCreateTweetRequestVariablesMediaMediaEntitiesInner, bool)`

GetMediaEntitiesOk returns a tuple with the MediaEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaEntities

`func (o *PostCreateTweetRequestVariablesMedia) SetMediaEntities(v []PostCreateTweetRequestVariablesMediaMediaEntitiesInner)`

SetMediaEntities sets MediaEntities field to given value.

### HasMediaEntities

`func (o *PostCreateTweetRequestVariablesMedia) HasMediaEntities() bool`

HasMediaEntities returns a boolean if a field has been set.

### GetPossiblySensitive

`func (o *PostCreateTweetRequestVariablesMedia) GetPossiblySensitive() bool`

GetPossiblySensitive returns the PossiblySensitive field if non-nil, zero value otherwise.

### GetPossiblySensitiveOk

`func (o *PostCreateTweetRequestVariablesMedia) GetPossiblySensitiveOk() (*bool, bool)`

GetPossiblySensitiveOk returns a tuple with the PossiblySensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossiblySensitive

`func (o *PostCreateTweetRequestVariablesMedia) SetPossiblySensitive(v bool)`

SetPossiblySensitive sets PossiblySensitive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


