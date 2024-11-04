# ArticleResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoverMedia** | [**ArticleCoverMedia**](ArticleCoverMedia.md) |  | 
**Id** | **string** |  | 
**LifecycleState** | Pointer to [**ArticleLifecycleState**](ArticleLifecycleState.md) |  | [optional] 
**Metadata** | [**ArticleMetadata**](ArticleMetadata.md) |  | 
**PreviewText** | **string** |  | 
**RestId** | **string** |  | 
**Title** | **string** |  | 

## Methods

### NewArticleResult

`func NewArticleResult(coverMedia ArticleCoverMedia, id string, metadata ArticleMetadata, previewText string, restId string, title string, ) *ArticleResult`

NewArticleResult instantiates a new ArticleResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleResultWithDefaults

`func NewArticleResultWithDefaults() *ArticleResult`

NewArticleResultWithDefaults instantiates a new ArticleResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoverMedia

`func (o *ArticleResult) GetCoverMedia() ArticleCoverMedia`

GetCoverMedia returns the CoverMedia field if non-nil, zero value otherwise.

### GetCoverMediaOk

`func (o *ArticleResult) GetCoverMediaOk() (*ArticleCoverMedia, bool)`

GetCoverMediaOk returns a tuple with the CoverMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverMedia

`func (o *ArticleResult) SetCoverMedia(v ArticleCoverMedia)`

SetCoverMedia sets CoverMedia field to given value.


### GetId

`func (o *ArticleResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArticleResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArticleResult) SetId(v string)`

SetId sets Id field to given value.


### GetLifecycleState

`func (o *ArticleResult) GetLifecycleState() ArticleLifecycleState`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *ArticleResult) GetLifecycleStateOk() (*ArticleLifecycleState, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *ArticleResult) SetLifecycleState(v ArticleLifecycleState)`

SetLifecycleState sets LifecycleState field to given value.

### HasLifecycleState

`func (o *ArticleResult) HasLifecycleState() bool`

HasLifecycleState returns a boolean if a field has been set.

### GetMetadata

`func (o *ArticleResult) GetMetadata() ArticleMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArticleResult) GetMetadataOk() (*ArticleMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArticleResult) SetMetadata(v ArticleMetadata)`

SetMetadata sets Metadata field to given value.


### GetPreviewText

`func (o *ArticleResult) GetPreviewText() string`

GetPreviewText returns the PreviewText field if non-nil, zero value otherwise.

### GetPreviewTextOk

`func (o *ArticleResult) GetPreviewTextOk() (*string, bool)`

GetPreviewTextOk returns a tuple with the PreviewText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewText

`func (o *ArticleResult) SetPreviewText(v string)`

SetPreviewText sets PreviewText field to given value.


### GetRestId

`func (o *ArticleResult) GetRestId() string`

GetRestId returns the RestId field if non-nil, zero value otherwise.

### GetRestIdOk

`func (o *ArticleResult) GetRestIdOk() (*string, bool)`

GetRestIdOk returns a tuple with the RestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestId

`func (o *ArticleResult) SetRestId(v string)`

SetRestId sets RestId field to given value.


### GetTitle

`func (o *ArticleResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ArticleResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ArticleResult) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


