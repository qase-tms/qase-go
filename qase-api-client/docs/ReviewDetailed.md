# ReviewDetailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Review ID, unique within the project. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | &#x60;create&#x60; — the review proposes a new test case; &#x60;edit&#x60; — the review proposes changes to an existing test case. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CaseId** | Pointer to **NullableInt64** | ID of the reviewed test case. Null for new-case draft reviews. | [optional] 
**AuthorUuid** | Pointer to **NullableString** | Author UUID of the review creator (see &#x60;GET /author&#x60;). | [optional] 
**Reviewers** | Pointer to [**[]ReviewReviewersInner**](ReviewReviewersInner.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**ProposedCase** | Pointer to **map[string]interface{}** | The proposed test case state. Merging the review applies it to the test case. | [optional] 

## Methods

### NewReviewDetailed

`func NewReviewDetailed() *ReviewDetailed`

NewReviewDetailed instantiates a new ReviewDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewDetailedWithDefaults

`func NewReviewDetailedWithDefaults() *ReviewDetailed`

NewReviewDetailedWithDefaults instantiates a new ReviewDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewDetailed) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewDetailed) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewDetailed) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ReviewDetailed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ReviewDetailed) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReviewDetailed) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReviewDetailed) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReviewDetailed) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *ReviewDetailed) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReviewDetailed) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReviewDetailed) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReviewDetailed) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *ReviewDetailed) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReviewDetailed) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReviewDetailed) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReviewDetailed) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCaseId

`func (o *ReviewDetailed) GetCaseId() int64`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *ReviewDetailed) GetCaseIdOk() (*int64, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *ReviewDetailed) SetCaseId(v int64)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *ReviewDetailed) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### SetCaseIdNil

`func (o *ReviewDetailed) SetCaseIdNil(b bool)`

 SetCaseIdNil sets the value for CaseId to be an explicit nil

### UnsetCaseId
`func (o *ReviewDetailed) UnsetCaseId()`

UnsetCaseId ensures that no value is present for CaseId, not even an explicit nil
### GetAuthorUuid

`func (o *ReviewDetailed) GetAuthorUuid() string`

GetAuthorUuid returns the AuthorUuid field if non-nil, zero value otherwise.

### GetAuthorUuidOk

`func (o *ReviewDetailed) GetAuthorUuidOk() (*string, bool)`

GetAuthorUuidOk returns a tuple with the AuthorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorUuid

`func (o *ReviewDetailed) SetAuthorUuid(v string)`

SetAuthorUuid sets AuthorUuid field to given value.

### HasAuthorUuid

`func (o *ReviewDetailed) HasAuthorUuid() bool`

HasAuthorUuid returns a boolean if a field has been set.

### SetAuthorUuidNil

`func (o *ReviewDetailed) SetAuthorUuidNil(b bool)`

 SetAuthorUuidNil sets the value for AuthorUuid to be an explicit nil

### UnsetAuthorUuid
`func (o *ReviewDetailed) UnsetAuthorUuid()`

UnsetAuthorUuid ensures that no value is present for AuthorUuid, not even an explicit nil
### GetReviewers

`func (o *ReviewDetailed) GetReviewers() []ReviewReviewersInner`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *ReviewDetailed) GetReviewersOk() (*[]ReviewReviewersInner, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *ReviewDetailed) SetReviewers(v []ReviewReviewersInner)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *ReviewDetailed) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ReviewDetailed) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReviewDetailed) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReviewDetailed) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReviewDetailed) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ReviewDetailed) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReviewDetailed) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReviewDetailed) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ReviewDetailed) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ReviewDetailed) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ReviewDetailed) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetProposedCase

`func (o *ReviewDetailed) GetProposedCase() map[string]interface{}`

GetProposedCase returns the ProposedCase field if non-nil, zero value otherwise.

### GetProposedCaseOk

`func (o *ReviewDetailed) GetProposedCaseOk() (*map[string]interface{}, bool)`

GetProposedCaseOk returns a tuple with the ProposedCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedCase

`func (o *ReviewDetailed) SetProposedCase(v map[string]interface{})`

SetProposedCase sets ProposedCase field to given value.

### HasProposedCase

`func (o *ReviewDetailed) HasProposedCase() bool`

HasProposedCase returns a boolean if a field has been set.

### SetProposedCaseNil

`func (o *ReviewDetailed) SetProposedCaseNil(b bool)`

 SetProposedCaseNil sets the value for ProposedCase to be an explicit nil

### UnsetProposedCase
`func (o *ReviewDetailed) UnsetProposedCase()`

UnsetProposedCase ensures that no value is present for ProposedCase, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


