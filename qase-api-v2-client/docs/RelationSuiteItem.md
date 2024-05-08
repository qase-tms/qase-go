# RelationSuiteItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**PublicId** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewRelationSuiteItem

`func NewRelationSuiteItem(title string, ) *RelationSuiteItem`

NewRelationSuiteItem instantiates a new RelationSuiteItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationSuiteItemWithDefaults

`func NewRelationSuiteItemWithDefaults() *RelationSuiteItem`

NewRelationSuiteItemWithDefaults instantiates a new RelationSuiteItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *RelationSuiteItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RelationSuiteItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RelationSuiteItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPublicId

`func (o *RelationSuiteItem) GetPublicId() int64`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *RelationSuiteItem) GetPublicIdOk() (*int64, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *RelationSuiteItem) SetPublicId(v int64)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *RelationSuiteItem) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### SetPublicIdNil

`func (o *RelationSuiteItem) SetPublicIdNil(b bool)`

 SetPublicIdNil sets the value for PublicId to be an explicit nil

### UnsetPublicId
`func (o *RelationSuiteItem) UnsetPublicId()`

UnsetPublicId ensures that no value is present for PublicId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


