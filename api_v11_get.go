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


// V11GetAPIService V11GetAPI service
type V11GetAPIService service

type ApiGetFriendsFollowingListRequest struct {
	ctx context.Context
	ApiService *V11GetAPIService
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
	cursor *int32
	userId *string
	count *int32
	withTotalCount *bool
}

func (r ApiGetFriendsFollowingListRequest) IncludeProfileInterstitialType(includeProfileInterstitialType int32) ApiGetFriendsFollowingListRequest {
	r.includeProfileInterstitialType = &includeProfileInterstitialType
	return r
}

func (r ApiGetFriendsFollowingListRequest) IncludeBlocking(includeBlocking int32) ApiGetFriendsFollowingListRequest {
	r.includeBlocking = &includeBlocking
	return r
}

func (r ApiGetFriendsFollowingListRequest) IncludeBlockedBy(includeBlockedBy int32) ApiGetFriendsFollowingListRequest {
	r.includeBlockedBy = &includeBlockedBy
	return r
}

func (r ApiGetFriendsFollowingListRequest) IncludeFollowedBy(includeFollowedBy int32) ApiGetFriendsFollowingListRequest {
	r.includeFollowedBy = &includeFollowedBy
	return r
}

func (r ApiGetFriendsFollowingListRequest) IncludeWantRetweets(includeWantRetweets int32) ApiGetFriendsFollowingListRequest {
	r.includeWantRetweets = &includeWantRetweets
	return r
}

func (r ApiGetFriendsFollowingListRequest) IncludeMuteEdge(includeMuteEdge int32) ApiGetFriendsFollowingListRequest {
	r.includeMuteEdge = &includeMuteEdge
	return r
}

func (r ApiGetFriendsFollowingListRequest) IncludeCanDm(includeCanDm int32) ApiGetFriendsFollowingListRequest {
	r.includeCanDm = &includeCanDm
	return r
}

func (r ApiGetFriendsFollowingListRequest) IncludeCanMediaTag(includeCanMediaTag int32) ApiGetFriendsFollowingListRequest {
	r.includeCanMediaTag = &includeCanMediaTag
	return r
}

func (r ApiGetFriendsFollowingListRequest) IncludeExtHasNftAvatar(includeExtHasNftAvatar int32) ApiGetFriendsFollowingListRequest {
	r.includeExtHasNftAvatar = &includeExtHasNftAvatar
	return r
}

func (r ApiGetFriendsFollowingListRequest) IncludeExtIsBlueVerified(includeExtIsBlueVerified int32) ApiGetFriendsFollowingListRequest {
	r.includeExtIsBlueVerified = &includeExtIsBlueVerified
	return r
}

func (r ApiGetFriendsFollowingListRequest) IncludeExtVerifiedType(includeExtVerifiedType int32) ApiGetFriendsFollowingListRequest {
	r.includeExtVerifiedType = &includeExtVerifiedType
	return r
}

func (r ApiGetFriendsFollowingListRequest) IncludeExtProfileImageShape(includeExtProfileImageShape int32) ApiGetFriendsFollowingListRequest {
	r.includeExtProfileImageShape = &includeExtProfileImageShape
	return r
}

func (r ApiGetFriendsFollowingListRequest) SkipStatus(skipStatus int32) ApiGetFriendsFollowingListRequest {
	r.skipStatus = &skipStatus
	return r
}

func (r ApiGetFriendsFollowingListRequest) Cursor(cursor int32) ApiGetFriendsFollowingListRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetFriendsFollowingListRequest) UserId(userId string) ApiGetFriendsFollowingListRequest {
	r.userId = &userId
	return r
}

func (r ApiGetFriendsFollowingListRequest) Count(count int32) ApiGetFriendsFollowingListRequest {
	r.count = &count
	return r
}

func (r ApiGetFriendsFollowingListRequest) WithTotalCount(withTotalCount bool) ApiGetFriendsFollowingListRequest {
	r.withTotalCount = &withTotalCount
	return r
}

func (r ApiGetFriendsFollowingListRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetFriendsFollowingListExecute(r)
}

/*
GetFriendsFollowingList Method for GetFriendsFollowingList

get friends following list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFriendsFollowingListRequest
*/
func (a *V11GetAPIService) GetFriendsFollowingList(ctx context.Context) ApiGetFriendsFollowingListRequest {
	return ApiGetFriendsFollowingListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *V11GetAPIService) GetFriendsFollowingListExecute(r ApiGetFriendsFollowingListRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V11GetAPIService.GetFriendsFollowingList")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1.1/friends/following/list.json"

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
	if r.cursor == nil {
		return nil, reportError("cursor is required and must be specified")
	}
	if r.userId == nil {
		return nil, reportError("userId is required and must be specified")
	}
	if r.count == nil {
		return nil, reportError("count is required and must be specified")
	}
	if r.withTotalCount == nil {
		return nil, reportError("withTotalCount is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "user_id", r.userId, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "with_total_count", r.withTotalCount, "form", "")
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

type ApiGetSearchTypeaheadRequest struct {
	ctx context.Context
	ApiService *V11GetAPIService
	includeExtIsBlueVerified *int32
	includeExtVerifiedType *int32
	includeExtProfileImageShape *int32
	q *string
	src *string
	resultType *string
}

func (r ApiGetSearchTypeaheadRequest) IncludeExtIsBlueVerified(includeExtIsBlueVerified int32) ApiGetSearchTypeaheadRequest {
	r.includeExtIsBlueVerified = &includeExtIsBlueVerified
	return r
}

func (r ApiGetSearchTypeaheadRequest) IncludeExtVerifiedType(includeExtVerifiedType int32) ApiGetSearchTypeaheadRequest {
	r.includeExtVerifiedType = &includeExtVerifiedType
	return r
}

func (r ApiGetSearchTypeaheadRequest) IncludeExtProfileImageShape(includeExtProfileImageShape int32) ApiGetSearchTypeaheadRequest {
	r.includeExtProfileImageShape = &includeExtProfileImageShape
	return r
}

func (r ApiGetSearchTypeaheadRequest) Q(q string) ApiGetSearchTypeaheadRequest {
	r.q = &q
	return r
}

func (r ApiGetSearchTypeaheadRequest) Src(src string) ApiGetSearchTypeaheadRequest {
	r.src = &src
	return r
}

func (r ApiGetSearchTypeaheadRequest) ResultType(resultType string) ApiGetSearchTypeaheadRequest {
	r.resultType = &resultType
	return r
}

func (r ApiGetSearchTypeaheadRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetSearchTypeaheadExecute(r)
}

/*
GetSearchTypeahead Method for GetSearchTypeahead

get search typeahead

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSearchTypeaheadRequest
*/
func (a *V11GetAPIService) GetSearchTypeahead(ctx context.Context) ApiGetSearchTypeaheadRequest {
	return ApiGetSearchTypeaheadRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *V11GetAPIService) GetSearchTypeaheadExecute(r ApiGetSearchTypeaheadRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V11GetAPIService.GetSearchTypeahead")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1.1/search/typeahead.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.includeExtIsBlueVerified == nil {
		return nil, reportError("includeExtIsBlueVerified is required and must be specified")
	}
	if r.includeExtVerifiedType == nil {
		return nil, reportError("includeExtVerifiedType is required and must be specified")
	}
	if r.includeExtProfileImageShape == nil {
		return nil, reportError("includeExtProfileImageShape is required and must be specified")
	}
	if r.q == nil {
		return nil, reportError("q is required and must be specified")
	}
	if r.src == nil {
		return nil, reportError("src is required and must be specified")
	}
	if r.resultType == nil {
		return nil, reportError("resultType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "include_ext_is_blue_verified", r.includeExtIsBlueVerified, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_ext_verified_type", r.includeExtVerifiedType, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "include_ext_profile_image_shape", r.includeExtProfileImageShape, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "src", r.src, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "result_type", r.resultType, "form", "")
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
