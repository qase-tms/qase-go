# ReviewUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reviewers** | Pointer to **[]string** | Author UUIDs of team members assigned as reviewers (see &#x60;GET /author&#x60;). When provided, replaces the current reviewer list; an empty array removes all reviewers. Omit to leave reviewers unchanged. | [optional] 
**ProposedCase** | Pointer to [**ReviewCaseData**](ReviewCaseData.md) | Sent fields are merged into the stored proposal. Changing the proposal resets all existing approvals; updating only the reviewers keeps them. | [optional] 

## Methods

### NewReviewUpdate

`func NewReviewUpdate() *ReviewUpdate`

NewReviewUpdate instantiates a new ReviewUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewUpdateWithDefaults

`func NewReviewUpdateWithDefaults() *ReviewUpdate`

NewReviewUpdateWithDefaults instantiates a new ReviewUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviewers

`func (o *ReviewUpdate) GetReviewers() []string`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *ReviewUpdate) GetReviewersOk() (*[]string, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *ReviewUpdate) SetReviewers(v []string)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *ReviewUpdate) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetProposedCase

`func (o *ReviewUpdate) GetProposedCase() ReviewCaseData`

GetProposedCase returns the ProposedCase field if non-nil, zero value otherwise.

### GetProposedCaseOk

`func (o *ReviewUpdate) GetProposedCaseOk() (*ReviewCaseData, bool)`

GetProposedCaseOk returns a tuple with the ProposedCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedCase

`func (o *ReviewUpdate) SetProposedCase(v ReviewCaseData)`

SetProposedCase sets ProposedCase field to given value.

### HasProposedCase

`func (o *ReviewUpdate) HasProposedCase() bool`

HasProposedCase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


