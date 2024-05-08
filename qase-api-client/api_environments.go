/*
Qase.io TestOps API v1

Qase TestOps API v1 Specification.

API version: 1.0.0
Contact: support@qase.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v1_client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// EnvironmentsAPIService EnvironmentsAPI service
type EnvironmentsAPIService service

type ApiCreateEnvironmentRequest struct {
	ctx               context.Context
	ApiService        *EnvironmentsAPIService
	code              string
	environmentCreate *EnvironmentCreate
}

func (r ApiCreateEnvironmentRequest) EnvironmentCreate(environmentCreate EnvironmentCreate) ApiCreateEnvironmentRequest {
	r.environmentCreate = &environmentCreate
	return r
}

func (r ApiCreateEnvironmentRequest) Execute() (*IdResponse, *http.Response, error) {
	return r.ApiService.CreateEnvironmentExecute(r)
}

/*
CreateEnvironment Create a new environment

This method allows to create an environment in selected project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param code Code of project, where to search entities.
	@return ApiCreateEnvironmentRequest
*/
func (a *EnvironmentsAPIService) CreateEnvironment(ctx context.Context, code string) ApiCreateEnvironmentRequest {
	return ApiCreateEnvironmentRequest{
		ApiService: a,
		ctx:        ctx,
		code:       code,
	}
}

// Execute executes the request
//
//	@return IdResponse
func (a *EnvironmentsAPIService) CreateEnvironmentExecute(r ApiCreateEnvironmentRequest) (*IdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsAPIService.CreateEnvironment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{code}"
	localVarPath = strings.Replace(localVarPath, "{"+"code"+"}", url.PathEscape(parameterValueToString(r.code, "code")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.code) < 2 {
		return localVarReturnValue, nil, reportError("code must have at least 2 elements")
	}
	if strlen(r.code) > 10 {
		return localVarReturnValue, nil, reportError("code must have less than 10 elements")
	}
	if r.environmentCreate == nil {
		return localVarReturnValue, nil, reportError("environmentCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.environmentCreate
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["TokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteEnvironmentRequest struct {
	ctx        context.Context
	ApiService *EnvironmentsAPIService
	code       string
	id         int32
}

func (r ApiDeleteEnvironmentRequest) Execute() (*IdResponse, *http.Response, error) {
	return r.ApiService.DeleteEnvironmentExecute(r)
}

/*
DeleteEnvironment Delete environment

This method completely deletes an environment from repository.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param code Code of project, where to search entities.
	@param id Identifier.
	@return ApiDeleteEnvironmentRequest
*/
func (a *EnvironmentsAPIService) DeleteEnvironment(ctx context.Context, code string, id int32) ApiDeleteEnvironmentRequest {
	return ApiDeleteEnvironmentRequest{
		ApiService: a,
		ctx:        ctx,
		code:       code,
		id:         id,
	}
}

// Execute executes the request
//
//	@return IdResponse
func (a *EnvironmentsAPIService) DeleteEnvironmentExecute(r ApiDeleteEnvironmentRequest) (*IdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsAPIService.DeleteEnvironment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{code}/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"code"+"}", url.PathEscape(parameterValueToString(r.code, "code")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.code) < 2 {
		return localVarReturnValue, nil, reportError("code must have at least 2 elements")
	}
	if strlen(r.code) > 10 {
		return localVarReturnValue, nil, reportError("code must have less than 10 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["TokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetEnvironmentRequest struct {
	ctx        context.Context
	ApiService *EnvironmentsAPIService
	code       string
	id         int32
}

func (r ApiGetEnvironmentRequest) Execute() (*EnvironmentResponse, *http.Response, error) {
	return r.ApiService.GetEnvironmentExecute(r)
}

/*
GetEnvironment Get a specific environment

This method allows to retrieve a specific environment.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param code Code of project, where to search entities.
	@param id Identifier.
	@return ApiGetEnvironmentRequest
*/
func (a *EnvironmentsAPIService) GetEnvironment(ctx context.Context, code string, id int32) ApiGetEnvironmentRequest {
	return ApiGetEnvironmentRequest{
		ApiService: a,
		ctx:        ctx,
		code:       code,
		id:         id,
	}
}

// Execute executes the request
//
//	@return EnvironmentResponse
func (a *EnvironmentsAPIService) GetEnvironmentExecute(r ApiGetEnvironmentRequest) (*EnvironmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsAPIService.GetEnvironment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{code}/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"code"+"}", url.PathEscape(parameterValueToString(r.code, "code")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.code) < 2 {
		return localVarReturnValue, nil, reportError("code must have at least 2 elements")
	}
	if strlen(r.code) > 10 {
		return localVarReturnValue, nil, reportError("code must have less than 10 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["TokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetEnvironmentsRequest struct {
	ctx        context.Context
	ApiService *EnvironmentsAPIService
	code       string
	limit      *int32
	offset     *int32
}

// A number of entities in result set.
func (r ApiGetEnvironmentsRequest) Limit(limit int32) ApiGetEnvironmentsRequest {
	r.limit = &limit
	return r
}

// How many entities should be skipped.
func (r ApiGetEnvironmentsRequest) Offset(offset int32) ApiGetEnvironmentsRequest {
	r.offset = &offset
	return r
}

func (r ApiGetEnvironmentsRequest) Execute() (*EnvironmentListResponse, *http.Response, error) {
	return r.ApiService.GetEnvironmentsExecute(r)
}

/*
GetEnvironments Get all environments

This method allows to retrieve all environments stored in selected project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param code Code of project, where to search entities.
	@return ApiGetEnvironmentsRequest
*/
func (a *EnvironmentsAPIService) GetEnvironments(ctx context.Context, code string) ApiGetEnvironmentsRequest {
	return ApiGetEnvironmentsRequest{
		ApiService: a,
		ctx:        ctx,
		code:       code,
	}
}

// Execute executes the request
//
//	@return EnvironmentListResponse
func (a *EnvironmentsAPIService) GetEnvironmentsExecute(r ApiGetEnvironmentsRequest) (*EnvironmentListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsAPIService.GetEnvironments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{code}"
	localVarPath = strings.Replace(localVarPath, "{"+"code"+"}", url.PathEscape(parameterValueToString(r.code, "code")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.code) < 2 {
		return localVarReturnValue, nil, reportError("code must have at least 2 elements")
	}
	if strlen(r.code) > 10 {
		return localVarReturnValue, nil, reportError("code must have less than 10 elements")
	}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 10
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	} else {
		var defaultValue int32 = 0
		r.offset = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["TokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateEnvironmentRequest struct {
	ctx               context.Context
	ApiService        *EnvironmentsAPIService
	code              string
	id                int32
	environmentUpdate *EnvironmentUpdate
}

func (r ApiUpdateEnvironmentRequest) EnvironmentUpdate(environmentUpdate EnvironmentUpdate) ApiUpdateEnvironmentRequest {
	r.environmentUpdate = &environmentUpdate
	return r
}

func (r ApiUpdateEnvironmentRequest) Execute() (*IdResponse, *http.Response, error) {
	return r.ApiService.UpdateEnvironmentExecute(r)
}

/*
UpdateEnvironment Update environment

This method updates an environment.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param code Code of project, where to search entities.
	@param id Identifier.
	@return ApiUpdateEnvironmentRequest
*/
func (a *EnvironmentsAPIService) UpdateEnvironment(ctx context.Context, code string, id int32) ApiUpdateEnvironmentRequest {
	return ApiUpdateEnvironmentRequest{
		ApiService: a,
		ctx:        ctx,
		code:       code,
		id:         id,
	}
}

// Execute executes the request
//
//	@return IdResponse
func (a *EnvironmentsAPIService) UpdateEnvironmentExecute(r ApiUpdateEnvironmentRequest) (*IdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsAPIService.UpdateEnvironment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{code}/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"code"+"}", url.PathEscape(parameterValueToString(r.code, "code")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.code) < 2 {
		return localVarReturnValue, nil, reportError("code must have at least 2 elements")
	}
	if strlen(r.code) > 10 {
		return localVarReturnValue, nil, reportError("code must have less than 10 elements")
	}
	if r.environmentUpdate == nil {
		return localVarReturnValue, nil, reportError("environmentUpdate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.environmentUpdate
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["TokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
