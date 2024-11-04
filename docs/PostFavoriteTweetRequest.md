# PostFavoriteTweetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryId** | **string** |  | [default to "lI07N6Otwv1PhnEgXILM7A"]
**Variables** | [**PostCreateBookmarkRequestVariables**](PostCreateBookmarkRequestVariables.md) |  | 

## Methods

### NewPostFavoriteTweetRequest

`func NewPostFavoriteTweetRequest(queryId string, variables PostCreateBookmarkRequestVariables, ) *PostFavoriteTweetRequest`

NewPostFavoriteTweetRequest instantiates a new PostFavoriteTweetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostFavoriteTweetRequestWithDefaults

`func NewPostFavoriteTweetRequestWithDefaults() *PostFavoriteTweetRequest`

NewPostFavoriteTweetRequestWithDefaults instantiates a new PostFavoriteTweetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryId

`func (o *PostFavoriteTweetRequest) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *PostFavoriteTweetRequest) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *PostFavoriteTweetRequest) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.


### GetVariables

`func (o *PostFavoriteTweetRequest) GetVariables() PostCreateBookmarkRequestVariables`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PostFavoriteTweetRequest) GetVariablesOk() (*PostCreateBookmarkRequestVariables, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PostFavoriteTweetRequest) SetVariables(v PostCreateBookmarkRequestVariables)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


