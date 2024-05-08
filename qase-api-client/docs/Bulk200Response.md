# Bulk200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Result** | Pointer to [**Bulk200ResponseAllOfResult**](Bulk200ResponseAllOfResult.md) |  | [optional] 

## Methods

### NewBulk200Response

`func NewBulk200Response() *Bulk200Response`

NewBulk200Response instantiates a new Bulk200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulk200ResponseWithDefaults

`func NewBulk200ResponseWithDefaults() *Bulk200Response`

NewBulk200ResponseWithDefaults instantiates a new Bulk200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Bulk200Response) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Bulk200Response) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Bulk200Response) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Bulk200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *Bulk200Response) GetResult() Bulk200ResponseAllOfResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Bulk200Response) GetResultOk() (*Bulk200ResponseAllOfResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Bulk200Response) SetResult(v Bulk200ResponseAllOfResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *Bulk200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


