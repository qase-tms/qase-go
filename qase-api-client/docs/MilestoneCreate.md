# MilestoneCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DueDate** | Pointer to **int64** | unix timestamp | [optional] 

## Methods

### NewMilestoneCreate

`func NewMilestoneCreate(title string, ) *MilestoneCreate`

NewMilestoneCreate instantiates a new MilestoneCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMilestoneCreateWithDefaults

`func NewMilestoneCreateWithDefaults() *MilestoneCreate`

NewMilestoneCreateWithDefaults instantiates a new MilestoneCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *MilestoneCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MilestoneCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MilestoneCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *MilestoneCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MilestoneCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MilestoneCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MilestoneCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *MilestoneCreate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MilestoneCreate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MilestoneCreate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MilestoneCreate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDueDate

`func (o *MilestoneCreate) GetDueDate() int64`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *MilestoneCreate) GetDueDateOk() (*int64, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *MilestoneCreate) SetDueDate(v int64)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *MilestoneCreate) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


