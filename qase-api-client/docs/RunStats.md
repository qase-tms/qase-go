# RunStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Statuses** | Pointer to **map[string]int32** |  | [optional] 
**Untested** | Pointer to **int32** |  | [optional] 
**Passed** | Pointer to **int32** |  | [optional] 
**Failed** | Pointer to **int32** |  | [optional] 
**Blocked** | Pointer to **int32** |  | [optional] 
**Skipped** | Pointer to **int32** |  | [optional] 
**Retest** | Pointer to **int32** |  | [optional] 
**InProgress** | Pointer to **int32** |  | [optional] 
**Invalid** | Pointer to **int32** |  | [optional] 

## Methods

### NewRunStats

`func NewRunStats() *RunStats`

NewRunStats instantiates a new RunStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunStatsWithDefaults

`func NewRunStatsWithDefaults() *RunStats`

NewRunStatsWithDefaults instantiates a new RunStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *RunStats) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RunStats) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RunStats) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RunStats) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetStatuses

`func (o *RunStats) GetStatuses() map[string]int32`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *RunStats) GetStatusesOk() (*map[string]int32, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *RunStats) SetStatuses(v map[string]int32)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *RunStats) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetUntested

`func (o *RunStats) GetUntested() int32`

GetUntested returns the Untested field if non-nil, zero value otherwise.

### GetUntestedOk

`func (o *RunStats) GetUntestedOk() (*int32, bool)`

GetUntestedOk returns a tuple with the Untested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntested

`func (o *RunStats) SetUntested(v int32)`

SetUntested sets Untested field to given value.

### HasUntested

`func (o *RunStats) HasUntested() bool`

HasUntested returns a boolean if a field has been set.

### GetPassed

`func (o *RunStats) GetPassed() int32`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *RunStats) GetPassedOk() (*int32, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *RunStats) SetPassed(v int32)`

SetPassed sets Passed field to given value.

### HasPassed

`func (o *RunStats) HasPassed() bool`

HasPassed returns a boolean if a field has been set.

### GetFailed

`func (o *RunStats) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *RunStats) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *RunStats) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *RunStats) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetBlocked

`func (o *RunStats) GetBlocked() int32`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *RunStats) GetBlockedOk() (*int32, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *RunStats) SetBlocked(v int32)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *RunStats) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetSkipped

`func (o *RunStats) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *RunStats) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *RunStats) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *RunStats) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetRetest

`func (o *RunStats) GetRetest() int32`

GetRetest returns the Retest field if non-nil, zero value otherwise.

### GetRetestOk

`func (o *RunStats) GetRetestOk() (*int32, bool)`

GetRetestOk returns a tuple with the Retest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetest

`func (o *RunStats) SetRetest(v int32)`

SetRetest sets Retest field to given value.

### HasRetest

`func (o *RunStats) HasRetest() bool`

HasRetest returns a boolean if a field has been set.

### GetInProgress

`func (o *RunStats) GetInProgress() int32`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *RunStats) GetInProgressOk() (*int32, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *RunStats) SetInProgress(v int32)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *RunStats) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetInvalid

`func (o *RunStats) GetInvalid() int32`

GetInvalid returns the Invalid field if non-nil, zero value otherwise.

### GetInvalidOk

`func (o *RunStats) GetInvalidOk() (*int32, bool)`

GetInvalidOk returns a tuple with the Invalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalid

`func (o *RunStats) SetInvalid(v int32)`

SetInvalid sets Invalid field to given value.

### HasInvalid

`func (o *RunStats) HasInvalid() bool`

HasInvalid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


