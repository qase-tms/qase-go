# ReviewBulkResponseAllOfResultItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReviewId** | Pointer to **NullableInt64** | ID of the created review. Null when the item failed. | [optional] 
**CaseId** | Pointer to **NullableInt64** | The &#x60;case_id&#x60; submitted with the item, echoed back for correlation. Null for new-case draft reviews. | [optional] 
**Error** | Pointer to **NullableString** | Failure reason. Null when the item was created. | [optional] 

## Methods

### NewReviewBulkResponseAllOfResultItems

`func NewReviewBulkResponseAllOfResultItems() *ReviewBulkResponseAllOfResultItems`

NewReviewBulkResponseAllOfResultItems instantiates a new ReviewBulkResponseAllOfResultItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewBulkResponseAllOfResultItemsWithDefaults

`func NewReviewBulkResponseAllOfResultItemsWithDefaults() *ReviewBulkResponseAllOfResultItems`

NewReviewBulkResponseAllOfResultItemsWithDefaults instantiates a new ReviewBulkResponseAllOfResultItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviewId

`func (o *ReviewBulkResponseAllOfResultItems) GetReviewId() int64`

GetReviewId returns the ReviewId field if non-nil, zero value otherwise.

### GetReviewIdOk

`func (o *ReviewBulkResponseAllOfResultItems) GetReviewIdOk() (*int64, bool)`

GetReviewIdOk returns a tuple with the ReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewId

`func (o *ReviewBulkResponseAllOfResultItems) SetReviewId(v int64)`

SetReviewId sets ReviewId field to given value.

### HasReviewId

`func (o *ReviewBulkResponseAllOfResultItems) HasReviewId() bool`

HasReviewId returns a boolean if a field has been set.

### SetReviewIdNil

`func (o *ReviewBulkResponseAllOfResultItems) SetReviewIdNil(b bool)`

 SetReviewIdNil sets the value for ReviewId to be an explicit nil

### UnsetReviewId
`func (o *ReviewBulkResponseAllOfResultItems) UnsetReviewId()`

UnsetReviewId ensures that no value is present for ReviewId, not even an explicit nil
### GetCaseId

`func (o *ReviewBulkResponseAllOfResultItems) GetCaseId() int64`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *ReviewBulkResponseAllOfResultItems) GetCaseIdOk() (*int64, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *ReviewBulkResponseAllOfResultItems) SetCaseId(v int64)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *ReviewBulkResponseAllOfResultItems) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### SetCaseIdNil

`func (o *ReviewBulkResponseAllOfResultItems) SetCaseIdNil(b bool)`

 SetCaseIdNil sets the value for CaseId to be an explicit nil

### UnsetCaseId
`func (o *ReviewBulkResponseAllOfResultItems) UnsetCaseId()`

UnsetCaseId ensures that no value is present for CaseId, not even an explicit nil
### GetError

`func (o *ReviewBulkResponseAllOfResultItems) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ReviewBulkResponseAllOfResultItems) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ReviewBulkResponseAllOfResultItems) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ReviewBulkResponseAllOfResultItems) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ReviewBulkResponseAllOfResultItems) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ReviewBulkResponseAllOfResultItems) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


