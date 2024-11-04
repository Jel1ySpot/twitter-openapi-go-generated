/*
Twitter OpenAPI

Testing UserAPIService

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

func Test_openapi_UserAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserAPIService GetUserByRestId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathQueryId string

		resp, httpRes, err := apiClient.UserAPI.GetUserByRestId(context.Background(), pathQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService GetUserByScreenName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathQueryId string

		resp, httpRes, err := apiClient.UserAPI.GetUserByScreenName(context.Background(), pathQueryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}