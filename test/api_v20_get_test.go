/*
Twitter OpenAPI

Testing V20GetAPIService

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

func Test_openapi_V20GetAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test V20GetAPIService GetSearchAdaptive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.V20GetAPI.GetSearchAdaptive(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}