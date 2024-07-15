# Defect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**ActualResult** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**MilestoneId** | Pointer to **NullableInt64** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 
**ResolvedAt** | Pointer to **NullableTime** |  | [optional] 
**MemberId** | Pointer to **int64** | Deprecated, use &#x60;author_id&#x60; instead. | [optional] 
**AuthorId** | Pointer to **int64** |  | [optional] 
**ExternalData** | Pointer to **string** |  | [optional] 
**Runs** | Pointer to **[]int64** |  | [optional] 
**Results** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to [**[]TagValue**](TagValue.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Created** | Pointer to **string** | Deprecated, use the &#x60;created_at&#x60; property instead. | [optional] 
**Updated** | Pointer to **string** | Deprecated, use the &#x60;updated_at&#x60; property instead. | [optional] 

## Methods

### NewDefect

`func NewDefect() *Defect`

NewDefect instantiates a new Defect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefectWithDefaults

`func NewDefectWithDefaults() *Defect`

NewDefectWithDefaults instantiates a new Defect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Defect) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Defect) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Defect) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Defect) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Defect) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Defect) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Defect) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Defect) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetActualResult

`func (o *Defect) GetActualResult() string`

GetActualResult returns the ActualResult field if non-nil, zero value otherwise.

### GetActualResultOk

`func (o *Defect) GetActualResultOk() (*string, bool)`

GetActualResultOk returns a tuple with the ActualResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualResult

`func (o *Defect) SetActualResult(v string)`

SetActualResult sets ActualResult field to given value.

### HasActualResult

`func (o *Defect) HasActualResult() bool`

HasActualResult returns a boolean if a field has been set.

### GetSeverity

`func (o *Defect) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Defect) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Defect) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Defect) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStatus

`func (o *Defect) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Defect) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Defect) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Defect) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMilestoneId

`func (o *Defect) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *Defect) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *Defect) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *Defect) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### SetMilestoneIdNil

`func (o *Defect) SetMilestoneIdNil(b bool)`

 SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil

### UnsetMilestoneId
`func (o *Defect) UnsetMilestoneId()`

UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
### GetCustomFields

`func (o *Defect) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Defect) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Defect) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Defect) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetAttachments

`func (o *Defect) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Defect) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Defect) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *Defect) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetResolvedAt

`func (o *Defect) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *Defect) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *Defect) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *Defect) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### SetResolvedAtNil

`func (o *Defect) SetResolvedAtNil(b bool)`

 SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil

### UnsetResolvedAt
`func (o *Defect) UnsetResolvedAt()`

UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil
### GetMemberId

`func (o *Defect) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *Defect) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *Defect) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *Defect) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetAuthorId

`func (o *Defect) GetAuthorId() int64`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *Defect) GetAuthorIdOk() (*int64, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *Defect) SetAuthorId(v int64)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *Defect) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetExternalData

`func (o *Defect) GetExternalData() string`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *Defect) GetExternalDataOk() (*string, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *Defect) SetExternalData(v string)`

SetExternalData sets ExternalData field to given value.

### HasExternalData

`func (o *Defect) HasExternalData() bool`

HasExternalData returns a boolean if a field has been set.

### GetRuns

`func (o *Defect) GetRuns() []int64`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *Defect) GetRunsOk() (*[]int64, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *Defect) SetRuns(v []int64)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *Defect) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetResults

`func (o *Defect) GetResults() []string`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Defect) GetResultsOk() (*[]string, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Defect) SetResults(v []string)`

SetResults sets Results field to given value.

### HasResults

`func (o *Defect) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTags

`func (o *Defect) GetTags() []TagValue`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Defect) GetTagsOk() (*[]TagValue, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Defect) SetTags(v []TagValue)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Defect) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Defect) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Defect) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Defect) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Defect) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Defect) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Defect) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Defect) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Defect) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreated

`func (o *Defect) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Defect) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Defect) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Defect) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *Defect) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Defect) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Defect) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Defect) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


