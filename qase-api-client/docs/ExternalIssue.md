# ExternalIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Issues** | Pointer to [**[]ExternalIssueIssuesInner**](ExternalIssueIssuesInner.md) |  | [optional] 

## Methods

### NewExternalIssue

`func NewExternalIssue() *ExternalIssue`

NewExternalIssue instantiates a new ExternalIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalIssueWithDefaults

`func NewExternalIssueWithDefaults() *ExternalIssue`

NewExternalIssueWithDefaults instantiates a new ExternalIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalIssue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalIssue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalIssue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExternalIssue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIssues

`func (o *ExternalIssue) GetIssues() []ExternalIssueIssuesInner`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *ExternalIssue) GetIssuesOk() (*[]ExternalIssueIssuesInner, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *ExternalIssue) SetIssues(v []ExternalIssueIssuesInner)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *ExternalIssue) HasIssues() bool`

HasIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


