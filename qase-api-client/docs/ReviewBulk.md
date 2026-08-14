# ReviewBulk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reviews** | [**[]ReviewCreate**](ReviewCreate.md) | Validated as a whole: if any item is invalid nothing is created. Otherwise each item is processed on its own and reported in the response. | 

## Methods

### NewReviewBulk

`func NewReviewBulk(reviews []ReviewCreate, ) *ReviewBulk`

NewReviewBulk instantiates a new ReviewBulk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewBulkWithDefaults

`func NewReviewBulkWithDefaults() *ReviewBulk`

NewReviewBulkWithDefaults instantiates a new ReviewBulk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviews

`func (o *ReviewBulk) GetReviews() []ReviewCreate`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *ReviewBulk) GetReviewsOk() (*[]ReviewCreate, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *ReviewBulk) SetReviews(v []ReviewCreate)`

SetReviews sets Reviews field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


