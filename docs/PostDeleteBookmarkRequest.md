# PostDeleteBookmarkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryId** | **string** |  | [default to "Wlmlj2-xzyS1GN3a6cj-mQ"]
**Variables** | [**PostCreateBookmarkRequestVariables**](PostCreateBookmarkRequestVariables.md) |  | 

## Methods

### NewPostDeleteBookmarkRequest

`func NewPostDeleteBookmarkRequest(queryId string, variables PostCreateBookmarkRequestVariables, ) *PostDeleteBookmarkRequest`

NewPostDeleteBookmarkRequest instantiates a new PostDeleteBookmarkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostDeleteBookmarkRequestWithDefaults

`func NewPostDeleteBookmarkRequestWithDefaults() *PostDeleteBookmarkRequest`

NewPostDeleteBookmarkRequestWithDefaults instantiates a new PostDeleteBookmarkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryId

`func (o *PostDeleteBookmarkRequest) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *PostDeleteBookmarkRequest) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *PostDeleteBookmarkRequest) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.


### GetVariables

`func (o *PostDeleteBookmarkRequest) GetVariables() PostCreateBookmarkRequestVariables`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PostDeleteBookmarkRequest) GetVariablesOk() (*PostCreateBookmarkRequestVariables, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PostDeleteBookmarkRequest) SetVariables(v PostCreateBookmarkRequestVariables)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


