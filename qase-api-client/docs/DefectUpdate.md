# DefectUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**ActualResult** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **int32** |  | [optional] 
**MilestoneId** | Pointer to **NullableInt64** |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**CustomField** | Pointer to **map[string]string** | A map of custom fields values (id &#x3D;&gt; value) | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDefectUpdate

`func NewDefectUpdate() *DefectUpdate`

NewDefectUpdate instantiates a new DefectUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefectUpdateWithDefaults

`func NewDefectUpdateWithDefaults() *DefectUpdate`

NewDefectUpdateWithDefaults instantiates a new DefectUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *DefectUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DefectUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DefectUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DefectUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetActualResult

`func (o *DefectUpdate) GetActualResult() string`

GetActualResult returns the ActualResult field if non-nil, zero value otherwise.

### GetActualResultOk

`func (o *DefectUpdate) GetActualResultOk() (*string, bool)`

GetActualResultOk returns a tuple with the ActualResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualResult

`func (o *DefectUpdate) SetActualResult(v string)`

SetActualResult sets ActualResult field to given value.

### HasActualResult

`func (o *DefectUpdate) HasActualResult() bool`

HasActualResult returns a boolean if a field has been set.

### GetSeverity

`func (o *DefectUpdate) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *DefectUpdate) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *DefectUpdate) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *DefectUpdate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetMilestoneId

`func (o *DefectUpdate) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *DefectUpdate) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *DefectUpdate) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *DefectUpdate) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### SetMilestoneIdNil

`func (o *DefectUpdate) SetMilestoneIdNil(b bool)`

 SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil

### UnsetMilestoneId
`func (o *DefectUpdate) UnsetMilestoneId()`

UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
### GetAttachments

`func (o *DefectUpdate) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *DefectUpdate) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *DefectUpdate) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *DefectUpdate) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCustomField

`func (o *DefectUpdate) GetCustomField() map[string]string`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *DefectUpdate) GetCustomFieldOk() (*map[string]string, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *DefectUpdate) SetCustomField(v map[string]string)`

SetCustomField sets CustomField field to given value.

### HasCustomField

`func (o *DefectUpdate) HasCustomField() bool`

HasCustomField returns a boolean if a field has been set.

### GetTags

`func (o *DefectUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DefectUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DefectUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DefectUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


