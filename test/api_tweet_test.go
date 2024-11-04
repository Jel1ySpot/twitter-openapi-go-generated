/*
Twitter OpenAPI

Testing TweetAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
    "context"
    openapiclient "github.com/Jel1ySpot/twitter-openapi-go-generated"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
)

func Test_openapi_TweetAPIService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test TweetAPIService GetBookmarks", func(t *testing.T) {

        t.Skip("skip test") // remove to run test

        var pathQueryId string

        resp, httpRes, err := apiClient.TweetAPI.GetBookmarks(context.Background(), pathQueryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TweetAPIService GetHomeLatestTimeline", func(t *testing.T) {

        t.Skip("skip test") // remove to run test

        var pathQueryId string

        resp, httpRes, err := apiClient.TweetAPI.GetHomeLatestTimeline(context.Background(), pathQueryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TweetAPIService GetHomeTimeline", func(t *testing.T) {

        t.Skip("skip test") // remove to run test

        var pathQueryId string

        resp, httpRes, err := apiClient.TweetAPI.GetHomeTimeline(context.Background(), pathQueryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TweetAPIService GetLikes", func(t *testing.T) {

        t.Skip("skip test") // remove to run test

        var pathQueryId string

        resp, httpRes, err := apiClient.TweetAPI.GetLikes(context.Background(), pathQueryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TweetAPIService GetListLatestTweetsTimeline", func(t *testing.T) {

        t.Skip("skip test") // remove to run test

        var pathQueryId string

        resp, httpRes, err := apiClient.TweetAPI.GetListLatestTweetsTimeline(context.Background(), pathQueryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TweetAPIService GetSearchTimeline", func(t *testing.T) {

        t.Skip("skip test") // remove to run test

        var pathQueryId string

        resp, httpRes, err := apiClient.TweetAPI.GetSearchTimeline(context.Background(), pathQueryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TweetAPIService GetTweetDetail", func(t *testing.T) {

        t.Skip("skip test") // remove to run test

        var pathQueryId string

        resp, httpRes, err := apiClient.TweetAPI.GetTweetDetail(context.Background(), pathQueryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TweetAPIService GetUserHighlightsTweets", func(t *testing.T) {

        t.Skip("skip test") // remove to run test

        var pathQueryId string

        resp, httpRes, err := apiClient.TweetAPI.GetUserHighlightsTweets(context.Background(), pathQueryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TweetAPIService GetUserMedia", func(t *testing.T) {

        t.Skip("skip test") // remove to run test

        var pathQueryId string

        resp, httpRes, err := apiClient.TweetAPI.GetUserMedia(context.Background(), pathQueryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TweetAPIService GetUserTweets", func(t *testing.T) {

        t.Skip("skip test") // remove to run test

        var pathQueryId string

        resp, httpRes, err := apiClient.TweetAPI.GetUserTweets(context.Background(), pathQueryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TweetAPIService GetUserTweetsAndReplies", func(t *testing.T) {

        t.Skip("skip test") // remove to run test

        var pathQueryId string

        resp, httpRes, err := apiClient.TweetAPI.GetUserTweetsAndReplies(context.Background(), pathQueryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
