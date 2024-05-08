# Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewConfiguration

`func NewConfiguration() *Configuration`

NewConfiguration instantiates a new Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationWithDefaults

`func NewConfigurationWithDefaults() *Configuration`

NewConfigurationWithDefaults instantiates a new Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Configuration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Configuration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Configuration) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Configuration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Configuration) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Configuration) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Configuration) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Configuration) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


