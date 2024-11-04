# Media

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalMediaInfo** | Pointer to [**AdditionalMediaInfo**](AdditionalMediaInfo.md) |  | [optional] 
**AllowDownloadStatus** | Pointer to [**AllowDownloadStatus**](AllowDownloadStatus.md) |  | [optional] 
**DisplayUrl** | **string** |  | 
**ExpandedUrl** | **string** |  | 
**ExtAltText** | Pointer to **string** |  | [optional] 
**ExtMediaAvailability** | [**ExtMediaAvailability**](ExtMediaAvailability.md) |  | 
**Features** | Pointer to **map[string]interface{}** |  | [optional] 
**IdStr** | **string** |  | 
**Indices** | **[]int32** |  | 
**MediaKey** | **string** |  | 
**MediaResults** | Pointer to [**MediaResults**](MediaResults.md) |  | [optional] 
**MediaUrlHttps** | **string** |  | 
**OriginalInfo** | [**MediaOriginalInfo**](MediaOriginalInfo.md) |  | 
**SensitiveMediaWarning** | Pointer to [**SensitiveMediaWarning**](SensitiveMediaWarning.md) |  | [optional] 
**Sizes** | [**MediaSizes**](MediaSizes.md) |  | 
**SourceStatusIdStr** | Pointer to **string** |  | [optional] 
**SourceUserIdStr** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Url** | **string** |  | 
**VideoInfo** | Pointer to [**MediaVideoInfo**](MediaVideoInfo.md) |  | [optional] 

## Methods

### NewMedia

`func NewMedia(displayUrl string, expandedUrl string, extMediaAvailability ExtMediaAvailability, idStr string, indices []int32, mediaKey string, mediaUrlHttps string, originalInfo MediaOriginalInfo, sizes MediaSizes, type_ string, url string, ) *Media`

NewMedia instantiates a new Media object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaWithDefaults

`func NewMediaWithDefaults() *Media`

NewMediaWithDefaults instantiates a new Media object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalMediaInfo

`func (o *Media) GetAdditionalMediaInfo() AdditionalMediaInfo`

GetAdditionalMediaInfo returns the AdditionalMediaInfo field if non-nil, zero value otherwise.

### GetAdditionalMediaInfoOk

`func (o *Media) GetAdditionalMediaInfoOk() (*AdditionalMediaInfo, bool)`

GetAdditionalMediaInfoOk returns a tuple with the AdditionalMediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMediaInfo

`func (o *Media) SetAdditionalMediaInfo(v AdditionalMediaInfo)`

SetAdditionalMediaInfo sets AdditionalMediaInfo field to given value.

### HasAdditionalMediaInfo

`func (o *Media) HasAdditionalMediaInfo() bool`

HasAdditionalMediaInfo returns a boolean if a field has been set.

### GetAllowDownloadStatus

`func (o *Media) GetAllowDownloadStatus() AllowDownloadStatus`

GetAllowDownloadStatus returns the AllowDownloadStatus field if non-nil, zero value otherwise.

### GetAllowDownloadStatusOk

`func (o *Media) GetAllowDownloadStatusOk() (*AllowDownloadStatus, bool)`

GetAllowDownloadStatusOk returns a tuple with the AllowDownloadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDownloadStatus

`func (o *Media) SetAllowDownloadStatus(v AllowDownloadStatus)`

SetAllowDownloadStatus sets AllowDownloadStatus field to given value.

### HasAllowDownloadStatus

`func (o *Media) HasAllowDownloadStatus() bool`

HasAllowDownloadStatus returns a boolean if a field has been set.

### GetDisplayUrl

`func (o *Media) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *Media) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *Media) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetExpandedUrl

`func (o *Media) GetExpandedUrl() string`

GetExpandedUrl returns the ExpandedUrl field if non-nil, zero value otherwise.

### GetExpandedUrlOk

`func (o *Media) GetExpandedUrlOk() (*string, bool)`

GetExpandedUrlOk returns a tuple with the ExpandedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandedUrl

`func (o *Media) SetExpandedUrl(v string)`

SetExpandedUrl sets ExpandedUrl field to given value.


### GetExtAltText

`func (o *Media) GetExtAltText() string`

GetExtAltText returns the ExtAltText field if non-nil, zero value otherwise.

### GetExtAltTextOk

`func (o *Media) GetExtAltTextOk() (*string, bool)`

GetExtAltTextOk returns a tuple with the ExtAltText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAltText

`func (o *Media) SetExtAltText(v string)`

SetExtAltText sets ExtAltText field to given value.

### HasExtAltText

`func (o *Media) HasExtAltText() bool`

HasExtAltText returns a boolean if a field has been set.

### GetExtMediaAvailability

`func (o *Media) GetExtMediaAvailability() ExtMediaAvailability`

GetExtMediaAvailability returns the ExtMediaAvailability field if non-nil, zero value otherwise.

### GetExtMediaAvailabilityOk

`func (o *Media) GetExtMediaAvailabilityOk() (*ExtMediaAvailability, bool)`

GetExtMediaAvailabilityOk returns a tuple with the ExtMediaAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtMediaAvailability

`func (o *Media) SetExtMediaAvailability(v ExtMediaAvailability)`

SetExtMediaAvailability sets ExtMediaAvailability field to given value.


### GetFeatures

`func (o *Media) GetFeatures() map[string]interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Media) GetFeaturesOk() (*map[string]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Media) SetFeatures(v map[string]interface{})`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Media) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetIdStr

`func (o *Media) GetIdStr() string`

GetIdStr returns the IdStr field if non-nil, zero value otherwise.

### GetIdStrOk

`func (o *Media) GetIdStrOk() (*string, bool)`

GetIdStrOk returns a tuple with the IdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdStr

`func (o *Media) SetIdStr(v string)`

SetIdStr sets IdStr field to given value.


### GetIndices

`func (o *Media) GetIndices() []int32`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *Media) GetIndicesOk() (*[]int32, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *Media) SetIndices(v []int32)`

SetIndices sets Indices field to given value.


### GetMediaKey

`func (o *Media) GetMediaKey() string`

GetMediaKey returns the MediaKey field if non-nil, zero value otherwise.

### GetMediaKeyOk

`func (o *Media) GetMediaKeyOk() (*string, bool)`

GetMediaKeyOk returns a tuple with the MediaKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaKey

`func (o *Media) SetMediaKey(v string)`

SetMediaKey sets MediaKey field to given value.


### GetMediaResults

`func (o *Media) GetMediaResults() MediaResults`

GetMediaResults returns the MediaResults field if non-nil, zero value otherwise.

### GetMediaResultsOk

`func (o *Media) GetMediaResultsOk() (*MediaResults, bool)`

GetMediaResultsOk returns a tuple with the MediaResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaResults

`func (o *Media) SetMediaResults(v MediaResults)`

SetMediaResults sets MediaResults field to given value.

### HasMediaResults

`func (o *Media) HasMediaResults() bool`

HasMediaResults returns a boolean if a field has been set.

### GetMediaUrlHttps

`func (o *Media) GetMediaUrlHttps() string`

GetMediaUrlHttps returns the MediaUrlHttps field if non-nil, zero value otherwise.

### GetMediaUrlHttpsOk

`func (o *Media) GetMediaUrlHttpsOk() (*string, bool)`

GetMediaUrlHttpsOk returns a tuple with the MediaUrlHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrlHttps

`func (o *Media) SetMediaUrlHttps(v string)`

SetMediaUrlHttps sets MediaUrlHttps field to given value.


### GetOriginalInfo

`func (o *Media) GetOriginalInfo() MediaOriginalInfo`

GetOriginalInfo returns the OriginalInfo field if non-nil, zero value otherwise.

### GetOriginalInfoOk

`func (o *Media) GetOriginalInfoOk() (*MediaOriginalInfo, bool)`

GetOriginalInfoOk returns a tuple with the OriginalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalInfo

`func (o *Media) SetOriginalInfo(v MediaOriginalInfo)`

SetOriginalInfo sets OriginalInfo field to given value.


### GetSensitiveMediaWarning

`func (o *Media) GetSensitiveMediaWarning() SensitiveMediaWarning`

GetSensitiveMediaWarning returns the SensitiveMediaWarning field if non-nil, zero value otherwise.

### GetSensitiveMediaWarningOk

`func (o *Media) GetSensitiveMediaWarningOk() (*SensitiveMediaWarning, bool)`

GetSensitiveMediaWarningOk returns a tuple with the SensitiveMediaWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveMediaWarning

`func (o *Media) SetSensitiveMediaWarning(v SensitiveMediaWarning)`

SetSensitiveMediaWarning sets SensitiveMediaWarning field to given value.

### HasSensitiveMediaWarning

`func (o *Media) HasSensitiveMediaWarning() bool`

HasSensitiveMediaWarning returns a boolean if a field has been set.

### GetSizes

`func (o *Media) GetSizes() MediaSizes`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *Media) GetSizesOk() (*MediaSizes, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *Media) SetSizes(v MediaSizes)`

SetSizes sets Sizes field to given value.


### GetSourceStatusIdStr

`func (o *Media) GetSourceStatusIdStr() string`

GetSourceStatusIdStr returns the SourceStatusIdStr field if non-nil, zero value otherwise.

### GetSourceStatusIdStrOk

`func (o *Media) GetSourceStatusIdStrOk() (*string, bool)`

GetSourceStatusIdStrOk returns a tuple with the SourceStatusIdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStatusIdStr

`func (o *Media) SetSourceStatusIdStr(v string)`

SetSourceStatusIdStr sets SourceStatusIdStr field to given value.

### HasSourceStatusIdStr

`func (o *Media) HasSourceStatusIdStr() bool`

HasSourceStatusIdStr returns a boolean if a field has been set.

### GetSourceUserIdStr

`func (o *Media) GetSourceUserIdStr() string`

GetSourceUserIdStr returns the SourceUserIdStr field if non-nil, zero value otherwise.

### GetSourceUserIdStrOk

`func (o *Media) GetSourceUserIdStrOk() (*string, bool)`

GetSourceUserIdStrOk returns a tuple with the SourceUserIdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUserIdStr

`func (o *Media) SetSourceUserIdStr(v string)`

SetSourceUserIdStr sets SourceUserIdStr field to given value.

### HasSourceUserIdStr

`func (o *Media) HasSourceUserIdStr() bool`

HasSourceUserIdStr returns a boolean if a field has been set.

### GetType

`func (o *Media) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Media) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Media) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *Media) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Media) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Media) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVideoInfo

`func (o *Media) GetVideoInfo() MediaVideoInfo`

GetVideoInfo returns the VideoInfo field if non-nil, zero value otherwise.

### GetVideoInfoOk

`func (o *Media) GetVideoInfoOk() (*MediaVideoInfo, bool)`

GetVideoInfoOk returns a tuple with the VideoInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoInfo

`func (o *Media) SetVideoInfo(v MediaVideoInfo)`

SetVideoInfo sets VideoInfo field to given value.

### HasVideoInfo

`func (o *Media) HasVideoInfo() bool`

HasVideoInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


