# ProjectListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Result** | Pointer to [**ProjectListResponseAllOfResult**](ProjectListResponseAllOfResult.md) |  | [optional] 

## Methods

### NewProjectListResponse

`func NewProjectListResponse() *ProjectListResponse`

NewProjectListResponse instantiates a new ProjectListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectListResponseWithDefaults

`func NewProjectListResponseWithDefaults() *ProjectListResponse`

NewProjectListResponseWithDefaults instantiates a new ProjectListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ProjectListResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectListResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectListResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *ProjectListResponse) GetResult() ProjectListResponseAllOfResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ProjectListResponse) GetResultOk() (*ProjectListResponseAllOfResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ProjectListResponse) SetResult(v ProjectListResponseAllOfResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ProjectListResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


