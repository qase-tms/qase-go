# ReviewCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseId** | Pointer to **NullableInt64** | ID of the reviewed test case. When present an &#x60;edit&#x60; review is created, otherwise a &#x60;create&#x60; review with a new-case draft. | [optional] 
**Reviewers** | Pointer to **[]string** | Author UUIDs of team members to assign as reviewers (see &#x60;GET /author&#x60;). | [optional] 
**ProposedCase** | [**ReviewCaseData**](ReviewCaseData.md) | For &#x60;create&#x60; reviews &#x60;title&#x60; and all required project fields are required. For &#x60;edit&#x60; reviews send only the fields the proposal changes. | 

## Methods

### NewReviewCreate

`func NewReviewCreate(proposedCase ReviewCaseData, ) *ReviewCreate`

NewReviewCreate instantiates a new ReviewCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewCreateWithDefaults

`func NewReviewCreateWithDefaults() *ReviewCreate`

NewReviewCreateWithDefaults instantiates a new ReviewCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseId

`func (o *ReviewCreate) GetCaseId() int64`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *ReviewCreate) GetCaseIdOk() (*int64, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *ReviewCreate) SetCaseId(v int64)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *ReviewCreate) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### SetCaseIdNil

`func (o *ReviewCreate) SetCaseIdNil(b bool)`

 SetCaseIdNil sets the value for CaseId to be an explicit nil

### UnsetCaseId
`func (o *ReviewCreate) UnsetCaseId()`

UnsetCaseId ensures that no value is present for CaseId, not even an explicit nil
### GetReviewers

`func (o *ReviewCreate) GetReviewers() []string`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *ReviewCreate) GetReviewersOk() (*[]string, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *ReviewCreate) SetReviewers(v []string)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *ReviewCreate) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetProposedCase

`func (o *ReviewCreate) GetProposedCase() ReviewCaseData`

GetProposedCase returns the ProposedCase field if non-nil, zero value otherwise.

### GetProposedCaseOk

`func (o *ReviewCreate) GetProposedCaseOk() (*ReviewCaseData, bool)`

GetProposedCaseOk returns a tuple with the ProposedCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedCase

`func (o *ReviewCreate) SetProposedCase(v ReviewCaseData)`

SetProposedCase sets ProposedCase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


