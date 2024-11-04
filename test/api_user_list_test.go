/*
Twitter OpenAPI

Testing UserListAPIService

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

func Test_openapi_UserListAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserListAPIService GetFavoriters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathQueryId string

		resp, httpRes, err := apiClient.UserListAPI.GetFavoriters(context.Background(), pathQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserListAPIService GetFollowers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathQueryId string

		resp, httpRes, err := apiClient.UserListAPI.GetFollowers(context.Background(), pathQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserListAPIService GetFollowersYouKnow", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathQueryId string

		resp, httpRes, err := apiClient.UserListAPI.GetFollowersYouKnow(context.Background(), pathQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserListAPIService GetFollowing", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathQueryId string

		resp, httpRes, err := apiClient.UserListAPI.GetFollowing(context.Background(), pathQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserListAPIService GetRetweeters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathQueryId string

		resp, httpRes, err := apiClient.UserListAPI.GetRetweeters(context.Background(), pathQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}