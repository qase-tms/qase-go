# ReviewReviewersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorUuid** | Pointer to **NullableString** | Author UUID of the reviewer (see &#x60;GET /author&#x60;). | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewReviewReviewersInner

`func NewReviewReviewersInner() *ReviewReviewersInner`

NewReviewReviewersInner instantiates a new ReviewReviewersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewReviewersInnerWithDefaults

`func NewReviewReviewersInnerWithDefaults() *ReviewReviewersInner`

NewReviewReviewersInnerWithDefaults instantiates a new ReviewReviewersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorUuid

`func (o *ReviewReviewersInner) GetAuthorUuid() string`

GetAuthorUuid returns the AuthorUuid field if non-nil, zero value otherwise.

### GetAuthorUuidOk

`func (o *ReviewReviewersInner) GetAuthorUuidOk() (*string, bool)`

GetAuthorUuidOk returns a tuple with the AuthorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorUuid

`func (o *ReviewReviewersInner) SetAuthorUuid(v string)`

SetAuthorUuid sets AuthorUuid field to given value.

### HasAuthorUuid

`func (o *ReviewReviewersInner) HasAuthorUuid() bool`

HasAuthorUuid returns a boolean if a field has been set.

### SetAuthorUuidNil

`func (o *ReviewReviewersInner) SetAuthorUuidNil(b bool)`

 SetAuthorUuidNil sets the value for AuthorUuid to be an explicit nil

### UnsetAuthorUuid
`func (o *ReviewReviewersInner) UnsetAuthorUuid()`

UnsetAuthorUuid ensures that no value is present for AuthorUuid, not even an explicit nil
### GetStatus

`func (o *ReviewReviewersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReviewReviewersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReviewReviewersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReviewReviewersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


