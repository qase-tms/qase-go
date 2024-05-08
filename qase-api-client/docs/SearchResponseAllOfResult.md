# SearchResponseAllOfResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]SearchResponseAllOfResultEntities**](SearchResponseAllOfResultEntities.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewSearchResponseAllOfResult

`func NewSearchResponseAllOfResult(entities []SearchResponseAllOfResultEntities, total int32, ) *SearchResponseAllOfResult`

NewSearchResponseAllOfResult instantiates a new SearchResponseAllOfResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseAllOfResultWithDefaults

`func NewSearchResponseAllOfResultWithDefaults() *SearchResponseAllOfResult`

NewSearchResponseAllOfResultWithDefaults instantiates a new SearchResponseAllOfResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *SearchResponseAllOfResult) GetEntities() []SearchResponseAllOfResultEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *SearchResponseAllOfResult) GetEntitiesOk() (*[]SearchResponseAllOfResultEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *SearchResponseAllOfResult) SetEntities(v []SearchResponseAllOfResultEntities)`

SetEntities sets Entities field to given value.


### GetTotal

`func (o *SearchResponseAllOfResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchResponseAllOfResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchResponseAllOfResult) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


