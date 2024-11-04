# PostCreateBookmarkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryId** | **string** |  | [default to "aoDbu3RHznuiSkQ9aNM67Q"]
**Variables** | [**PostCreateBookmarkRequestVariables**](PostCreateBookmarkRequestVariables.md) |  | 

## Methods

### NewPostCreateBookmarkRequest

`func NewPostCreateBookmarkRequest(queryId string, variables PostCreateBookmarkRequestVariables, ) *PostCreateBookmarkRequest`

NewPostCreateBookmarkRequest instantiates a new PostCreateBookmarkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCreateBookmarkRequestWithDefaults

`func NewPostCreateBookmarkRequestWithDefaults() *PostCreateBookmarkRequest`

NewPostCreateBookmarkRequestWithDefaults instantiates a new PostCreateBookmarkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryId

`func (o *PostCreateBookmarkRequest) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *PostCreateBookmarkRequest) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *PostCreateBookmarkRequest) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.


### GetVariables

`func (o *PostCreateBookmarkRequest) GetVariables() PostCreateBookmarkRequestVariables`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PostCreateBookmarkRequest) GetVariablesOk() (*PostCreateBookmarkRequestVariables, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PostCreateBookmarkRequest) SetVariables(v PostCreateBookmarkRequestVariables)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


