# PostUnfavoriteTweetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryId** | **string** |  | [default to "ZYKSe-w7KEslx3JhSIk5LA"]
**Variables** | [**PostCreateRetweetRequestVariables**](PostCreateRetweetRequestVariables.md) |  | 

## Methods

### NewPostUnfavoriteTweetRequest

`func NewPostUnfavoriteTweetRequest(queryId string, variables PostCreateRetweetRequestVariables, ) *PostUnfavoriteTweetRequest`

NewPostUnfavoriteTweetRequest instantiates a new PostUnfavoriteTweetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostUnfavoriteTweetRequestWithDefaults

`func NewPostUnfavoriteTweetRequestWithDefaults() *PostUnfavoriteTweetRequest`

NewPostUnfavoriteTweetRequestWithDefaults instantiates a new PostUnfavoriteTweetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryId

`func (o *PostUnfavoriteTweetRequest) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *PostUnfavoriteTweetRequest) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *PostUnfavoriteTweetRequest) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.


### GetVariables

`func (o *PostUnfavoriteTweetRequest) GetVariables() PostCreateRetweetRequestVariables`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PostUnfavoriteTweetRequest) GetVariablesOk() (*PostCreateRetweetRequestVariables, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PostUnfavoriteTweetRequest) SetVariables(v PostCreateRetweetRequestVariables)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


