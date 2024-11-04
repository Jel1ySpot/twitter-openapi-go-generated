# TweetLegacy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BookmarkCount** | **int32** |  | 
**Bookmarked** | **bool** |  | 
**ConversationControl** | Pointer to **map[string]interface{}** |  | [optional] 
**ConversationIdStr** | **string** |  | 
**CreatedAt** | **string** |  | 
**DisplayTextRange** | **[]int32** |  | 
**Entities** | [**Entities**](Entities.md) |  | 
**ExtendedEntities** | Pointer to [**ExtendedEntities**](ExtendedEntities.md) |  | [optional] 
**FavoriteCount** | **int32** |  | 
**Favorited** | **bool** |  | 
**FullText** | **string** |  | 
**IdStr** | **string** |  | 
**InReplyToScreenName** | Pointer to **string** |  | [optional] 
**InReplyToStatusIdStr** | Pointer to **string** |  | [optional] 
**InReplyToUserIdStr** | Pointer to **string** |  | [optional] 
**IsQuoteStatus** | **bool** |  | 
**Lang** | **string** |  | 
**LimitedActions** | Pointer to **string** |  | [optional] 
**Place** | Pointer to **map[string]interface{}** |  | [optional] 
**PossiblySensitive** | Pointer to **bool** |  | [optional] 
**PossiblySensitiveEditable** | Pointer to **bool** |  | [optional] 
**QuoteCount** | **int32** |  | 
**QuotedStatusIdStr** | Pointer to **string** |  | [optional] 
**QuotedStatusPermalink** | Pointer to [**QuotedStatusPermalink**](QuotedStatusPermalink.md) |  | [optional] 
**ReplyCount** | **int32** |  | 
**RetweetCount** | **int32** |  | 
**Retweeted** | **bool** |  | 
**RetweetedStatusResult** | Pointer to [**ItemResult**](ItemResult.md) |  | [optional] 
**Scopes** | Pointer to [**TweetLegacyScopes**](TweetLegacyScopes.md) |  | [optional] 
**SelfThread** | Pointer to [**SelfThread**](SelfThread.md) |  | [optional] 
**UserIdStr** | **string** |  | 

## Methods

### NewTweetLegacy

`func NewTweetLegacy(bookmarkCount int32, bookmarked bool, conversationIdStr string, createdAt string, displayTextRange []int32, entities Entities, favoriteCount int32, favorited bool, fullText string, idStr string, isQuoteStatus bool, lang string, quoteCount int32, replyCount int32, retweetCount int32, retweeted bool, userIdStr string, ) *TweetLegacy`

NewTweetLegacy instantiates a new TweetLegacy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTweetLegacyWithDefaults

`func NewTweetLegacyWithDefaults() *TweetLegacy`

NewTweetLegacyWithDefaults instantiates a new TweetLegacy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookmarkCount

`func (o *TweetLegacy) GetBookmarkCount() int32`

GetBookmarkCount returns the BookmarkCount field if non-nil, zero value otherwise.

### GetBookmarkCountOk

`func (o *TweetLegacy) GetBookmarkCountOk() (*int32, bool)`

GetBookmarkCountOk returns a tuple with the BookmarkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarkCount

`func (o *TweetLegacy) SetBookmarkCount(v int32)`

SetBookmarkCount sets BookmarkCount field to given value.


### GetBookmarked

`func (o *TweetLegacy) GetBookmarked() bool`

GetBookmarked returns the Bookmarked field if non-nil, zero value otherwise.

### GetBookmarkedOk

`func (o *TweetLegacy) GetBookmarkedOk() (*bool, bool)`

GetBookmarkedOk returns a tuple with the Bookmarked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarked

`func (o *TweetLegacy) SetBookmarked(v bool)`

SetBookmarked sets Bookmarked field to given value.


### GetConversationControl

`func (o *TweetLegacy) GetConversationControl() map[string]interface{}`

GetConversationControl returns the ConversationControl field if non-nil, zero value otherwise.

### GetConversationControlOk

`func (o *TweetLegacy) GetConversationControlOk() (*map[string]interface{}, bool)`

GetConversationControlOk returns a tuple with the ConversationControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationControl

`func (o *TweetLegacy) SetConversationControl(v map[string]interface{})`

SetConversationControl sets ConversationControl field to given value.

### HasConversationControl

`func (o *TweetLegacy) HasConversationControl() bool`

HasConversationControl returns a boolean if a field has been set.

### GetConversationIdStr

`func (o *TweetLegacy) GetConversationIdStr() string`

GetConversationIdStr returns the ConversationIdStr field if non-nil, zero value otherwise.

### GetConversationIdStrOk

`func (o *TweetLegacy) GetConversationIdStrOk() (*string, bool)`

GetConversationIdStrOk returns a tuple with the ConversationIdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationIdStr

`func (o *TweetLegacy) SetConversationIdStr(v string)`

SetConversationIdStr sets ConversationIdStr field to given value.


### GetCreatedAt

`func (o *TweetLegacy) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TweetLegacy) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TweetLegacy) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDisplayTextRange

`func (o *TweetLegacy) GetDisplayTextRange() []int32`

GetDisplayTextRange returns the DisplayTextRange field if non-nil, zero value otherwise.

### GetDisplayTextRangeOk

`func (o *TweetLegacy) GetDisplayTextRangeOk() (*[]int32, bool)`

GetDisplayTextRangeOk returns a tuple with the DisplayTextRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTextRange

`func (o *TweetLegacy) SetDisplayTextRange(v []int32)`

SetDisplayTextRange sets DisplayTextRange field to given value.


### GetEntities

`func (o *TweetLegacy) GetEntities() Entities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *TweetLegacy) GetEntitiesOk() (*Entities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *TweetLegacy) SetEntities(v Entities)`

SetEntities sets Entities field to given value.


### GetExtendedEntities

`func (o *TweetLegacy) GetExtendedEntities() ExtendedEntities`

GetExtendedEntities returns the ExtendedEntities field if non-nil, zero value otherwise.

### GetExtendedEntitiesOk

`func (o *TweetLegacy) GetExtendedEntitiesOk() (*ExtendedEntities, bool)`

GetExtendedEntitiesOk returns a tuple with the ExtendedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedEntities

`func (o *TweetLegacy) SetExtendedEntities(v ExtendedEntities)`

SetExtendedEntities sets ExtendedEntities field to given value.

### HasExtendedEntities

`func (o *TweetLegacy) HasExtendedEntities() bool`

HasExtendedEntities returns a boolean if a field has been set.

### GetFavoriteCount

`func (o *TweetLegacy) GetFavoriteCount() int32`

GetFavoriteCount returns the FavoriteCount field if non-nil, zero value otherwise.

### GetFavoriteCountOk

`func (o *TweetLegacy) GetFavoriteCountOk() (*int32, bool)`

GetFavoriteCountOk returns a tuple with the FavoriteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteCount

`func (o *TweetLegacy) SetFavoriteCount(v int32)`

SetFavoriteCount sets FavoriteCount field to given value.


### GetFavorited

`func (o *TweetLegacy) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *TweetLegacy) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *TweetLegacy) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.


### GetFullText

`func (o *TweetLegacy) GetFullText() string`

GetFullText returns the FullText field if non-nil, zero value otherwise.

### GetFullTextOk

`func (o *TweetLegacy) GetFullTextOk() (*string, bool)`

GetFullTextOk returns a tuple with the FullText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullText

`func (o *TweetLegacy) SetFullText(v string)`

SetFullText sets FullText field to given value.


### GetIdStr

`func (o *TweetLegacy) GetIdStr() string`

GetIdStr returns the IdStr field if non-nil, zero value otherwise.

### GetIdStrOk

`func (o *TweetLegacy) GetIdStrOk() (*string, bool)`

GetIdStrOk returns a tuple with the IdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdStr

`func (o *TweetLegacy) SetIdStr(v string)`

SetIdStr sets IdStr field to given value.


### GetInReplyToScreenName

`func (o *TweetLegacy) GetInReplyToScreenName() string`

GetInReplyToScreenName returns the InReplyToScreenName field if non-nil, zero value otherwise.

### GetInReplyToScreenNameOk

`func (o *TweetLegacy) GetInReplyToScreenNameOk() (*string, bool)`

GetInReplyToScreenNameOk returns a tuple with the InReplyToScreenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyToScreenName

`func (o *TweetLegacy) SetInReplyToScreenName(v string)`

SetInReplyToScreenName sets InReplyToScreenName field to given value.

### HasInReplyToScreenName

`func (o *TweetLegacy) HasInReplyToScreenName() bool`

HasInReplyToScreenName returns a boolean if a field has been set.

### GetInReplyToStatusIdStr

`func (o *TweetLegacy) GetInReplyToStatusIdStr() string`

GetInReplyToStatusIdStr returns the InReplyToStatusIdStr field if non-nil, zero value otherwise.

### GetInReplyToStatusIdStrOk

`func (o *TweetLegacy) GetInReplyToStatusIdStrOk() (*string, bool)`

GetInReplyToStatusIdStrOk returns a tuple with the InReplyToStatusIdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyToStatusIdStr

`func (o *TweetLegacy) SetInReplyToStatusIdStr(v string)`

SetInReplyToStatusIdStr sets InReplyToStatusIdStr field to given value.

### HasInReplyToStatusIdStr

`func (o *TweetLegacy) HasInReplyToStatusIdStr() bool`

HasInReplyToStatusIdStr returns a boolean if a field has been set.

### GetInReplyToUserIdStr

`func (o *TweetLegacy) GetInReplyToUserIdStr() string`

GetInReplyToUserIdStr returns the InReplyToUserIdStr field if non-nil, zero value otherwise.

### GetInReplyToUserIdStrOk

`func (o *TweetLegacy) GetInReplyToUserIdStrOk() (*string, bool)`

GetInReplyToUserIdStrOk returns a tuple with the InReplyToUserIdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyToUserIdStr

`func (o *TweetLegacy) SetInReplyToUserIdStr(v string)`

SetInReplyToUserIdStr sets InReplyToUserIdStr field to given value.

### HasInReplyToUserIdStr

`func (o *TweetLegacy) HasInReplyToUserIdStr() bool`

HasInReplyToUserIdStr returns a boolean if a field has been set.

### GetIsQuoteStatus

`func (o *TweetLegacy) GetIsQuoteStatus() bool`

GetIsQuoteStatus returns the IsQuoteStatus field if non-nil, zero value otherwise.

### GetIsQuoteStatusOk

`func (o *TweetLegacy) GetIsQuoteStatusOk() (*bool, bool)`

GetIsQuoteStatusOk returns a tuple with the IsQuoteStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuoteStatus

`func (o *TweetLegacy) SetIsQuoteStatus(v bool)`

SetIsQuoteStatus sets IsQuoteStatus field to given value.


### GetLang

`func (o *TweetLegacy) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *TweetLegacy) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *TweetLegacy) SetLang(v string)`

SetLang sets Lang field to given value.


### GetLimitedActions

`func (o *TweetLegacy) GetLimitedActions() string`

GetLimitedActions returns the LimitedActions field if non-nil, zero value otherwise.

### GetLimitedActionsOk

`func (o *TweetLegacy) GetLimitedActionsOk() (*string, bool)`

GetLimitedActionsOk returns a tuple with the LimitedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedActions

`func (o *TweetLegacy) SetLimitedActions(v string)`

SetLimitedActions sets LimitedActions field to given value.

### HasLimitedActions

`func (o *TweetLegacy) HasLimitedActions() bool`

HasLimitedActions returns a boolean if a field has been set.

### GetPlace

`func (o *TweetLegacy) GetPlace() map[string]interface{}`

GetPlace returns the Place field if non-nil, zero value otherwise.

### GetPlaceOk

`func (o *TweetLegacy) GetPlaceOk() (*map[string]interface{}, bool)`

GetPlaceOk returns a tuple with the Place field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlace

`func (o *TweetLegacy) SetPlace(v map[string]interface{})`

SetPlace sets Place field to given value.

### HasPlace

`func (o *TweetLegacy) HasPlace() bool`

HasPlace returns a boolean if a field has been set.

### GetPossiblySensitive

`func (o *TweetLegacy) GetPossiblySensitive() bool`

GetPossiblySensitive returns the PossiblySensitive field if non-nil, zero value otherwise.

### GetPossiblySensitiveOk

`func (o *TweetLegacy) GetPossiblySensitiveOk() (*bool, bool)`

GetPossiblySensitiveOk returns a tuple with the PossiblySensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossiblySensitive

`func (o *TweetLegacy) SetPossiblySensitive(v bool)`

SetPossiblySensitive sets PossiblySensitive field to given value.

### HasPossiblySensitive

`func (o *TweetLegacy) HasPossiblySensitive() bool`

HasPossiblySensitive returns a boolean if a field has been set.

### GetPossiblySensitiveEditable

`func (o *TweetLegacy) GetPossiblySensitiveEditable() bool`

GetPossiblySensitiveEditable returns the PossiblySensitiveEditable field if non-nil, zero value otherwise.

### GetPossiblySensitiveEditableOk

`func (o *TweetLegacy) GetPossiblySensitiveEditableOk() (*bool, bool)`

GetPossiblySensitiveEditableOk returns a tuple with the PossiblySensitiveEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossiblySensitiveEditable

`func (o *TweetLegacy) SetPossiblySensitiveEditable(v bool)`

SetPossiblySensitiveEditable sets PossiblySensitiveEditable field to given value.

### HasPossiblySensitiveEditable

`func (o *TweetLegacy) HasPossiblySensitiveEditable() bool`

HasPossiblySensitiveEditable returns a boolean if a field has been set.

### GetQuoteCount

`func (o *TweetLegacy) GetQuoteCount() int32`

GetQuoteCount returns the QuoteCount field if non-nil, zero value otherwise.

### GetQuoteCountOk

`func (o *TweetLegacy) GetQuoteCountOk() (*int32, bool)`

GetQuoteCountOk returns a tuple with the QuoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCount

`func (o *TweetLegacy) SetQuoteCount(v int32)`

SetQuoteCount sets QuoteCount field to given value.


### GetQuotedStatusIdStr

`func (o *TweetLegacy) GetQuotedStatusIdStr() string`

GetQuotedStatusIdStr returns the QuotedStatusIdStr field if non-nil, zero value otherwise.

### GetQuotedStatusIdStrOk

`func (o *TweetLegacy) GetQuotedStatusIdStrOk() (*string, bool)`

GetQuotedStatusIdStrOk returns a tuple with the QuotedStatusIdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedStatusIdStr

`func (o *TweetLegacy) SetQuotedStatusIdStr(v string)`

SetQuotedStatusIdStr sets QuotedStatusIdStr field to given value.

### HasQuotedStatusIdStr

`func (o *TweetLegacy) HasQuotedStatusIdStr() bool`

HasQuotedStatusIdStr returns a boolean if a field has been set.

### GetQuotedStatusPermalink

`func (o *TweetLegacy) GetQuotedStatusPermalink() QuotedStatusPermalink`

GetQuotedStatusPermalink returns the QuotedStatusPermalink field if non-nil, zero value otherwise.

### GetQuotedStatusPermalinkOk

`func (o *TweetLegacy) GetQuotedStatusPermalinkOk() (*QuotedStatusPermalink, bool)`

GetQuotedStatusPermalinkOk returns a tuple with the QuotedStatusPermalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedStatusPermalink

`func (o *TweetLegacy) SetQuotedStatusPermalink(v QuotedStatusPermalink)`

SetQuotedStatusPermalink sets QuotedStatusPermalink field to given value.

### HasQuotedStatusPermalink

`func (o *TweetLegacy) HasQuotedStatusPermalink() bool`

HasQuotedStatusPermalink returns a boolean if a field has been set.

### GetReplyCount

`func (o *TweetLegacy) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *TweetLegacy) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *TweetLegacy) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.


### GetRetweetCount

`func (o *TweetLegacy) GetRetweetCount() int32`

GetRetweetCount returns the RetweetCount field if non-nil, zero value otherwise.

### GetRetweetCountOk

`func (o *TweetLegacy) GetRetweetCountOk() (*int32, bool)`

GetRetweetCountOk returns a tuple with the RetweetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetweetCount

`func (o *TweetLegacy) SetRetweetCount(v int32)`

SetRetweetCount sets RetweetCount field to given value.


### GetRetweeted

`func (o *TweetLegacy) GetRetweeted() bool`

GetRetweeted returns the Retweeted field if non-nil, zero value otherwise.

### GetRetweetedOk

`func (o *TweetLegacy) GetRetweetedOk() (*bool, bool)`

GetRetweetedOk returns a tuple with the Retweeted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetweeted

`func (o *TweetLegacy) SetRetweeted(v bool)`

SetRetweeted sets Retweeted field to given value.


### GetRetweetedStatusResult

`func (o *TweetLegacy) GetRetweetedStatusResult() ItemResult`

GetRetweetedStatusResult returns the RetweetedStatusResult field if non-nil, zero value otherwise.

### GetRetweetedStatusResultOk

`func (o *TweetLegacy) GetRetweetedStatusResultOk() (*ItemResult, bool)`

GetRetweetedStatusResultOk returns a tuple with the RetweetedStatusResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetweetedStatusResult

`func (o *TweetLegacy) SetRetweetedStatusResult(v ItemResult)`

SetRetweetedStatusResult sets RetweetedStatusResult field to given value.

### HasRetweetedStatusResult

`func (o *TweetLegacy) HasRetweetedStatusResult() bool`

HasRetweetedStatusResult returns a boolean if a field has been set.

### GetScopes

`func (o *TweetLegacy) GetScopes() TweetLegacyScopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TweetLegacy) GetScopesOk() (*TweetLegacyScopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TweetLegacy) SetScopes(v TweetLegacyScopes)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TweetLegacy) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSelfThread

`func (o *TweetLegacy) GetSelfThread() SelfThread`

GetSelfThread returns the SelfThread field if non-nil, zero value otherwise.

### GetSelfThreadOk

`func (o *TweetLegacy) GetSelfThreadOk() (*SelfThread, bool)`

GetSelfThreadOk returns a tuple with the SelfThread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfThread

`func (o *TweetLegacy) SetSelfThread(v SelfThread)`

SetSelfThread sets SelfThread field to given value.

### HasSelfThread

`func (o *TweetLegacy) HasSelfThread() bool`

HasSelfThread returns a boolean if a field has been set.

### GetUserIdStr

`func (o *TweetLegacy) GetUserIdStr() string`

GetUserIdStr returns the UserIdStr field if non-nil, zero value otherwise.

### GetUserIdStrOk

`func (o *TweetLegacy) GetUserIdStrOk() (*string, bool)`

GetUserIdStrOk returns a tuple with the UserIdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdStr

`func (o *TweetLegacy) SetUserIdStr(v string)`

SetUserIdStr sets UserIdStr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


