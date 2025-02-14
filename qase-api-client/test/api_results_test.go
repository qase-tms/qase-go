/*
Qase.io TestOps API v1

Testing ResultsAPIService

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

func Test_openapi_ResultsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ResultsAPIService CreateResult", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var id int32

		resp, httpRes, err := apiClient.ResultsAPI.CreateResult(context.Background(), code, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ResultsAPIService CreateResultBulk", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var id int32

		resp, httpRes, err := apiClient.ResultsAPI.CreateResultBulk(context.Background(), code, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ResultsAPIService DeleteResult", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var id int32
		var hash string

		resp, httpRes, err := apiClient.ResultsAPI.DeleteResult(context.Background(), code, id, hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ResultsAPIService GetResult", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var hash string

		resp, httpRes, err := apiClient.ResultsAPI.GetResult(context.Background(), code, hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ResultsAPIService GetResults", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string

		resp, httpRes, err := apiClient.ResultsAPI.GetResults(context.Background(), code).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ResultsAPIService UpdateResult", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var id int32
		var hash string

		resp, httpRes, err := apiClient.ResultsAPI.UpdateResult(context.Background(), code, id, hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
