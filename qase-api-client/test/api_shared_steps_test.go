/*
Qase.io TestOps API v1

Testing SharedStepsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/qase-tms/qase-go/qase-api-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_SharedStepsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SharedStepsAPIService CreateSharedStep", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string

		resp, httpRes, err := apiClient.SharedStepsAPI.CreateSharedStep(context.Background(), code).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SharedStepsAPIService DeleteSharedStep", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var hash string

		resp, httpRes, err := apiClient.SharedStepsAPI.DeleteSharedStep(context.Background(), code, hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SharedStepsAPIService GetSharedStep", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var hash string

		resp, httpRes, err := apiClient.SharedStepsAPI.GetSharedStep(context.Background(), code, hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SharedStepsAPIService GetSharedSteps", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string

		resp, httpRes, err := apiClient.SharedStepsAPI.GetSharedSteps(context.Background(), code).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SharedStepsAPIService UpdateSharedStep", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var hash string

		resp, httpRes, err := apiClient.SharedStepsAPI.UpdateSharedStep(context.Background(), code, hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
