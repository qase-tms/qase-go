# ProjectCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cases** | Pointer to **int32** |  | [optional] 
**Suites** | Pointer to **int32** |  | [optional] 
**Milestones** | Pointer to **int32** |  | [optional] 
**Runs** | Pointer to [**ProjectCountsRuns**](ProjectCountsRuns.md) |  | [optional] 
**Defects** | Pointer to [**ProjectCountsDefects**](ProjectCountsDefects.md) |  | [optional] 

## Methods

### NewProjectCounts

`func NewProjectCounts() *ProjectCounts`

NewProjectCounts instantiates a new ProjectCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCountsWithDefaults

`func NewProjectCountsWithDefaults() *ProjectCounts`

NewProjectCountsWithDefaults instantiates a new ProjectCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCases

`func (o *ProjectCounts) GetCases() int32`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *ProjectCounts) GetCasesOk() (*int32, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *ProjectCounts) SetCases(v int32)`

SetCases sets Cases field to given value.

### HasCases

`func (o *ProjectCounts) HasCases() bool`

HasCases returns a boolean if a field has been set.

### GetSuites

`func (o *ProjectCounts) GetSuites() int32`

GetSuites returns the Suites field if non-nil, zero value otherwise.

### GetSuitesOk

`func (o *ProjectCounts) GetSuitesOk() (*int32, bool)`

GetSuitesOk returns a tuple with the Suites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuites

`func (o *ProjectCounts) SetSuites(v int32)`

SetSuites sets Suites field to given value.

### HasSuites

`func (o *ProjectCounts) HasSuites() bool`

HasSuites returns a boolean if a field has been set.

### GetMilestones

`func (o *ProjectCounts) GetMilestones() int32`

GetMilestones returns the Milestones field if non-nil, zero value otherwise.

### GetMilestonesOk

`func (o *ProjectCounts) GetMilestonesOk() (*int32, bool)`

GetMilestonesOk returns a tuple with the Milestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestones

`func (o *ProjectCounts) SetMilestones(v int32)`

SetMilestones sets Milestones field to given value.

### HasMilestones

`func (o *ProjectCounts) HasMilestones() bool`

HasMilestones returns a boolean if a field has been set.

### GetRuns

`func (o *ProjectCounts) GetRuns() ProjectCountsRuns`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *ProjectCounts) GetRunsOk() (*ProjectCountsRuns, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *ProjectCounts) SetRuns(v ProjectCountsRuns)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *ProjectCounts) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetDefects

`func (o *ProjectCounts) GetDefects() ProjectCountsDefects`

GetDefects returns the Defects field if non-nil, zero value otherwise.

### GetDefectsOk

`func (o *ProjectCounts) GetDefectsOk() (*ProjectCountsDefects, bool)`

GetDefectsOk returns a tuple with the Defects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefects

`func (o *ProjectCounts) SetDefects(v ProjectCountsDefects)`

SetDefects sets Defects field to given value.

### HasDefects

`func (o *ProjectCounts) HasDefects() bool`

HasDefects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


