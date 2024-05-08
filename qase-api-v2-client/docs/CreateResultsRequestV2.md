# CreateResultsRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]ResultCreate**](ResultCreate.md) |  | [optional] 

## Methods

### NewCreateResultsRequestV2

`func NewCreateResultsRequestV2() *CreateResultsRequestV2`

NewCreateResultsRequestV2 instantiates a new CreateResultsRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResultsRequestV2WithDefaults

`func NewCreateResultsRequestV2WithDefaults() *CreateResultsRequestV2`

NewCreateResultsRequestV2WithDefaults instantiates a new CreateResultsRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *CreateResultsRequestV2) GetResults() []ResultCreate`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CreateResultsRequestV2) GetResultsOk() (*[]ResultCreate, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CreateResultsRequestV2) SetResults(v []ResultCreate)`

SetResults sets Results field to given value.

### HasResults

`func (o *CreateResultsRequestV2) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


