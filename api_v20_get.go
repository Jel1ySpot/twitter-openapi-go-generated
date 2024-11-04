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
	"context"
	"io"
	"net/http"
	"net/url"
)


// V20GetAPIService V20GetAPI service
type V20GetAPIService service

type ApiGetSearchAdaptiveRequest struct {
	ctx context.Context
	ApiService *V20GetAPIService
	includeProfileInterstitialType *int32
	includeBlocking *int32
	includeBlockedBy *int32
	includeFollowedBy *int32
	includeWantRetweets *int32
	includeMuteEdge *int32
	includeCanDm *int32
	includeCanMediaTag *int32
	includeExtHasNftAvatar *int32
	includeExtIsBlueVerified *int32
	includeExtVerifiedType *int32
	includeExtProfileImageShape *int32
	skipStatus *int32
	cardsPlatform *string
	includeCards *int32
	includeExtAltText *bool
	includeExtLimitedActionResults *bool
	includeQuoteCount *bool
	includeReplyCount *int32
	tweetMode *string
	includeExtViews *bool
	includeEntities *bool
	includeUserEntities *bool
	includeExtMediaColor *bool
	includeExtMediaAvailability *bool
	includeExtSensitiveMediaWarning *bool
	includeExtTrustedFriendsMetadata *bool
	sendErrorCodes *bool
	simpleQuotedTweet *bool
	q *string
	querySource *string
	count *int32
	requestContext *string
	pc *int32
	spellingCorrections *int32
	includeExtEditControl *bool
	ext *string
}

func (r ApiGetSearchAdaptiveRequest) IncludeProfileInterstitialType(includeProfileInterstitialType int32) ApiGetSearchAdaptiveRequest {
	r.includeProfileInterstitialType = &includeProfileInterstitialType
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeBlocking(includeBlocking int32) ApiGetSearchAdaptiveRequest {
	r.includeBlocking = &includeBlocking
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeBlockedBy(includeBlockedBy int32) ApiGetSearchAdaptiveRequest {
	r.includeBlockedBy = &includeBlockedBy
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeFollowedBy(includeFollowedBy int32) ApiGetSearchAdaptiveRequest {
	r.includeFollowedBy = &includeFollowedBy
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeWantRetweets(includeWantRetweets int32) ApiGetSearchAdaptiveRequest {
	r.includeWantRetweets = &includeWantRetweets
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeMuteEdge(includeMuteEdge int32) ApiGetSearchAdaptiveRequest {
	r.includeMuteEdge = &includeMuteEdge
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeCanDm(includeCanDm int32) ApiGetSearchAdaptiveRequest {
	r.includeCanDm = &includeCanDm
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeCanMediaTag(includeCanMediaTag int32) ApiGetSearchAdaptiveRequest {
	r.includeCanMediaTag = &includeCanMediaTag
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeExtHasNftAvatar(includeExtHasNftAvatar int32) ApiGetSearchAdaptiveRequest {
	r.includeExtHasNftAvatar = &includeExtHasNftAvatar
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeExtIsBlueVerified(includeExtIsBlueVerified int32) ApiGetSearchAdaptiveRequest {
	r.includeExtIsBlueVerified = &includeExtIsBlueVerified
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeExtVerifiedType(includeExtVerifiedType int32) ApiGetSearchAdaptiveRequest {
	r.includeExtVerifiedType = &includeExtVerifiedType
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeExtProfileImageShape(includeExtProfileImageShape int32) ApiGetSearchAdaptiveRequest {
	r.includeExtProfileImageShape = &includeExtProfileImageShape
	return r
}

func (r ApiGetSearchAdaptiveRequest) SkipStatus(skipStatus int32) ApiGetSearchAdaptiveRequest {
	r.skipStatus = &skipStatus
	return r
}

func (r ApiGetSearchAdaptiveRequest) CardsPlatform(cardsPlatform string) ApiGetSearchAdaptiveRequest {
	r.cardsPlatform = &cardsPlatform
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeCards(includeCards int32) ApiGetSearchAdaptiveRequest {
	r.includeCards = &includeCards
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeExtAltText(includeExtAltText bool) ApiGetSearchAdaptiveRequest {
	r.includeExtAltText = &includeExtAltText
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeExtLimitedActionResults(includeExtLimitedActionResults bool) ApiGetSearchAdaptiveRequest {
	r.includeExtLimitedActionResults = &includeExtLimitedActionResults
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeQuoteCount(includeQuoteCount bool) ApiGetSearchAdaptiveRequest {
	r.includeQuoteCount = &includeQuoteCount
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeReplyCount(includeReplyCount int32) ApiGetSearchAdaptiveRequest {
	r.includeReplyCount = &includeReplyCount
	return r
}

func (r ApiGetSearchAdaptiveRequest) TweetMode(tweetMode string) ApiGetSearchAdaptiveRequest {
	r.tweetMode = &tweetMode
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeExtViews(includeExtViews bool) ApiGetSearchAdaptiveRequest {
	r.includeExtViews = &includeExtViews
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeEntities(includeEntities bool) ApiGetSearchAdaptiveRequest {
	r.includeEntities = &includeEntities
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeUserEntities(includeUserEntities bool) ApiGetSearchAdaptiveRequest {
	r.includeUserEntities = &includeUserEntities
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeExtMediaColor(includeExtMediaColor bool) ApiGetSearchAdaptiveRequest {
	r.includeExtMediaColor = &includeExtMediaColor
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeExtMediaAvailability(includeExtMediaAvailability bool) ApiGetSearchAdaptiveRequest {
	r.includeExtMediaAvailability = &includeExtMediaAvailability
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeExtSensitiveMediaWarning(includeExtSensitiveMediaWarning bool) ApiGetSearchAdaptiveRequest {
	r.includeExtSensitiveMediaWarning = &includeExtSensitiveMediaWarning
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeExtTrustedFriendsMetadata(includeExtTrustedFriendsMetadata bool) ApiGetSearchAdaptiveRequest {
	r.includeExtTrustedFriendsMetadata = &includeExtTrustedFriendsMetadata
	return r
}

func (r ApiGetSearchAdaptiveRequest) SendErrorCodes(sendErrorCodes bool) ApiGetSearchAdaptiveRequest {
	r.sendErrorCodes = &sendErrorCodes
	return r
}

func (r ApiGetSearchAdaptiveRequest) SimpleQuotedTweet(simpleQuotedTweet bool) ApiGetSearchAdaptiveRequest {
	r.simpleQuotedTweet = &simpleQuotedTweet
	return r
}

func (r ApiGetSearchAdaptiveRequest) Q(q string) ApiGetSearchAdaptiveRequest {
	r.q = &q
	return r
}

func (r ApiGetSearchAdaptiveRequest) QuerySource(querySource string) ApiGetSearchAdaptiveRequest {
	r.querySource = &querySource
	return r
}

func (r ApiGetSearchAdaptiveRequest) Count(count int32) ApiGetSearchAdaptiveRequest {
	r.count = &count
	return r
}

func (r ApiGetSearchAdaptiveRequest) RequestContext(requestContext string) ApiGetSearchAdaptiveRequest {
	r.requestContext = &requestContext
	return r
}

func (r ApiGetSearchAdaptiveRequest) Pc(pc int32) ApiGetSearchAdaptiveRequest {
	r.pc = &pc
	return r
}

func (r ApiGetSearchAdaptiveRequest) SpellingCorrections(spellingCorrections int32) ApiGetSearchAdaptiveRequest {
	r.spellingCorrections = &spellingCorrections
	return r
}

func (r ApiGetSearchAdaptiveRequest) IncludeExtEditControl(includeExtEditControl bool) ApiGetSearchAdaptiveRequest {
	r.includeExtEditControl = &includeExtEditControl
	return r
}

func (r ApiGetSearchAdaptiveRequest) Ext(ext string) ApiGetSearchAdaptiveRequest {
	r.ext = &ext
	return r
}

func (r ApiGetSearchAdaptiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetSearchAdaptiveExecute(r)
}

/*
GetSearchAdaptive Method for GetSearchAdaptive

get search adaptive

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSearchAdaptiveRequest
*/
func (a *V20GetAPIService) GetSearchAdaptive(ctx context.Context) ApiGetSearchAdaptiveRequest {
	return ApiGetSearchAdaptiveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *V20GetAPIService) GetSearchAdaptiveExecute(r ApiGetSearchAdaptiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V20GetAPIService.GetSearchAdaptive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2/search/adaptive.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.includeProfileInterstitialType == nil {
		return nil, reportError("includeProfileInterstitialType is required and must be specified")
	}
	if r.includeBlocking == nil {
		return nil, reportError("includeBlocking is required and must be specified")
	}
	if r.includeBlockedBy == nil {
		return nil, reportError("includeBlockedBy is required and must be specified")
	}
	if r.includeFollowedBy == nil {
		return nil, reportError("includeFollowedBy is required and must be specified")
	}
	if r.includeWantRetweets == nil {
		return nil, reportError("includeWantRetweets is required and must be specified")
	}
	if r.includeMuteEdge == nil {
		return nil, reportError("includeMuteEdge is required and must be specified")
	}
	if r.includeCanDm == nil {
		return nil, reportError("includeCanDm is required and must be specified")
	}
	if r.includeCanMediaTag == nil {
		return nil, reportError("includeCanMediaTag is required and must be specified")
	}
	if r.includeExtHasNftAvatar == nil {
		return nil, reportError("includeExtHasNftAvatar is required and must be specified")
	}
	if r.includeExtIsBlueVerified == nil {
		return nil, reportError("includeExtIsBlueVerified is required and must be specified")
	}
	if r.includeExtVerifiedType == nil {
		return nil, reportError("includeExtVerifiedType is required and must be specified")
	}
	if r.includeExtProfileImageShape == nil {
		return nil, reportError("includeExtProfileImageShape is required and must be specified")
	}
	if r.skipStatus == nil {
		return nil, reportError("skipStatus is required and must be specified")
	}
	if r.cardsPlatform == nil {
		return nil, reportError("cardsPlatform is required and must be specified")
	}
	if r.includeCards == nil {
		return nil, reportError("includeCards is required and must be specified")
	}
	if r.includeExtAltText == nil {
		return nil, reportError("includeExtAltText is required and must be specified")
	}
	if r.includeExtLimitedActionResults == nil {
		return nil, reportError("includeExtLimitedActionResults is required and must be specified")
	}
	if r.includeQuoteCount == nil {
		return nil, reportError("includeQuoteCount is required and must be specified")
	}
	if r.includeReplyCount == nil {
		return nil, reportError("includeReplyCount is required and must be specified")
	}
	if r.tweetMode == nil {
		return nil, reportError("tweetMode is required and must be specified")
	}
	if r.includeExtViews == nil {
		return nil, reportError("includeExtViews is required and must be specified")
	}
	if r.includeEntities == nil {
		return nil, reportError("includeEntities is required and must be specified")
	}
	if r.includeUserEntities == nil {
		return nil, reportError("includeUserEntities is required and must be specified")
	}
	if r.includeExtMediaColor == nil {
		return nil, reportError("includeExtMediaColor is required and must be specified")
	}
	if r.includeExtMediaAvailability == nil {
		return nil, reportError("includeExtMediaAvailability is required and must be specified")
	}
	if r.includeExtSensitiveMediaWarning == nil {
		return nil, reportError("includeExtSensitiveMediaWarning is required and must be specified")
	}
	if r.includeExtTrustedFriendsMetadata == nil {
		return nil, reportError("includeExtTrustedFriendsMetadata is required and must be specified")
	}
	if r.sendErrorCodes == nil {
		return nil, reportError("sendErrorCodes is required and must be specified")
	}
	if r.simpleQuotedTweet == nil {
		return nil, reportError("simpleQuotedTweet is required and must be specified")
	}
	if r.q == nil {
		return nil, reportError("q is required and must be specified")
	}
	if r.querySource == nil {
		return nil, reportError("querySource is required and must be specified")
	}
	if r.count == nil {
		return nil, reportError("count is required and must be specified")
	}
	if r.requestContext == nil {
		return nil, reportError("requestContext is required and must be specified")
	}
	if r.pc == nil {
		return nil, reportError("pc is required and must be specified")
	}
	if r.spellingCorrections == nil {
		return nil, reportError("spellingCorrections is required and must be specified")
	}
	if r.includeExtEditControl == nil {
		return nil, reportError("includeExtEditControl is required and must be specified")
	}
	if r.ext == nil {
		return nil, reportError("ext is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "include_profile_interstitial_type", r.includeProfileInterstitialType, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_blocking", r.includeBlocking, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_blocked_by", r.includeBlockedBy, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_followed_by", r.includeFollowedBy, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_want_retweets", r.includeWantRetweets, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_mute_edge", r.includeMuteEdge, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_can_dm", r.includeCanDm, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_can_media_tag", r.includeCanMediaTag, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_ext_has_nft_avatar", r.includeExtHasNftAvatar, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_ext_is_blue_verified", r.includeExtIsBlueVerified, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_ext_verified_type", r.includeExtVerifiedType, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_ext_profile_image_shape", r.includeExtProfileImageShape, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "skip_status", r.skipStatus, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "cards_platform", r.cardsPlatform, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_cards", r.includeCards, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_ext_alt_text", r.includeExtAltText, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_ext_limited_action_results", r.includeExtLimitedActionResults, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_quote_count", r.includeQuoteCount, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_reply_count", r.includeReplyCount, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "tweet_mode", r.tweetMode, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_ext_views", r.includeExtViews, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_entities", r.includeEntities, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_user_entities", r.includeUserEntities, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_ext_media_color", r.includeExtMediaColor, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_ext_media_availability", r.includeExtMediaAvailability, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_ext_sensitive_media_warning", r.includeExtSensitiveMediaWarning, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_ext_trusted_friends_metadata", r.includeExtTrustedFriendsMetadata, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "send_error_codes", r.sendErrorCodes, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "simple_quoted_tweet", r.simpleQuotedTweet, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "query_source", r.querySource, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "requestContext", r.requestContext, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "pc", r.pc, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "spelling_corrections", r.spellingCorrections, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_ext_edit_control", r.includeExtEditControl, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "ext", r.ext, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Accept"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Accept"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ClientLanguage"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-twitter-client-language"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Referer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Referer"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["SecFetchDest"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Sec-Fetch-Dest"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["SecChUaPlatform"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Sec-Ch-Ua-Platform"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["SecFetchMode"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Sec-Fetch-Mode"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CsrfToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-csrf-token"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ClientUuid"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-client-uuid"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["GuestToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-guest-token"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["SecChUa"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Sec-Ch-Ua"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ClientTransactionId"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-client-transaction-id"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ActiveUser"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-twitter-active-user"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["UserAgent"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["user-agent"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["AcceptLanguage"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Accept-Language"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["SecFetchSite"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Sec-Fetch-Site"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["AuthType"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-twitter-auth-type"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["SecChUaMobile"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Sec-Ch-Ua-Mobile"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["AcceptEncoding"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Accept-Encoding"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}