# PostFavoriteTweet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ErrorsData**](ErrorsData.md) |  | 
**Errors** | [**[]Error**](Error.md) |  | 

## Methods

### NewPostFavoriteTweet200Response

`func NewPostFavoriteTweet200Response(data ErrorsData, errors []Error, ) *PostFavoriteTweet200Response`

NewPostFavoriteTweet200Response instantiates a new PostFavoriteTweet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostFavoriteTweet200ResponseWithDefaults

`func NewPostFavoriteTweet200ResponseWithDefaults() *PostFavoriteTweet200Response`

NewPostFavoriteTweet200ResponseWithDefaults instantiates a new PostFavoriteTweet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PostFavoriteTweet200Response) GetData() ErrorsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PostFavoriteTweet200Response) GetDataOk() (*ErrorsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PostFavoriteTweet200Response) SetData(v ErrorsData)`

SetData sets Data field to given value.


### GetErrors

`func (o *PostFavoriteTweet200Response) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PostFavoriteTweet200Response) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PostFavoriteTweet200Response) SetErrors(v []Error)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


