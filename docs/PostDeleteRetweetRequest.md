# PostDeleteRetweetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryId** | **string** |  | [default to "iQtK4dl5hBmXewYZuEOKVw"]
**Variables** | [**PostDeleteRetweetRequestVariables**](PostDeleteRetweetRequestVariables.md) |  | 

## Methods

### NewPostDeleteRetweetRequest

`func NewPostDeleteRetweetRequest(queryId string, variables PostDeleteRetweetRequestVariables, ) *PostDeleteRetweetRequest`

NewPostDeleteRetweetRequest instantiates a new PostDeleteRetweetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostDeleteRetweetRequestWithDefaults

`func NewPostDeleteRetweetRequestWithDefaults() *PostDeleteRetweetRequest`

NewPostDeleteRetweetRequestWithDefaults instantiates a new PostDeleteRetweetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryId

`func (o *PostDeleteRetweetRequest) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *PostDeleteRetweetRequest) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *PostDeleteRetweetRequest) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.


### GetVariables

`func (o *PostDeleteRetweetRequest) GetVariables() PostDeleteRetweetRequestVariables`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PostDeleteRetweetRequest) GetVariablesOk() (*PostDeleteRetweetRequestVariables, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PostDeleteRetweetRequest) SetVariables(v PostDeleteRetweetRequestVariables)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


