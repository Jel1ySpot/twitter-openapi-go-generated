/*
Twitter OpenAPI

Testing PostAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/Jel1ySpot/twitter-openapi-go"
)

func Test_openapi_PostAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PostAPIService PostCreateBookmark", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathQueryId string

		resp, httpRes, err := apiClient.PostAPI.PostCreateBookmark(context.Background(), pathQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostAPIService PostCreateRetweet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathQueryId string

		resp, httpRes, err := apiClient.PostAPI.PostCreateRetweet(context.Background(), pathQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostAPIService PostCreateTweet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathQueryId string

		resp, httpRes, err := apiClient.PostAPI.PostCreateTweet(context.Background(), pathQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostAPIService PostDeleteBookmark", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathQueryId string

		resp, httpRes, err := apiClient.PostAPI.PostDeleteBookmark(context.Background(), pathQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostAPIService PostDeleteRetweet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathQueryId string

		resp, httpRes, err := apiClient.PostAPI.PostDeleteRetweet(context.Background(), pathQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostAPIService PostDeleteTweet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathQueryId string

		resp, httpRes, err := apiClient.PostAPI.PostDeleteTweet(context.Background(), pathQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostAPIService PostFavoriteTweet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathQueryId string

		resp, httpRes, err := apiClient.PostAPI.PostFavoriteTweet(context.Background(), pathQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostAPIService PostUnfavoriteTweet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathQueryId string

		resp, httpRes, err := apiClient.PostAPI.PostUnfavoriteTweet(context.Background(), pathQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}