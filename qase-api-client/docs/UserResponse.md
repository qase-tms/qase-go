# UserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Result** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewUserResponse

`func NewUserResponse() *UserResponse`

NewUserResponse instantiates a new UserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseWithDefaults

`func NewUserResponseWithDefaults() *UserResponse`

NewUserResponseWithDefaults instantiates a new UserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UserResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *UserResponse) GetResult() User`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UserResponse) GetResultOk() (*User, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UserResponse) SetResult(v User)`

SetResult sets Result field to given value.

### HasResult

`func (o *UserResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


