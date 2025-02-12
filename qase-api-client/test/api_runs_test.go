/*
Qase.io TestOps API v1

Testing RunsAPIService

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

func Test_openapi_RunsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RunsAPIService CompleteRun", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var id int32

		resp, httpRes, err := apiClient.RunsAPI.CompleteRun(context.Background(), code, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RunsAPIService CreateRun", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string

		resp, httpRes, err := apiClient.RunsAPI.CreateRun(context.Background(), code).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RunsAPIService DeleteRun", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var id int32

		resp, httpRes, err := apiClient.RunsAPI.DeleteRun(context.Background(), code, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RunsAPIService GetRun", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var id int32

		resp, httpRes, err := apiClient.RunsAPI.GetRun(context.Background(), code, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RunsAPIService GetRuns", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string

		resp, httpRes, err := apiClient.RunsAPI.GetRuns(context.Background(), code).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RunsAPIService UpdateRunPublicity", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var id int32

		resp, httpRes, err := apiClient.RunsAPI.UpdateRunPublicity(context.Background(), code, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
