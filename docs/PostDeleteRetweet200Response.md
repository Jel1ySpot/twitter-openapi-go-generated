# PostDeleteRetweet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ErrorsData**](ErrorsData.md) |  | 
**Errors** | [**[]Error**](Error.md) |  | 

## Methods

### NewPostDeleteRetweet200Response

`func NewPostDeleteRetweet200Response(data ErrorsData, errors []Error, ) *PostDeleteRetweet200Response`

NewPostDeleteRetweet200Response instantiates a new PostDeleteRetweet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostDeleteRetweet200ResponseWithDefaults

`func NewPostDeleteRetweet200ResponseWithDefaults() *PostDeleteRetweet200Response`

NewPostDeleteRetweet200ResponseWithDefaults instantiates a new PostDeleteRetweet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PostDeleteRetweet200Response) GetData() ErrorsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PostDeleteRetweet200Response) GetDataOk() (*ErrorsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PostDeleteRetweet200Response) SetData(v ErrorsData)`

SetData sets Data field to given value.


### GetErrors

`func (o *PostDeleteRetweet200Response) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PostDeleteRetweet200Response) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PostDeleteRetweet200Response) SetErrors(v []Error)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


