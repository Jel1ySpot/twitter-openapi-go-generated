# PostCreateTweetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | [**PostCreateTweetRequestFeatures**](PostCreateTweetRequestFeatures.md) |  | 
**QueryId** | **string** |  | [default to "xT36w0XM3A8jDynpkram2A"]
**Variables** | [**PostCreateTweetRequestVariables**](PostCreateTweetRequestVariables.md) |  | 

## Methods

### NewPostCreateTweetRequest

`func NewPostCreateTweetRequest(features PostCreateTweetRequestFeatures, queryId string, variables PostCreateTweetRequestVariables, ) *PostCreateTweetRequest`

NewPostCreateTweetRequest instantiates a new PostCreateTweetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCreateTweetRequestWithDefaults

`func NewPostCreateTweetRequestWithDefaults() *PostCreateTweetRequest`

NewPostCreateTweetRequestWithDefaults instantiates a new PostCreateTweetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *PostCreateTweetRequest) GetFeatures() PostCreateTweetRequestFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PostCreateTweetRequest) GetFeaturesOk() (*PostCreateTweetRequestFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PostCreateTweetRequest) SetFeatures(v PostCreateTweetRequestFeatures)`

SetFeatures sets Features field to given value.


### GetQueryId

`func (o *PostCreateTweetRequest) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *PostCreateTweetRequest) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *PostCreateTweetRequest) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.


### GetVariables

`func (o *PostCreateTweetRequest) GetVariables() PostCreateTweetRequestVariables`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PostCreateTweetRequest) GetVariablesOk() (*PostCreateTweetRequestVariables, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PostCreateTweetRequest) SetVariables(v PostCreateTweetRequestVariables)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


