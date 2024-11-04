# ItemResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Typename** | Pointer to [**TypeName**](TypeName.md) |  | [optional] 
**Result** | Pointer to [**TweetUnion**](TweetUnion.md) |  | [optional] 

## Methods

### NewItemResult

`func NewItemResult() *ItemResult`

NewItemResult instantiates a new ItemResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemResultWithDefaults

`func NewItemResultWithDefaults() *ItemResult`

NewItemResultWithDefaults instantiates a new ItemResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypename

`func (o *ItemResult) GetTypename() TypeName`

GetTypename returns the Typename field if non-nil, zero value otherwise.

### GetTypenameOk

`func (o *ItemResult) GetTypenameOk() (*TypeName, bool)`

GetTypenameOk returns a tuple with the Typename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypename

`func (o *ItemResult) SetTypename(v TypeName)`

SetTypename sets Typename field to given value.

### HasTypename

`func (o *ItemResult) HasTypename() bool`

HasTypename returns a boolean if a field has been set.

### GetResult

`func (o *ItemResult) GetResult() TweetUnion`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ItemResult) GetResultOk() (*TweetUnion, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ItemResult) SetResult(v TweetUnion)`

SetResult sets Result field to given value.

### HasResult

`func (o *ItemResult) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


