# PostCreateRetweetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryId** | **string** |  | [default to "ojPdsZsimiJrUGLR1sjUtA"]
**Variables** | [**PostCreateRetweetRequestVariables**](PostCreateRetweetRequestVariables.md) |  | 

## Methods

### NewPostCreateRetweetRequest

`func NewPostCreateRetweetRequest(queryId string, variables PostCreateRetweetRequestVariables, ) *PostCreateRetweetRequest`

NewPostCreateRetweetRequest instantiates a new PostCreateRetweetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCreateRetweetRequestWithDefaults

`func NewPostCreateRetweetRequestWithDefaults() *PostCreateRetweetRequest`

NewPostCreateRetweetRequestWithDefaults instantiates a new PostCreateRetweetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryId

`func (o *PostCreateRetweetRequest) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *PostCreateRetweetRequest) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *PostCreateRetweetRequest) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.


### GetVariables

`func (o *PostCreateRetweetRequest) GetVariables() PostCreateRetweetRequestVariables`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PostCreateRetweetRequest) GetVariablesOk() (*PostCreateRetweetRequestVariables, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PostCreateRetweetRequest) SetVariables(v PostCreateRetweetRequestVariables)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


