# ConfigurationGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Configurations** | Pointer to [**[]Configuration**](Configuration.md) |  | [optional] 

## Methods

### NewConfigurationGroup

`func NewConfigurationGroup() *ConfigurationGroup`

NewConfigurationGroup instantiates a new ConfigurationGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationGroupWithDefaults

`func NewConfigurationGroupWithDefaults() *ConfigurationGroup`

NewConfigurationGroupWithDefaults instantiates a new ConfigurationGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigurationGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ConfigurationGroup) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConfigurationGroup) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConfigurationGroup) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ConfigurationGroup) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetConfigurations

`func (o *ConfigurationGroup) GetConfigurations() []Configuration`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *ConfigurationGroup) GetConfigurationsOk() (*[]Configuration, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *ConfigurationGroup) SetConfigurations(v []Configuration)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *ConfigurationGroup) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


