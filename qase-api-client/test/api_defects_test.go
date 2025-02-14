/*
Qase.io TestOps API v1

Testing DefectsAPIService

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

func Test_openapi_DefectsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefectsAPIService CreateDefect", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string

		resp, httpRes, err := apiClient.DefectsAPI.CreateDefect(context.Background(), code).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefectsAPIService DeleteDefect", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var id int32

		resp, httpRes, err := apiClient.DefectsAPI.DeleteDefect(context.Background(), code, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefectsAPIService GetDefect", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var id int32

		resp, httpRes, err := apiClient.DefectsAPI.GetDefect(context.Background(), code, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefectsAPIService GetDefects", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string

		resp, httpRes, err := apiClient.DefectsAPI.GetDefects(context.Background(), code).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefectsAPIService ResolveDefect", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var id int32

		resp, httpRes, err := apiClient.DefectsAPI.ResolveDefect(context.Background(), code, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefectsAPIService UpdateDefect", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var id int32

		resp, httpRes, err := apiClient.DefectsAPI.UpdateDefect(context.Background(), code, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefectsAPIService UpdateDefectStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var code string
		var id int32

		resp, httpRes, err := apiClient.DefectsAPI.UpdateDefectStatus(context.Background(), code, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
