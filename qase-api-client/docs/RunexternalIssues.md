# RunexternalIssues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Links** | [**[]RunexternalIssuesLinksInner**](RunexternalIssuesLinksInner.md) | Array of external issue links. Each test run (run_id) can have only one external issue link. | 

## Methods

### NewRunexternalIssues

`func NewRunexternalIssues(type_ string, links []RunexternalIssuesLinksInner, ) *RunexternalIssues`

NewRunexternalIssues instantiates a new RunexternalIssues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunexternalIssuesWithDefaults

`func NewRunexternalIssuesWithDefaults() *RunexternalIssues`

NewRunexternalIssuesWithDefaults instantiates a new RunexternalIssues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RunexternalIssues) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RunexternalIssues) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RunexternalIssues) SetType(v string)`

SetType sets Type field to given value.


### GetLinks

`func (o *RunexternalIssues) GetLinks() []RunexternalIssuesLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RunexternalIssues) GetLinksOk() (*[]RunexternalIssuesLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RunexternalIssues) SetLinks(v []RunexternalIssuesLinksInner)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


