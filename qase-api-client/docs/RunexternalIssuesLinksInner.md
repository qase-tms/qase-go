# RunexternalIssuesLinksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | **int64** |  | 
**ExternalIssue** | Pointer to **NullableString** | An external issue identifier, e.g. \&quot;PROJ-1234\&quot;. Or null if you want to remove the link. | [optional] 

## Methods

### NewRunexternalIssuesLinksInner

`func NewRunexternalIssuesLinksInner(runId int64, ) *RunexternalIssuesLinksInner`

NewRunexternalIssuesLinksInner instantiates a new RunexternalIssuesLinksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunexternalIssuesLinksInnerWithDefaults

`func NewRunexternalIssuesLinksInnerWithDefaults() *RunexternalIssuesLinksInner`

NewRunexternalIssuesLinksInnerWithDefaults instantiates a new RunexternalIssuesLinksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *RunexternalIssuesLinksInner) GetRunId() int64`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *RunexternalIssuesLinksInner) GetRunIdOk() (*int64, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *RunexternalIssuesLinksInner) SetRunId(v int64)`

SetRunId sets RunId field to given value.


### GetExternalIssue

`func (o *RunexternalIssuesLinksInner) GetExternalIssue() string`

GetExternalIssue returns the ExternalIssue field if non-nil, zero value otherwise.

### GetExternalIssueOk

`func (o *RunexternalIssuesLinksInner) GetExternalIssueOk() (*string, bool)`

GetExternalIssueOk returns a tuple with the ExternalIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIssue

`func (o *RunexternalIssuesLinksInner) SetExternalIssue(v string)`

SetExternalIssue sets ExternalIssue field to given value.

### HasExternalIssue

`func (o *RunexternalIssuesLinksInner) HasExternalIssue() bool`

HasExternalIssue returns a boolean if a field has been set.

### SetExternalIssueNil

`func (o *RunexternalIssuesLinksInner) SetExternalIssueNil(b bool)`

 SetExternalIssueNil sets the value for ExternalIssue to be an explicit nil

### UnsetExternalIssue
`func (o *RunexternalIssuesLinksInner) UnsetExternalIssue()`

UnsetExternalIssue ensures that no value is present for ExternalIssue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


