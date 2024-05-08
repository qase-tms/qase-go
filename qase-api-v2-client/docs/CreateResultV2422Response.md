# CreateResultV2422Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **string** |  | [optional] 
**ErrorFields** | Pointer to [**[]BaseErrorFieldResponseErrorFieldsInner**](BaseErrorFieldResponseErrorFieldsInner.md) |  | [optional] 

## Methods

### NewCreateResultV2422Response

`func NewCreateResultV2422Response() *CreateResultV2422Response`

NewCreateResultV2422Response instantiates a new CreateResultV2422Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResultV2422ResponseWithDefaults

`func NewCreateResultV2422ResponseWithDefaults() *CreateResultV2422Response`

NewCreateResultV2422ResponseWithDefaults instantiates a new CreateResultV2422Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *CreateResultV2422Response) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CreateResultV2422Response) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CreateResultV2422Response) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CreateResultV2422Response) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetErrorFields

`func (o *CreateResultV2422Response) GetErrorFields() []BaseErrorFieldResponseErrorFieldsInner`

GetErrorFields returns the ErrorFields field if non-nil, zero value otherwise.

### GetErrorFieldsOk

`func (o *CreateResultV2422Response) GetErrorFieldsOk() (*[]BaseErrorFieldResponseErrorFieldsInner, bool)`

GetErrorFieldsOk returns a tuple with the ErrorFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFields

`func (o *CreateResultV2422Response) SetErrorFields(v []BaseErrorFieldResponseErrorFieldsInner)`

SetErrorFields sets ErrorFields field to given value.

### HasErrorFields

`func (o *CreateResultV2422Response) HasErrorFields() bool`

HasErrorFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


