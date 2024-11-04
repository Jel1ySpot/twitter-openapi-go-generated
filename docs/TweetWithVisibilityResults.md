# TweetWithVisibilityResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Typename** | [**TypeName**](TypeName.md) |  | 
**LimitedActionResults** | Pointer to **map[string]interface{}** |  | [optional] 
**MediaVisibilityResults** | Pointer to [**MediaVisibilityResults**](MediaVisibilityResults.md) |  | [optional] 
**Tweet** | [**Tweet**](Tweet.md) |  | 
**TweetInterstitial** | Pointer to [**TweetInterstitial**](TweetInterstitial.md) |  | [optional] 

## Methods

### NewTweetWithVisibilityResults

`func NewTweetWithVisibilityResults(typename TypeName, tweet Tweet, ) *TweetWithVisibilityResults`

NewTweetWithVisibilityResults instantiates a new TweetWithVisibilityResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTweetWithVisibilityResultsWithDefaults

`func NewTweetWithVisibilityResultsWithDefaults() *TweetWithVisibilityResults`

NewTweetWithVisibilityResultsWithDefaults instantiates a new TweetWithVisibilityResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypename

`func (o *TweetWithVisibilityResults) GetTypename() TypeName`

GetTypename returns the Typename field if non-nil, zero value otherwise.

### GetTypenameOk

`func (o *TweetWithVisibilityResults) GetTypenameOk() (*TypeName, bool)`

GetTypenameOk returns a tuple with the Typename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypename

`func (o *TweetWithVisibilityResults) SetTypename(v TypeName)`

SetTypename sets Typename field to given value.


### GetLimitedActionResults

`func (o *TweetWithVisibilityResults) GetLimitedActionResults() map[string]interface{}`

GetLimitedActionResults returns the LimitedActionResults field if non-nil, zero value otherwise.

### GetLimitedActionResultsOk

`func (o *TweetWithVisibilityResults) GetLimitedActionResultsOk() (*map[string]interface{}, bool)`

GetLimitedActionResultsOk returns a tuple with the LimitedActionResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedActionResults

`func (o *TweetWithVisibilityResults) SetLimitedActionResults(v map[string]interface{})`

SetLimitedActionResults sets LimitedActionResults field to given value.

### HasLimitedActionResults

`func (o *TweetWithVisibilityResults) HasLimitedActionResults() bool`

HasLimitedActionResults returns a boolean if a field has been set.

### GetMediaVisibilityResults

`func (o *TweetWithVisibilityResults) GetMediaVisibilityResults() MediaVisibilityResults`

GetMediaVisibilityResults returns the MediaVisibilityResults field if non-nil, zero value otherwise.

### GetMediaVisibilityResultsOk

`func (o *TweetWithVisibilityResults) GetMediaVisibilityResultsOk() (*MediaVisibilityResults, bool)`

GetMediaVisibilityResultsOk returns a tuple with the MediaVisibilityResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaVisibilityResults

`func (o *TweetWithVisibilityResults) SetMediaVisibilityResults(v MediaVisibilityResults)`

SetMediaVisibilityResults sets MediaVisibilityResults field to given value.

### HasMediaVisibilityResults

`func (o *TweetWithVisibilityResults) HasMediaVisibilityResults() bool`

HasMediaVisibilityResults returns a boolean if a field has been set.

### GetTweet

`func (o *TweetWithVisibilityResults) GetTweet() Tweet`

GetTweet returns the Tweet field if non-nil, zero value otherwise.

### GetTweetOk

`func (o *TweetWithVisibilityResults) GetTweetOk() (*Tweet, bool)`

GetTweetOk returns a tuple with the Tweet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweet

`func (o *TweetWithVisibilityResults) SetTweet(v Tweet)`

SetTweet sets Tweet field to given value.


### GetTweetInterstitial

`func (o *TweetWithVisibilityResults) GetTweetInterstitial() TweetInterstitial`

GetTweetInterstitial returns the TweetInterstitial field if non-nil, zero value otherwise.

### GetTweetInterstitialOk

`func (o *TweetWithVisibilityResults) GetTweetInterstitialOk() (*TweetInterstitial, bool)`

GetTweetInterstitialOk returns a tuple with the TweetInterstitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweetInterstitial

`func (o *TweetWithVisibilityResults) SetTweetInterstitial(v TweetInterstitial)`

SetTweetInterstitial sets TweetInterstitial field to given value.

### HasTweetInterstitial

`func (o *TweetWithVisibilityResults) HasTweetInterstitial() bool`

HasTweetInterstitial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


