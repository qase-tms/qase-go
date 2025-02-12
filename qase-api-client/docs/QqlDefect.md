# QqlDefect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**DefectId** | **int64** |  | 
**Title** | Pointer to **string** |  | [optional] 
**ActualResult** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**MilestoneId** | Pointer to **NullableInt64** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 
**Resolved** | Pointer to **NullableTime** |  | [optional] 
**MemberId** | Pointer to **int64** | Deprecated, use &#x60;author_id&#x60; instead. | [optional] 
**AuthorId** | Pointer to **int64** |  | [optional] 
**ExternalData** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagValue**](TagValue.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewQqlDefect

`func NewQqlDefect(defectId int64, ) *QqlDefect`

NewQqlDefect instantiates a new QqlDefect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQqlDefectWithDefaults

`func NewQqlDefectWithDefaults() *QqlDefect`

NewQqlDefectWithDefaults instantiates a new QqlDefect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QqlDefect) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QqlDefect) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QqlDefect) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *QqlDefect) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDefectId

`func (o *QqlDefect) GetDefectId() int64`

GetDefectId returns the DefectId field if non-nil, zero value otherwise.

### GetDefectIdOk

`func (o *QqlDefect) GetDefectIdOk() (*int64, bool)`

GetDefectIdOk returns a tuple with the DefectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectId

`func (o *QqlDefect) SetDefectId(v int64)`

SetDefectId sets DefectId field to given value.


### GetTitle

`func (o *QqlDefect) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *QqlDefect) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *QqlDefect) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *QqlDefect) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetActualResult

`func (o *QqlDefect) GetActualResult() string`

GetActualResult returns the ActualResult field if non-nil, zero value otherwise.

### GetActualResultOk

`func (o *QqlDefect) GetActualResultOk() (*string, bool)`

GetActualResultOk returns a tuple with the ActualResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualResult

`func (o *QqlDefect) SetActualResult(v string)`

SetActualResult sets ActualResult field to given value.

### HasActualResult

`func (o *QqlDefect) HasActualResult() bool`

HasActualResult returns a boolean if a field has been set.

### GetSeverity

`func (o *QqlDefect) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *QqlDefect) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *QqlDefect) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *QqlDefect) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStatus

`func (o *QqlDefect) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QqlDefect) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QqlDefect) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QqlDefect) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMilestoneId

`func (o *QqlDefect) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *QqlDefect) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *QqlDefect) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *QqlDefect) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### SetMilestoneIdNil

`func (o *QqlDefect) SetMilestoneIdNil(b bool)`

 SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil

### UnsetMilestoneId
`func (o *QqlDefect) UnsetMilestoneId()`

UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
### GetCustomFields

`func (o *QqlDefect) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QqlDefect) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QqlDefect) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *QqlDefect) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetAttachments

`func (o *QqlDefect) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *QqlDefect) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *QqlDefect) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *QqlDefect) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetResolved

`func (o *QqlDefect) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *QqlDefect) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *QqlDefect) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *QqlDefect) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### SetResolvedNil

`func (o *QqlDefect) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *QqlDefect) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetMemberId

`func (o *QqlDefect) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *QqlDefect) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *QqlDefect) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *QqlDefect) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetAuthorId

`func (o *QqlDefect) GetAuthorId() int64`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *QqlDefect) GetAuthorIdOk() (*int64, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *QqlDefect) SetAuthorId(v int64)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *QqlDefect) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetExternalData

`func (o *QqlDefect) GetExternalData() string`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *QqlDefect) GetExternalDataOk() (*string, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *QqlDefect) SetExternalData(v string)`

SetExternalData sets ExternalData field to given value.

### HasExternalData

`func (o *QqlDefect) HasExternalData() bool`

HasExternalData returns a boolean if a field has been set.

### GetTags

`func (o *QqlDefect) GetTags() []TagValue`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *QqlDefect) GetTagsOk() (*[]TagValue, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *QqlDefect) SetTags(v []TagValue)`

SetTags sets Tags field to given value.

### HasTags

`func (o *QqlDefect) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatedAt

`func (o *QqlDefect) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QqlDefect) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QqlDefect) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *QqlDefect) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *QqlDefect) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QqlDefect) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QqlDefect) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *QqlDefect) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


