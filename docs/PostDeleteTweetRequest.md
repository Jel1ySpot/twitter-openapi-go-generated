# PostDeleteTweetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryId** | **string** |  | [default to "VaenaVgh5q5ih7kvyVjgtg"]
**Variables** | [**PostCreateRetweetRequestVariables**](PostCreateRetweetRequestVariables.md) |  | 

## Methods

### NewPostDeleteTweetRequest

`func NewPostDeleteTweetRequest(queryId string, variables PostCreateRetweetRequestVariables, ) *PostDeleteTweetRequest`

NewPostDeleteTweetRequest instantiates a new PostDeleteTweetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostDeleteTweetRequestWithDefaults

`func NewPostDeleteTweetRequestWithDefaults() *PostDeleteTweetRequest`

NewPostDeleteTweetRequestWithDefaults instantiates a new PostDeleteTweetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryId

`func (o *PostDeleteTweetRequest) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *PostDeleteTweetRequest) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *PostDeleteTweetRequest) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.


### GetVariables

`func (o *PostDeleteTweetRequest) GetVariables() PostCreateRetweetRequestVariables`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PostDeleteTweetRequest) GetVariablesOk() (*PostCreateRetweetRequestVariables, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PostDeleteTweetRequest) SetVariables(v PostCreateRetweetRequestVariables)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


