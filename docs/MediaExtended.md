# MediaExtended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalMediaInfo** | Pointer to [**AdditionalMediaInfo**](AdditionalMediaInfo.md) |  | [optional] 
**AllowDownloadStatus** | Pointer to [**AllowDownloadStatus**](AllowDownloadStatus.md) |  | [optional] 
**DisplayUrl** | **string** |  | 
**ExpandedUrl** | **string** |  | 
**ExtAltText** | Pointer to **string** |  | [optional] 
**ExtMediaAvailability** | Pointer to [**ExtMediaAvailability**](ExtMediaAvailability.md) |  | [optional] 
**Features** | Pointer to **map[string]interface{}** |  | [optional] 
**IdStr** | **string** |  | 
**Indices** | **[]int32** |  | 
**MediaStats** | Pointer to [**MediaStats**](MediaStats.md) |  | [optional] 
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

### NewMediaExtended

`func NewMediaExtended(displayUrl string, expandedUrl string, idStr string, indices []int32, mediaKey string, mediaUrlHttps string, originalInfo MediaOriginalInfo, sizes MediaSizes, type_ string, url string, ) *MediaExtended`

NewMediaExtended instantiates a new MediaExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaExtendedWithDefaults

`func NewMediaExtendedWithDefaults() *MediaExtended`

NewMediaExtendedWithDefaults instantiates a new MediaExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalMediaInfo

`func (o *MediaExtended) GetAdditionalMediaInfo() AdditionalMediaInfo`

GetAdditionalMediaInfo returns the AdditionalMediaInfo field if non-nil, zero value otherwise.

### GetAdditionalMediaInfoOk

`func (o *MediaExtended) GetAdditionalMediaInfoOk() (*AdditionalMediaInfo, bool)`

GetAdditionalMediaInfoOk returns a tuple with the AdditionalMediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMediaInfo

`func (o *MediaExtended) SetAdditionalMediaInfo(v AdditionalMediaInfo)`

SetAdditionalMediaInfo sets AdditionalMediaInfo field to given value.

### HasAdditionalMediaInfo

`func (o *MediaExtended) HasAdditionalMediaInfo() bool`

HasAdditionalMediaInfo returns a boolean if a field has been set.

### GetAllowDownloadStatus

`func (o *MediaExtended) GetAllowDownloadStatus() AllowDownloadStatus`

GetAllowDownloadStatus returns the AllowDownloadStatus field if non-nil, zero value otherwise.

### GetAllowDownloadStatusOk

`func (o *MediaExtended) GetAllowDownloadStatusOk() (*AllowDownloadStatus, bool)`

GetAllowDownloadStatusOk returns a tuple with the AllowDownloadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDownloadStatus

`func (o *MediaExtended) SetAllowDownloadStatus(v AllowDownloadStatus)`

SetAllowDownloadStatus sets AllowDownloadStatus field to given value.

### HasAllowDownloadStatus

`func (o *MediaExtended) HasAllowDownloadStatus() bool`

HasAllowDownloadStatus returns a boolean if a field has been set.

### GetDisplayUrl

`func (o *MediaExtended) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *MediaExtended) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *MediaExtended) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetExpandedUrl

`func (o *MediaExtended) GetExpandedUrl() string`

GetExpandedUrl returns the ExpandedUrl field if non-nil, zero value otherwise.

### GetExpandedUrlOk

`func (o *MediaExtended) GetExpandedUrlOk() (*string, bool)`

GetExpandedUrlOk returns a tuple with the ExpandedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandedUrl

`func (o *MediaExtended) SetExpandedUrl(v string)`

SetExpandedUrl sets ExpandedUrl field to given value.


### GetExtAltText

`func (o *MediaExtended) GetExtAltText() string`

GetExtAltText returns the ExtAltText field if non-nil, zero value otherwise.

### GetExtAltTextOk

`func (o *MediaExtended) GetExtAltTextOk() (*string, bool)`

GetExtAltTextOk returns a tuple with the ExtAltText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAltText

`func (o *MediaExtended) SetExtAltText(v string)`

SetExtAltText sets ExtAltText field to given value.

### HasExtAltText

`func (o *MediaExtended) HasExtAltText() bool`

HasExtAltText returns a boolean if a field has been set.

### GetExtMediaAvailability

`func (o *MediaExtended) GetExtMediaAvailability() ExtMediaAvailability`

GetExtMediaAvailability returns the ExtMediaAvailability field if non-nil, zero value otherwise.

### GetExtMediaAvailabilityOk

`func (o *MediaExtended) GetExtMediaAvailabilityOk() (*ExtMediaAvailability, bool)`

GetExtMediaAvailabilityOk returns a tuple with the ExtMediaAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtMediaAvailability

`func (o *MediaExtended) SetExtMediaAvailability(v ExtMediaAvailability)`

SetExtMediaAvailability sets ExtMediaAvailability field to given value.

### HasExtMediaAvailability

`func (o *MediaExtended) HasExtMediaAvailability() bool`

HasExtMediaAvailability returns a boolean if a field has been set.

### GetFeatures

`func (o *MediaExtended) GetFeatures() map[string]interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *MediaExtended) GetFeaturesOk() (*map[string]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *MediaExtended) SetFeatures(v map[string]interface{})`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *MediaExtended) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetIdStr

`func (o *MediaExtended) GetIdStr() string`

GetIdStr returns the IdStr field if non-nil, zero value otherwise.

### GetIdStrOk

`func (o *MediaExtended) GetIdStrOk() (*string, bool)`

GetIdStrOk returns a tuple with the IdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdStr

`func (o *MediaExtended) SetIdStr(v string)`

SetIdStr sets IdStr field to given value.


### GetIndices

`func (o *MediaExtended) GetIndices() []int32`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *MediaExtended) GetIndicesOk() (*[]int32, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *MediaExtended) SetIndices(v []int32)`

SetIndices sets Indices field to given value.


### GetMediaStats

`func (o *MediaExtended) GetMediaStats() MediaStats`

GetMediaStats returns the MediaStats field if non-nil, zero value otherwise.

### GetMediaStatsOk

`func (o *MediaExtended) GetMediaStatsOk() (*MediaStats, bool)`

GetMediaStatsOk returns a tuple with the MediaStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaStats

`func (o *MediaExtended) SetMediaStats(v MediaStats)`

SetMediaStats sets MediaStats field to given value.

### HasMediaStats

`func (o *MediaExtended) HasMediaStats() bool`

HasMediaStats returns a boolean if a field has been set.

### GetMediaKey

`func (o *MediaExtended) GetMediaKey() string`

GetMediaKey returns the MediaKey field if non-nil, zero value otherwise.

### GetMediaKeyOk

`func (o *MediaExtended) GetMediaKeyOk() (*string, bool)`

GetMediaKeyOk returns a tuple with the MediaKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaKey

`func (o *MediaExtended) SetMediaKey(v string)`

SetMediaKey sets MediaKey field to given value.


### GetMediaResults

`func (o *MediaExtended) GetMediaResults() MediaResults`

GetMediaResults returns the MediaResults field if non-nil, zero value otherwise.

### GetMediaResultsOk

`func (o *MediaExtended) GetMediaResultsOk() (*MediaResults, bool)`

GetMediaResultsOk returns a tuple with the MediaResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaResults

`func (o *MediaExtended) SetMediaResults(v MediaResults)`

SetMediaResults sets MediaResults field to given value.

### HasMediaResults

`func (o *MediaExtended) HasMediaResults() bool`

HasMediaResults returns a boolean if a field has been set.

### GetMediaUrlHttps

`func (o *MediaExtended) GetMediaUrlHttps() string`

GetMediaUrlHttps returns the MediaUrlHttps field if non-nil, zero value otherwise.

### GetMediaUrlHttpsOk

`func (o *MediaExtended) GetMediaUrlHttpsOk() (*string, bool)`

GetMediaUrlHttpsOk returns a tuple with the MediaUrlHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrlHttps

`func (o *MediaExtended) SetMediaUrlHttps(v string)`

SetMediaUrlHttps sets MediaUrlHttps field to given value.


### GetOriginalInfo

`func (o *MediaExtended) GetOriginalInfo() MediaOriginalInfo`

GetOriginalInfo returns the OriginalInfo field if non-nil, zero value otherwise.

### GetOriginalInfoOk

`func (o *MediaExtended) GetOriginalInfoOk() (*MediaOriginalInfo, bool)`

GetOriginalInfoOk returns a tuple with the OriginalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalInfo

`func (o *MediaExtended) SetOriginalInfo(v MediaOriginalInfo)`

SetOriginalInfo sets OriginalInfo field to given value.


### GetSensitiveMediaWarning

`func (o *MediaExtended) GetSensitiveMediaWarning() SensitiveMediaWarning`

GetSensitiveMediaWarning returns the SensitiveMediaWarning field if non-nil, zero value otherwise.

### GetSensitiveMediaWarningOk

`func (o *MediaExtended) GetSensitiveMediaWarningOk() (*SensitiveMediaWarning, bool)`

GetSensitiveMediaWarningOk returns a tuple with the SensitiveMediaWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveMediaWarning

`func (o *MediaExtended) SetSensitiveMediaWarning(v SensitiveMediaWarning)`

SetSensitiveMediaWarning sets SensitiveMediaWarning field to given value.

### HasSensitiveMediaWarning

`func (o *MediaExtended) HasSensitiveMediaWarning() bool`

HasSensitiveMediaWarning returns a boolean if a field has been set.

### GetSizes

`func (o *MediaExtended) GetSizes() MediaSizes`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *MediaExtended) GetSizesOk() (*MediaSizes, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *MediaExtended) SetSizes(v MediaSizes)`

SetSizes sets Sizes field to given value.


### GetSourceStatusIdStr

`func (o *MediaExtended) GetSourceStatusIdStr() string`

GetSourceStatusIdStr returns the SourceStatusIdStr field if non-nil, zero value otherwise.

### GetSourceStatusIdStrOk

`func (o *MediaExtended) GetSourceStatusIdStrOk() (*string, bool)`

GetSourceStatusIdStrOk returns a tuple with the SourceStatusIdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStatusIdStr

`func (o *MediaExtended) SetSourceStatusIdStr(v string)`

SetSourceStatusIdStr sets SourceStatusIdStr field to given value.

### HasSourceStatusIdStr

`func (o *MediaExtended) HasSourceStatusIdStr() bool`

HasSourceStatusIdStr returns a boolean if a field has been set.

### GetSourceUserIdStr

`func (o *MediaExtended) GetSourceUserIdStr() string`

GetSourceUserIdStr returns the SourceUserIdStr field if non-nil, zero value otherwise.

### GetSourceUserIdStrOk

`func (o *MediaExtended) GetSourceUserIdStrOk() (*string, bool)`

GetSourceUserIdStrOk returns a tuple with the SourceUserIdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUserIdStr

`func (o *MediaExtended) SetSourceUserIdStr(v string)`

SetSourceUserIdStr sets SourceUserIdStr field to given value.

### HasSourceUserIdStr

`func (o *MediaExtended) HasSourceUserIdStr() bool`

HasSourceUserIdStr returns a boolean if a field has been set.

### GetType

`func (o *MediaExtended) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MediaExtended) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MediaExtended) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *MediaExtended) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MediaExtended) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MediaExtended) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVideoInfo

`func (o *MediaExtended) GetVideoInfo() MediaVideoInfo`

GetVideoInfo returns the VideoInfo field if non-nil, zero value otherwise.

### GetVideoInfoOk

`func (o *MediaExtended) GetVideoInfoOk() (*MediaVideoInfo, bool)`

GetVideoInfoOk returns a tuple with the VideoInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoInfo

`func (o *MediaExtended) SetVideoInfo(v MediaVideoInfo)`

SetVideoInfo sets VideoInfo field to given value.

### HasVideoInfo

`func (o *MediaExtended) HasVideoInfo() bool`

HasVideoInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


