/*
Twitter OpenAPI

Twitter OpenAPI(Swagger) specification

API version: 0.0.1
Contact: yuki@yuki0311.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TweetLegacy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetLegacy{}

// TweetLegacy struct for TweetLegacy
type TweetLegacy struct {
	BookmarkCount             int32                  `json:"bookmark_count"`
	Bookmarked                bool                   `json:"bookmarked"`
	ConversationControl       map[string]interface{} `json:"conversation_control,omitempty"`
	ConversationIdStr         string                 `json:"conversation_id_str" validate:"regexp=^[0-9]+$"`
	CreatedAt                 string                 `json:"created_at" validate:"regexp=^(Sun|Mon|Tue|Wed|Thu|Fri|Sat) (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) (0[1-9]|[12][0-9]|3[01]) (0[0-9]|1[0-9]|2[0-3])(: ?)([0-5][0-9])(: ?)([0-5][0-9]) ([+-][0-9]{4}) ([0-9]{4})$"`
	DisplayTextRange          []int32                `json:"display_text_range"`
	Entities                  Entities               `json:"entities"`
	ExtendedEntities          *ExtendedEntities      `json:"extended_entities,omitempty"`
	FavoriteCount             int32                  `json:"favorite_count"`
	Favorited                 bool                   `json:"favorited"`
	FullText                  string                 `json:"full_text"`
	IdStr                     string                 `json:"id_str" validate:"regexp=^[0-9]+$"`
	InReplyToScreenName       *string                `json:"in_reply_to_screen_name,omitempty"`
	InReplyToStatusIdStr      *string                `json:"in_reply_to_status_id_str,omitempty" validate:"regexp=^[0-9]+$"`
	InReplyToUserIdStr        *string                `json:"in_reply_to_user_id_str,omitempty" validate:"regexp=^[0-9]+$"`
	IsQuoteStatus             bool                   `json:"is_quote_status"`
	Lang                      string                 `json:"lang"`
	LimitedActions            *string                `json:"limited_actions,omitempty"`
	Place                     map[string]interface{} `json:"place,omitempty"`
	PossiblySensitive         *bool                  `json:"possibly_sensitive,omitempty"`
	PossiblySensitiveEditable *bool                  `json:"possibly_sensitive_editable,omitempty"`
	QuoteCount                int32                  `json:"quote_count"`
	QuotedStatusIdStr         *string                `json:"quoted_status_id_str,omitempty" validate:"regexp=^[0-9]+$"`
	QuotedStatusPermalink     *QuotedStatusPermalink `json:"quoted_status_permalink,omitempty"`
	ReplyCount                int32                  `json:"reply_count"`
	RetweetCount              int32                  `json:"retweet_count"`
	Retweeted                 bool                   `json:"retweeted"`
	RetweetedStatusResult     *ItemResult            `json:"retweeted_status_result,omitempty"`
	Scopes                    *TweetLegacyScopes     `json:"scopes,omitempty"`
	SelfThread                *SelfThread            `json:"self_thread,omitempty"`
	UserIdStr                 string                 `json:"user_id_str" validate:"regexp=^[0-9]+$"`
}

type _TweetLegacy TweetLegacy

// NewTweetLegacy instantiates a new TweetLegacy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetLegacy(bookmarkCount int32, bookmarked bool, conversationIdStr string, createdAt string, displayTextRange []int32, entities Entities, favoriteCount int32, favorited bool, fullText string, idStr string, isQuoteStatus bool, lang string, quoteCount int32, replyCount int32, retweetCount int32, retweeted bool, userIdStr string) *TweetLegacy {
	this := TweetLegacy{}
	this.BookmarkCount = bookmarkCount
	this.Bookmarked = bookmarked
	this.ConversationIdStr = conversationIdStr
	this.CreatedAt = createdAt
	this.DisplayTextRange = displayTextRange
	this.Entities = entities
	this.FavoriteCount = favoriteCount
	this.Favorited = favorited
	this.FullText = fullText
	this.IdStr = idStr
	this.IsQuoteStatus = isQuoteStatus
	this.Lang = lang
	this.QuoteCount = quoteCount
	this.ReplyCount = replyCount
	this.RetweetCount = retweetCount
	this.Retweeted = retweeted
	this.UserIdStr = userIdStr
	return &this
}

// NewTweetLegacyWithDefaults instantiates a new TweetLegacy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetLegacyWithDefaults() *TweetLegacy {
	this := TweetLegacy{}
	return &this
}

// GetBookmarkCount returns the BookmarkCount field value
func (o *TweetLegacy) GetBookmarkCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BookmarkCount
}

// GetBookmarkCountOk returns a tuple with the BookmarkCount field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetBookmarkCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BookmarkCount, true
}

// SetBookmarkCount sets field value
func (o *TweetLegacy) SetBookmarkCount(v int32) {
	o.BookmarkCount = v
}

// GetBookmarked returns the Bookmarked field value
func (o *TweetLegacy) GetBookmarked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Bookmarked
}

// GetBookmarkedOk returns a tuple with the Bookmarked field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetBookmarkedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bookmarked, true
}

// SetBookmarked sets field value
func (o *TweetLegacy) SetBookmarked(v bool) {
	o.Bookmarked = v
}

// GetConversationControl returns the ConversationControl field value if set, zero value otherwise.
func (o *TweetLegacy) GetConversationControl() map[string]interface{} {
	if o == nil || IsNil(o.ConversationControl) {
		var ret map[string]interface{}
		return ret
	}
	return o.ConversationControl
}

// GetConversationControlOk returns a tuple with the ConversationControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetConversationControlOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ConversationControl) {
		return map[string]interface{}{}, false
	}
	return o.ConversationControl, true
}

// HasConversationControl returns a boolean if a field has been set.
func (o *TweetLegacy) HasConversationControl() bool {
	if o != nil && !IsNil(o.ConversationControl) {
		return true
	}

	return false
}

// SetConversationControl gets a reference to the given map[string]interface{} and assigns it to the ConversationControl field.
func (o *TweetLegacy) SetConversationControl(v map[string]interface{}) {
	o.ConversationControl = v
}

// GetConversationIdStr returns the ConversationIdStr field value
func (o *TweetLegacy) GetConversationIdStr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConversationIdStr
}

// GetConversationIdStrOk returns a tuple with the ConversationIdStr field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetConversationIdStrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationIdStr, true
}

// SetConversationIdStr sets field value
func (o *TweetLegacy) SetConversationIdStr(v string) {
	o.ConversationIdStr = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TweetLegacy) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TweetLegacy) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDisplayTextRange returns the DisplayTextRange field value
func (o *TweetLegacy) GetDisplayTextRange() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.DisplayTextRange
}

// GetDisplayTextRangeOk returns a tuple with the DisplayTextRange field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetDisplayTextRangeOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayTextRange, true
}

// SetDisplayTextRange sets field value
func (o *TweetLegacy) SetDisplayTextRange(v []int32) {
	o.DisplayTextRange = v
}

// GetEntities returns the Entities field value
func (o *TweetLegacy) GetEntities() Entities {
	if o == nil {
		var ret Entities
		return ret
	}

	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetEntitiesOk() (*Entities, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entities, true
}

// SetEntities sets field value
func (o *TweetLegacy) SetEntities(v Entities) {
	o.Entities = v
}

// GetExtendedEntities returns the ExtendedEntities field value if set, zero value otherwise.
func (o *TweetLegacy) GetExtendedEntities() ExtendedEntities {
	if o == nil || IsNil(o.ExtendedEntities) {
		var ret ExtendedEntities
		return ret
	}
	return *o.ExtendedEntities
}

// GetExtendedEntitiesOk returns a tuple with the ExtendedEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetExtendedEntitiesOk() (*ExtendedEntities, bool) {
	if o == nil || IsNil(o.ExtendedEntities) {
		return nil, false
	}
	return o.ExtendedEntities, true
}

// HasExtendedEntities returns a boolean if a field has been set.
func (o *TweetLegacy) HasExtendedEntities() bool {
	if o != nil && !IsNil(o.ExtendedEntities) {
		return true
	}

	return false
}

// SetExtendedEntities gets a reference to the given ExtendedEntities and assigns it to the ExtendedEntities field.
func (o *TweetLegacy) SetExtendedEntities(v ExtendedEntities) {
	o.ExtendedEntities = &v
}

// GetFavoriteCount returns the FavoriteCount field value
func (o *TweetLegacy) GetFavoriteCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FavoriteCount
}

// GetFavoriteCountOk returns a tuple with the FavoriteCount field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetFavoriteCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FavoriteCount, true
}

// SetFavoriteCount sets field value
func (o *TweetLegacy) SetFavoriteCount(v int32) {
	o.FavoriteCount = v
}

// GetFavorited returns the Favorited field value
func (o *TweetLegacy) GetFavorited() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Favorited
}

// GetFavoritedOk returns a tuple with the Favorited field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetFavoritedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Favorited, true
}

// SetFavorited sets field value
func (o *TweetLegacy) SetFavorited(v bool) {
	o.Favorited = v
}

// GetFullText returns the FullText field value
func (o *TweetLegacy) GetFullText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullText
}

// GetFullTextOk returns a tuple with the FullText field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetFullTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullText, true
}

// SetFullText sets field value
func (o *TweetLegacy) SetFullText(v string) {
	o.FullText = v
}

// GetIdStr returns the IdStr field value
func (o *TweetLegacy) GetIdStr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdStr
}

// GetIdStrOk returns a tuple with the IdStr field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetIdStrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdStr, true
}

// SetIdStr sets field value
func (o *TweetLegacy) SetIdStr(v string) {
	o.IdStr = v
}

// GetInReplyToScreenName returns the InReplyToScreenName field value if set, zero value otherwise.
func (o *TweetLegacy) GetInReplyToScreenName() string {
	if o == nil || IsNil(o.InReplyToScreenName) {
		var ret string
		return ret
	}
	return *o.InReplyToScreenName
}

// GetInReplyToScreenNameOk returns a tuple with the InReplyToScreenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetInReplyToScreenNameOk() (*string, bool) {
	if o == nil || IsNil(o.InReplyToScreenName) {
		return nil, false
	}
	return o.InReplyToScreenName, true
}

// HasInReplyToScreenName returns a boolean if a field has been set.
func (o *TweetLegacy) HasInReplyToScreenName() bool {
	if o != nil && !IsNil(o.InReplyToScreenName) {
		return true
	}

	return false
}

// SetInReplyToScreenName gets a reference to the given string and assigns it to the InReplyToScreenName field.
func (o *TweetLegacy) SetInReplyToScreenName(v string) {
	o.InReplyToScreenName = &v
}

// GetInReplyToStatusIdStr returns the InReplyToStatusIdStr field value if set, zero value otherwise.
func (o *TweetLegacy) GetInReplyToStatusIdStr() string {
	if o == nil || IsNil(o.InReplyToStatusIdStr) {
		var ret string
		return ret
	}
	return *o.InReplyToStatusIdStr
}

// GetInReplyToStatusIdStrOk returns a tuple with the InReplyToStatusIdStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetInReplyToStatusIdStrOk() (*string, bool) {
	if o == nil || IsNil(o.InReplyToStatusIdStr) {
		return nil, false
	}
	return o.InReplyToStatusIdStr, true
}

// HasInReplyToStatusIdStr returns a boolean if a field has been set.
func (o *TweetLegacy) HasInReplyToStatusIdStr() bool {
	if o != nil && !IsNil(o.InReplyToStatusIdStr) {
		return true
	}

	return false
}

// SetInReplyToStatusIdStr gets a reference to the given string and assigns it to the InReplyToStatusIdStr field.
func (o *TweetLegacy) SetInReplyToStatusIdStr(v string) {
	o.InReplyToStatusIdStr = &v
}

// GetInReplyToUserIdStr returns the InReplyToUserIdStr field value if set, zero value otherwise.
func (o *TweetLegacy) GetInReplyToUserIdStr() string {
	if o == nil || IsNil(o.InReplyToUserIdStr) {
		var ret string
		return ret
	}
	return *o.InReplyToUserIdStr
}

// GetInReplyToUserIdStrOk returns a tuple with the InReplyToUserIdStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetInReplyToUserIdStrOk() (*string, bool) {
	if o == nil || IsNil(o.InReplyToUserIdStr) {
		return nil, false
	}
	return o.InReplyToUserIdStr, true
}

// HasInReplyToUserIdStr returns a boolean if a field has been set.
func (o *TweetLegacy) HasInReplyToUserIdStr() bool {
	if o != nil && !IsNil(o.InReplyToUserIdStr) {
		return true
	}

	return false
}

// SetInReplyToUserIdStr gets a reference to the given string and assigns it to the InReplyToUserIdStr field.
func (o *TweetLegacy) SetInReplyToUserIdStr(v string) {
	o.InReplyToUserIdStr = &v
}

// GetIsQuoteStatus returns the IsQuoteStatus field value
func (o *TweetLegacy) GetIsQuoteStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsQuoteStatus
}

// GetIsQuoteStatusOk returns a tuple with the IsQuoteStatus field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetIsQuoteStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsQuoteStatus, true
}

// SetIsQuoteStatus sets field value
func (o *TweetLegacy) SetIsQuoteStatus(v bool) {
	o.IsQuoteStatus = v
}

// GetLang returns the Lang field value
func (o *TweetLegacy) GetLang() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Lang
}

// GetLangOk returns a tuple with the Lang field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetLangOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lang, true
}

// SetLang sets field value
func (o *TweetLegacy) SetLang(v string) {
	o.Lang = v
}

// GetLimitedActions returns the LimitedActions field value if set, zero value otherwise.
func (o *TweetLegacy) GetLimitedActions() string {
	if o == nil || IsNil(o.LimitedActions) {
		var ret string
		return ret
	}
	return *o.LimitedActions
}

// GetLimitedActionsOk returns a tuple with the LimitedActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetLimitedActionsOk() (*string, bool) {
	if o == nil || IsNil(o.LimitedActions) {
		return nil, false
	}
	return o.LimitedActions, true
}

// HasLimitedActions returns a boolean if a field has been set.
func (o *TweetLegacy) HasLimitedActions() bool {
	if o != nil && !IsNil(o.LimitedActions) {
		return true
	}

	return false
}

// SetLimitedActions gets a reference to the given string and assigns it to the LimitedActions field.
func (o *TweetLegacy) SetLimitedActions(v string) {
	o.LimitedActions = &v
}

// GetPlace returns the Place field value if set, zero value otherwise.
func (o *TweetLegacy) GetPlace() map[string]interface{} {
	if o == nil || IsNil(o.Place) {
		var ret map[string]interface{}
		return ret
	}
	return o.Place
}

// GetPlaceOk returns a tuple with the Place field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetPlaceOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Place) {
		return map[string]interface{}{}, false
	}
	return o.Place, true
}

// HasPlace returns a boolean if a field has been set.
func (o *TweetLegacy) HasPlace() bool {
	if o != nil && !IsNil(o.Place) {
		return true
	}

	return false
}

// SetPlace gets a reference to the given map[string]interface{} and assigns it to the Place field.
func (o *TweetLegacy) SetPlace(v map[string]interface{}) {
	o.Place = v
}

// GetPossiblySensitive returns the PossiblySensitive field value if set, zero value otherwise.
func (o *TweetLegacy) GetPossiblySensitive() bool {
	if o == nil || IsNil(o.PossiblySensitive) {
		var ret bool
		return ret
	}
	return *o.PossiblySensitive
}

// GetPossiblySensitiveOk returns a tuple with the PossiblySensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetPossiblySensitiveOk() (*bool, bool) {
	if o == nil || IsNil(o.PossiblySensitive) {
		return nil, false
	}
	return o.PossiblySensitive, true
}

// HasPossiblySensitive returns a boolean if a field has been set.
func (o *TweetLegacy) HasPossiblySensitive() bool {
	if o != nil && !IsNil(o.PossiblySensitive) {
		return true
	}

	return false
}

// SetPossiblySensitive gets a reference to the given bool and assigns it to the PossiblySensitive field.
func (o *TweetLegacy) SetPossiblySensitive(v bool) {
	o.PossiblySensitive = &v
}

// GetPossiblySensitiveEditable returns the PossiblySensitiveEditable field value if set, zero value otherwise.
func (o *TweetLegacy) GetPossiblySensitiveEditable() bool {
	if o == nil || IsNil(o.PossiblySensitiveEditable) {
		var ret bool
		return ret
	}
	return *o.PossiblySensitiveEditable
}

// GetPossiblySensitiveEditableOk returns a tuple with the PossiblySensitiveEditable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetPossiblySensitiveEditableOk() (*bool, bool) {
	if o == nil || IsNil(o.PossiblySensitiveEditable) {
		return nil, false
	}
	return o.PossiblySensitiveEditable, true
}

// HasPossiblySensitiveEditable returns a boolean if a field has been set.
func (o *TweetLegacy) HasPossiblySensitiveEditable() bool {
	if o != nil && !IsNil(o.PossiblySensitiveEditable) {
		return true
	}

	return false
}

// SetPossiblySensitiveEditable gets a reference to the given bool and assigns it to the PossiblySensitiveEditable field.
func (o *TweetLegacy) SetPossiblySensitiveEditable(v bool) {
	o.PossiblySensitiveEditable = &v
}

// GetQuoteCount returns the QuoteCount field value
func (o *TweetLegacy) GetQuoteCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QuoteCount
}

// GetQuoteCountOk returns a tuple with the QuoteCount field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetQuoteCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteCount, true
}

// SetQuoteCount sets field value
func (o *TweetLegacy) SetQuoteCount(v int32) {
	o.QuoteCount = v
}

// GetQuotedStatusIdStr returns the QuotedStatusIdStr field value if set, zero value otherwise.
func (o *TweetLegacy) GetQuotedStatusIdStr() string {
	if o == nil || IsNil(o.QuotedStatusIdStr) {
		var ret string
		return ret
	}
	return *o.QuotedStatusIdStr
}

// GetQuotedStatusIdStrOk returns a tuple with the QuotedStatusIdStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetQuotedStatusIdStrOk() (*string, bool) {
	if o == nil || IsNil(o.QuotedStatusIdStr) {
		return nil, false
	}
	return o.QuotedStatusIdStr, true
}

// HasQuotedStatusIdStr returns a boolean if a field has been set.
func (o *TweetLegacy) HasQuotedStatusIdStr() bool {
	if o != nil && !IsNil(o.QuotedStatusIdStr) {
		return true
	}

	return false
}

// SetQuotedStatusIdStr gets a reference to the given string and assigns it to the QuotedStatusIdStr field.
func (o *TweetLegacy) SetQuotedStatusIdStr(v string) {
	o.QuotedStatusIdStr = &v
}

// GetQuotedStatusPermalink returns the QuotedStatusPermalink field value if set, zero value otherwise.
func (o *TweetLegacy) GetQuotedStatusPermalink() QuotedStatusPermalink {
	if o == nil || IsNil(o.QuotedStatusPermalink) {
		var ret QuotedStatusPermalink
		return ret
	}
	return *o.QuotedStatusPermalink
}

// GetQuotedStatusPermalinkOk returns a tuple with the QuotedStatusPermalink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetQuotedStatusPermalinkOk() (*QuotedStatusPermalink, bool) {
	if o == nil || IsNil(o.QuotedStatusPermalink) {
		return nil, false
	}
	return o.QuotedStatusPermalink, true
}

// HasQuotedStatusPermalink returns a boolean if a field has been set.
func (o *TweetLegacy) HasQuotedStatusPermalink() bool {
	if o != nil && !IsNil(o.QuotedStatusPermalink) {
		return true
	}

	return false
}

// SetQuotedStatusPermalink gets a reference to the given QuotedStatusPermalink and assigns it to the QuotedStatusPermalink field.
func (o *TweetLegacy) SetQuotedStatusPermalink(v QuotedStatusPermalink) {
	o.QuotedStatusPermalink = &v
}

// GetReplyCount returns the ReplyCount field value
func (o *TweetLegacy) GetReplyCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReplyCount
}

// GetReplyCountOk returns a tuple with the ReplyCount field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetReplyCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplyCount, true
}

// SetReplyCount sets field value
func (o *TweetLegacy) SetReplyCount(v int32) {
	o.ReplyCount = v
}

// GetRetweetCount returns the RetweetCount field value
func (o *TweetLegacy) GetRetweetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RetweetCount
}

// GetRetweetCountOk returns a tuple with the RetweetCount field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetRetweetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetweetCount, true
}

// SetRetweetCount sets field value
func (o *TweetLegacy) SetRetweetCount(v int32) {
	o.RetweetCount = v
}

// GetRetweeted returns the Retweeted field value
func (o *TweetLegacy) GetRetweeted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Retweeted
}

// GetRetweetedOk returns a tuple with the Retweeted field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetRetweetedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retweeted, true
}

// SetRetweeted sets field value
func (o *TweetLegacy) SetRetweeted(v bool) {
	o.Retweeted = v
}

// GetRetweetedStatusResult returns the RetweetedStatusResult field value if set, zero value otherwise.
func (o *TweetLegacy) GetRetweetedStatusResult() ItemResult {
	if o == nil || IsNil(o.RetweetedStatusResult) {
		var ret ItemResult
		return ret
	}
	return *o.RetweetedStatusResult
}

// GetRetweetedStatusResultOk returns a tuple with the RetweetedStatusResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetRetweetedStatusResultOk() (*ItemResult, bool) {
	if o == nil || IsNil(o.RetweetedStatusResult) {
		return nil, false
	}
	return o.RetweetedStatusResult, true
}

// HasRetweetedStatusResult returns a boolean if a field has been set.
func (o *TweetLegacy) HasRetweetedStatusResult() bool {
	if o != nil && !IsNil(o.RetweetedStatusResult) {
		return true
	}

	return false
}

// SetRetweetedStatusResult gets a reference to the given ItemResult and assigns it to the RetweetedStatusResult field.
func (o *TweetLegacy) SetRetweetedStatusResult(v ItemResult) {
	o.RetweetedStatusResult = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *TweetLegacy) GetScopes() TweetLegacyScopes {
	if o == nil || IsNil(o.Scopes) {
		var ret TweetLegacyScopes
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetScopesOk() (*TweetLegacyScopes, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *TweetLegacy) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given TweetLegacyScopes and assigns it to the Scopes field.
func (o *TweetLegacy) SetScopes(v TweetLegacyScopes) {
	o.Scopes = &v
}

// GetSelfThread returns the SelfThread field value if set, zero value otherwise.
func (o *TweetLegacy) GetSelfThread() SelfThread {
	if o == nil || IsNil(o.SelfThread) {
		var ret SelfThread
		return ret
	}
	return *o.SelfThread
}

// GetSelfThreadOk returns a tuple with the SelfThread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetSelfThreadOk() (*SelfThread, bool) {
	if o == nil || IsNil(o.SelfThread) {
		return nil, false
	}
	return o.SelfThread, true
}

// HasSelfThread returns a boolean if a field has been set.
func (o *TweetLegacy) HasSelfThread() bool {
	if o != nil && !IsNil(o.SelfThread) {
		return true
	}

	return false
}

// SetSelfThread gets a reference to the given SelfThread and assigns it to the SelfThread field.
func (o *TweetLegacy) SetSelfThread(v SelfThread) {
	o.SelfThread = &v
}

// GetUserIdStr returns the UserIdStr field value
func (o *TweetLegacy) GetUserIdStr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserIdStr
}

// GetUserIdStrOk returns a tuple with the UserIdStr field value
// and a boolean to check if the value has been set.
func (o *TweetLegacy) GetUserIdStrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserIdStr, true
}

// SetUserIdStr sets field value
func (o *TweetLegacy) SetUserIdStr(v string) {
	o.UserIdStr = v
}

func (o TweetLegacy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetLegacy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bookmark_count"] = o.BookmarkCount
	toSerialize["bookmarked"] = o.Bookmarked
	if !IsNil(o.ConversationControl) {
		toSerialize["conversation_control"] = o.ConversationControl
	}
	toSerialize["conversation_id_str"] = o.ConversationIdStr
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["display_text_range"] = o.DisplayTextRange
	toSerialize["entities"] = o.Entities
	if !IsNil(o.ExtendedEntities) {
		toSerialize["extended_entities"] = o.ExtendedEntities
	}
	toSerialize["favorite_count"] = o.FavoriteCount
	toSerialize["favorited"] = o.Favorited
	toSerialize["full_text"] = o.FullText
	toSerialize["id_str"] = o.IdStr
	if !IsNil(o.InReplyToScreenName) {
		toSerialize["in_reply_to_screen_name"] = o.InReplyToScreenName
	}
	if !IsNil(o.InReplyToStatusIdStr) {
		toSerialize["in_reply_to_status_id_str"] = o.InReplyToStatusIdStr
	}
	if !IsNil(o.InReplyToUserIdStr) {
		toSerialize["in_reply_to_user_id_str"] = o.InReplyToUserIdStr
	}
	toSerialize["is_quote_status"] = o.IsQuoteStatus
	toSerialize["lang"] = o.Lang
	if !IsNil(o.LimitedActions) {
		toSerialize["limited_actions"] = o.LimitedActions
	}
	if !IsNil(o.Place) {
		toSerialize["place"] = o.Place
	}
	if !IsNil(o.PossiblySensitive) {
		toSerialize["possibly_sensitive"] = o.PossiblySensitive
	}
	if !IsNil(o.PossiblySensitiveEditable) {
		toSerialize["possibly_sensitive_editable"] = o.PossiblySensitiveEditable
	}
	toSerialize["quote_count"] = o.QuoteCount
	if !IsNil(o.QuotedStatusIdStr) {
		toSerialize["quoted_status_id_str"] = o.QuotedStatusIdStr
	}
	if !IsNil(o.QuotedStatusPermalink) {
		toSerialize["quoted_status_permalink"] = o.QuotedStatusPermalink
	}
	toSerialize["reply_count"] = o.ReplyCount
	toSerialize["retweet_count"] = o.RetweetCount
	toSerialize["retweeted"] = o.Retweeted
	if !IsNil(o.RetweetedStatusResult) {
		toSerialize["retweeted_status_result"] = o.RetweetedStatusResult
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.SelfThread) {
		toSerialize["self_thread"] = o.SelfThread
	}
	toSerialize["user_id_str"] = o.UserIdStr
	return toSerialize, nil
}

func (o *TweetLegacy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bookmark_count",
		"bookmarked",
		"conversation_id_str",
		"created_at",
		"display_text_range",
		"entities",
		"favorite_count",
		"favorited",
		"full_text",
		"id_str",
		"is_quote_status",
		"lang",
		"quote_count",
		"reply_count",
		"retweet_count",
		"retweeted",
		"user_id_str",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTweetLegacy := _TweetLegacy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTweetLegacy)

	if err != nil {
		return err
	}

	*o = TweetLegacy(varTweetLegacy)

	return err
}

type NullableTweetLegacy struct {
	value *TweetLegacy
	isSet bool
}

func (v NullableTweetLegacy) Get() *TweetLegacy {
	return v.value
}

func (v *NullableTweetLegacy) Set(val *TweetLegacy) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetLegacy) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetLegacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetLegacy(val *TweetLegacy) *NullableTweetLegacy {
	return &NullableTweetLegacy{value: val, isSet: true}
}

func (v NullableTweetLegacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetLegacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
