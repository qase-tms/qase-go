# DefectCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**ActualResult** | **string** |  | 
**Severity** | **int32** |  | 
**MilestoneId** | Pointer to **NullableInt64** |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**CustomField** | Pointer to **map[string]string** | A map of custom fields values (id &#x3D;&gt; value) | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDefectCreate

`func NewDefectCreate(title string, actualResult string, severity int32, ) *DefectCreate`

NewDefectCreate instantiates a new DefectCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefectCreateWithDefaults

`func NewDefectCreateWithDefaults() *DefectCreate`

NewDefectCreateWithDefaults instantiates a new DefectCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *DefectCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DefectCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DefectCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetActualResult

`func (o *DefectCreate) GetActualResult() string`

GetActualResult returns the ActualResult field if non-nil, zero value otherwise.

### GetActualResultOk

`func (o *DefectCreate) GetActualResultOk() (*string, bool)`

GetActualResultOk returns a tuple with the ActualResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualResult

`func (o *DefectCreate) SetActualResult(v string)`

SetActualResult sets ActualResult field to given value.


### GetSeverity

`func (o *DefectCreate) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *DefectCreate) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *DefectCreate) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.


### GetMilestoneId

`func (o *DefectCreate) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *DefectCreate) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *DefectCreate) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *DefectCreate) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### SetMilestoneIdNil

`func (o *DefectCreate) SetMilestoneIdNil(b bool)`

 SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil

### UnsetMilestoneId
`func (o *DefectCreate) UnsetMilestoneId()`

UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
### GetAttachments

`func (o *DefectCreate) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *DefectCreate) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *DefectCreate) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *DefectCreate) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCustomField

`func (o *DefectCreate) GetCustomField() map[string]string`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *DefectCreate) GetCustomFieldOk() (*map[string]string, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *DefectCreate) SetCustomField(v map[string]string)`

SetCustomField sets CustomField field to given value.

### HasCustomField

`func (o *DefectCreate) HasCustomField() bool`

HasCustomField returns a boolean if a field has been set.

### GetTags

`func (o *DefectCreate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DefectCreate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DefectCreate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DefectCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


