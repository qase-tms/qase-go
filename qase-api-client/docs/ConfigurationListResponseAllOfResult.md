# ConfigurationListResponseAllOfResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Filtered** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Entities** | Pointer to [**[]ConfigurationGroup**](ConfigurationGroup.md) |  | [optional] 

## Methods

### NewConfigurationListResponseAllOfResult

`func NewConfigurationListResponseAllOfResult() *ConfigurationListResponseAllOfResult`

NewConfigurationListResponseAllOfResult instantiates a new ConfigurationListResponseAllOfResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationListResponseAllOfResultWithDefaults

`func NewConfigurationListResponseAllOfResultWithDefaults() *ConfigurationListResponseAllOfResult`

NewConfigurationListResponseAllOfResultWithDefaults instantiates a new ConfigurationListResponseAllOfResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ConfigurationListResponseAllOfResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ConfigurationListResponseAllOfResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ConfigurationListResponseAllOfResult) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ConfigurationListResponseAllOfResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetFiltered

`func (o *ConfigurationListResponseAllOfResult) GetFiltered() int32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *ConfigurationListResponseAllOfResult) GetFilteredOk() (*int32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *ConfigurationListResponseAllOfResult) SetFiltered(v int32)`

SetFiltered sets Filtered field to given value.

### HasFiltered

`func (o *ConfigurationListResponseAllOfResult) HasFiltered() bool`

HasFiltered returns a boolean if a field has been set.

### GetCount

`func (o *ConfigurationListResponseAllOfResult) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ConfigurationListResponseAllOfResult) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ConfigurationListResponseAllOfResult) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ConfigurationListResponseAllOfResult) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetEntities

`func (o *ConfigurationListResponseAllOfResult) GetEntities() []ConfigurationGroup`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *ConfigurationListResponseAllOfResult) GetEntitiesOk() (*[]ConfigurationGroup, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *ConfigurationListResponseAllOfResult) SetEntities(v []ConfigurationGroup)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *ConfigurationListResponseAllOfResult) HasEntities() bool`

HasEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


