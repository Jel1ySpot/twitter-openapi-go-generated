# PostCreateBookmark200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ErrorsData**](ErrorsData.md) |  | 
**Errors** | [**[]Error**](Error.md) |  | 

## Methods

### NewPostCreateBookmark200Response

`func NewPostCreateBookmark200Response(data ErrorsData, errors []Error, ) *PostCreateBookmark200Response`

NewPostCreateBookmark200Response instantiates a new PostCreateBookmark200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCreateBookmark200ResponseWithDefaults

`func NewPostCreateBookmark200ResponseWithDefaults() *PostCreateBookmark200Response`

NewPostCreateBookmark200ResponseWithDefaults instantiates a new PostCreateBookmark200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PostCreateBookmark200Response) GetData() ErrorsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PostCreateBookmark200Response) GetDataOk() (*ErrorsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PostCreateBookmark200Response) SetData(v ErrorsData)`

SetData sets Data field to given value.


### GetErrors

`func (o *PostCreateBookmark200Response) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PostCreateBookmark200Response) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PostCreateBookmark200Response) SetErrors(v []Error)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

