/*
Qase.io TestOps API v1

Testing SuitesAPIService

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

func Test_openapi_SuitesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SuitesAPIService CreateSuite", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string

		resp, httpRes, err := apiClient.SuitesAPI.CreateSuite(context.Background(), code).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuitesAPIService DeleteSuite", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var id int32

		resp, httpRes, err := apiClient.SuitesAPI.DeleteSuite(context.Background(), code, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuitesAPIService GetSuite", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var id int32

		resp, httpRes, err := apiClient.SuitesAPI.GetSuite(context.Background(), code, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuitesAPIService GetSuites", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string

		resp, httpRes, err := apiClient.SuitesAPI.GetSuites(context.Background(), code).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuitesAPIService UpdateSuite", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var id int32

		resp, httpRes, err := apiClient.SuitesAPI.UpdateSuite(context.Background(), code, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
